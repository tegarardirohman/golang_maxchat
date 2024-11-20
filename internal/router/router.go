package router

import (
	"encoding/json"
	"maxchat/internal/services"
	"net/http"
)

// SetupRoutes mengonfigurasi dan mengembalikan routing handler
func SetupRoutes() {
	http.HandleFunc("/datas", getAllData)
	http.HandleFunc("/data/filter", filterData)
	http.HandleFunc("/data/", getDataByCode)
}

// getAllData mengembalikan semua data dalam format JSON
func getAllData(w http.ResponseWriter, r *http.Request) {
	data := service.GetAllData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// getDataByCode mengembalikan data berdasarkan kode dalam format JSON
func getDataByCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/data/"):]
	data, err := service.GetDataByCode(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// filterData memfilter data berdasarkan model dan tech
func filterData(w http.ResponseWriter, r *http.Request) {
	modelFilter := r.URL.Query().Get("model")
	techFilter := r.URL.Query().Get("tech")
	data := service.FilterDataByModelAndTech(modelFilter, techFilter)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
