package controllers

import (
	"eloizhang/danmu/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type request struct {
	RoomId  string  `json:"room_id"`
	UserId  string  `json:"user_id"`
	Time    float64 `json:"time"`
	Message `json:"msg"`
}

type Message struct {
	Color   string `json:"color"`
	Content string `json:"content"`
	Speed   string `json:"speed"`
}

func WriterDanmu(c *gin.Context) {
	roomId := c.DefaultQuery("roomid", "房间1")
	userid := c.DefaultQuery("userid", "zbl935754045")
	//time := c.DefaultQuery("time", "3")

	req := request{
		RoomId: roomId,
		UserId: userid,
		Time:   123456,
		Message: Message{
			Color:   "323",
			Content: "dsf",
			Speed:   "trt",
		},
	}

	// 将结构体解析为字符串
	str, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	fmt.Printf("str:%v", string(str))

	// 初始化生产
	utils.InitProducer("localhost:9092")
	// 关闭
	defer utils.Close()

	// 发送测试消息
	utils.Send("Test", string(str))
	fmt.Println("发送成功")
}
