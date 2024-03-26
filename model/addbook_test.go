package model_test

import (
	"fmt"
	"programbuku-v3/config"
	"programbuku-v3/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}

func TestCreateCar(t *testing.T) {
	Init()
	carData := model.DaftarBuku{
		ISBN:    "test",
		Penulis: "penulis1",
		Tahun:   1010,
		Judul:   "judul",
		Gambar:  "1021",
		Stok:    100,
	}
	err := carData.CreateBuku(config.Mysql.Db)
	assert.Nil(t, err)
}
func TestGetCarByID(t *testing.T) {
	Init()
	carData := model.DaftarBuku{
		ID: 73201,
	}
	data, err := carData.GetByID(config.Mysql.Db, carData.ID)
	assert.Nil(t, err)
	fmt.Println(data)
}

func TestGetAll(t *testing.T) {
	Init()
	carData := model.DaftarBuku{
		ISBN:    "test",
		Penulis: "penulis1",
		Tahun:   1010,
		Judul:   "judul",
		Gambar:  "1021",
		Stok:    100,
	}
	err := carData.CreateBuku(config.Mysql.Db)
	assert.Nil(t, err)

	res, err := carData.GetAll(config.Mysql.Db)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)
	// assert.Len(t, res, 2)

	fmt.Println(res)
}

func TestUpdate(t *testing.T) {
	Init()
	var err error
	carInsert := model.DaftarBuku{
		ISBN:    "test",
		Penulis: "penulis1",
		Tahun:   1010,
		Judul:   "judul",
		Gambar:  "1021",
		Stok:    100,
	}
	err = carInsert.CreateBuku(config.Mysql.Db)
	assert.Nil(t, err)
	carData := model.DaftarBuku{
		ID:      73201,
		ISBN:    "test updated version",
		Penulis: "penulis1 updated version",
		Tahun:   1010,
		Judul:   "judul",
		Gambar:  "1021",
		Stok:    100,
	}
	err = carData.UpdateOne(config.Mysql.Db)
	assert.Nil(t, err)
}

func TestDeleteById(t *testing.T) {
	Init()
	var err error
	carInsert := model.DaftarBuku{
		ISBN:    "test",
		Penulis: "penulis1",
		Tahun:   1010,
		Judul:   "judul",
		Gambar:  "1021",
		Stok:    100,
	}
	err = carInsert.CreateBuku(config.Mysql.Db)
	assert.Nil(t, err)
	carData := model.DaftarBuku{
		ID: 73201,
	}
	err = carData.DeleteById(config.Mysql.Db)
	assert.Nil(t, err)
}
