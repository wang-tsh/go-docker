package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-docker/db"
	accessSystem "go-docker/db/entity/system"
	"go-docker/model/vo"
	"net/http"
)

func GetSystems(c *gin.Context) {

	log.Infoln("to get system")
	var AccessSystem accessSystem.AccessSystem
	id := c.Query("id")
	db.DB.Find(&AccessSystem, id)

	resultVo := vo.BaseResultVo{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    AccessSystem,
	}
	c.JSON(http.StatusOK, resultVo)
}
