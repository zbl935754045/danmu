package controllers

import (
	"eloizhang/danmu/kafkautils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type message struct {
	RoomId string `json:"room_id"`
	UserId string `json:"user_id"`
	Time   string `json:"time"`
	Msg    string `json:"msg"`
}

func WriterDanmu(c *gin.Context) {
	roomId := c.DefaultQuery("roomid", "1")
	userid := c.DefaultQuery("userid", "2")
	time := c.DefaultQuery("time", "3")
	msg := c.DefaultQuery("msg", "4")

	messa := message{
		RoomId: roomId,
		UserId: userid,
		Time:   time,
		Msg:    msg,
	}

	// 将结构体解析为字符串
	str, err := json.Marshal(messa)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	fmt.Printf("str:%v", string(str))

	// 初始化生产
	kafka.InitProducer("localhost:9092")
	// 关闭
	defer kafka.Close()

	// 发送测试消息
	kafka.Send("Test", string(str))
	fmt.Println("发送成功")
}
