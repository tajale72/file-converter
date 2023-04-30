package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/tajale72/file-converter/internal/router"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.GET("/systeminfo", SystemInfo)
	r.POST("/upload", router.FileUpload)
	r.Run()

}

type SystemConfig struct {
	Host string   `json:"host" default:"localhost"`
	Port int      `json:"port" default:"9000"`
	Env  []string `json:"env"`
}

func SystemInfo(c *gin.Context) {
	var system SystemConfig
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}
	env := os.Environ()
	system.Env = env
	system.Host = hostname

	c.JSON(http.StatusAccepted, system)
}
