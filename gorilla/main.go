package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/KingBean4903/BenchHTTPRouters/models"
)

func main() {

	r := mux.NewRouter()
	
	r.HandleFunc("/stocks/{symbol}", func(w http.ResponseWriter, r *http.Request) {
				vars := mux.Vars(r)
				json.NewEncoder(w).Encode(models.Stock{vars["symbol"], 182.3})
			}).Methods("GET")

	r.HandleFunc("/stocks/{symbol}/history", func(w http.ResponseWriter, r *http.Request) {
			data := make([]models.HistoricalData, 30)

			for i:= 0; i < 30; i++ {
				data[i] = models.HistoricalData{time.Now().AddDate(0, 0, -i).Format("2006-01-02"),180 + float64(i)}
			}
			json.NewEncoder(w).Encode(data)
	}).Methods("GET")

	r.HandleFunc("/market/trending", func(w http.ResponseWriter, r *http.Request) {
		
			json.NewEncoder(w).Encode(map[string][]string{
					"trending" : {"AAPL", "TSLA", "GOOGL", "AMZN", "MSFT"},
			})
				
	}).Methods("GET")

	http.ListenAndServe(":8900", r)

}

