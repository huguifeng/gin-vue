package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/regster", func(c *gin.Context) {
		tel := c.PostForm("tel")
		if len(tel) != 11 {
			c.JSON(422, gin.H{"code": 422, "msg": "手机号错误"})
			return
		}
		pasword := c.PostForm("pasword")
		if len(pasword) < 6 {
			c.JSON(422, gin.H{"code": 422, "msg": "密码格式错误"})
			return
		}
		name := c.PostForm("name")
		if len(name) == 0 {
			name = RandStr(2)
		}
		fmt.Println(tel, pasword, name)
	})
	r.Run()
}
func RandStr(n int) string {
	var leter = []byte("abcGEFG")
	result := make([]byte, n)
	// fmt.Println(string(result))
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = leter[rand.Intn(len(leter))]
	}
	return string(result)
}
