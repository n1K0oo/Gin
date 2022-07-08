package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

///////////////////跳转新增中介
func Jump1admin(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin": v,
		"tag":   1,
	})
}

//////////////插入中介
func J1adminInsert(c *gin.Context) {
	var es Username_table
	es.Username = c.PostForm("username")
	es.Password = c.PostForm("password")
	es.Name = c.PostForm("name")
	es.Tel = c.PostForm("tel")
	es.Authority = 2
	es.Area = c.PostForm("area")
	es.Street = c.PostForm("street")

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&es)
	session := sessions.Default(c)
	v := session.Get("nowlogin")
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin": v,
		"tag":   1,
	})
}

//////////////跳转查看中介
func Jump2admin(c *gin.Context) {

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var users []Username_table
	session := sessions.Default(c)
	v := session.Get("nowlogin")
	db.Where("authority = ?", 2).Find(&users)
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin": v,
		"tag":   2,
		"users": users,
	})
}

//////////////跳转查看所有房产
func Jump3admin(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var estates []Estatemess_table
	db.Find(&estates)
	//fmt.Println(estates)
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin":   v,
		"tag":     3,
		"estates": estates,
	})
}

////////////////超管删除中介
func Delmid(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	name := c.Query("username")
	var users []Username_table
	db.Where("username = ?", name).Delete(&users)
	db.Where("authority = ?", 2).Find(&users)
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin": v,
		"tag":   2,
		"users": users,
	})
}

//////////////超管更新房屋
func Estateupdate(c *gin.Context) {

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var t Estatemess_table
	esno := c.Query("esno")
	num, _ := strconv.Atoi(esno)
	db.Where("estateno = ?", num).Find(&t)
	c.HTML(http.StatusOK, "update.html", gin.H{
		"estate": t,
	})
}
func Estateupdate2(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	esno := c.PostForm("esno")
	num, _ := strconv.Atoi(esno)
	var tmp Estatemess_table
	db.Where("estateno = ?", num).Find(&tmp)
	var et Estatemess_table
	et.Estateno = num
	et.Estatename = c.PostForm("esname")
	et.Type1 = c.PostForm("type1")
	et.Type2 = c.PostForm("type2")
	et.Btime = c.PostForm("time1")
	et.Etime = c.PostForm("time2")
	et.Estateagent = c.PostForm("estateagent")
	et.Street = c.PostForm("street")
	et.Area = c.PostForm("area")
	et.Address = c.PostForm("address")
	et.Communityname = c.PostForm("communityname")
	et.Roomno = c.PostForm("roomno")
	et.Layer = c.PostForm("layer")
	et.Toward = c.PostForm("toward")
	et.Mianji, _ = strconv.ParseFloat(c.PostForm("mianji"), 64)
	et.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	et.Describes = c.PostForm("describe")
	et.Username = tmp.Username
	et.Flag = tmp.Flag
	et.Flag2 = tmp.Flag2
	var estates []Estatemess_table
	db.Where("estateno = ?", num).Delete(&estates)
	db.Create(&et)
	db.Find(&estates)
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin":   v,
		"tag":     3,
		"estates": estates,
	})
}

//////////////超管删除房产
func Estatedel(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var t Estatemess_table
	esno := c.Query("esno")
	num, _ := strconv.Atoi(esno)
	db.Where("estateno = ?", num).Delete(&t)
	var estates []Estatemess_table
	db.Find(&estates)
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin":   v,
		"tag":     3,
		"estates": estates,
	})
}
