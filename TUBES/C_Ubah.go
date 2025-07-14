package main

import "fmt"

func ubahData(A *tabMahasiswa, N *int, matkul tabMataKuliah) {
	var keluar bool = false
	var menuinput int
	var nim int
	var idxUbah int
	fmt.Print("Masukkan NIM (Nomor Induk Mahasiswa) yang ingin diubah datanya: ")
	fmt.Scan(&nim)
	idxUbah = searchingMahasiswa(*A, *N, nim)

	if idxUbah != -1 {
		fmt.Println("----------------------------------------------")
		fmt.Println("Nama\t:", A[idxUbah].nama)
		fmt.Println("Kelas\t:", A[idxUbah].kelas)
		fmt.Println("NIM\t:", A[idxUbah].nim)
		fmt.Println("+--------------------------------------------+")
		fmt.Println("|            MENU UBAH/HAPUS DATA            |")
		fmt.Println("+--------------------------------------------+")
		fmt.Println("| 1. Ubah/hapus data mahasiswa               |")
		fmt.Println("| 2. Ubah/hapus data mata kuliah             |")
		fmt.Println("| 0. Kembali                                 |")
		fmt.Println("+--------------------------------------------+")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menuinput)
		for !keluar {
			switch menuinput {
			case 1:
				fmt.Println("+------------------------------------------------+")
				fmt.Println("|            MENU UBAH DATA MAHASISWA            |")
				fmt.Println("+------------------------------------------------+")
				fmt.Println("| 1. Edit data mahasiswa                         |")
				fmt.Println("| 2. Hapus data mahasiswa                        |")
				fmt.Println("| 0. Kembali                                     |")
				fmt.Println("+------------------------------------------------+")
				fmt.Print("Pilih menu: ")
				fmt.Scan(&menuinput)
				switch menuinput {
				case 1:
					editDataMahasiswa(A, *N, idxUbah)
				case 2:
					hapusDataMahasiswa(A, N, idxUbah)
				case 0:
					keluar = true
				default:
					fmt.Println("Menu tidak valid.")
				}
			case 2:
				fmt.Println("+------------------------------------------------------------+")
				fmt.Println("|            MENU UBAH DATA MATA KULIAH MAHASISWA            |")
				fmt.Println("+------------------------------------------------------------+")
				fmt.Println("| 1. Edit data mata kuliah                                   |")
				fmt.Println("| 2. Hapus data mata kuliah                                  |")
				fmt.Println("| 0. Kembali                                                 |")
				fmt.Println("+------------------------------------------------------------+")
				fmt.Print("Pilih menu: ")
				fmt.Scan(&menuinput)
				switch menuinput {
				case 1:
					editMatkul(A, *N, idxUbah, nim)
				case 2:
					hapusMataKuliah(A, *N, idxUbah, nim)
				case 0:
					keluar = true
				default:
					fmt.Println("Menu tidak valid.")
				}
			case 0:
				keluar = true
			default:
				fmt.Println("Menu tidak valid, silakan input ulang.")
				keluar = false
			}
		}
	} else {
		fmt.Println("+--------------------------+")
		fmt.Println("|           MENU           |")
		fmt.Println("+--------------------------+")
		fmt.Println("| 1. Masukkan NIM ulang    |")
		fmt.Println("| 0. Kembali ke main menu  |")
		fmt.Println("+--------------------------+")
		fmt.Print("Pilih menu: [1/0]? ")
		fmt.Scan(&menuinput)
		switch menuinput {
		case 1:
			ubahData(A, N, matkul)
		case 0:
			keluar = true
		default:
			fmt.Println("Menu tidak valid, otomatis kembali")
		}
	}

}

