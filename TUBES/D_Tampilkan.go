package main

import "fmt"

func tampilkanData(mhsw tabMahasiswa, matkul tabMataKuliah, nMahasiswa int) {
	var menuinput int
	var keluar bool = false
	for !keluar {
		fmt.Println("+--------------------------------------------------------------+")
		fmt.Println("|                      MENU TAMPILKAN DATA                     |")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("| 1. Tampilkan data mahasiswa                                  |")
		fmt.Println("| 2. Tampilkan data mata kuliah                                |")
		fmt.Println("| 3. Tampilkan data mata kuliah yang diambil oleh mahasiswa    |")
		fmt.Println("| 4. Tampilkan data mahasiswa yang mengambil suatu mata kuliah |")
		fmt.Println("| 5. Tampilkan transkrip nilai mahasiswa                       |")
		fmt.Println("| 0. Kembali                                                   |")
		fmt.Println("+--------------------------------------------------------------+")
		fmt.Print("Pilih data yang ingin ditampilkan: [1/2/3../6/0]? ")
		fmt.Scan(&menuinput)
		switch menuinput {
		case 1:
			fmt.Println("+----------------------------------------------------+")
			fmt.Println("|            MENU TAMPILAN DATA MAHASISWA            |")
			fmt.Println("+----------------------------------------------------+")
			fmt.Println("| 1. Urut berdasarkan Nilai                          |")
			fmt.Println("| 2. Urut berdasarkan SKS                            |")
			fmt.Println("| 3. Kembali                                         |")
			fmt.Println("| 0. Kembali ke main menu                            |")
			fmt.Println("+----------------------------------------------------+")
			fmt.Print("Pilih menu: [1/2/3/0]? ")
			fmt.Scan(&menuinput)
			switch menuinput {
			case 1:
				urutNilai(&mhsw, nMahasiswa)
				fmt.Println("Data Mahasiswa Berdasarkan Nilai:")
				tampilkanDataMahasiswa(mhsw, nMahasiswa)
				fmt.Println("+--------------------------+")
				fmt.Println("|           MENU           |")
				fmt.Println("+--------------------------+")
				fmt.Println("| 1. Kembali               |")
				fmt.Println("| 0. Kembali ke main menu  |")
				fmt.Println("+--------------------------+")
				fmt.Print("Pilih menu: [1/0]? ")
				fmt.Scan(&menuinput)
				switch menuinput {
				case 1:
					keluar = false
				case 0:
					keluar = true
				default:
					fmt.Println("Menu tidak valid, otomatis kembali")
				}
			case 2:
				urutSKS(&mhsw, matkul, nMahasiswa)
				fmt.Println("Data Mahasiswa Berdasarkan SKS:")
				tampilkanDataMahasiswa(mhsw, nMahasiswa)
				fmt.Println("+--------------------------+")
				fmt.Println("|           MENU           |")
				fmt.Println("+--------------------------+")
				fmt.Println("| 1. Kembali               |")
				fmt.Println("| 0. Kembali ke main menu  |")
				fmt.Println("+--------------------------+")
				fmt.Print("Pilih menu: [1/0]? ")
				fmt.Scan(&menuinput)
				switch menuinput {
				case 1:
					keluar = false
				case 2:
					keluar = true
				default:
					fmt.Println("Menu tidak valid, otomatis kembali")
				}
			case 3:
				keluar = false
			case 4:
				keluar = true
			default:
				fmt.Println("Menu tidak valid, otomatis kembali")
			}
		case 2:
			tampilkanDataMatkul(matkul)
			fmt.Println("Data berhasil ditampilkan.")
			fmt.Println("+--------------------------+")
			fmt.Println("|           MENU           |")
			fmt.Println("+--------------------------+")
			fmt.Println("| 1. Kembali               |")
			fmt.Println("| 0. Kembali ke main menu  |")
			fmt.Println("+--------------------------+")
			fmt.Print("Pilih menu: [1/0]? ")
			fmt.Scan(&menuinput)
			switch menuinput {
			case 1:
				keluar = false
			case 0:
				keluar = true
			default:
				fmt.Println("Menu tidak valid, otomatis kembali")
			}
		case 3:
			var nim int
			fmt.Print("Masukkan NIM (Nomor Induk Mahasiswa) yang ingin ditampilkan mata kuliahnya: ")
			fmt.Scan(&nim)
			tampilkanDataMatakuliahDiambilMahasiswa(mhsw, nMahasiswa, nim)
			fmt.Println("+--------------------------+")
			fmt.Println("|           MENU           |")
			fmt.Println("+--------------------------+")
			fmt.Println("| 1. Kembali               |")
			fmt.Println("| 0. Kembali ke main menu  |")
			fmt.Println("+--------------------------+")
			fmt.Print("Pilih menu: [1/0]? ")
			fmt.Scan(&menuinput)
			switch menuinput {
			case 1:
				keluar = false
			case 0:
				keluar = true
			default:
				fmt.Println("Menu tidak valid, otomatis kembali")
			}

		case 4:
			var no int
			tampilkanDataMatkul(matkul)
			fmt.Print("Pilih mata kuliah yang ingin ditampilkan daftar mahasiswanya: [1/2/3.../10]: ")
			fmt.Scan(&no)
			for !keluar {
				if no >= 1 && no <= 10 {
					keluar = true
				} else {
					tampilkanDataMatkul(matkul)
					fmt.Println("Mata kuliah tidak valid, silakan ambil ulang.")
					fmt.Print("Pilih mata kuliah yang ingin ditampilkan daftar mahasiswanya: [1/2/3.../10]: ")
					fmt.Scan(&no)
				}
			}
			tampilkanMahasiswaMengambilMatakuliah(mhsw, matkul, nMahasiswa, no)
			fmt.Println("+--------------------------+")
			fmt.Println("|           MENU           |")
			fmt.Println("+--------------------------+")
			fmt.Println("| 1. Kembali               |")
			fmt.Println("| 0. Kembali ke main menu  |")
			fmt.Println("+--------------------------+")
			fmt.Print("Pilih menu: [1/0]? ")
			fmt.Scan(&menuinput)
			switch menuinput {
			case 1:
				keluar = false
			case 0:
				keluar = true
			default:
				fmt.Println("Menu tidak valid, otomatis kembali")
			}
		case 5:
			var nim int
			fmt.Print("Masukkan NIM (Nomor Induk Mahasiswa) yang ingin ditampilkan transkripnya: ")
			fmt.Scan(&nim)
			tampilkanTranskrip(mhsw, nMahasiswa, nim)
			fmt.Println("+--------------------------+")
			fmt.Println("|           MENU           |")
			fmt.Println("+--------------------------+")
			fmt.Println("| 1. Kembali               |")
			fmt.Println("| 0. Kembali ke main menu  |")
			fmt.Println("+--------------------------+")
			fmt.Print("Pilih menu: [1/2]? ")
			fmt.Scan(&menuinput)
			switch menuinput {
			case 1:
				keluar = false
			case 2:
				keluar = true
			default:
				fmt.Println("Menu tidak valid, otomatis kembali")
			}

		case 0:
			keluar = true
		default:
			fmt.Println("Menu tidak valid, silakan input ulang")
		}
	}
}

