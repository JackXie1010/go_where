package controllers

import (
	"github.com/astaxie/beego"
	"go_where/models"
	"strconv"
)

type GetListController struct{
	beego.Controller
}

func (c *GetListController) Get() {
	curr_page := c.GetString("curr_page", "1")
	page_num := c.GetString("page_num", "10")
	i, err := strconv.Atoi(curr_page)
	n, _ := strconv.Atoi(page_num)
	numCount, pageCount := models.PageCount(n)
	list, err := models.QueryAll(i, n)
	var code int
	if err != nil{
		code = 204
	}else{
		code = 200
	}
	res := map[string]interface{}{"pageCount": pageCount, "numCount": numCount, "list": list}
	ret := map[string]interface{}{"code": code, "msg": "success ", "data": res}
	c.Data["json"] = ret
	c.ServeJSON()
}

type GetCommentController struct {
	beego.Controller
}

func (c *GetCommentController) Get() {
	var lst [3] interface{}
	lst[0] = map[string] string{"head":"http://img1.qunarzz.com/sight/p0/1707/f4/f44d291b76de7b14a3.img.jpg_710x360_0df65faf.jpg", "nick_name":"aaa", "comment":"好玩"}
	lst[1] = map[string] string{"head":"http://img1.qunarzz.com/sight/p0/1707/f4/f44d291b76de7b14a3.img.jpg_710x360_0df65faf.jpg", "nick_name":"aaa", "comment":"好玩"}
	lst[2] = map[string] string{"head":"http://img1.qunarzz.com/sight/p0/1707/f4/f44d291b76de7b14a3.img.jpg_710x360_0df65faf.jpg", "nick_name":"aaa", "comment":"好玩"}
	ret := map[string]interface{}{"code": 200, "msg": "success ", "data": lst}
	c.Data["json"] = ret
	c.ServeJSON()
}

