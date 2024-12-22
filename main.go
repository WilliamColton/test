package main

import (
	"github.com/gin-gonic/gin"
	"test/dao"
	"test/service"
)

func main() {
	d := dao.NewDao()
	s := service.NewService(d)

	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/config", s.GetConfig)
		v1.PUT("/config", s.UpdateConfig)
		v1.GET("/alivetime", s.GetAliveTime)
	}
}
