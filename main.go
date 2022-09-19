package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

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
	if err := c.BindJSON(&ord); err != nil {
		ret := map[string]any{
			"status": http.StatusBadRequest,
			"msg":    errors.Wrap(err, "print order failed"),
		}
		c.JSON(http.StatusBadRequest, ret)
	}

	printAddrFrom(ord.From)
	printAddrTo(ord.To)

	ret := map[string]any{
		"status": http.StatusOK,
	}
	c.JSON(http.StatusOK, ret)
}
