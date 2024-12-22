package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"test/dao"
	"test/service"
)

import (
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic("Config file not found; please make sure config.json is in the current directory.")
		}
	}

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
