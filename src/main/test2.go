package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("aaaaa")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func testSy(str string, i int){
	fmt.Println(str, i)
}