func editDataMahasiswa(A *tabMahasiswa, N int, idx int) {
	var menuinput int
	var keluar bool = false
	for !keluar {
		fmt.Println("+------------------------------------------------+")
		fmt.Println("|            MENU EDIT DATA MAHASISWA            |")
		fmt.Println("+------------------------------------------------+")
		fmt.Println("| 1. Edit Nama Mahasiswa                         |")
		fmt.Println("| 2. Edit NIM Mahasiswa                          |")
		fmt.Println("| 3. Edit Jurusan Mahasiswa                      |")
		fmt.Println("| 4. Edit Kelas Mahasiswa                        |")
		fmt.Println("| 0. Kembali                                     |")
		fmt.Println("+------------------------------------------------+")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menuinput)
		switch menuinput {
		case 1:
			fmt.Println("Nama sebelumnya:", A[idx].nama)
			fmt.Print("Masukkan Nama baru (pisahkan dengan '_'): ")
			fmt.Scan(&A[idx].nama)
			fmt.Println("Data berhasil diubah.")
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
			var nim int
			fmt.Println("NIM sebelumnya:", A[idx].nim)
			fmt.Print("Masukkan NIM baru: ")
			fmt.Scan(&nim)
			for !keluar {
				if searchingMahasiswa(*A, N, nim) == -1 {
					keluar = true
					A[idx].nim = nim
				} else {
					fmt.Printf("Mahasiswa dengan NIM %d sudah ada, masukkan nim ulang.\n", nim)
					fmt.Print("Masukkan NIM: ")
					fmt.Scan(&nim)
				}
			}
			fmt.Println("Data berhasil diubah.")
			fmt.Println("+--------------------------+")
			fmt.Println("|           MENU           |")
			fmt.Println("+--------------------------+")
			fmt.Println("| 1. Kembali               |")
			fmt.Println("| 2. Kembali ke main menu  |")
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
			fmt.Println("Jurusan sebelumnya:", A[idx].jurusan)
			fmt.Print("Masukkan Jurusan baru: ")
			fmt.Scan(&A[idx].jurusan)
			fmt.Println("Data berhasil diubah.")
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
			fmt.Println("Kelas sebelumnya:", A[idx].kelas)
			fmt.Print("Masukkan Kelas baru: ")
			fmt.Scan(&A[idx].kelas)
			fmt.Println("Data berhasil diubah.")
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
		case 0:
			keluar = true
		default:
			fmt.Println("Menu tidak valid, silakan input ulang.")
		}
	}
}
func editMatkul(A *tabMahasiswa, N, idx, nim int) {
	var i int
	var uas, uts, quiz float64
	tampilkanDataMatakuliahDiambilMahasiswa(*A, N, nim)
	if idx != -1 {
		fmt.Printf("Pilih mata kuliah yang ingin diubah datanya: [1/../%d] ", A[idx].Nmatkul)
		fmt.Scan(&i)
		i--
		var menuinput int
		var keluar bool = false
		for !keluar {
			fmt.Println("+--------------------------------------------------+")
			fmt.Println("|            MENU EDIT DATA MATA KULIAH            |")
			fmt.Println("+--------------------------------------------------+")
			fmt.Println("| 1. Edit Nilai UTS Mahasiswa                      |")
			fmt.Println("| 2. Edit Nilai UAS Mahasiswa                      |")
			fmt.Println("| 3. Edit Nilai Quiz Mahasiswa                     |")
			fmt.Println("| 0. Kembali                                       |")
			fmt.Println("+--------------------------------------------------+")
			fmt.Print("Pilih menu: [1/2/3/0]? ")
			fmt.Scan(&menuinput)
			switch menuinput {
			case 1:
				fmt.Print("Masukkan nilai UTS baru: ")
				fmt.Scan(&uts)
				keluar = false
				for !keluar {
					if uts > 100 || uts < 0 {
						fmt.Println("Nilai UTS tidak valid, masukkan nilai UTS ulang.")
						fmt.Print("Masukkan nilai UTS baru: ")
						fmt.Scan(&uts)
					} else {
						A[idx].listUts[i] = uts
						keluar = true
					}
				}
				fmt.Println("Data berhasil diubah.")
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
				fmt.Print("Masukkan nilai UAS baru: ")
				fmt.Scan(&uas)
				keluar = false
				for !keluar {
					if uas > 100 || uas < 0 {
						fmt.Println("Nilai UAS tidak valid, masukkan nilai UAS ulang.")
						fmt.Print("Masukkan Nilai UAS baru: ")
						fmt.Scan(&uas)
					} else {
						A[idx].listUas[i] = uas
						keluar = true
					}
				}
				fmt.Println("Data berhasil diubah.")
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
				fmt.Print("Masukkan Nilai Quiz baru: ")
				fmt.Scan(&quiz)
				keluar = false
				for !keluar {
					if quiz > 100 || quiz < 0 {
						fmt.Println("Nilai Quiz tidak valid, masukkan nilai Quiz ulang.")
						fmt.Print("Masukkan Nilai Quiz baru: ")
						fmt.Scan(&quiz)
					} else {
						A[idx].listQuiz[i] = quiz
						keluar = true
					}
				}
				fmt.Println("Data berhasil diubah.")
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
			case 0:
				keluar = true
			default:
				fmt.Println("Menu tidak valid, silakan input ulang.")
			}
		}
		A[idx].listTotal[i] = (A[idx].listUts[i] * 40 / 100) + (A[idx].listUas[i] * 40 / 100) + (A[idx].listQuiz[i] * 20 / 100)
		A[idx].GradeMatkul[i] = hitungMatkulGrade(A[idx].listTotal[i])
		A[idx].IPMatkul[i] = hitungIPMatkul(A[idx].listTotal[i])
		totalSKS(A, idx)
		nilaiAkhir(A, idx)
	}

}

