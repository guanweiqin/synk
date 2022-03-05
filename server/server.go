package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/guanweiqin/synk/server/config"
	"github.com/guanweiqin/synk/server/controllers"
	"github.com/guanweiqin/synk/server/initializers"
	"github.com/guanweiqin/synk/server/ws"
)

//go:embed frontend/dist/*
var FS embed.FS

func Run() {
	gin.SetMode(gin.DebugMode)
	gin.DisableConsoleColor()
	router := gin.Default()
	initializers.InitCors(router)
	hub := ws.NewHub()
	go hub.Run()
	router.GET("/ws", func(c *gin.Context) {
		ws.HttpController(c, hub)
	})
	router.GET("/uploads/:path", controllers.UploadsController)
	router.GET("/api/v1/addresses", controllers.AddressesController)
	router.GET("/api/v1/qrcodes", controllers.QrcodesController)
	router.POST("/api/v1/files", controllers.FilesController)
	router.POST("/api/v1/texts", controllers.TextsController)
	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	router.StaticFS("/static", http.FS(staticFiles))
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/static/") {
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			stat, err := reader.Stat()
			if err != nil {
				log.Fatal(err)
			}
			c.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
		} else {
			c.Status(http.StatusNotFound)
		}
	})
	router.Run(":" + config.GetPort())
}
