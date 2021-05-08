package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	server.GET("/test",test)

	server.Run(":7000")
}

func test(c *gin.Context){
	fmt.Println("ABC")
	c.JSON(200,gin.H{
		"message":"請求成功",
		"code":503,
	})
}
