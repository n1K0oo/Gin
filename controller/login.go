package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

////////登录验证
func Verify(context *gin.Context) {
	name := context.PostForm("username")
	pass := context.PostForm("password")
	//dst := "root:2000109@tcp(127.0.0.1:3306)/gin"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	var user Username_table
	db.Where("username = ?", name).Find(&user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)) //验证（对比）
	if err != nil {
		context.HTML(http.StatusOK, "login.html", gin.H{
			"fail": 1,
		})
	} else {
		session := sessions.Default(context)
		session.Set("nowlogin", user.Name)
		session.Save()
		if user.Authority == 1 {
			var users []Username_table
			db.Where("authority = ?", 2).Find(&users)
			context.HTML(http.StatusOK, "houtai.html", gin.H{
				"admin": user.Name,
				"tag":   2,
				"users": users,
			})

		} else if user.Authority == 2 {
			var estates []Estatemess_table
			fmt.Println(user.Area, user.Street)
			db.Where("area = ? AND street = ?", user.Area, user.Street).Find(&estates)
			context.HTML(http.StatusOK, "houtai.html", gin.H{
				"estates": estates,
				"admin":   user.Name,
				"tag":     4,
			})
		} else {
			context.HTML(http.StatusOK, "homepage.html", gin.H{
				"admin": user.Name,
			})
		}
	}
	//if pass == user.Password {
	//	session := sessions.Default(context)
	//	session.Set("nowlogin", user.Name)
	//	session.Save()
	//	if user.Authority == 1 {
	//		var users []Username_table
	//		db.Where("authority = ?", 2).Find(&users)
	//		context.HTML(http.StatusOK, "houtai.html", gin.H{
	//			"admin": user.Name,
	//			"tag":   2,
	//			"users": users,
	//		})
	//
	//	} else if user.Authority == 2 {
	//		var estates []Estatemess_table
	//		fmt.Println(user.Area, user.Street)
	//		db.Where("area = ? AND street = ?", user.Area, user.Street).Find(&estates)
	//		context.HTML(http.StatusOK, "houtai.html", gin.H{
	//			"estates": estates,
	//			"admin":   user.Name,
	//			"tag":     4,
	//		})
	//	} else {
	//		context.HTML(http.StatusOK, "homepage.html", gin.H{
	//			"admin": user.Name,
	//		})
	//	}
	//} else {
	//	context.HTML(http.StatusOK, "login.html", gin.H{
	//		"fail": 1,
	//	})
	//}
}

///////注册 sql  insert
func Register(context *gin.Context) {
	var user Username_table
	user.Username = context.PostForm("username")
	pass := context.PostForm("password")
	user.Authority = 3
	user.Name = context.PostForm("name")
	user.Tel = context.PostForm("tel")
	user.ID = context.PostForm("ID")
	//dst := "root:2000109@tcp(127.0.0.1:3306)/gin"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	password, err2 := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err2 != nil {
		fmt.Println(err2)
	}
	user.Password = string(password)
	db.Create(&user)
	context.HTML(http.StatusOK, "login.html", gin.H{})
}

//////////退出当前登录
func Exitlogin(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("nowlogin")
	session.Save()
	v := session.Get("nowlogin")
	c.HTML(http.StatusOK, "homepage.html", gin.H{
		"admin": v,
	})
}
