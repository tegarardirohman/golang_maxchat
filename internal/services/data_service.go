package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"maxchat/internal/model"
	"os"
)

// LoadDataFromFile membaca data dari file teks dan mengisi data di memori.
func LoadDataFromFile(filePath string) error {
	// Baca file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v", filePath, err)
	}
	defer file.Close()

	// Baca isi file
	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v", filePath, err)
	}

	// Parsing JSON ke dalam model.Datas
	var datas []model.Data
	if err := json.Unmarshal(dataBytes, &datas); err != nil {
		return fmt.Errorf("could not unmarshal data: %v", err)
	}

	model.Datas = datas
	return nil
}

// GetAllData mengembalikan semua data yang ada.
func GetAllData() []model.Data {
	return model.Datas
}

// GetDataByCode mencari data berdasarkan code.
func GetDataByCode(code string) (*model.Data, error) {
	for _, data := range model.Datas {
		if data.Code == code {
			return &data, nil
		}
	}
	return nil, errors.New("data not found")
}

// CreateData menambahkan data baru ke dalam array Datas.
func CreateData(newData model.Data) {
	model.Datas = append(model.Datas, newData)
}

// UpdateData memperbarui data yang sudah ada berdasarkan code.
func UpdateData(code string, updatedData model.Data) error {
	for i, data := range model.Datas {
		if data.Code == code {
			model.Datas[i] = updatedData
			return nil
		}
	}
	return errors.New("data not found")
}

// DeleteData menghapus data berdasarkan code.
func DeleteData(code string) error {
	for i, data := range model.Datas {
		if data.Code == code {
			model.Datas = append(model.Datas[:i], model.Datas[i+1:]...)
			return nil
		}
	}
	return errors.New("data not found")
}

// FilterDataByModelAndTech memfilter data berdasarkan model dan teknologi.
func FilterDataByModelAndTech(modelFilter, techFilter string) []model.Data {
	var filteredData []model.Data
	for _, data := range model.Datas {
		if modelFilter == "" || (modelFilter != "" && data.Model == modelFilter) {
			if techFilter == "" || containsTech(data.Tech, techFilter) {
				filteredData = append(filteredData, data)
			}
		}
	}
	return filteredData
}

// containsTech memeriksa apakah data memiliki teknologi yang sesuai.
func containsTech(tech []string, filter string) bool {
	for _, t := range tech {
		if t == filter {
			return true
		}
	}
	return false
}
