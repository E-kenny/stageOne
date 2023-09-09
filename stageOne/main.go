package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
  now := time.Now()
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "slack_name": "E_kenny",
      "current_day": now.Format("Monday"),
      "utc_time": now.Format(time.RFC3339),
      "track": "backend",
      "github_file_url": "https://github.com/E-kenny/stageOne/blob/main/file_name.ext",
      "github_repo_url": "https://github.com/E-kenny/stageOne",
      "status_code": http.StatusOK,
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}