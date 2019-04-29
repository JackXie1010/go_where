# go_where
爬取去哪网的旅游数据，持久化到MySQL并使用接口请求数据展示到前端页面   

访问：启动main.go文件后，浏览器输入http://localhost:8080/static/index.html 访问页面   
   
页面效果:   
![Image text](https://github.com/JackXie1010/go_where/blob/master/static/img/a.PNG)  
![Image text](https://github.com/JackXie1010/go_where/blob/master/static/img/b.PNG)  
  
前端：   
Vue + ElementUi中的分页组件，模态框 + fly进行接口HTTP请求  
  
后端：  
GO语言中的beego框架，数据库使用的是MySQL   
