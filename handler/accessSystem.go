package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-docker/db"
	accessSystem "go-docker/db/entity/system"
	"go-docker/model/vo"
	"net/http"
	"strconv"
)

func GetSystems(c *gin.Context) {

	log.Infoln("to get system")
	var AccessSystems []accessSystem.AccessSystem
	id := c.Query("id")
	pageSize, _ := strconv.ParseInt(c.Query("pageSize"), 10, 64)
	pageNum, _ := strconv.ParseInt(c.Query("pageNum"), 10, 64)
	page := vo.NewPage(pageSize, pageNum)
	var err error
	if "" != id {
		err = db.DB.First(&AccessSystems, id).Error
	} else {
		err = db.DB.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&AccessSystems).Error

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