func tampilkanDataMahasiswa(A tabMahasiswa, N int) {
	if N > 0 {
		fmt.Println("+-------------------------------------------------------------------------------------------------------------+")
		fmt.Println("|                                               DAFTAR MAHASISWA                                              |")
		fmt.Println("+-------------------------------------------------------------------------------------------------------------+")
		fmt.Printf("| %-2s | %-16s | %-18s | %-18s | %-14s | %-4s | %-3s | %-5s |\n", "No", "Nama", "NIM", "Jurusan", "Kelas", "IPS", "Total SKS", "Grade")
		fmt.Println("--------------------------------------------------------------------------------------------------------------+")
		for i := 0; i < N; i++ {
			fmt.Printf("| %-2d | %-16s | %-18d | %-18s | %-14s | %-4.2f | %-9d | %-5s |\n", i+1, A[i].nama, A[i].nim, A[i].jurusan, A[i].kelas, A[i].listIPS, A[i].listTotalSKS, A[i].IPgrade)
		}
		fmt.Println("--------------------------------------------------------------------------------------------------------------+")
		fmt.Println("Data berhasil ditampilkan.")
	} else {
		fmt.Println("Data mahasiswa tidak ada.")
	}
}

func tampilkanDataMatkul(A tabMataKuliah) {
	fmt.Println("+---------------------------------------------------------+")
	fmt.Println("|                   DAFTAR MATA KULIAH                    |")
	fmt.Println("+---------------------------------------------------------+")
	fmt.Printf("| %-2s | %-11s | %-30s | %-3s |\n", "No", "Kode MK", "Mata Kuliah", "SKS")
	fmt.Println("+---------------------------------------------------------+")
	for i := 0; i < 10; i++ {
		if A[i].KodeMK != "" {
			fmt.Printf("| %-2d | %-11s | %-30s | %-3d |\n", i+1, A[i].KodeMK, A[i].NamaMK, A[i].SKS)
		}
	}
	fmt.Println("+---------------------------------------------------------+")
}

