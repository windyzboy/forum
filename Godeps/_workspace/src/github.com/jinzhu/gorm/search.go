package gorm

import "fmt"

type search struct {
	db              *DB
	whereConditions []map[string]interface{}
	orConditions    []map[string]interface{}
	notConditions   []map[string]interface{}
	havingCondition map[string]interface{}
	initAttrs       []interface{}
	assignAttrs     []interface{}
	selects         map[string]interface{}
	omits           []string
	orders          []string
	joins           string
	preload         map[string][]interface{}
	offset          string
	limit           string
	group           string
	tableName       string
	raw             bool
	Unscoped        bool
}

func (s *search) clone() *search {
	return &search{
		preload:         s.preload,
		whereConditions: s.whereConditions,
		orConditions:    s.orConditions,
		notConditions:   s.notConditions,
		havingCondition: s.havingCondition,
		initAttrs:       s.initAttrs,
		assignAttrs:     s.assignAttrs,
		selects:         s.selects,
		omits:           s.omits,
		orders:          s.orders,
		joins:           s.joins,
		offset:          s.offset,
		limit:           s.limit,
		group:           s.group,
		tableName:       s.tableName,
		raw:             s.raw,
		Unscoped:        s.Unscoped,
	}
}

func (s *search) Where(query interface{}, values ...interface{}) *search {
	s.whereConditions = append(s.whereConditions, map[string]interface{}{"query": query, "args": values})
	return s
}

func (s *search) Not(query interface{}, values ...interface{}) *search {
	s.notConditions = append(s.notConditions, map[string]interface{}{"query": query, "args": values})
	return s
}

func (s *search) Or(query interface{}, values ...interface{}) *search {
	s.orConditions = append(s.orConditions, map[string]interface{}{"query": query, "args": values})
	return s
}

func (s *search) Attrs(attrs ...interface{}) *search {
	s.initAttrs = append(s.initAttrs, toSearchableMap(attrs...))
	return s
}

func (s *search) Assign(attrs ...interface{}) *search {
	s.assignAttrs = append(s.assignAttrs, toSearchableMap(attrs...))
	return s
}

func (s *search) Order(value string, reorder ...bool) *search {
	if len(reorder) > 0 && reorder[0] {
		s.orders = []string{value}
	} else {
		s.orders = append(s.orders, value)
	}
	return s
}

func (s *search) Select(query interface{}, args ...interface{}) *search {
	s.selects = map[string]interface{}{"query": query, "args": args}
	return s
}

func (s *search) Omit(columns ...string) *search {
	s.omits = columns
	return s
}

func (s *search) Limit(value interface{}) *search {
	s.limit = s.getInterfaceAsSql(value)
	return s
}

func (s *search) Offset(value interface{}) *search {
	s.offset = s.getInterfaceAsSql(value)
	return s
}

func (s *search) Group(query string) *search {
	s.group = s.getInterfaceAsSql(query)
	return s
}

func (s *search) Having(query string, values ...interface{}) *search {
	s.havingCondition = map[string]interface{}{"query": query, "args": values}
	return s
}

func (s *search) Joins(query string) *search {
	s.joins = query
	return s
}

func (s *search) Preload(column string, values ...interface{}) *search {
	if s.preload == nil {
		s.preload = map[string][]interface{}{}
	}
	s.preload[column] = values
	return s
}

func (s *search) Raw(b bool) *search {
	s.raw = b
	return s
}

func (s *search) unscoped() *search {
	s.Unscoped = true
	return s
}

func (s *search) Table(name string) *search {
	s.tableName = name
	return s
}

func (s *search) getInterfaceAsSql(value interface{}) (str string) {
	switch value.(type) {
	case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		str = fmt.Sprintf("%v", value)
	default:
		s.db.err(InvalidSql)
	}

	if str == "-1" {
		return ""
	}
	return
}
