package yesql

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

var (
	_ PrepareScanner = (*prepareScanner)(nil)
)

type prepareScanner struct {
	prepareHook PrepareHook
}

func (s *prepareScanner) SetPrepareHook(hook PrepareHook) {
	if hook != nil {
		s.prepareHook = hook
	}
}

func (s *prepareScanner) Scan(obj any, query SQLQuery) error {
	return s.scan(nil, obj, query)
}

func (s *prepareScanner) ScanContext(ctx context.Context, obj any, query SQLQuery) error {
	return s.scan(ctx, obj, query)
}

func (s *prepareScanner) scan(ctx context.Context, obj any, query SQLQuery) error {
	ob := reflect.ValueOf(obj)
	if ob.Kind() == reflect.Ptr {
		ob = ob.Elem()
	}
	if ob.Kind() != reflect.Struct {
		return fmt.Errorf("Failed to apply SQL statements to struct. Non struct type: %T", ob)
	}
	// Go through every field in the struct and look for it in the Args map.
	var namespace string
	fieldValues := make(map[string]reflect.Value, ob.NumField())
	for i := 0; i < ob.NumField(); i++ {
		if f := ob.Field(i); f.IsValid() {
			field := ob.Type().Field(i)
			if tag := field.Tag.Get(_structTag); tag != "" {
				// Extract the value of the `query` tag.
				var (
					tg       = strings.Split(tag, ",")
					tagVaule string
				)
				if len(tg) == 2 {
					if tg[0] != "-" && tg[0] != "" {
						tagVaule = tg[0]
					}
				} else {
					tagVaule = tg[0]
				}
				if field.Type.Name() == "Namespace" {
					namespace = tagVaule
					continue
				}
				if f.CanSet() {
					fieldValues[tagVaule] = f
				} else {
					return fmt.Errorf("query field '%s' is unexported", ob.Type().Field(i).Name)
				}
			}
		}
	}
	qs, err := query.ListQuery(namespace)
	if err != nil {
		return err
	}
	for name, value := range fieldValues {
		qv, exist := qs[name]
		if !exist {
			return fmt.Errorf("query '%s' not found in query map with namespace: %s", name, namespace)
		}
		var v any
		if ctx == nil {
			v, err = s.prepareHook.Prepare(value.Type(), qv.Query)
		} else {
			v, err = s.prepareHook.PrepareContext(ctx, value.Type(), qv.Query)
		}
		if err != nil {
			return err
		}
		value.Set(reflect.ValueOf(v))
	}
	return nil
}

// NewPrepareScanner create prepare scnanner instance
func NewPrepareScanner(prepareHook PrepareHook) PrepareScanner {
	return &prepareScanner{
		prepareHook: prepareHook,
	}
}
