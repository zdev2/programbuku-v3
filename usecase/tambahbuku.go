package usecase

import (
	"bufio"
	"fmt"
	"os"
	"programbuku-v3/config"
	"programbuku-v3/model"
	"strings"
	"sync"

	"gorm.io/gorm"
)

var inputUser = bufio.NewReader(os.Stdin)

func TambahBuku(db *gorm.DB) {
	var id, stok, tahun uint
	penulis := ""
	judul := ""
	// penulis := ""
	gambar := ""
	fmt.Println("=================================")
	fmt.Println("Tambah Pesanan")
	fmt.Println("=================================")

	draftBuku := []model.DaftarBuku{}

	for {
		fmt.Print("Silahkan Masukan Id Buku : ")
		_, err := fmt.Scanln(&id)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
		// if kodeBukuExists(id) {
		// 	fmt.Println("Kode buku sudah digunakan. Masukkan kode buku yang berbeda.")
		// } else {
		// 	break
		// }
		fmt.Print("Silahkan Masukan ISBN : ")
		// untuk user Windows, gunakan yang dicomment (\r) :
		// menuPelanggan, err := inputanUser.ReadString('\r')
		isbn, err := inputUser.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		isbn = strings.Replace(
			isbn,
			"\n",
			"",
			1)

		// special treatment untuk windows
		isbn = strings.Replace(
			isbn,
			"\r",
			"",
			1)

		fmt.Print("Silahkan Masukan Penulis : ")
		_, err = fmt.Scanln(&penulis)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		fmt.Print("Silahkan Masukan Tahun : ")
		_, err = fmt.Scanln(&tahun)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
		fmt.Print("Silahkan Masukan Judul : ")
		_, err = fmt.Scanln(&judul)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		fmt.Print("Silahkan Masukan Gambar : ")
		_, err = fmt.Scanln(&gambar)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
		fmt.Print("Silahkan Masukan Stok : ")
		_, err = fmt.Scanln(&stok)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
		// Simpan ID dan Tanggal
		draftBuku = append(draftBuku, model.DaftarBuku{
			ID:      id,
			ISBN:    isbn,
			Penulis: penulis,
			Tahun:   tahun,
			Judul:   judul,
			Gambar:  gambar,
			Stok:    stok,
		})
		var pilMenuBuku = 0
		fmt.Println("Ketik 1 untuk tambah Buku, ketik 0 untuk keluar")
		_, err = fmt.Scanln(&pilMenuBuku)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		if pilMenuBuku == 0 {
			break
		}
	}

	fmt.Println("Menambah buku...")

	ch := make(chan model.DaftarBuku)

	wg := sync.WaitGroup{}

	jumlahAntrian := 5

	// Mendaftarkan receiver/pemroses data
	for i := 0; i < jumlahAntrian; i++ {
		wg.Add(1)
		go simpanBuku(ch, db, &wg, i)
	}

	// Mengirimkan data ke channel
	for _, book := range draftBuku {
		ch <- book
	}
	// fmt.Println("tes thasil : ", ch)
	close(ch)

	wg.Wait()

	fmt.Println("Berhasil Menambah Buku!")
}

func simpanBuku(ch <-chan model.DaftarBuku, db *gorm.DB, wg *sync.WaitGroup, noAntrian int) {

	for buku := range ch {
		err := buku.CreateBuku(config.Mysql.Db)
		if err != nil {
			fmt.Printf("Terjadi error saat menyimpan buku: %v\n", err)
			continue
		}

		fmt.Printf("Antrian No %d Memproses Kode Buku : %d!\n", noAntrian, buku.ID)
	}
	wg.Done()
}
