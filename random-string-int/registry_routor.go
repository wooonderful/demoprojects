package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 创建一个默认的路由引擎
var router = gin.Default()

func Register() {
	// 定义一个GET请求的路由处理函数
	router.GET("/", randomInt)

}

type RandomResp struct {
	Val  string `json:"val"`
	Type int    `json:"type"`
}

func randomInt(c *gin.Context) {
	rand.NewSource(time.Now().UnixNano())
	resp := &RandomResp{}

	typ := rand.Intn(3)
	resp.Type = typ + 1
	if typ == 2 {
		randomFloat := rand.Float32()
		resp.Val = fmt.Sprintf("%f", randomFloat)
	} else if typ == 1 {
		randomNum := 0 + rand.Int63n(100)
		resp.Val = strconv.FormatInt(randomNum, 10)
	} else {
		randomStr := generateRandomString(10)
		resp.Val = randomStr
	}

	c.JSON(200, resp)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result []byte
	for i := 0; i < length; i++ {
		index := rand.Intn(len(charset))
		result = append(result, charset[index])
	}
	return string(result)
}
