package mydb

import (
	"database/sql"
	"fmt"
	"gensql/model"
)

type Shop struct {
	itemId	int
	name	string
	price	float32
}

func makeShop() model.Record {
	return new(Shop)
}

func NewShop() *Shop {
	return new(Shop)
}

func (this Shop) TableName() string {
	return "Shop"
}

func (this Shop) KeyName() string {
	return "itemId"
}

func (this Shop) Key() int {
	return this.itemId
}

func (this Shop) GetItemId() int {
	return this.itemId
}

func (this Shop) GetName() string {
	return this.name
}

func (this Shop) GetPrice() float32 {
	return this.price
}

func (this *Shop) SetItemId(itemId_ int) {
	this.itemId = itemId_
}

func (this *Shop) SetName(name_ string) {
	this.name = name_
}

func (this *Shop) SetPrice(price_ float32) {
	this.price = price_
}

func (this Shop) InsertSql() string {
	return fmt.Sprintf("insert into Shop values(%d,\"%s\",%d)", this.itemId,this.name,this.price)
}

func (this Shop) UpdateSql() string {
	return fmt.Sprintf("update Shop set itemId=%d,name=\"%s\",price=%d where itemId=%d",this.itemId,this.name,this.price,this.Key())
}

func (this Shop) DeleteSql() string {
	return fmt.Sprintf("delete from Shop where itemId=%d",this.Key())
}

func (this *Shop) Scan(rows *sql.Rows) error {
	if err := rows.Scan(&(this.itemId),&(this.name),&(this.price)); err != nil {
		return err
	}
	return nil
}

