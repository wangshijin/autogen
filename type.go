package main

/**** type map
+------------+------------+------------+
+   table    +    sql     +     go     +
+------------+------------+------------+
+ int8       + tinyint    + int8       +
+ int16      + smallint   + int16      +
+ int32      + int        + int32      +
+ int64      + bigint     + int64      +
+ int        + int        + int        +
+ float      + float      + float32    +
+ double     + double     + float64    +
+ string     + varchar(30)+ string     +
+ varcharn   + varchar(n) + string     +
+ text       + text       + string     +
+------------+------------+------------+
*/

import (
	"regexp"
	"strconv"
)

var namereg = `[a-zA-Z_0-9]+`
var typereg = `int8|int16|int32|int64|int|float|double|string|varchar[1-9][0-9]*|text`

func go_type(table_type string) string {
	switch table_type {
	case "int8":
		return "int8"
	case "int16":
		return "int16"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "int":
		return "int"
	case "float":
		return "float32"
	case "double":
		return "float64"
	default:
		return "string"
	}
}

func sql_type(table_type string) string {
	switch table_type {
	case "int8":
		return "tinyint"
	case "int16":
		return "smallint"
	case "int32":
		return "int"
	case "int64":
		return "bigint"
	case "int":
		return "int"
	case "float":
		return "float"
	case "double":
		return "double"
	case "string":
		return "varchar(30)"
	case "text":
		return "text"
	default:
		return parse_varchar(table_type)
	}
}

func parse_varchar(table_type string) string {
	reg := regexp.MustCompile("varchar([1-9][0-9]*)")
	if !reg.MatchString(table_type) {
		return ""
	}
	strs := reg.FindStringSubmatch(table_type)
	if len(strs) != 2 {
		return ""
	}
	numstr := strs[1]
	num, err := strconv.Atoi(numstr)
	if err != nil {
		return ""
	}
	if num > 255 {
		return "text"
	}
	return "varchar(" + numstr + ")"
}

func check_field_type(table_type string) bool {
	switch table_type {
	case "int8":
		return true
	case "int16":
		return true
	case "int32":
		return true
	case "int64":
		return true
	case "int":
		return true
	case "float":
		return true
	case "double":
		return true
	case "string":
		return true
	case "text":
		return true
	default:
		reg := regexp.MustCompile("varchar([1-9][0-9]*)")
		if reg.MatchString(table_type) {
			return true
		} else {
			return false
		}
	}
}
