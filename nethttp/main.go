package main

import (
	"encoding/json"
	"os"
	"flag"
	"runtime/pprof"
	"net/http"
	"time"
	"github.com/KingBean4903/BenchHTTPRouters/models"
)


var (
	 cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	 memprofile = flag.String("memprofile", "", "write memory profile to file")
	 port       = flag.String("port", "8700", "HTTP server port")
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
	
	http.HandleFunc("/stocks", func(w http.ResponseWriter, r *http.Request) {
			symbol := r.URL.Path[len("/stocks/"):]
			json.NewEncoder(w).Encode(models.Stock{symbol, 182.3})
	})



	http.HandleFunc("/stocks/", func(w http.ResponseWriter, r *http.Request) {
			symbol := r.URL.Path[len("/stocks"):]

			if r.URL.Path == "/stocks/"+symbol+"/history" {
					data := make([]models.HistoricalData, 30)
					for i := 0 ; i < 30; i++ {
							data[i] = models.HistoricalData{time.Now().AddDate(0, 0, -i).Format("2006-01-02"), 180 + float64(i)}
					}
					json.NewEncoder(w).Encode(data)
				}	
	})

	http.HandleFunc("/market/trending", func(w http.ResponseWriter, r *http.Request) {
			
			json.NewEncoder(w).Encode(map[string][]string{
					"trending": {"AAPL", "TSLA", "GOOGL", "AMZN", "MSFT"},
			})

	})

	http.ListenAndServe(":8700", nil)

}
