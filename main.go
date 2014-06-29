package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var table_dir = flag.String("table-dir", "./table", "input dir of .tables files")
var sql_dir = flag.String("sql-dir", "./sql", "output dir of .sql files")
var go_dir = flag.String("go-dir", "./model", "output dir of .go files")
var model_pkg = flag.String("model-pkg", "model", "model pkg")

func main() {
	flag.Parse()

	fmt.Println("table_dir =", *table_dir)
	fmt.Println("sql_dir =", *sql_dir)
	fmt.Println("go_dir =", *go_dir)
	fmt.Println("model_pkg =", *model_pkg)

	table_files := GetTableFiles(*table_dir)
	fmt.Println("table_files =", table_files)

	for _, fn := range table_files {
		buf, err := ioutil.ReadFile(fn)
		if err != nil {
			continue
		}
		s := string(buf)

		data, err := gen_data(s)
		if err != nil {
			fmt.Printf("gen_data error:%v\n", err)
			return
		}

		go_data := data.CopyTo(go_type)
		sql_data := data.CopyTo(sql_type)

		gen_go(go_data, *go_dir, *model_pkg)
		gen_sql(sql_data, *sql_dir)
	}
}
