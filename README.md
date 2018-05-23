# go-sqlr [![GoDoc](https://godoc.org/github.com/m1kc/go-sqlr?status.svg)](https://godoc.org/github.com/m1kc/go-sqlr) [![Go Report Card](https://goreportcard.com/badge/github.com/m1kc/go-sqlr)](https://goreportcard.com/report/github.com/m1kc/go-sqlr)

SQL query composer that does not stand in your way. Alpha quality, Postgres only, API not stabilized, use with care.

## Implemented features

### [SELECT](https://www.postgresql.org/docs/9.5/static/sql-select.html)

- [x] `WITH`
- [ ] `WITH RECURSIVE`
- [ ] `ALL, DISTINCT`
- [x] `FROM`
- [x] `WHERE`
- [x] `GROUP BY`
- [ ] `HAVING`
- [ ] `WINDOW`
- [ ] `UNION, INTERSECT, EXCEPT`
- [x] `ORDER BY`
- [x] `ORDER ... ASC/DESC`
- [x] `ORDER ... USING`
- [x] `ORDER ... NULLS`
- [x] `LIMIT`
- [ ] `OFFSET`
- [ ] `FETCH ROWS`
- [ ] `FOR UPDATE`
