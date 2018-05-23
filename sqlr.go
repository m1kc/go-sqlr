package sqlr

import (
	"fmt"
)

func sample() {
	q := Select()
	q.With("random_number", "5")
	q.Select("id")
	q.Select("field")
	q.Select("fio")
	q.From("data.contacts")
	q.From("something_else")
	q.WhereAnd("id < $%v OR id > $%v", 10, 20)
	q.WhereAnd("1 = 2")
	q.WhereAnd("id = $%v", 666)
	q.OrderBy("id")
	q.OrderDirection("DESC")
	q.Limit(50)

	sql, args := q.Get()
	fmt.Println(sql)
	fmt.Println("")
	fmt.Println(args)
}

func Select() SelectQuery {
	return &SelectQueryImpl{
		witcher:   make(map[string]string),
		selectr:   make([]string, 0, 20),
		from:      make([]string, 0, 20),
		where:     make([]string, 0, 20),
		whereArgs: make([]interface{}, 0, 20),
	}
}

type SelectQuery interface {
	With(string, string)
	Select(string)
	From(string)
	WhereAnd(string, ...interface{})
	GroupBy(string)
	OrderBy(string)
	OrderDirection(string)
	Limit(uint64)

	SQL() string
	Args() []interface{}
	Get() (string, []interface{})

	Tail() string
}

type SelectQueryImpl struct {
	witcher        map[string]string
	selectr        []string
	from           []string
	where          []string
	whereArgs      []interface{}
	groupBy        string
	orderBy        string
	orderDirection string
	limit          uint64
}

func (s *SelectQueryImpl) With(name string, value string) {
	s.witcher[name] = value
}

func (s *SelectQueryImpl) Select(x string) {
	s.selectr = append(s.selectr, x)
}

func (s *SelectQueryImpl) From(x string) {
	s.from = append(s.from, x)
}

func (s *SelectQueryImpl) WhereAnd(sql string, args ...interface{}) {
	s.where = append(s.where, sql)
	// TODO: check number of $%v
	for _, arg := range args {
		s.whereArgs = append(s.whereArgs, arg)
	}
}

func (s *SelectQueryImpl) GroupBy(x string) {
	s.groupBy = x
}

func (s *SelectQueryImpl) OrderBy(x string) {
	s.orderBy = x
}

func (s *SelectQueryImpl) OrderDirection(x string) {
	s.orderDirection = x
}

func (s *SelectQueryImpl) Limit(x uint64) {
	s.limit = x
}

func (s *SelectQueryImpl) formatWith() (ret string) {
	first := true
	for name, value := range s.witcher {
		if first {
			first = false
		} else {
			ret += ", "
		}
		ret += fmt.Sprintf("WITH %s AS ( %s )", name, value)
	}

	if len(s.witcher) > 0 {
		ret += "\n"
	}

	return
}

func (s *SelectQueryImpl) formatSelect() (ret string) {
	ret += "SELECT "
	first := true
	for _, value := range s.selectr {
		if first {
			first = false
		} else {
			ret += ", "
		}
		ret += value
	}
	return
}

func (s *SelectQueryImpl) formatFrom() (ret string) {
	if len(s.from) > 0 {
		ret += "\n"
		ret += "FROM "
		first := true
		for _, value := range s.from {
			if first {
				first = false
			} else {
				ret += ", "
			}
			ret += value
		}
	}
	return
}

func (s *SelectQueryImpl) formatWhere() (ret string) {
	if len(s.where) > 0 {
		ret += "\n"
		ret += "WHERE "
		firstWhere := true
		for _, value := range s.where {
			if firstWhere {
				firstWhere = false
			} else {
				ret += "\n"
				ret += "  AND "
			}
			ret += fmt.Sprintf("( %s )", value)
		}
	}
	return
}

func (s *SelectQueryImpl) SQL() (ret string) {
	ret += s.formatWith()
	ret += s.formatSelect()
	ret += s.Tail()
	return
}

func (s *SelectQueryImpl) Tail() (ret string) {
	ret += s.formatFrom()
	ret += s.formatWhere()

	if s.groupBy != "" {
		ret += "\n"
		ret += fmt.Sprintf("GROUP BY %s", s.groupBy)
	}

	if s.orderBy != "" {
		ret += "\n"
		ret += fmt.Sprintf("ORDER BY %s %s", s.orderBy, s.orderDirection)
	}

	if s.limit != uint64(0) {
		ret += "\n"
		ret += fmt.Sprintf("LIMIT %d", s.limit)
	}

	// $%v -> $1, $2, $3...
	n := len(s.Args())
	numbers := make([]interface{}, 0, n)
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}
	ret = fmt.Sprintf(ret, numbers...)

	return
}

func (s *SelectQueryImpl) Args() []interface{} {
	//ret := make([]interface{}, 0, 0)
	return s.whereArgs
	//return ret
}

func (s *SelectQueryImpl) Get() (string, []interface{}) {
	return s.SQL(), s.Args()
}
