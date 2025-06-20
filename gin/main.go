package main

import (
	"time"
	"github.com/KingBean4903/BenchHTTPRouters/models"
	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()

	r.GET("/stocks/:symbol", func(c *gin.Context) {
			symbol := c.Param("symbol")
			c.JSON(200, models.Stock{symbol, 182.3})
	})

	r.GET("/stocks/:symbol/history", func(c *gin.Context) {
			data := make([]models.HistoricalData, 30)
			for i := 0; i < 30; i++ {
					data[i] = models.HistoricalData{time.Now().AddDate(0, 0, -i).Format("2006-01-02"), 189 + float64(i)}
			}
			c.JSON(200, data)
	})

	r.GET("/market/trending", func(c *gin.Context) {
			
		c.JSON(200, gin.H{
				"trending" : []string{ "AAPL", "TSLA", "GOOGL", "AMZN", "MSFT", },
		})
	})
	
	r.Run(":8800")
}
