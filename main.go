package main

import (
	"fmt"
)

func main() {
	s := `/****************
autogen .sql files and .go files
*/

// database name
database MyDB;

// tables

struct Test {
	int8 field1;
	int16 field2;
	int32 field3;
	int64 field4;
	int field5;
	float field6;
	double field7;
	string field8;
	varchar5 field9;
	varchar35 field10;
	text field11;
};

struct Shop {
	int itemId;
	string name;
	float price;
};
`

	data, err := gen_data(s)

	if err != nil {
		fmt.Printf("gen_data error:%v\n", err)
		return
	}

	go_data := data.CopyTo(go_type)
	sql_data := data.CopyTo(sql_type)

	gen_go(go_data, "./model", "gensql/model")
	gen_sql(sql_data, "./sql")
}
