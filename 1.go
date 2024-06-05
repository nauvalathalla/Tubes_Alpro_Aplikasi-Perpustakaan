package main

import (
	"fmt"
	"sort"
)

type buku struct {
	Kode                              int
	Judul, Jenis, Pengarang, Penerbit string
	TahunTerbit                       int
	totalPeminjaman                   int
	totalPeminjam                     int
}

const NMAX int = 100

type Tabbuku [NMAX]buku
type Tabpinjam [NMAX]buku

func main() {
	var buku Tabbuku
	var b Tabpinjam
	var pilih int
	var jumlahBuku int
	var judul, judulbaru string

	var n int
	// var jumlahPeminjaman int

	jumlahBuku = 0
	// jumlahPeminjaman = 0
	for {
		menu()
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahBuku(&buku, &jumlahBuku)
		} else if pilih == 2 {
			editBuku(&buku, jumlahBuku)
		} else if pilih == 3 {
			hapusBuku(&buku, &jumlahBuku)
		} else if pilih == 4 {
			cetakBuku(buku, jumlahBuku)

		} else if pilih == 5 {
			var Judul string
			fmt.Println("Masukkan Judul Buku yang ingin dicari: ")
			fmt.Scan(&Judul)
			idx := cariBuku(buku, jumlahBuku, Judul) // Mencari buku dan menyimpan indeksnya

			if idx == -1 {
				fmt.Println("Buku tidak ditemukan.")
			} else {
				fmt.Println("Buku ditemukan:", buku[idx].Judul) // Menampilkan data buku
			}

		} else if pilih == 12 {
			Denda(buku, n)
		} else if pilih == 6 {
			fmt.Scan(&judul)
			tambahBukuPinjam(&buku, &b, &n, judul, judulbaru)
		} else if pilih == 7 {
			editBukuPinjam(&b, n)
		} else if pilih == 8 {
			hapusBukuPinjam(&b, &n)
		} else if pilih == 9 {
			cetakBukuPinjam(b, n, buku)
		} else if pilih == 10 {
			cetakBukuSeringDipinjam(buku, jumlahBuku)
		} else if pilih == 11 {
			cetakBukuJarangDipinjam(buku, jumlahBuku)
		}
	}
}

func (b *buku) tambahPeminjaman() {
	b.totalPeminjaman++
}

func (b *buku) kurangiPeminjaman() {
	b.totalPeminjaman--
}

func cetakBukuPinjam(b Tabpinjam, n int, a Tabbuku) {
	urutBuku(&a, n)
	for i := 0; i < n; i++ {
		fmt.Println(b[i].Kode, b[i].Judul, b[i].Pengarang, b[i].Penerbit, b[i].TahunTerbit)
	}
}

func cetakBukuSeringDipinjam(b Tabbuku, n int) {
	sort.Slice(b[:n], func(i, j int) bool {
		return b[i].totalPeminjaman > b[j].totalPeminjaman
	})
	for i := 0; i < n; i++ {
		fmt.Println(b[i].Kode, b[i].Judul, b[i].Pengarang, b[i].Penerbit, b[i].TahunTerbit, "Total Peminjaman:", b[i].totalPeminjaman)
	}
}

func cetakBukuJarangDipinjam(b Tabbuku, n int) {
	sort.Slice(b[:n], func(i, j int) bool {
		return b[i].totalPeminjaman < b[j].totalPeminjaman
	})
	for i := 0; i < n; i++ {
		fmt.Println(b[i].Kode, b[i].Judul, b[i].Pengarang, b[i].Penerbit, b[i].TahunTerbit, "Total Peminjaman:", b[i].totalPeminjaman)
	}
}

func menu() {
	fmt.Println("\n=== Aplikasi Perpustakaan ===")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Edit Buku")
	fmt.Println("3. Hapus Buku")
	fmt.Println("4. Tampilkan Buku")
	fmt.Println("5. Cari Buku")
	fmt.Println("6.Tambah Peminjaman")
	fmt.Println("7. Edit Peminjaman")
	fmt.Println("8. Hapus Peminjaman")
	fmt.Println("9. Tampilkan Peminjaman")
	fmt.Println("10. Buku Sering Dipinjam")
	fmt.Println("11. Buku Jarang Dipinjam")
	fmt.Println("12. Denda")
	fmt.Print("Pilih menu: ")

}

func tambahBuku(buku *Tabbuku, jumlahBuku *int) {

	fmt.Println("Masukkan Kode Buku: ")
	fmt.Scan(&buku[*jumlahBuku].Kode)
	fmt.Println("Masukkan Judul Buku: ")
	fmt.Scan(&buku[*jumlahBuku].Judul)
	fmt.Println("Masukkan Pengarang Buku: ")
	fmt.Scan(&buku[*jumlahBuku].Pengarang)
	fmt.Println("Masukkan Penerbit Buku: ")
	fmt.Scan(&buku[*jumlahBuku].Penerbit)
	fmt.Println("Masukkan Tahun Terbit: ")
	fmt.Scan(&buku[*jumlahBuku].TahunTerbit)

	*jumlahBuku++
	fmt.Println("Data buku berhasil ditambahkan!")
}

