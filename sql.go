package main

type Sql struct {
	Name   string
	Tables []Table
}

type Table struct {
	Name   string
	Fields []Field
}

type Field struct {
	Type string
	Name string
}

func NewSql(name string) *Sql {
	ans := new(Sql)
	ans.Name = name
	ans.Tables = make([]Table, 0)
	return ans
}

func NewTable(name string) *Table {
	ans := new(Table)
	ans.Name = name
	ans.Fields = make([]Field, 0)
	return ans
}

func (this *Sql) AddTable(tb Table) {
	this.Tables = append(this.Tables, tb)
}

func (this *Table) AddField(fd Field) {
	this.Fields = append(this.Fields, fd)
}

func (this *Sql) CopyTo(type_conv func(string) string) *Sql {
	res := NewSql(UpperFirst(this.Name))
	res.Tables = make([]Table, len(this.Tables))

	for i, table := range this.Tables {
		tbi := &(res.Tables[i])
		tbi.Name = UpperFirst(table.Name)
		tbi.Fields = make([]Field, len(table.Fields))
		for j := 0; j < len(tbi.Fields); j++ {
			tbi.Fields[j].Name = LowerFirst(table.Fields[j].Name)
			tbi.Fields[j].Type = type_conv(table.Fields[j].Type)
		}
	}
	return res
}
