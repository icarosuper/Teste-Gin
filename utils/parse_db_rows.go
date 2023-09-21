package utils

import (
	"database/sql"
	"fmt"
)

func ParseDbRows[T any](rows *sql.Rows) ([]T, error) {
	var results []T

	columns, err := rows.Columns()
	if err != nil {
		return results, err
	}

	colNum := len(columns)

	var values = make([]interface{}, colNum)

	for i, _ := range values {
		var ii interface{}
		values[i] = &ii
	}

	for rows.Next() {
		err = rows.Scan(values...)

		fmt.Println(values)

		for i, colName := range columns {
			var rawValue = *(values[i].(*interface{}))

			fmt.Println(colName, rawValue)
		}
	}

	return results, err
}
