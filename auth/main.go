package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "timestamp": time.Now()})
	})
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
