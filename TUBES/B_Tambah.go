package main

import "fmt"

func tambahDataMahasiswa(A *tabMahasiswa, N *int, matkul tabMataKuliah) {
	var idx int
	var keluar bool = false
	var nim int
	var uts, uas, quiz float64
	idx = *N

	fmt.Println("+------------------------------------------------+")
	fmt.Println("|             TAMBAH DATA MAHASISWA              |")
	fmt.Println("+------------------------------------------------+")

	fmt.Print("Masukkan Nama mahasiswa (pisahkan dengan '_', contoh: Alif_Ihsan): ")
	fmt.Scan(&A[idx].nama)
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	for !keluar {
		if searchingMahasiswa(*A, *N, nim) == -1 {
			keluar = true
			A[idx].nim = nim
		} else {
			fmt.Printf("Mahasiswa dengan NIM %d sudah ada, masukkan nim ulang.\n", nim)
			fmt.Print("Masukkan NIM: ")
			fmt.Scan(&nim)
		}
	}
	fmt.Print("Masukkan Jurusan (pisahkan dengan '_', contoh: S-1_Informatika): ")
	fmt.Scan(&A[idx].jurusan)

	fmt.Print("Masukkan Kelas (contoh: IF-47-04): ")
	fmt.Scan(&A[idx].kelas)

	fmt.Println("Data Mahasiswa berhasil diinput.")

	fmt.Println("+--------------------------------------------------+")
	fmt.Println("|             TAMBAH DATA MATA KULIAH              |")
	fmt.Println("+--------------------------------------------------+")

	// Sebagai contoh, mahasiswa hanya bisa mengambil maksimal 8 mata kuliah
	fmt.Print("Masukkan berapa banyak mata kuliah yang ingin diambil (maksimal 8): ")
	fmt.Scan(&A[idx].Nmatkul)

	// Jika inputan Nmatkul (banyak mata kuliah yang diambil) lebih dari 8, maka akan otomatis Nmatkulnya menjadi 8
	if A[idx].Nmatkul > 8 {
		fmt.Println("Mata kuliah yang diambil harus tidak lebih dari 8")
		A[idx].Nmatkul = 8
	}

	fmt.Println("Banyak mata kuliah yang diambil:", A[idx].Nmatkul)

	for i := 0; i < A[idx].Nmatkul; i++ {
		var matkulIndex int
		tampilkanDataMatkul(matkul)
		fmt.Print("Pilih mata kuliah yang ingin diambil [1/2/3/.../10]: ")
		fmt.Scan(&matkulIndex)
		keluar = false
		for !keluar {
			if matkulIndex >= 1 && matkulIndex <= 10 {
				keluar = true
			} else {
				tampilkanDataMatkul(matkul)
				fmt.Println("Mata kuliah tidak valid, silakan ambil ulang.")
				fmt.Print("Pilih mata kuliah yang ingin diambil [1/2/3/.../10]: ")
				fmt.Scan(&matkulIndex)
			}
		}
		matkulIndex -= 1
		if cekMatkul(*A, matkul, idx, matkulIndex) != 1 {
			A[idx].listMataKuliah[i] = matkul[matkulIndex]
			fmt.Print("Masukkan Nilai UTS: ")
			fmt.Scan(&uts)
			keluar = false
			for !keluar {
				if uts > 100 || uts < 0 {
					fmt.Println("Nilai UTS tidak valid, masukkan nilai UTS ulang.")
					fmt.Print("Masukkan Nilai UTS: ")
					fmt.Scan(&uts)
				} else {
					A[idx].listUts[i] = uts
					keluar = true
				}
			}
			fmt.Print("Masukkan Nilai UAS: ")
			fmt.Scan(&uas)
			keluar = false
			for !keluar {
				if uas > 100 || uas < 0 {
					fmt.Println("Nilai UAS tidak valid, masukkan nilai UAS ulang.")
					fmt.Print("Masukkan Nilai UAS: ")
					fmt.Scan(&uas)
				} else {
					A[idx].listUas[i] = uas
					keluar = true
				}
			}
			fmt.Print("Masukkan Nilai Quiz: ")
			fmt.Scan(&quiz)
			keluar = false
			for !keluar {
				if quiz > 100 || quiz < 0 {
					fmt.Println("Nilai Quiz tidak valid, masukkan nilai Quiz ulang.")
					fmt.Print("Masukkan Nilai Quiz: ")
					fmt.Scan(&quiz)
				} else {
					A[idx].listQuiz[i] = quiz
					keluar = true
				}
			}
			A[idx].listTotal[i] = (A[idx].listUts[i] * 40 / 100) + (A[idx].listUas[i] * 40 / 100) + (A[idx].listQuiz[i] * 20 / 100)
			A[idx].GradeMatkul[i] = hitungMatkulGrade(A[idx].listTotal[i])
			A[idx].IPMatkul[i] = hitungIPMatkul(A[idx].listTotal[i])
			totalSKS(A, idx)
			nilaiAkhir(A, idx)
		} else {
			fmt.Println("Mata kuliah sudah pernah diambil, silakan pilih mata kuliah lain.")
			i--
		}
	}
	fmt.Println("Data Mahasiswa berhasil diinput, kembali otomatis ke main menu.")
	*N++
}

