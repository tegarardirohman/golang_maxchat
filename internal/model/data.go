package model

// Data struct untuk menyimpan informasi tentang model.
type Data struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Model       string   `json:"model"`
	Tech        []string `json:"tech"`
	Status      string   `json:"status"`
}

// In-memory storage untuk data yang akan dimuat dari file.
var Datas []Data
