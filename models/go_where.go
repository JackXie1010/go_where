package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
)


type GoWhere struct {
	Id int `json:"id"`
	Tag_id int `json:"tag_id"`
	Tag string `json:"tag"`
	Name string `json:"name"`
	Level string `json:"level"`
	Address string `json:"address"`
	Price float64 `json:"price"`
	Star string `json:"star"`
	Num int `json:"num"`
	Relate string `json:"relate"`
	Introduce string `json:"introduce"`
}

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:xzjroot@tcp(127.0.0.1:3306)/beego?charset=utf8", 30)
	orm.RegisterModel(new(GoWhere))
}

func PageCount(page_num int)(int64, int64){
	o := orm.NewOrm()
	qs := o.QueryTable("go_where")
	numCount, _ := qs.Count()
	size := int64(page_num)
	quotient := numCount / size
	remainder := numCount % size
	if remainder > 0{
		quotient = quotient + 1
	}
	pageCount := quotient
	return numCount, pageCount
}

func QueryAll(curr_page, page_num int)(list[] GoWhere, err error)  {
	o := orm.NewOrm()
	start := (curr_page - 1)*page_num
	sql := fmt.Sprintf("select tag, name,level,address,price,star,num,relate,introduce from go_where limit %d, %d",start, page_num)
	o.Raw(sql).QueryRows(&list)
	return list, err
}
