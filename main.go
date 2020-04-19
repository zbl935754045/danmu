package main

import (
	"eloizhang/danmu/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := routers.InitRouter()
	router.Run(":8081")
}
