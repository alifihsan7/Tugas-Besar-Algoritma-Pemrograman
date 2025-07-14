package main

import "fmt"

type mataKuliah struct { // data mata kuliah mahasiswa
	No     int
	KodeMK string
	NamaMK string
	SKS    int
}

type mahasiswa struct { // data mahasiswa
	listMataKuliah                                  [8]mataKuliah
	listUts, listUas, listQuiz, listTotal, IPMatkul [8]float64
	GradeMatkul                                     [8]string
	listTotalSKS                                    int
	listIPS                                         float64
	IPgrade                                         string
	Nmatkul                                         int
	nama                                            string
	kelas                                           string
	nim                                             int
	jurusan                                         string
}

type tabMahasiswa [100]mahasiswa  // array mahasiswa
type tabMataKuliah [10]mataKuliah // array mata kuliah

func main() {
	// deklarasi variabel
	var mhsw tabMahasiswa
	var matkul tabMataKuliah
	var menuinput, nMahasiswa int
	var keluar bool
	nMahasiswa = 0
	keluar = false

	// Daftar Mata Kuliah yang tersedia
	matkul = tabMataKuliah{
		{No: 1, KodeMK: "CII1G3", NamaMK: "Matematika Diskrit", SKS: 3},
		{No: 2, KodeMK: "CII1H3", NamaMK: "Kalkulus Lanjut", SKS: 4},
		{No: 3, KodeMK: "CII1F4", NamaMK: "Algoritma Pemrograman", SKS: 3},
		{No: 4, KodeMK: "CII2D3", NamaMK: "Matriks dan Ruang Vektor", SKS: 3},
		{No: 5, KodeMK: "UWJXA2", NamaMK: "Bahasa Inggris", SKS: 3},
		{No: 6, KodeMK: "UKJXC2", NamaMK: "Bahasa Indonesia", SKS: 3},
		{No: 7, KodeMK: "CII1J3", NamaMK: "Pemodelan Basis Data", SKS: 2},
		{No: 8, KodeMK: "CII1I3", NamaMK: "Sistem Digital", SKS: 2},
		{No: 9, KodeMK: "CII1B3", NamaMK: "Logika Matematika", SKS: 3},
		{No: 10, KodeMK: "CII1C2", NamaMK: "Statistika", SKS: 3},
	}

	//login
	login()

	//menu utama
	for !keluar {
		mainmenu()
		fmt.Scan(&menuinput)
		switch menuinput {
		case 1:
			//menu tambah data mahasiswa
			tambahDataMahasiswa(&mhsw, &nMahasiswa, matkul)
		case 2:
			//menu tambah data mata kuliah mahasiswa
			tambahMatakuliah(&mhsw, nMahasiswa, matkul)
		case 3:
			// mengubah/menghapus data mahasiswa serta mata kuliah mahasiswa
			ubahData(&mhsw, &nMahasiswa, matkul)
		case 4:
			// menampilkan data mahasiswa, mata kuliah, dll
			tampilkanData(mhsw, matkul, nMahasiswa)
		case 0:
			fmt.Println("Keluar")
			keluar = true
		default:
			fmt.Println("Pilihan tidak valid, silakan input ulang")
			keluar = false
		}
	}
	fmt.Println("Terima kasih! Program berakhir.")

}

func login() {
	var username, password string
	var keluar bool = false

	fmt.Println("+---------------------------------------------+")
	fmt.Println("| SELAMAT DATANG DI APLIKASI NILAI MAHASISWA! |")
	fmt.Println("| SILAKAN LOGIN TERLEBIH DAHULU               |")
	fmt.Println("+---------------------------------------------+")

	// Contoh login sederhana
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)
	for !keluar {
		if username == "DOSEN" && password == "TELKOM" {
			fmt.Println("--------------------------------------------")
			fmt.Println("LOGIN BERHASIL!")
			fmt.Println("SELAMAT DATANG")
			fmt.Println("--------------------------------------------")
			keluar = true
		} else {
			fmt.Println("Username atau password salah. Silakan coba lagi.")
			login()
		}
	}
}

func mainmenu() {
	fmt.Println("+---------------------------------------+")
	fmt.Println("|               MAIN MENU               |")
	fmt.Println("+---------------------------------------+")
	fmt.Println("| 1. Tambah data mahasiswa              |")
	fmt.Println("| 2. Tambah data mata kuliah mahasiswa  |")
	fmt.Println("| 3. Ubah/hapus data                    |")
	fmt.Println("| 4. Tampilkan data                     |")
	fmt.Println("| 0. Exit                               |")
	fmt.Println("+---------------------------------------+")
	fmt.Print("Pilih menu: [1/2/3.../4/0]? ")
}

func searchingMahasiswa(A tabMahasiswa, N int, cariNim int) int {
	for i := 0; i < N; i++ {
		if A[i].nim == cariNim {
			return i
		}
	}
	return -1
}

func cekMatkul(A tabMahasiswa, matkul tabMataKuliah, idx, matkulIndex int) int {
	for i := 0; i < A[idx].Nmatkul; i++ {
		if A[idx].listMataKuliah[i].KodeMK == matkul[matkulIndex].KodeMK {
			return 1
		}
	}
	return -1
}

func hitungIPMatkul(total float64) float64 {
	if total > 80 && total <= 100 {
		return 4.00
	} else if total > 70 && total <= 80 {
		return 3.50
	} else if total > 65 && total <= 70 {
		return 3.00
	} else if total >= 60 && total <= 65 {
		return 2.5
	} else if total > 50 && total <= 60 {
		return 2
	} else if total > 40 && total <= 50 {
		return 1
	} else {
		return 0
	}
}
func nilaiAkhir(A *tabMahasiswa, idx int) {
	A[idx].listIPS = 0
	for j := 0; j < A[idx].Nmatkul; j++ {
		A[idx].listIPS = A[idx].listIPS + (A[idx].IPMatkul[j] * float64(A[idx].listMataKuliah[j].SKS))
	}
	A[idx].listIPS = A[idx].listIPS / float64(A[idx].listTotalSKS)
	A[idx].IPgrade = hitungIPGrade(A[idx].listIPS)
}

func totalSKS(A *tabMahasiswa, idx int) {
	A[idx].listTotalSKS = 0
	for j := 0; j < A[idx].Nmatkul; j++ {
		A[idx].listTotalSKS = A[idx].listTotalSKS + A[idx].listMataKuliah[j].SKS
	}
}

func hitungMatkulGrade(total float64) string {
	if total > 80 && total <= 100 {
		return "A"
	} else if total > 70 && total <= 80 {
		return "AB"
	} else if total > 65 && total <= 70 {
		return "B"
	} else if total >= 60 && total <= 65 {
		return "BC"
	} else if total > 50 && total <= 60 {
		return "C"
	} else if total > 40 && total <= 50 {
		return "D"
	} else {
		return "E"
	}
}

func hitungIPGrade(total float64) string {
	if total > 3.5 && total <= 4 {
		return "A"
	} else if total > 3 && total <= 3.5 {
		return "AB"
	} else if total > 2.5 && total <= 3 {
		return "B"
	} else if total >= 2 && total <= 2.5 {
		return "BC"
	} else if total > 1.5 && total <= 1 {
		return "C"
	} else if total > 1 && total <= 1.5 {
		return "D"
	} else {
		return "E"
	}
}
