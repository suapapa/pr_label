package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.POST("/order", v1OrderHandler)
	}
	r.Run()
}

func v1OrderHandler(c *gin.Context) {
	var ord Order
	c.BindJSON(&ord)

	printAddrFrom(ord.From)
	printAddrTo(ord.To)
}
