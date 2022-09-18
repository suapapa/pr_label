package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.POST("/addr", v1AddrHandler)
	}
	r.Run()
}

type Addr struct {
	Line1       string `json:"line1"`
	Line2       string `json:"line2,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	PostNumber  string `json:"post_number,omitempty"`
}

type Envelope struct {
	From *Addr `json:"from"`
	To   *Addr `json:"to"`
}

func v1AddrHandler(c *gin.Context) {
	var env Envelope
	c.BindJSON(&env)

	printAddrFrom(env.From)
	printAddrTo(env.To)
}
