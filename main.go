package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/controller"
)

func main() {
	r := gin.Default()
	/////////////////加载
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFS("/statics", http.Dir("./templates/statics"))
	/////////////session设置
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	///////////////////直接跳转
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/register", func(context *gin.Context) {
		context.HTML(http.StatusOK, "register.html", gin.H{})
	})
	r.GET("/home", func(context *gin.Context) {
		context.HTML(http.StatusOK, "homepage.html", gin.H{})
	})
	r.GET("/sell", func(context *gin.Context) {
		session := sessions.Default(context)
		v := session.Get("nowlogin")
		if v == nil {
			context.HTML(http.StatusOK, "login.html", gin.H{})
		} else {
			context.HTML(http.StatusOK, "sell.html", gin.H{})
		}
	})
	////////////////////首页  入口
	r.GET("/homepage", func(context *gin.Context) {
		session := sessions.Default(context)
		v := session.Get("nowlogin")
		fmt.Println(v)
		context.HTML(http.StatusOK, "homepage.html", gin.H{
			"admin": v,
		})
	})
	////////买家相关功能
	r.POST("/search", controller.Search)
	r.GET("/esdetails", controller.Esdetails)
	r.GET("/buyorrent", controller.Buyorrent)
	////////卖家相关功能
	r.POST("/releasenew", controller.Releasenew)
	r.GET("/jumptosell", controller.Jumptosell)
	r.GET("/jumptorent", controller.Jumptorent)
	////////////////注册登录相关
	r.POST("/userlogin", controller.Verify)
	r.POST("/userregister", controller.Register)
	r.GET("/exitlogin", controller.Exitlogin)
	////////////////后台(超管)
	r.GET("/jump1admin", controller.Jump1admin)
	r.POST("/j1adminInsert", controller.J1adminInsert)
	r.GET("/jump2admin", controller.Jump2admin)
	r.GET("/delmid", controller.Delmid)
	r.GET("/jump3admin", controller.Jump3admin)
	r.GET("/estatedel", controller.Estatedel)
	r.GET("/estateupdate", controller.Estateupdate)
	r.POST("/estateupdate2", controller.Estateupdate2)
	////////////////后台(中介)
	r.GET("/jump4mid", controller.Jump4mid)
	r.GET("/approvalmid", controller.Approvalmid)
	r.GET("/jump5mid", controller.Jump5mid)
	r.GET("/jump6mid", controller.Jump6mid)
	////////////////端口
	r.Run(":8080")
}
