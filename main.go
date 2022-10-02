package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	flagPort int

	q *OrderQ
)

func main() {
	flag.IntVar(&flagPort, "p", 8080, "port to serve")
	flag.Parse()

	q = NewOrderQ(DefaultOrderQLen)
	go printLoop()

	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.POST("/order", v1OrderHandler)
	}
	if err := r.Run(fmt.Sprintf(":%d", flagPort)); err != nil {
		log.Fatal("fail on server run")
	}
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
	log.Println("printLoop started")
	tk := time.NewTicker(time.Second)
	defer tk.Stop()
	for range tk.C {
		if !q.IsEmpty() {
			// log.Println("somthin in q")
			ord, err := q.Pop()
			if err != nil {
				log.Printf("ERR: %v", errors.Wrap(err, "print will be stopped"))
				return
			}
			printItems(ord.ID, ord.Items)
			printAddrFrom(ord.ID, ord.From)
			printAddrTo(ord.ID, ord.To)
			time.Sleep(time.Second)
		}
	}
}
