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

//////////////////跳转查看所属房产
func Jump4mid(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")
	//dst := "root:2000109@tcp(127.0.0.1:3306)/gin"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var user Username_table
	db.Where("username = ?", v).Find(&user)
	var estates []Estatemess_table
	db.Where("area = ? AND street = ?", user.Area, user.Street).Find(&estates)
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin":   v,
		"tag":     4,
		"estates": estates,
	})
}

////////////////////上架待审核房产
func Jump5mid(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")
	//dst := "root:2000109@tcp(127.0.0.1:3306)/gin"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var user Username_table
	db.Where("username = ?", v).Find(&user)
	var estates []Estatemess_table
	db.Where("area = ? AND street = ? AND flag = 0", user.Area, user.Street).Find(&estates)
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin":   v,
		"tag":     5,
		"estates": estates,
	})
}

////////////////////交易待审核房产
func Jump6mid(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")
	//dst := "root:2000109@tcp(127.0.0.1:3306)/gin"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var user Username_table
	db.Where("username = ?", v).Find(&user)
	var estates []Estatemess_table
	db.Where("area = ? AND street = ? AND flag = 2", user.Area, user.Street).Find(&estates)
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin":   v,
		"tag":     6,
		"estates": estates,
	})
}

//////////////中介审批
func Approvalmid(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")
	//dst := "root:2000109@tcp(127.0.0.1:3306)/gin"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	esnum := c.Query("esnum")
	tag := c.Query("tag")
	t, _ := strconv.Atoi(tag)
	var user Username_table
	db.Where("username = ?", v).Find(&user)
	var estates []Estatemess_table
	var ttttmp Estatemess_table
	db.Where("estateno = ?", esnum).Find(&ttttmp)
	db.Model(&estates).Where("estateno = ?", esnum).Update("flag", ttttmp.Flag+1)
	db.Where("area = ? AND street = ? AND flag = ?", user.Area, user.Street, ttttmp.Flag).Find(&estates)
	c.HTML(http.StatusOK, "houtai.html", gin.H{
		"admin":   v,
		"tag":     t,
		"estates": estates,
	})
}
