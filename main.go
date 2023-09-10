package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
  now := time.Now()
  r := gin.Default()
  r.GET("/api", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "slack_name": c.Query("slack_name"),
      "current_day": now.Format("Monday"),
      "utc_time": now.Format(time.RFC3339),
      "track": c.Query("track"),
      "github_file_url": "https://github.com/E-kenny/stageOne/blob/main/main.go",
      "github_repo_url": "https://github.com/E-kenny/stageOne",
      "status_code": http.StatusOK,
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}