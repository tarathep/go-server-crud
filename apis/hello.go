package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tarathep/go-server-crud/db"
	"github.com/tarathep/go-server-crud/model"
)

type HelloHandler struct {
	DB db.HelloRepository
}

func (h *HelloHandler) GetHello(c *gin.Context) {

	hellos, err := h.DB.AllHello()
	if err != nil {
		return
	}
	c.JSON(200, hellos)
}

func (h *HelloHandler) PostHello(c *gin.Context) {

	hello := model.Hello{}

	if err := c.ShouldBindJSON(&hello); err != nil {
		return
	}

	result, err := h.DB.InsertHello(hello)

	fmt.Println("Inserted a single document: ", result, err)

	c.String(200, "success")
}