func tampilkanDataMatakuliahDiambilMahasiswa(A tabMahasiswa, N int, nim int) {
	idx := searchingMahasiswa(A, N, nim)
	if idx != -1 {
		fmt.Println("\nMata Kuliah yang diambil oleh")
		fmt.Println("Nama\t:", A[idx].nama)
		fmt.Println("NIM\t:", A[idx].nim)
		fmt.Println("Jurusan\t:", A[idx].jurusan)
		fmt.Println("Kelas\t:", A[idx].kelas)

		fmt.Println("+-----------------------------------------------------------------------------------------------+")
		fmt.Printf("| %-2s | %-30s | %-6s | %-6s | %-6s | %-6s | %-5s | %-3s | %-5s |\n", "No", "Mata Kuliah", "UTS", "UAS", "Quiz", "Total", "IP", "SKS", "Grade")
		fmt.Println("+-----------------------------------------------------------------------------------------------+")

		for i := 0; i < A[idx].Nmatkul; i++ {
			fmt.Printf("| %-2d | %-30s | %-6.2f | %-6.2f | %-6.2f | %-6.2f | %-5.2f | %-3d | %-5s |\n", i+1, A[idx].listMataKuliah[i].NamaMK, A[idx].listUts[i], A[idx].listUas[i], A[idx].listQuiz[i], A[idx].listTotal[i], A[idx].IPMatkul[i], A[idx].listMataKuliah[i].SKS, A[idx].GradeMatkul[i])
		}

		fmt.Println("+-----------------------------------------------------------------------------------------------+")
		fmt.Println("Data berhasil ditampilkan.")
	} else {
		fmt.Println("+-----------------------------------------------------------------------------------------------+")
		fmt.Println("| Mahasiswa tidak ditemukan.                                                                    |")
		fmt.Println("+-----------------------------------------------------------------------------------------------+")
	}
}

func tampilkanMahasiswaMengambilMatakuliah(A tabMahasiswa, B tabMataKuliah, N int, no int) {
	fmt.Printf("\nMahasiswa yang mengambil mata kuliah:\n")
	fmt.Println("Kode MK			:", B[no-1].KodeMK)
	fmt.Println("Nama Mata Kuliah	:", B[no-1].NamaMK)
	fmt.Println("SKS			:", B[no-1].SKS)
	fmt.Println("+------------------------------------------------------------------------------------+")
	fmt.Println("| No |        Nama        |       NIM       |         Jurusan         |    Kelas     |")
	fmt.Println("+------------------------------------------------------------------------------------+")
	count := 1
	for i := 0; i < N; i++ {
		for j := 0; j < A[i].Nmatkul; j++ {
			if A[i].listMataKuliah[j].NamaMK == B[no-1].NamaMK {
				fmt.Printf("| %-2d | %-18s | %-15d | %-23s | %-12s |\n", count, A[i].nama, A[i].nim, A[i].jurusan, A[i].kelas)
				count++
			}
		}
	}
	if count != 1 {
		fmt.Println("+------------------------------------------------------------------------------------+")
		fmt.Println("Data berhasil ditampilkan.")
	} else {
		fmt.Println("| Tidak ada mahasiswa yang mengambil mata kuliah ini.                                |")
		fmt.Println("+------------------------------------------------------------------------------------+")
	}
}