func hapusDataMahasiswa(A *tabMahasiswa, N *int, idxHapus int) {
	if idxHapus != -1 {
		for i := idxHapus; i < *N-1; i++ {
			A[i], A[i+1] = A[i+1], A[i]
		}
		*N--
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Mahasiswa tidak ditemukan.")
	}
}
func hapusMataKuliah(A *tabMahasiswa, N int, idx int, nim int) {
	var idxMatkul int
	tampilkanDataMatakuliahDiambilMahasiswa(*A, N, nim)
	fmt.Print("Pilih mata kuliah yang ingin dihapus [1/2/3../10]: ")
	fmt.Scan(&idxMatkul)
	idxMatkul--
	for i := idxMatkul; i < A[idx].Nmatkul-1; i++ {
		//proses penghapusan mata kuliah, dengan melakukan penukaran index i dengan i+1
		A[idx].listMataKuliah[i], A[idx].listMataKuliah[i+1] = A[idx].listMataKuliah[i+1], A[idx].listMataKuliah[i]
		A[idx].listUts[i], A[idx].listUts[i+1] = A[idx].listUts[i+1], A[idx].listUts[i]
		A[idx].listUas[i], A[idx].listUas[i+1] = A[idx].listUas[i+1], A[idx].listUas[i]
		A[idx].listQuiz[i], A[idx].listQuiz[i+1] = A[idx].listQuiz[i+1], A[idx].listQuiz[i]
		A[idx].listTotal[i], A[idx].listTotal[i+1] = A[idx].listTotal[i+1], A[idx].listTotal[i]
		A[idx].GradeMatkul[i], A[idx].GradeMatkul[i+1] = A[idx].GradeMatkul[i+1], A[idx].GradeMatkul[i]

		A[idx].listTotal[i] = (A[idx].listUts[i] * 40 / 100) + (A[idx].listUas[i] * 40 / 100) + (A[idx].listQuiz[i] * 20 / 100)
		A[idx].GradeMatkul[i] = hitungMatkulGrade(A[idx].listTotal[i])
		A[idx].IPMatkul[i] = hitungIPMatkul(A[idx].listTotal[i])
		totalSKS(A, idx)
		nilaiAkhir(A, idx)
	}
	A[idx].Nmatkul--
	fmt.Println("Data berhasil dihapus.")
}
