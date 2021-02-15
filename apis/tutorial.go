package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/tarathep/go-server-crud/db"
	"github.com/tarathep/go-server-crud/model"
)

type TutorialHandler struct {
	DB db.TutorialRepository
}

func (h *TutorialHandler) CreateTutorial(c *gin.Context) {

	tutorial := model.Tutorial{}
	if err := c.ShouldBindJSON(&tutorial); err != nil {
		c.String(500, err.Error())
		return
	}

	if err := h.DB.Create(tutorial); err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "Inserted a single document Success")
}

func (h *TutorialHandler) AllTutorial(c *gin.Context) {
	title := c.Query("title")
	tutorials, err := h.DB.FindAll(title)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, tutorials)
}
