package yesql

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	tagName = "name"
)

var (
	_ SQLQuery = (*sqlQuery)(nil)
)

type sqlQuery struct {
	hooks           []func(query *Query) (*Query, error)
	defaultQueryMap QueryMap
	nsQueryMap      map[string]QueryMap
}

func (s *sqlQuery) AddHooks(hooks ...func(query *Query) (*Query, error)) {
	for _, h := range hooks {
		if h != nil {
			s.hooks = append(s.hooks, h)
		}
	}
}

func (s *sqlQuery) ListQuery(namespace string) (QueryMap, error) {
	if len(namespace) == 0 {
		return s.defaultQueryMap, nil
	}
	if qm, exist := s.nsQueryMap[namespace]; exist {
		return qm, nil
	}
	return nil, fmt.Errorf("no exist query list for namespace: %s", namespace)
}

// parseReader takes an io.Reader and returns Queries or an error.
func (s *sqlQuery) parseReader(reader io.Reader) error {
	var (
		nameTag   string
		namespace string
		queryLine string
		scanner   = bufio.NewScanner(reader)
	)

	for scanner.Scan() {
		line := parseLine(scanner.Text())

		switch line.Type {
		case lineBlank, lineComment:
			// Ignore.
			continue

		case lineQuery:
			// Got a query but no preceding name tag.
			if nameTag == "" {
				return fmt.Errorf("Query is missing the 'name' tag: %s", line.Value)
			}

			q := line.Value
			query := s.defaultQueryMap[nameTag]
			if len(namespace) > 0 {
				query = s.nsQueryMap[namespace][nameTag]
			}

			// If query is multiline.
			if query.Query != "" {
				q = " " + q
			}
			query.Query += q
			queryLine = query.Query
		case lineTag:
			// Has this name already been read?
			if line.Tag == tagName {
				// reset namespace and queryLine first
				namespace = ""
				queryLine = ""
				tagSlice := strings.Split(line.Value, "@")
				queries := s.defaultQueryMap
				if len(tagSlice) > 1 {
					nameTag = tagSlice[0]
					if len(tagSlice[1]) > 0 {
						namespace = tagSlice[1]
						if _, exist := s.nsQueryMap[namespace]; !exist {
							s.nsQueryMap[namespace] = make(QueryMap)
						}
						queries = s.nsQueryMap[namespace]
					}
				} else {
					nameTag = tagSlice[0]
				}

				if _, ok := queries[nameTag]; ok {
					return fmt.Errorf("Duplicate tag %s = %s ", line.Tag, line.Value)
				}

				newQuery := &Query{Tags: make(map[string]string)}
				if len(namespace) > 0 {
					s.nsQueryMap[namespace][nameTag] = newQuery
				} else {
					s.defaultQueryMap[nameTag] = newQuery
				}
			} else {
				// Is there a name tag for this query?
				if len(queryLine) > 0 {
					return errors.New("'name' should be the first tag")
				}

				queries := s.defaultQueryMap
				if len(namespace) > 0 {
					queries = s.nsQueryMap[namespace]
				}

				// Has this tag already been used on this query?
				if _, ok := queries[nameTag].Tags[line.Tag]; ok {
					return fmt.Errorf("Duplicate tag %s = %s ", line.Tag, line.Value)
				}

				if len(namespace) > 0 {
					s.nsQueryMap[namespace][nameTag].Tags[line.Tag] = line.Value
				} else {
					s.defaultQueryMap[nameTag].Tags[line.Tag] = line.Value
				}
			}
		}
	}

	err := scanner.Err()
	if err != nil {
		return err
	}
	if err = s.checkQuery(); err != nil {
		return err
	}
	if err = s.runHooks(); err != nil {
		return err
	}
	return nil
}

func (s *sqlQuery) checkQuery() error {
	for name, q := range s.defaultQueryMap {
		if q.Query == "" {
			return fmt.Errorf("'%s' is missing query", name)
		}
	}
	for ns, qs := range s.nsQueryMap {
		for name, q := range qs {
			if q.Query == "" {
				return fmt.Errorf("'%s@%s' is missing query", name, ns)
			}
		}
	}
	return nil
}

func (s *sqlQuery) runHooks() (err error) {
	for name, query := range s.defaultQueryMap {
		for _, hook := range s.hooks {
			if query, err = hook(query); err != nil {
				return fmt.Errorf("run hook failue name: %s err: %w", name, err)
			}
		}
		s.defaultQueryMap[name] = query
	}
	for ns, qs := range s.nsQueryMap {
		for name, query := range qs {
			for _, hook := range s.hooks {
				if query, err = hook(query); err != nil {
					return fmt.Errorf("run hook failue name: %s@%s err: %w", name, ns, err)
				}
			}
			s.nsQueryMap[ns][name] = query
		}
	}
	return
}

func newSqlQuery() *sqlQuery {
	return &sqlQuery{
		defaultQueryMap: make(QueryMap),
		nsQueryMap:      make(map[string]QueryMap),
	}
}
