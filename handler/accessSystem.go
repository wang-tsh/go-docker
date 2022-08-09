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
	var AccessSystems []accessSystem.AccessSystem
	id := c.Query("id")
	var err error
	if "" != id {
		err = db.DB.Find(&AccessSystems, id).Error
	} else {
		err = db.DB.Find(&AccessSystems).Error

	}
	if err != nil {
		log.Errorf("db has error %v", err.Error())
	}

	resultVo := vo.BaseResultVo{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    AccessSystems,
	}
	c.JSON(http.StatusOK, resultVo)
}
