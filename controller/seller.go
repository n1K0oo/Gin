package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

//////////////卖房租房跳转
func Jumptosell(c *gin.Context) {
	c.HTML(http.StatusOK, "sell.html", gin.H{
		"flag2": 2,
	})
}
func Jumptorent(c *gin.Context) {
	c.HTML(http.StatusOK, "sell.html", gin.H{
		"flag2": 1,
	})
}

//////////////卖家发布
func Releasenew(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")
	dst := "root:cyz2000109@tcp(127.0.0.1:3306)/big_work"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var et Estatemess_table
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
	et.Username = fmt.Sprint(v)
	et.Flag = 0
	et.Flag2, _ = strconv.Atoi(c.PostForm("flag2"))
	var tmp []Estatemess_table
	db.Order("estateno desc").Find(&tmp)
	if len(tmp) == 0 {
		et.Estateno = 0
	} else {
		et.Estateno = tmp[0].Estateno + 1
	}
	i1, err1 := c.FormFile("file1")
	if err1 == nil {
		et.Image1 = i1.Filename
		fmt.Println(et.Image1)
	}
	db.Create(&et)
	var insert []Indoor_table
	i2, err2 := c.FormFile("file2")
	if err2 == nil {
		insert = append(insert, Indoor_table{et.Estateno, i2.Filename})
	}
	i3, err3 := c.FormFile("file3")
	if err3 == nil {
		insert = append(insert, Indoor_table{et.Estateno, i3.Filename})
	}
	i4, err4 := c.FormFile("file4")
	if err4 == nil {
		insert = append(insert, Indoor_table{et.Estateno, i4.Filename})
	}
	db.Create(&insert)
	c.HTML(http.StatusOK, "homepage.html", gin.H{
		"admin": v,
	})
}