func tambahMatakuliah(A *tabMahasiswa, N int, matkul tabMataKuliah) {
	var nim int
	var matkulIndex int
	var ulang bool = false
	var keluar bool
	var uts, uas, quiz float64
	fmt.Print("Masukkan NIM (Nomor Induk Mahasiswa) yang ingin ditambah mata kuliahnya: ")
	fmt.Scan(&nim)
	idx := searchingMahasiswa(*A, N, nim)
	if idx != -1 {
		if A[idx].Nmatkul == 8 { //mahasiswa yang sudah mengambil 8 mata kuliah tidak dapat mengambil mata kuliah lagi.
			fmt.Println("Tidak bisa menambah mata kuliah lagi.")
			fmt.Println("Mahasiswa sudah mengambil 8 mata kuliah.")
		} else { // mahasiswa yang mengambil kurang dari 8 mata kuliah bisa mengambil mata kuliah lagi.
			A[idx].Nmatkul++
			if A[idx].Nmatkul > 8 {
				fmt.Println("Mata kuliah yang diinput harus tidak lebih dari 8")
				A[idx].Nmatkul = 8
			}
			idxMatkul := A[idx].Nmatkul - 1
			tampilkanDataMatkul(matkul)
			fmt.Print("Pilih mata kuliah yang ingin diambil [1/2/3/.../10]: ")
			fmt.Scan(&matkulIndex)
			keluar = false
			for !keluar {
				if matkulIndex >= 1 && matkulIndex <= 10 {
					keluar = true
				} else {
					tampilkanDataMatkul(matkul)
					fmt.Println("Mata kuliah tidak valid, silakan ambil ulang.")
					fmt.Print("Pilih mata kuliah yang ingin diambil [1/2/3/.../10]: ")
					fmt.Scan(&matkulIndex)
				}
			}
			matkulIndex -= 1
			for !ulang {
				if cekMatkul(*A, matkul, idx, matkulIndex) != 1 {
					A[idx].listMataKuliah[idxMatkul] = matkul[matkulIndex]
					fmt.Print("Masukkan Nilai UTS: ")
					fmt.Scan(&uts)
					keluar = false
					for !keluar {
						if uts > 100 || uts < 0 {
							fmt.Println("Nilai UTS tidak valid, masukkan nilai UTS ulang.")
							fmt.Print("Masukkan Nilai UTS: ")
							fmt.Scan(&uts)
						} else {
							A[idx].listUts[idxMatkul] = uts
							keluar = true
						}
					}
					fmt.Print("Masukkan Nilai UAS: ")
					fmt.Scan(&uas)
					keluar = false
					for !keluar {
						if uas > 100 || uas < 0 {
							fmt.Println("Nilai UAS tidak valid, masukkan nilai UAS ulang.")
							fmt.Print("Masukkan Nilai UAS: ")
							fmt.Scan(&uas)
						} else {
							A[idx].listUas[idxMatkul] = uas
							keluar = true
						}
					}
					fmt.Print("Masukkan Nilai Quiz: ")
					fmt.Scan(&quiz)
					keluar = false
					for !keluar {
						if quiz > 100 || quiz < 0 {
							fmt.Println("Nilai Quiz tidak valid, masukkan nilai Quiz ulang.")
							fmt.Print("Masukkan Nilai Quiz: ")
							fmt.Scan(&quiz)
						} else {
							A[idx].listQuiz[idxMatkul] = quiz
							keluar = true
						}
					}
					A[idx].listTotal[idxMatkul] = (A[idx].listUts[idxMatkul] * 40 / 100) + (A[idx].listUas[idxMatkul] * 40 / 100) + (A[idx].listQuiz[idxMatkul] * 20 / 100)
					A[idx].GradeMatkul[idxMatkul] = hitungMatkulGrade(A[idx].listTotal[idxMatkul])
					A[idx].IPMatkul[idxMatkul] = hitungIPMatkul(A[idx].listTotal[idxMatkul])
					totalSKS(A, idx)
					nilaiAkhir(A, idx)
					ulang = true
				} else {
					fmt.Println("Mata kuliah sudah pernah diambil, silakan pilih mata kuliah lain.")
					fmt.Print("Pilih nomor mata kuliah yang ingin diambil: ")
					fmt.Scan(&matkulIndex)
					matkulIndex -= 1
				}
			}
		}
	} else {
		var menuinput int
		fmt.Println("Mahasiswa tidak ditemukan.")
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
			tambahMatakuliah(A, N, matkul)
		case 0:
			keluar = true
		default:
			fmt.Println("Menu tidak valid, otomatis kembali")
		}
	}
}
