package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
)

type GetDataController struct{
	beego.Controller
}

type JSONStruct struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *GetDataController) Get() {
	//data := [...]string{"a", "b", "c"}
	arg := c.GetString("param")
	fmt.Println(arg)
	//ret := &JSONStruct{0, "hello", 1}
	//c.Data["json"] = ret
	ret := map[string]interface{}{"code": 200, "msg": "login success ", "data": 1}
	c.Data["json"] = ret
	c.ServeJSON()
}

type User struct {
	Username string
	Password string
}
func (c *GetDataController) PostFunc() {
	var user User
	data := c.Ctx.Input.RequestBody
	//json数据封装到user对象中
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	fmt.Println(user, user.Username)
	GetUser(user)
	ret := &JSONStruct{0, "hello", "aaa"}
	c.Data["json"] = ret
	c.ServeJSON()
}

func GetUser(user User)  {
	println(user.Username, user.Password)
}
