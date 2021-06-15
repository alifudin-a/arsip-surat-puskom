package helper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/structs"
)

// QueryBuilderUpdate query builder for updating data
func QueryBuilderUpdate(table string, arg interface{}) (string, []interface{}) {
	sql := "UPDATE " + table + " SET "
	s := structs.Fields(arg)
	var values []interface{}
	var where []string
	key := ""
	j := 0
	for _, f := range s {
		if f.IsExported() {
			if !f.IsZero() {
				j++
				if strings.ToLower(f.Name()) != "id" {
					field := f.Tag("db")
					values = append(values, f.Value())
					where = append(where, fmt.Sprintf("%s = $%d", field, j))
				} else {
					values = append(values, f.Value())
					key = " WHERE id " + " = $" + strconv.Itoa(j)
				}
			}
		}
	}

	sql = sql + strings.Join(where, ", ") + key
	sql = sql + " RETURNING *;"
	return sql, values
}
