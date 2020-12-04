package date_agent

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
)

func InitHttp(addr string, hub *Hub) *http.Server {
	router := gin.New()
	//router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: writer}), gin.RecoveryWithWriter(writer))
	router.Use(cors.Default())
	router.GET("/hello", func(c *gin.Context) {
		// todo handle request and return data by hub
		c.JSON(http.StatusOK, "hello world")
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
