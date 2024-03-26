package usecase

import (
	"fmt"
	"programbuku-v3/config"
	"programbuku-v3/model"

	"gorm.io/gorm"
)

var id uint

func DeleteBuku(db *gorm.DB) {
	fmt.Println("=================================")
	fmt.Println("Delete Buku")
	fmt.Println("=================================")
	ListBuku(db)

	fmt.Println("=================================")
	fmt.Print("Masukan ID Buku yang Ingin Dihapus : ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	var buku model.DaftarBuku
	buku, err = buku.GetByID(config.Mysql.Db, id)
	fmt.Println("ini test", buku)
	if err != nil {
		fmt.Println("Buku dengan ID", id, "tidak ditemukan")
		return
	}

	err = buku.DeleteById(config.Mysql.Db)
	if err != nil {
		fmt.Printf("Terjadi error saat menghapus buku: %v\n", err)
		return
	}
	fmt.Println("Buku dengan ID", buku.ID, " dengan judul : ", buku.Judul, " Berhasil Dihapus.")
}
