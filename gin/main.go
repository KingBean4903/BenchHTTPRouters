package main

import (
	"time"
	"runtime/pprof"
	"flag"
	"os"
	"github.com/KingBean4903/BenchHTTPRouters/models"
	"github.com/gin-gonic/gin"
)

var (
	 cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	 memprofile = flag.String("memprofile", "", "write memory profile to file")
	 port       = flag.String("port", "8800", "HTTP server port")
)

func main() {

	flag.Parse()
	
	if *cpuprofile != "" {
		
			f, err := os.Create(*cpuprofile)
			if err != nil {
					panic(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()

	}


	if *memprofile != " " {
			defer func() {
					f, err := os.Create(*memprofile)
					if err != nil {
							panic(err)
					}
					pprof.WriteHeapProfile(f)
					f.Close()
			}()
	}


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


	time.Sleep(5 * time.Second)
}
