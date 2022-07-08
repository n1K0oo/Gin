package controller

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"sort"
	"strconv"
)

var dst string = "root:2000109@tcp(127.0.0.1:3306)/gin"

type Sortnum struct {
	point int
	esno  int
}

type Sortnums []Sortnum

func (s Sortnums) Len() int {
	return len(s)
}
func (s Sortnums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Sortnums) Less(i, j int) bool {
	return s[i].point > s[j].point
}

/////////////搜索
func Search(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	estatename := c.PostForm("search")
	flag2 := c.PostForm("flag2")
	status := c.DefaultPostForm("status", "2")
	var estates []Estatemess_table
	var estatess []Estatemess_table
	db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ?", flag2, "%"+estatename+"%").Find(&estates)
	if status == "2" {
		var mp map[int]int
		mp = make(map[int]int)

		area := c.PostForm("area")
		street := c.PostForm("street")
		type2 := c.PostForm("type2")
		price1 := c.PostForm("price1")
		price2 := c.PostForm("price2")
		mianji1 := c.PostForm("mianji1")
		mianji2 := c.PostForm("mianji2")
		toward := c.PostForm("toward")
		layer := c.PostForm("layer")

		/////1
		var estates1 []Estatemess_table
		if area == "不限" {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ?", flag2, "%"+estatename+"%").Find(&estates1)
		} else {
			if street == "不限" {
				db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ? AND area = ?", flag2, "%"+estatename+"%", area).Find(&estates1)
			} else {
				db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ? AND area = ? AND street = ?", flag2, "%"+estatename+"%", area, street).Find(&estates1)
			}
		}
		for i := 0; i < len(estates1); i++ {
			mp[estates1[i].Estateno] += 1
		}
		////2
		var estates2 []Estatemess_table
		if type2 == "不限" {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ?", flag2, "%"+estatename+"%").Find(&estates2)
		} else {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ? AND street = ?", flag2, "%"+estatename+"%", type2).Find(&estates2)
		}
		for i := 0; i < len(estates2); i++ {
			mp[estates2[i].Estateno] += 1
		}
		////3
		var estates3 []Estatemess_table
		price11, _ := strconv.Atoi(price1)
		if price1 == "" {
			price11 = 0
		}
		price22, _ := strconv.Atoi(price2)
		if price2 == "" {
			price22 = 1000000
		}
		if street == "不限" {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ?", flag2, "%"+estatename+"%").Find(&estates3)
		} else {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ? AND price >= ? AND price <= ?", flag2, "%"+estatename+"%", price11, price22).Find(&estates3)
		}
		for i := 0; i < len(estates3); i++ {
			mp[estates3[i].Estateno] += 1
		}
		////4
		var estates4 []Estatemess_table
		mianji11, _ := strconv.Atoi(mianji1)
		mianji22, _ := strconv.Atoi(mianji2)
		if mianji1 == "" {
			mianji11 = 0
		}
		if mianji2 == "" {
			mianji22 = 1000000
		}
		if street == "不限" {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ?", flag2, "%"+estatename+"%").Find(&estates4)
		} else {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ? AND mianji >= ? AND mianji <= ?", flag2, "%"+estatename+"%", mianji11, mianji22).Find(&estates4)
		}
		for i := 0; i < len(estates4); i++ {
			mp[estates4[i].Estateno] += 1
		}
		////5
		var estates5 []Estatemess_table
		if toward == "不限" {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ?", flag2, "%"+estatename+"%").Find(&estates5)
		} else {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ? AND toward = ?", flag2, "%"+estatename+"%", toward).Find(&estates5)
		}
		for i := 0; i < len(estates5); i++ {
			mp[estates5[i].Estateno] += 1
		}
		////6
		var estates6 []Estatemess_table
		if layer == "不限" {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ?", flag2, "%"+estatename+"%").Find(&estates6)
		} else {
			db.Where("flag = 1 AND flag2 = ? AND estatename LIKE ? AND toward = ?", flag2, "%"+estatename+"%", layer).Find(&estates6)
		}
		for i := 0; i < len(estates6); i++ {
			mp[estates6[i].Estateno] += 1
		}
		estates = []Estatemess_table{}
		for k, v := range mp {
			var tmp Estatemess_table
			if v == 6 {
				db.Where("estateno = ?", k).Find(&tmp)
				if tmp.Flag2 == 2 {
					//fmt.Println(fmt.Sprintf("%.2f", tmp.Mianji*tmp.Price/10000))
					tmp.Allprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", tmp.Mianji*tmp.Price/10000), 64)

				}
				estates = append(estates, tmp)
			}
		}

		conn2, err4 := redis.Dial("tcp", "127.0.0.1:6379")
		if err4 != nil {
			fmt.Println(err4)
		}

		//fmt.Println(estates)
		Data := []Sortnum{}
		for i := 0; i < len(estates); i++ {

			r, err5 := redis.Int(conn2.Do("Get", estates[i].Estateno)) //类型断言接口
			if err5 != nil {
				_, err6 := conn2.Do("Set", estates[i].Estateno, 0) //redis写入数据
				if err6 != nil {
					fmt.Println(err6)
				}
				fmt.Println(err5)
				Data = append(Data, Sortnum{point: 0, esno: estates[i].Estateno})
			} else {
				Data = append(Data, Sortnum{point: r, esno: estates[i].Estateno})
			}

		}
		sort.Sort(Sortnums(Data))
		fmt.Println(Data)
		estatess = []Estatemess_table{}
		for i := 0; i < len(Data); i++ {
			for j := 0; j < len(estates); j++ {
				if estates[j].Estateno == Data[i].esno {
					estatess = append(estatess, estates[j])
					break
				}
			}
		}
		fmt.Println(estatess)
	}
	c.HTML(http.StatusOK, "main.html", gin.H{
		"admin":   v,
		"estates": estates,
		"search":  estatename,
		"flag2":   flag2,
	})
}

/////////点击条目进入房屋详情
func Esdetails(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")
	if v == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	esno := c.Query("esno")
	num, _ := strconv.Atoi(esno)
	/////////redis更新点击次数
	conn2, err4 := redis.Dial("tcp", "127.0.0.1:6379")
	if err4 != nil {
		fmt.Println(err4)
	}
	r, err5 := redis.Int(conn2.Do("Get", esno)) //类型断言接口
	if err5 != nil {
		_, err6 := conn2.Do("Set", esno, 1) //redis写入数据
		if err6 != nil {
			fmt.Println(err6)
		}
		fmt.Println(err5)
	} else {
		_, err7 := conn2.Do("Set", esno, r+1)
		if err7 != nil {
			fmt.Println(err7)
		}
	}
	fmt.Println(r)

	var tmp Estatemess_table
	db.Where("estateno = ?", num).Find(&tmp)
	tmp.Allprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", tmp.Mianji*tmp.Price/10000), 64)
	var igs []Indoor_table
	db.Where("estateno = ?", num).Find(&igs)
	c.HTML(http.StatusOK, "esdetails.html", gin.H{
		"admin":  v,
		"estate": tmp,
		"igs":    igs,
	})
}

//////////////买房租房
func Buyorrent(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("nowlogin")

	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	esno := c.Query("esno")
	num, _ := strconv.Atoi(esno)
	var tmp Estatemess_table
	db.Where("estateno = ?", num).Find(&tmp)
	db.Model(&tmp).Where("estateno = ?", num).Update("flag", tmp.Flag+1)
	db.Model(&tmp).Where("estateno = ?", num).Update("buyer", fmt.Sprint(v))
	c.HTML(http.StatusOK, "homepage.html", gin.H{
		"admin": v,
	})
}
