package usecase

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"programbuku-v3/config"
	"programbuku-v3/model"
	"strconv"

	"gorm.io/gorm"
)

func ImportDataFromCSV(db *gorm.DB, filePath string) error {
	var err error
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for cek, line := range lines {
		if cek == 0 {
			continue
		}
		id, err := strconv.ParseUint(line[0], 20, 64)
		if err != nil {
			return err
		}
		year, err := strconv.ParseUint(line[3], 20, 64)
		if err != nil {
			return err
		}

		stok, err := strconv.ParseUint(line[6], 20, 64)
		if err != nil {
			return err
		}
		book := model.DaftarBuku{
			ID:      uint(id),
			ISBN:    line[1],
			Penulis: line[2],
			Tahun:   uint(year),
			Judul:   line[4],
			Gambar:  line[5],
			Stok:    uint(stok),
		}
		err = book.UpsertBuku(config.Mysql.Db)
		if err != nil {
			fmt.Println("error creating book")
			return err
		}
	}

	return nil
}

func ImportFile() {
	var filePath string
	var err error
	fmt.Print("Masukkan path lokasi file CSV: ")
	fmt.Scanln(&filePath)

	filePath, err = filepath.Abs(filePath)
	if err != nil {
		panic(err)
	}

	err = ImportDataFromCSV(&gorm.DB{}, filePath)
	if err != nil {
		panic(err)
	}
}
