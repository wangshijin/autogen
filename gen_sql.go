package main

import (
	"fmt"
	"os"
	"strings"
)

func gen_sql(data *Sql, dir string) bool {
	dir = strings.TrimSuffix(dir, "/")
	path := dir + "/" + Lower(data.Name)
	if err := os.MkdirAll(path, 0777); err != nil {
		fmt.Errorf("mkdir %s error:%v\n", path, err)
		return false
	}

	// gen table .sql files
	for _, table := range data.Tables {
		str := gen_sql_table(&table)
		filename := path + "/" + Lower(table.Name) + ".sql"
		if err := write_to_file(str, filename); err != nil {
			continue
		}

	}

	return true
}

func gen_sql_table(tb *Table) string {
	str := fmt.Sprintf("create table %s (\n", tb.Name)
	for i, fd := range tb.Fields {
		str += fmt.Sprintf("\t%s\t%s", fd.Name, fd.Type)
		if i+1 != len(tb.Fields) {
			str += ",\n"
		} else {
			str += ");\n\n"
		}
	}
	return str
}
