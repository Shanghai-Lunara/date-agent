package date_agent

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
)

func InitHttp(addr string, hub *Hub) *http.Server {
	router := gin.New()
	//router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: writer}), gin.RecoveryWithWriter(writer))
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/*")
	router.StaticFS("/statics", http.Dir("./statics"))
	router.GET("/hello", func(c *gin.Context) {
		// todo handle request and return data by hub
		for index, value := range hub.tasks {
			fmt.Println("taskId", hub.taskId)
			fmt.Printf("%+v\n", index)
			fmt.Printf("%+v\n", value)
		}
		fmt.Println("hub.ret", hub.ret)
		c.JSON(http.StatusOK, hub.ret)
	})

	router.GET("/getJobs", func(c *gin.Context) {
		fmt.Printf("%+v\n", hub)
		c.JSON(http.StatusOK, hub.ret)
	})

	router.POST("/changeTime", func(c *gin.Context) {
		//name := c.PostForm("hostname")
		command := c.PostForm("command")
		go func() {
			//<-time.After(time.Second * 10)
			klog.Info("new task")
			hub.NewTask([]string{command})
		}()
		c.JSON(http.StatusOK, hub.ret)
	})

	router.POST("/getHub", func(c *gin.Context) {
		var l int
		if len(hub.tasks) > 5 {
			l = len(hub.tasks) - 5
		}
		c.JSONP(http.StatusOK, gin.H{"tasks": hub.tasks[l:], "ret": hub.ret})
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
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
