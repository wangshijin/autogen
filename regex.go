package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrInvalidDBName = errors.New("invalid db name")
var ErrInvalidTableName = errors.New("invalid table name")
var ErrInvalidFieldName = errors.New("invalid field name")
var ErrInvalidFieldType = errors.New("invalid field type")

// gen data
func gen_data(s string) (data *Sql, err error) {
	remove_comment(&s)

	dbname := take_database_name(&s)
	if !check_name(dbname) {
		return nil, ErrInvalidDBName
	}

	data = NewSql(dbname)

	for {
		table_str := take_table(&s)
		if len(table_str) != 2 {
			break
		}
		if !check_name(table_str[0]) {
			return nil, ErrInvalidTableName
		}

		table := NewTable(table_str[0])
		table.Fields = take_fields(&table_str[1])

		for _, fd := range table.Fields {
			if !check_name(fd.Name) {
				return nil, ErrInvalidFieldName
			}
			if !check_field_type(fd.Type) {
				return nil, ErrInvalidFieldType
			}
		}

		data.AddTable(*table)
	}

	return data, nil
}

// remove comment
func remove_comment(s *string) {
	// remove `/**/`
	reg0 := regexp.MustCompile(`(?s)(/\*.*?\*/)`)
	*s = reg0.ReplaceAllString(*s, "")

	// remove `//`
	reg1 := regexp.MustCompile(`(?m://.*$)`)
	*s = reg1.ReplaceAllString(*s, "")
}

// get database name
func take_database_name(s *string) string {
	regstr := fmt.Sprintf(`[ \t\n]*?database[ \t]+(%s)[ \t]*?;`, namereg)
	reg := regexp.MustCompile(regstr)
	res := reg.FindStringSubmatch(*s)
	if len(res) > 1 && res[1] != "" {
		*s = strings.TrimPrefix(*s, res[0])
		return res[1]
	}
	return ""
}

// take table
func take_table(s *string) []string {
	regstr := fmt.Sprintf(`[ \t\n]*?struct[ \t]+(%s)[ \t]*?\{(?s)(.*?)\};`, namereg)
	reg := regexp.MustCompile(regstr)
	res := reg.FindStringSubmatch(*s)
	if len(res) > 2 && res[1] != "" {
		*s = strings.TrimPrefix(*s, res[0])
		return res[1:3]
	}
	return nil
}

// take fields
func take_fields(s *string) []Field {
	braces := fmt.Sprintf(`[ \t\n]*?(%s)[ \t]+?(%s)[ \t]*?;[ \t\n]*?`, typereg, namereg)
	reg := regexp.MustCompile(braces)
	fields := make([]Field, 0)
	for {
		res := reg.FindStringSubmatch(*s)
		if len(res) > 2 && check_field_type(res[1]) && check_name(res[2]) {
			*s = strings.TrimPrefix(*s, res[0])
			fields = append(fields, Field{res[1], res[2]})
		} else {
			break
		}
	}
	return fields
}

// check
func check_name(name string) bool {
	return name != ""
}
