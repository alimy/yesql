package yesql

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// default tag name of sql clause in sql file
const tagName = "name"

// A line may be blank, a tag, a comment or a query
const (
	lineBlank = iota
	lineQuery
	lineComment
	lineTag
)

var (
	// -- tag: $value
	reTag = regexp.MustCompile(`^\s*--\s*(.+)\s*:\s*(.+)`)

	// -- $comment
	reComment = regexp.MustCompile(`^\s*--\s*(.*)`)

	_ SQLParser = (*sqlParser)(nil)
	_ SQLQuery  = (*sqlParser)(nil)
)

// ParsedLine stores line type and value
// For example: parsedLine{Type=lineTag, Value="foo"}
type parsedLine struct {
	Type  int
	Tag   string
	Value string
}

type sqlParser struct {
	hooks      []func(query *Query) (*Query, error)
	queryMap   QueryMap
	scopeQuery ScopeQuery
}

func (s *sqlParser) AddHooks(hooks ...func(query *Query) (*Query, error)) {
	for _, h := range hooks {
		if h != nil {
			s.hooks = append(s.hooks, h)
		}
	}
}

func (s *sqlParser) ListQuery(namespace ...string) (QueryMap, error) {
	if len(namespace) == 0 {
		return s.queryMap, nil
	}
	ns := namespace[0]
	if len(ns) == 0 {
		return s.queryMap, nil
	}
	if qm, exist := s.scopeQuery[ns]; exist {
		return qm, nil
	}
	return nil, fmt.Errorf("no exist query list for namespace: %s", ns)
}

func (s *sqlParser) ListScope() ScopeQuery {
	return s.scopeQuery
}

func (s *sqlParser) AllQuery() []*Query {
	allQuery := make([]*Query, 0, len(s.queryMap))
	for _, query := range s.queryMap {
		allQuery = append(allQuery, query)
	}
	for _, qm := range s.scopeQuery {
		for _, query := range qm {
			allQuery = append(allQuery, query)
		}
	}
	return allQuery
}

func (s *sqlParser) SqlQuery(namespace string) (QueryMap, QueryMap, error) {
	if len(namespace) == 0 {
		return s.queryMap, s.queryMap, nil
	}
	if nsQuery, exist := s.scopeQuery[namespace]; exist {
		return s.queryMap, nsQuery, nil
	}
	return nil, nil, fmt.Errorf("no exist query list for namespace: %s", namespace)
}

// parseReader takes an io.Reader and returns Queries or an error.
func (s *sqlParser) ParseReader(reader io.Reader) (SQLQuery, error) {
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
				return nil, fmt.Errorf("Query is missing the 'name' tag: %s", line.Value)
			}

			q := line.Value
			query := s.queryMap[nameTag]
			if len(namespace) > 0 {
				query = s.scopeQuery[namespace][nameTag]
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
				queries := s.queryMap
				if len(tagSlice) > 1 {
					nameTag = tagSlice[0]
					if len(tagSlice[1]) > 0 {
						namespace = tagSlice[1]
						if _, exist := s.scopeQuery[namespace]; !exist {
							s.scopeQuery[namespace] = make(QueryMap)
						}
						queries = s.scopeQuery[namespace]
					}
				} else {
					nameTag = tagSlice[0]
				}
				// trim special '$' in start
				nameTag = strings.TrimLeft(nameTag, "$")

				if _, ok := queries[nameTag]; ok {
					return nil, fmt.Errorf("Duplicate tag %s = %s ", line.Tag, line.Value)
				}

				newQuery := &Query{
					Scope: namespace,
					Tags:  make(map[string]string),
				}
				if len(namespace) > 0 {
					s.scopeQuery[namespace][nameTag] = newQuery
				} else {
					s.queryMap[nameTag] = newQuery
				}
			} else {
				// Is there a name tag for this query?
				if len(queryLine) > 0 {
					return nil, errors.New("'name' should be the first tag")
				}

				queries := s.queryMap
				if len(namespace) > 0 {
					queries = s.scopeQuery[namespace]
				}

				// Has this tag already been used on this query?
				if _, ok := queries[nameTag].Tags[line.Tag]; ok {
					return nil, fmt.Errorf("Duplicate tag %s = %s ", line.Tag, line.Value)
				}

				if len(namespace) > 0 {
					s.scopeQuery[namespace][nameTag].Tags[line.Tag] = line.Value
				} else {
					s.queryMap[nameTag].Tags[line.Tag] = line.Value
				}
			}
		}
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	if err = s.checkQuery(); err != nil {
		return nil, err
	}
	if err = s.runHooks(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *sqlParser) checkQuery() error {
	for name, q := range s.queryMap {
		if q.Query == "" {
			return fmt.Errorf("'%s' is missing query", name)
		}
	}
	for ns, qs := range s.scopeQuery {
		for name, q := range qs {
			if q.Query == "" {
				return fmt.Errorf("'%s@%s' is missing query", name, ns)
			}
		}
	}
	return nil
}

func (s *sqlParser) runHooks() (err error) {
	for name, query := range s.queryMap {
		for _, hook := range s.hooks {
			if query, err = hook(query); err != nil {
				return fmt.Errorf("run hook failue name: %s err: %w", name, err)
			}
		}
		s.queryMap[name] = query
	}
	for ns, qs := range s.scopeQuery {
		for name, query := range qs {
			for _, hook := range s.hooks {
				if query, err = hook(query); err != nil {
					return fmt.Errorf("run hook failue name: %s@%s err: %w", name, ns, err)
				}
			}
			s.scopeQuery[ns][name] = query
		}
	}
	return
}

func parseLine(line string) parsedLine {
	line = strings.Trim(line, " ")

	if line == "" {
		return parsedLine{lineBlank, "", ""}
	} else if matches := reTag.FindStringSubmatch(line); len(matches) > 1 {
		return parsedLine{lineTag, matches[1], matches[2]}
	} else if matches := reComment.FindStringSubmatch(line); len(matches) > 0 {
		return parsedLine{lineComment, "", matches[1]}
	}

	return parsedLine{lineQuery, "", line}
}

func newSQLParser(hooks ...func(query *Query) (*Query, error)) SQLParser {
	obj := &sqlParser{
		queryMap:   make(QueryMap),
		scopeQuery: make(ScopeQuery),
	}
	obj.hooks = append(obj.hooks, hooks...)
	return obj
}