func urutNilai(A *tabMahasiswa, N int) { // insertion
	var urut int
	var pass, i int
	var temp mahasiswa
	pass = 1
	fmt.Println("+--------------------------------------+")
	fmt.Println("|           MENU SORTING NILAI         |")
	fmt.Println("+--------------------------------------+")
	fmt.Println("| 1. Menaik                            |")
	fmt.Println("| 2. Menurun                           |")
	fmt.Println("+--------------------------------------+")
	fmt.Print("Pilih menu: [1/2]? ")
	fmt.Scan(&urut)
	if urut == 1 {
		for pass < N {
			i = pass
			temp = A[pass]
			for i > 0 && temp.listIPS < A[i-1].listIPS {
				A[i] = A[i-1]
				i--
			}
			A[i] = temp
			pass++
		}
	} else if urut == 2 {
		for pass < N {
			i = pass
			temp = A[pass]
			for i > 0 && temp.listIPS > A[i-1].listIPS {
				A[i] = A[i-1]
				i--
			}
			A[i] = temp
			pass++
		}
	}

}

func urutSKS(A *tabMahasiswa, matkul tabMataKuliah, N int) { //selection sort
	var urut int
	var pass, i, idx int
	var temp mahasiswa
	fmt.Println("+------------------------------------+")
	fmt.Println("|           MENU SORTING SKS         |")
	fmt.Println("+------------------------------------+")
	fmt.Println("| 1. Menaik                          |")
	fmt.Println("| 2. Menurun                         |")
	fmt.Println("+------------------------------------+")
	fmt.Print("Pilih menu: [1/2]? ")
	fmt.Scan(&urut)
	pass = 1
	if urut == 1 {
		for pass < N {
			idx = pass - 1
			i = pass
			for i < N {
				if A[idx].listTotalSKS > A[i].listTotalSKS {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
	} else if urut == 2 {
		for pass < N {
			idx = pass - 1
			i = pass
			for i < N {
				if A[idx].listTotalSKS < A[i].listTotalSKS {
					idx = i
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		}
	}
}

func tampilkanTranskrip(A tabMahasiswa, N int, nim int) {
	idx := searchingMahasiswa(A, N, nim)
	if idx != -1 {
		fmt.Println("\nTranskrip Mahasiswa milik:")
		fmt.Println("Nama\t:", A[idx].nama)
		fmt.Println("NIM\t:", A[idx].nim)
		fmt.Println("Jurusan\t:", A[idx].jurusan)
		fmt.Println("Kelas\t:", A[idx].kelas)
		fmt.Println("+------------------------------------------------------------------------------------------------------------------+")
		fmt.Printf("| %-2s | %-15s | %-30s | %-3s | %-6s | %-6s | %-6s | %-6s | %-5s | %-5s |\n", "No", "Kode Mata Kuliah", "Mata Kuliah", "SKS", "UTS", "UAS", "Quiz", "Total", "IP", "Grade")
		fmt.Println("-------------------------------------------------------------------------------------------------------------------|")

		for i := 0; i < A[idx].Nmatkul; i++ {
			fmt.Printf("| %-2d | %-15s  | %-30s | %-3d | %-6.2f | %-6.2f | %-6.2f | %-6.2f | %-5.2f | %-5s |\n", i+1, A[idx].listMataKuliah[i].KodeMK, A[idx].listMataKuliah[i].NamaMK, A[idx].listMataKuliah[i].SKS, A[idx].listUts[i], A[idx].listUas[i], A[idx].listQuiz[i], A[idx].listTotal[i], A[idx].IPMatkul[i], A[idx].GradeMatkul[i])
		}

		fmt.Println("+------------------------------------------------------------------------------------------------------------------+")
		fmt.Printf("|         Jumlah SKS terselesaikan semester ini          | %-3d |                                                   |\n", A[idx].listTotalSKS)
		fmt.Println("+------------------------------------------------------------------------------------------------------------------+")
		fmt.Printf("|                          IPS                           |                           %.2f                          |\n", A[idx].listIPS)
		fmt.Println("+------------------------------------------------------------------------------------------------------------------+")
		fmt.Println("Data berhasil ditampilkan.")
	} else {
		fmt.Println("+------------------------------------------------------------------------------------+")
		fmt.Println("| Mahasiswa tidak ditemukan.                                                         |")
		fmt.Println("+------------------------------------------------------------------------------------+")
	}
}
