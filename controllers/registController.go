package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

var db *sql.DB

type Response struct {
	code int
	msg  string
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:199794@tcp(localhost:3306)/danmu?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	// 设置最大连接数
	db.SetMaxOpenConns(300)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(100)
	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(time.Second * 5)
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

func Regist(c *gin.Context) {
	userName := c.Request.FormValue("username")
	passWd := c.Request.FormValue("password")

	//查询是否存在该user：
	//存在则返回已存在
	//不存在则注册
	_, isExist := isExistUser(userName)
	if isExist {
		//存在这条记录
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "注册失败：用户已存在",
		})
		return
	} else {
		//不存在这条记录，插入db
		err := insertUser(userName, passWd)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "注册成功",
			})
			return
		}
	}
	return
}

func insertUser(userName string, passWd string) error {
	r, err := db.Exec("insert into users(username, password)values(?, ?)", userName, passWd)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return fmt.Errorf("sql执行失败")
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return fmt.Errorf("获取id失败")
	}
	fmt.Println("insert succ:", id)
	return nil
}

func isExistUser(userName string) (error, bool) {
	var id int
	err := db.QueryRow("SELECT id FROM users where username = ?", userName).Scan(&id)
	if err != nil {
		fmt.Println("不存在该记录")
		return nil, false
	}
	return nil, true
}
