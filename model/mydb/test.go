package mydb

import (
	"database/sql"
	"fmt"
	"gensql/model"
)

type Test struct {
	field1	int8
	field2	int16
	field3	int32
	field4	int64
	field5	int
	field6	float32
	field7	float64
	field8	string
	field9	string
	field10	string
	field11	string
}

func makeTest() model.Record {
	return new(Test)
}

func NewTest() *Test {
	return new(Test)
}

func (this Test) TableName() string {
	return "Test"
}

func (this Test) KeyName() string {
	return "field1"
}

func (this Test) Key() int8 {
	return this.field1
}

func (this Test) GetField1() int8 {
	return this.field1
}

func (this Test) GetField2() int16 {
	return this.field2
}

func (this Test) GetField3() int32 {
	return this.field3
}

func (this Test) GetField4() int64 {
	return this.field4
}

func (this Test) GetField5() int {
	return this.field5
}

func (this Test) GetField6() float32 {
	return this.field6
}

func (this Test) GetField7() float64 {
	return this.field7
}

func (this Test) GetField8() string {
	return this.field8
}

func (this Test) GetField9() string {
	return this.field9
}

func (this Test) GetField10() string {
	return this.field10
}

func (this Test) GetField11() string {
	return this.field11
}

func (this *Test) SetField1(field1_ int8) {
	this.field1 = field1_
}

func (this *Test) SetField2(field2_ int16) {
	this.field2 = field2_
}

func (this *Test) SetField3(field3_ int32) {
	this.field3 = field3_
}

func (this *Test) SetField4(field4_ int64) {
	this.field4 = field4_
}

func (this *Test) SetField5(field5_ int) {
	this.field5 = field5_
}

func (this *Test) SetField6(field6_ float32) {
	this.field6 = field6_
}

func (this *Test) SetField7(field7_ float64) {
	this.field7 = field7_
}

func (this *Test) SetField8(field8_ string) {
	this.field8 = field8_
}

func (this *Test) SetField9(field9_ string) {
	this.field9 = field9_
}

func (this *Test) SetField10(field10_ string) {
	this.field10 = field10_
}

func (this *Test) SetField11(field11_ string) {
	this.field11 = field11_
}

func (this Test) InsertSql() string {
	return fmt.Sprintf("insert into Test values(%d,%d,%d,%d,%d,%d,%d,\"%s\",\"%s\",\"%s\",\"%s\")", this.field1,this.field2,this.field3,this.field4,this.field5,this.field6,this.field7,this.field8,this.field9,this.field10,this.field11)
}

func (this Test) UpdateSql() string {
	return fmt.Sprintf("update Test set field1=%d,field2=%d,field3=%d,field4=%d,field5=%d,field6=%d,field7=%d,field8=\"%s\",field9=\"%s\",field10=\"%s\",field11=\"%s\" where field1=%d",this.field1,this.field2,this.field3,this.field4,this.field5,this.field6,this.field7,this.field8,this.field9,this.field10,this.field11,this.Key())
}

func (this Test) DeleteSql() string {
	return fmt.Sprintf("delete from Test where field1=%d",this.Key())
}

func (this *Test) Scan(rows *sql.Rows) error {
	if err := rows.Scan(&(this.field1),&(this.field2),&(this.field3),&(this.field4),&(this.field5),&(this.field6),&(this.field7),&(this.field8),&(this.field9),&(this.field10),&(this.field11)); err != nil {
		return err
	}
	return nil
}

