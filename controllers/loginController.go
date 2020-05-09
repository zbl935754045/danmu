package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	userName := c.Request.FormValue("username")
	passWd := c.Request.FormValue("password")
	fmt.Println(userName + passWd)
	if userName == "" || passWd == "" {
		//用户名或者密码为空
		c.JSON(http.StatusOK, gin.H{
			"code": 1002,
			"msg":  "用户名或者密码为空",
		})
	}

	err := getUser(userName, passWd)
	if err != nil {
		//用户名或者密码错误
		c.JSON(http.StatusOK, gin.H{
			"code": 1003,
			"msg":  "用户名或者密码错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登陆成功",
	})
	return

	//查询是否存在该user：
	//不存在则返回未注册提示
	//存在则验证账号密码
	//验证通过则登陆成功
	//验证不通过则登录失败

}

func getUser(userName string, password string) error {
	var id int
	err := db.QueryRow("SELECT id FROM users where username = ? and password = ?", userName, password).Scan(&id)
	if err != nil {
		fmt.Println("不存在该记录")
		return fmt.Errorf("不存在该记录")
	}
	return nil
}
