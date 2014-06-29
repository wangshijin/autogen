package main

import (
	"fmt"
	"os"
	"strings"
)

func gen_go(data *Sql, dir, model_pkg string) bool {
	dir = strings.TrimSuffix(dir, "/")
	path := dir + "/" + Lower(data.Name)
	fmt.Printf("dest path of .go files:%s\n", path)
	if err := os.MkdirAll(path, 0777); err != nil {
		fmt.Errorf("mkdir %s error:%v\n", path, err)
		return false
	}

	// gen record.go
	str := gen_go_record(dir)
	filename := dir + "/record.go"
	if err := write_to_file(str, filename); err != nil {
		return false
	}

	// gen table .go files
	for _, table := range data.Tables {
		str = gen_go_table(data.Name, model_pkg, &table)
		filename = path + "/" + Lower(table.Name) + ".go"
		if err := write_to_file(str, filename); err != nil {
			continue
		}

	}

	return true
}

func gen_go_record(dir string) string {
	pkg := FileName(dir)

	// gen head
	str := fmt.Sprintf("package %s\n\n", pkg)
	imports := []string{`"database/sql"`}
	str += "import (\n"
	for _, v := range imports {
		str += fmt.Sprintf("\t%s\n", v)
	}
	str += ")\n\n"

	// gen type Record interface
	str += "type Record interface {\n"
	str += "	InsertSql() string\n"
	str += "	UpdateSql() string\n"
	str += "	DeleteSql() string\n"
	str += "	Scan(*sql.Rows) error\n"
	// add more ......

	str += "}\n\n"

	return str
}

func gen_go_table(pkg, model_pkg string, tb *Table) string {
	str := gen_go_head(Lower(pkg), model_pkg)
	str += gen_go_typedef(tb)
	str += gen_go_new_func(tb)
	if len(tb.Fields) == 0 {
		return str
	}
	str += gen_go_get_set_func(tb)
	str += gen_go_insert_sql_func(tb)
	str += gen_go_update_sql_func(tb)
	str += gen_go_delete_sql_func(tb)
	str += gen_go_scan_func(tb)
	// add more ......
	return str
}

func gen_go_head(pkg, model_pkg string) string {
	str := fmt.Sprintf("package %s\n\n", pkg)

	imports := []string{`"database/sql"`, `"fmt"`, fmt.Sprintf(`"%s"`, model_pkg)}

	str += "import (\n"
	for _, v := range imports {
		str += fmt.Sprintf("\t%s\n", v)
	}
	str += ")\n\n"
	return str
}

func gen_go_typedef(tb *Table) string {
	str := fmt.Sprintf("type %s struct {\n", tb.Name)

	for _, v := range tb.Fields {
		str += fmt.Sprintf("\t%s\t%s\n", v.Name, v.Type)
	}
	str += "}\n\n"
	return str
}

func gen_go_new_func(tb *Table) string {
	// func makeTable() model.Record
	str := fmt.Sprintf("func make%s() model.Record {\n", tb.Name)
	str += fmt.Sprintf("\treturn new(%s)\n}\n\n", tb.Name)

	// func NewTable() *Table
	str += fmt.Sprintf("func New%s() *%s {\n", tb.Name, tb.Name)
	str += fmt.Sprintf("\treturn new(%s)\n}\n\n", tb.Name)

	return str
}

