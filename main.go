package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	q *OrderQ
)

func main() {
	q = NewOrderQ(DefaultOrderQLen)
	go printLoop()

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
		return
	}

	if err := q.Push(&ord); err != nil {
		ret := map[string]any{
			"status": http.StatusInternalServerError,
			"msg":    errors.Wrap(err, "print order failed"),
		}
		c.JSON(http.StatusInternalServerError, ret)
		return
	}

	ret := map[string]any{
		"status": http.StatusOK,
	}
	c.JSON(http.StatusOK, ret)
}

func printLoop() {
	for !q.IsEmpty() {
		ord, err := q.Pop()
		if err != nil {
			log.Printf("ERR: %v", errors.Wrap(err, "print will be stopped"))
			return
		}
		printAddrFrom(ord.From)
		printAddrTo(ord.To)
		time.Sleep(time.Second)
	}
}