func editBuku(b *Tabbuku, n int) {
	var judul string
	var idx int
	fmt.Scan(&judul)
	idx = cariBuku(*b, n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		fmt.Scan(&b[idx].Judul)
		fmt.Println("Data buku berhasil diedit!")
	}

}
func hapusBuku(b *Tabbuku, n *int) {
	var judul string
	var idx int
	fmt.Scan(&judul)
	idx = cariBuku(*b, *n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		for i := idx; i < *n-1; i++ {
			b[i] = b[i+1]
		}
		*n--
		fmt.Println("Data buku berhasil dihapus!")
	}

}
func urutBuku(b *Tabbuku, n int) {

	var i int
	var j int
	var idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if b[idx_min].Judul > b[j].Judul {
				idx_min = j
			}
			j = j + 1
		}
		t := b[idx_min]
		b[idx_min] = b[i-1]
		b[i-1] = t
		i = i + 1
	}

}
func cetakBuku(b Tabbuku, n int) {
	urutBuku(&b, n)
	for i := 0; i < n; i++ {
		fmt.Println(b[i].Kode, b[i].Judul, b[i].Pengarang, b[i].Penerbit, b[i].TahunTerbit)
	}
}

func tambahBukuPinjam(b *Tabbuku, p *Tabpinjam, n *int, judul, judulbaru string) {
	var idx int
	cariBuku(*b, *n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		p[*n] = b[idx]
		b[idx].tambahPeminjaman() // Tambah jumlah peminjaman
		*n++
		fmt.Println("Data buku berhasil ditambahkan!")
	}
}

func editBukuPinjam(b *Tabpinjam, n int) {
	var judul string
	var idx int
	fmt.Scan(&judul)
	idx = cariBukuPinjam(*b, n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		fmt.Scan(&b[idx].Judul)
		fmt.Println("Data buku berhasil diedit!")
	}

}

func hapusBukuPinjam(b *Tabpinjam, n *int) {
	var judul string
	var idx int
	fmt.Scan(&judul)
	idx = cariBukuPinjam(*b, *n, judul)
	if idx == -1 {
		fmt.Println("buku tidak ditemukan")
	} else {
		for i := idx; i < *n-1; i++ {
			b[i] = b[i+1]
		}
		*n--
		fmt.Println("Data buku berhasil dihapus!")
	}

}

func cariBuku(b Tabbuku, n int, judul string) int {
	/* mengembalikan indeks dari X apabila X ditemukan di dalam array T yang berisi
	n buah teks, atau -1 apabilaX tidak ditemukan */
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if b[j].Judul == judul {
			found = j
		}
		j = j + 1
	}
	return found
}

func cariBukuPinjam(b Tabpinjam, n int, judul string) int {
	/* mengembalikan indeks dari X apabila X ditemukan di dalam array T yang berisi
	n buah teks, atau -1 apabila X tidak ditemukan */
	var found int = -1
	var j int = 0
	for j < n && found == -1 {
		if b[j].Judul == judul {
			found = j
		}
		j = j + 1
	}
	return found
}

func hitungDenda(a, b, c, d, e, f int) int {
	var denda int
	var selisihHari int
	// Hitung denda
	a = a * 3600
	b = b * 60
	hasilMasuk := a + b + c
	d = d * 3600
	e = e * 60
	HasilKeluar := d + e + f
	selisihHari = (HasilKeluar - hasilMasuk)
	selisihHari = selisihHari / 3600

	denda = selisihHari * 1000

	if denda > 10000 {
		denda = 10000
	}
	return denda
}

func Denda(buku Tabbuku, n int) {
	var x int
	var temp int
	var masukJam, masukDetik, masukMenit int
	var keluarJam, keluarMenit, keluarDetik int
	fmt.Println("Masukkan kode Buku Yang ingin di cari")
	fmt.Scan(&x)
	for i := 0; i < n; i++ {
		if buku[i].Kode == x {
			temp = i
		}
	}
	fmt.Println("Masukkan Waktu Masuk dan Keluar")
	fmt.Scan(&masukJam, &masukDetik, &masukMenit)
	fmt.Scan(&keluarJam, &keluarMenit, &keluarDetik)
	fmt.Println(buku[temp].Kode)
	fmt.Println(buku[temp].Judul)
	fmt.Println(buku[temp].Pengarang)
	fmt.Println(buku[temp].Penerbit)
	fmt.Println(buku[temp].TahunTerbit)
	fmt.Println("Denda :", hitungDenda(masukJam, masukMenit, masukDetik, keluarJam, keluarMenit, keluarDetik))
}