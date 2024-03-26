package usecase

import (
	"fmt"
	"programbuku-v3/config"
	"programbuku-v3/model"
	"sort"
	"sync"

	"gorm.io/gorm"
)

func LihatBuku(ch <-chan model.DaftarBuku, db *gorm.DB, chPesanan chan model.DaftarBuku, wg *sync.WaitGroup) {
	for buku := range ch {
		_, err := buku.GetAll(config.Mysql.Db)
		if err != nil {
			fmt.Printf("Terjadi error saat menyimpan buku: %v\n", err)
			continue
		}
		chPesanan <- buku
	}
	wg.Done()
}

func ListBuku(db *gorm.DB) {
	fmt.Println("Lihat List Buku")
	fmt.Println("=================================")
	fmt.Println("Memuat data ...")
	var listBook []model.DaftarBuku
	daftarBuku := &model.DaftarBuku{}

	books, err := daftarBuku.GetAll(config.Mysql.Db)
	if err != nil {
		fmt.Printf("Gagal memuat data buku: %v\n", err)
		return
	}
	listBook = books
	wg := sync.WaitGroup{}
	chPesanan := make(chan model.DaftarBuku, len(listBook))
	ch := make(chan model.DaftarBuku)

	wg.Add(len(listBook))

	for _, book := range listBook {
		go LihatBuku(ch, db, chPesanan, &wg)
		ch <- book
	}
	close(ch)

	wg.Wait()

	close(chPesanan)

	var orderedBooks []model.DaftarBuku
	for dataPesanan := range chPesanan {
		orderedBooks = append(orderedBooks, dataPesanan)
	}

	// Urutkan buku berdasarkan waktu pembuatan
	sort.Slice(orderedBooks, func(i, j int) bool {
		return orderedBooks[i].CreatedAt.Before(orderedBooks[j].CreatedAt)
	})

	// Tampilkan buku yang telah diurutkan
	for urutan, book := range orderedBooks {
		fmt.Printf("%d. ID Buku : %d, ISBN : %s, Penulis : %s, Tahun : %d, Judul : %s, Gambar : %s, Stok : %d\n",
			urutan+1,
			book.ID,
			book.ISBN,
			book.Penulis,
			book.Tahun,
			book.Judul,
			book.Gambar,
			book.Stok,
		)
	}
}
