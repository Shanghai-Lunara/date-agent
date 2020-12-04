package date_agent

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"k8s.io/klog/v2"
	"net/http"
	"os"
)

func InitHttp(addr string, hub *Hub) *http.Server {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	//router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: writer}), gin.RecoveryWithWriter(writer))
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/*")
	router.StaticFS("/statics", http.Dir("./statics"))
	router.GET("/hello", func(c *gin.Context) {
		// todo handle request and return data by hub
		c.JSON(http.StatusOK, hub.nodes)
	})

	router.GET("/getJobs", func(c *gin.Context) {
		c.JSONP(http.StatusOK, hub.nodes)
	})

	router.POST("/changeTime", func(c *gin.Context) {
		//name := c.PostForm("hostname")
		command := c.PostForm("command")
		go func() {
			//<-time.After(time.Second * 10)
			klog.Info("new task")
			hub.NewTask([]string{command})
		}()
		c.JSONP(http.StatusOK, hub.nodes)
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "test",
			"value": hub.nodes,
		})
		fmt.Printf("%+v\n", &hub.nodes)
	})

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				klog.Info("Server closed under request")
			} else {
				klog.V(2).Info("Server closed unexpected err:", err)
			}
		}
	}()
	return server
}