func gen_go_get_set_func(tb *Table) string {
	str := ""

	// func TableName() string
	str += fmt.Sprintf("func (this %s) TableName() string {\n",
		tb.Name)
	str += fmt.Sprintf("\treturn \"%s\"\n", tb.Name)
	str += "}\n\n"

	// func KeyName() string
	str += fmt.Sprintf("func (this %s) KeyName() string {\n",
		tb.Name)
	str += fmt.Sprintf("\treturn \"%s\"\n", (tb.Fields[0].Name))
	str += "}\n\n"

	// func Key
	str += fmt.Sprintf("func (this %s) Key() %s {\n",
		tb.Name, tb.Fields[0].Type)
	str += fmt.Sprintf("\treturn this.%s\n", tb.Fields[0].Name)
	str += "}\n\n"

	// get
	for _, v := range tb.Fields {
		str += fmt.Sprintf("func (this %s) Get%s() %s {\n",
			tb.Name, UpperFirst(v.Name), v.Type)
		str += fmt.Sprintf("\treturn this.%s\n", v.Name)
		str += "}\n\n"
	}

	// set
	for _, v := range tb.Fields {
		str += fmt.Sprintf("func (this *%s) Set%s(%s_ %s) {\n",
			tb.Name, UpperFirst(v.Name), v.Name, v.Type)
		str += fmt.Sprintf("\tthis.%s = %s_\n", v.Name, v.Name)
		str += "}\n\n"
	}

	return str
}

func gen_go_insert_sql_func(tb *Table) string {
	str := fmt.Sprintf("func (this %s) InsertSql() string {\n", tb.Name)
	str += fmt.Sprintf(`	return fmt.Sprintf("insert into %s values(`, tb.Name)

	synx := ""
	values := ""
	for i, field := range tb.Fields {
		if i != 0 {
			synx += ","
			values += ","
		}
		if is_string(field.Type) {
			synx += "\\\"%s\\\""
		} else {
			synx += "%d"
		}
		values += "this." + field.Name
	}

	str += fmt.Sprintf("%s)\", %s)\n", synx, values)
	str += "}\n\n"
	return str
}

func gen_go_update_sql_func(tb *Table) string {
	str := fmt.Sprintf("func (this %s) UpdateSql() string {\n", tb.Name)

	kv := ""
	values := ""
	first_kv := ""
	for i, field := range tb.Fields {
		if i != 0 {
			kv += ","
			values += ","
		}
		new_kv := ""
		if is_string(field.Type) {
			new_kv += field.Name + "=\\\"%s\\\""
		} else {
			new_kv += field.Name + "=%d"
		}
		kv += new_kv
		values += "this." + field.Name
		if i == 0 {
			first_kv = new_kv
		}
	}

	values += ",this.Key()"
	str += fmt.Sprintf(`	return fmt.Sprintf("update %s set %s where %s`,
		tb.Name, kv, first_kv)
	str += fmt.Sprintf("\",%s)\n}\n\n", values)
	return str
}

func gen_go_delete_sql_func(tb *Table) string {
	str := fmt.Sprintf("func (this %s) DeleteSql() string {\n", tb.Name)
	first_kv := ""
	if is_string(tb.Fields[0].Type) {
		first_kv += tb.Fields[0].Name + "=\\\"%s\\\""
	} else {
		first_kv += tb.Fields[0].Name + "=%d"
	}
	str += fmt.Sprintf(`	return fmt.Sprintf("delete from %s where %s"`,
		tb.Name, first_kv)
	str += fmt.Sprintf(",this.Key())\n}\n\n")
	return str
}

func gen_go_scan_func(tb *Table) string {
	str := fmt.Sprintf("func (this *%s) Scan(rows *sql.Rows) error {\n", tb.Name)

	arglist := ""
	for i, fd := range tb.Fields {
		if i != 0 {
			arglist += ","
		}
		arglist += fmt.Sprintf("&(this.%s)", fd.Name)
	}

	str += fmt.Sprintf("\tif err := rows.Scan(%s); err != nil {\n", arglist)
	str += fmt.Sprintf("\t\treturn err\n\t}\n")
	str += "\treturn nil\n}\n\n"
	return str
}

func is_string(tp string) bool {
	switch tp {
	case "string":
		return true
	default:
		return false
	}
}

func write_to_file(src string, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Errorf(`write file "%s" error:%v\n`, filename, err)
		return err
	}
	file.WriteString(src)
	file.Close()
	return nil
}
