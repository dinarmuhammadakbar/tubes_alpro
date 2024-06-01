package main

import (
	"fmt"
	"time"
)

const NMAX int = 10

type Akun struct {
	nama, email, password string
	noRek                 int
	saldo                 int
	countTransaksi        int
	transaksi             [NMAX]transaksi
}

type transaksi struct {
	saldo, pemasukan, pengeluaran int
	tipe                          string
	waktu                         time.Time
}

type tabAkun [NMAX]Akun

func main() {
	var akun tabAkun
	var idAkun int

	idAkun = 0

	for {
		fmt.Println("______________________________________________")
		fmt.Println("****  TUGAS BESAR ALGORITMA PEMROGRAMAN   ****")
		fmt.Println("****         APLIKASI E- MONEY            ****")
		fmt.Println("______________________________________________")
		fmt.Println()
		fmt.Println("Menu Utama\n1.Registrasi\n2.Login\n3.Admin\n4.Keluar")
		fmt.Print("Pilih (1/2/3/4): ")

		var pilihan int
		fmt.Scanf("%d\n",&pilihan)

		if pilihan == 1 {
			registrasi(&akun, &idAkun)
		} else if pilihan == 2 {
			login(&akun)
		} else if pilihan == 3 {
			admin(&akun, idAkun)
		} else if pilihan == 4 {
			fmt.Println()
			fmt.Println("Terima kasih telah mengunjungi aplikasi ini!")
			return
		}
	}
}
func registrasi(A *tabAkun, idAkun *int) {
	// IS : Terdefinisi array A dan idAkun bilangan bulat
	// FS : menampilkan kalimat berupa registrasi berhasil dan jika gagal maka user diminta memasukan data kembali
	var password string
	var nama string
	var email string
	var norek, pilih int

	fmt.Println()
	fmt.Println("______________________________________________")
	fmt.Println("Registrasi")
	fmt.Println()

	if *idAkun >= NMAX {
		fmt.Println("Kapasitas akun sudah penuh")
		return
	}

	fmt.Print("Nama       : ")
	fmt.Scanln(&nama)
	fmt.Print("Email      : ")
	fmt.Scanln(&email)
	fmt.Print("Password   : ")
	fmt.Scanln(&password)
	fmt.Println()

	if nama == "" || email == ""{
		fmt.Println("Masukan inputan tidak boleh kosong!")
		fmt.Println()
		return
	} else if len(password) < 8 {
		fmt.Println("Password minimal 8 karakter")
		fmt.Println()
		return
	} 

	fmt.Println("1.Simpan  2.Batal")
	fmt.Print("Pilih (1/2) : ")
	fmt.Scanf("%d\n",&pilih )
	fmt.Println()

	if pilih == 1{
		norek = 51458094 + *idAkun
		fmt.Println("______________________________________________")
		fmt.Println()
		fmt.Println("Registrasi berhasil")
		fmt.Printf("Nomor rekening : %d\n", norek)
		A[*idAkun].noRek = norek
		A[*idAkun].nama = nama
		A[*idAkun].email = email
		A[*idAkun].password = password
		*idAkun++
	
		login(A)
	}else if pilih == 2{
		return
	}
}
func login(A *tabAkun) {
	// I.S : Terdefinisi array A
	// F.S : menampilkan tampilan login
	var noRek int
	var password string
	var i int
	fmt.Println("______________________________________________")
	fmt.Println("Login")
	fmt.Println()
	fmt.Print("NoRek       : ")
	fmt.Scanln(&noRek)
	fmt.Print("Password    : ")
	fmt.Scanln(&password)

	for i < NMAX {
		if noRek == A[i].noRek && password == A[i].password {
			fmt.Println()
			fmt.Println("Login berhasil!")
			home(A, i)
			return
		}
		i++
	}
	fmt.Println()
	fmt.Println("NoRek atau password salah.")
	fmt.Println()
}
func admin(A *tabAkun, idAkun int) {
	// I.S. Terdefinisi array A dan idAkun bilangan bulat
	// F.S. Menampilkan tampilan admin terdapat beberapa pilihan jika memilih angka 1 akan ke function cetak akun,
	//      jika memilih angka 2 akan ke function hapus akun, jika memilih angka 3 maka keluar, jika memilih selain
	//      angka 1,2 dan 3 akan menampilkan tampilan Masukan tidak valid.
	var pilih int

	fmt.Println("______________________________________________")
	fmt.Println("> > >             A D M I N              < < <")
	fmt.Println("______________________________________________")
	fmt.Println()
	fmt.Println("Pilih Menu : ")
	fmt.Println("1. Cetak Akun")
	fmt.Println("2. Hapus Akun")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih (1/2/3) : ")
	fmt.Scanf("%d\n", &pilih)
	if pilih == 1 {
		cetakAkun(*A, idAkun)
	} else if pilih == 2 {
		hapusAkun(A, &idAkun)
	}
	fmt.Println()
}
func cetakAkun(A tabAkun, idAkun int) {
	// I.S. Terdefinisi array A dan idAkun bilangan bulat
	// F.S. Menampilkan tampilan cetak akun
	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("              C E T A K   A K U N             ")
	fmt.Println("==============================================")
	fmt.Println()
	if idAkun == 0 {
		fmt.Println("Belum ada akun")
	} else {
		fmt.Printf("%-19s  %-15s  %-15s %-12s \n", "Nama Akun", "Nomor Rekening", "Password", "Saldo")
		for i := 0; i < idAkun; i++ {
			if A[i].nama != "" {
				fmt.Printf("%-19s  %-15d  %-15s %12d\n", A[i].nama, A[i].noRek, A[i].password, A[i].saldo)
			}
		}
	}
}
func hapusAkun(A *tabAkun, idAkun *int) {
	// I.S. Terdefinisi array A dan idAkun bilangan bulat
	// F.S. Menampilkan tampilan hapus akun
	var input, i int
	var konfirmasi string

	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("              H A P U S    A K U N            ")
	fmt.Println("==============================================")
	fmt.Println()
	fmt.Print("Masukan nomor rekening    : ")
	fmt.Scanf("%d\n", &input)

	ketemu := false
	for i < *idAkun && !ketemu {
		if input == (*A)[i].noRek {
			ketemu = true
			fmt.Printf("Yakin hapus akun %s? (y/n) : ", (*A)[i].nama)
			fmt.Scanf("%s\n", &konfirmasi)
			if konfirmasi == "y" || konfirmasi == "Y" {
				fmt.Println()
				fmt.Printf("Akun %s berhasil dihapus.\n", (*A)[i].nama)
				for j := i; j < *idAkun-1; j++ {
					(*A)[j] = (*A)[j+1]
				}
				(*A)[*idAkun-1] = Akun{}
				*idAkun--
			} else if konfirmasi == "n" || konfirmasi == "N" {
				fmt.Println("Penghapusan akun dibatalkan.")
			}
		}
		i++
	}

	if !ketemu {
		fmt.Println()
		fmt.Println("Nomor rekening tidak ditemukan.")
	}
}
func home(A *tabAkun, idAkun int) {
	//I.S : Terdefinisi array A Dan idAkun bilangan bulat
	//F.S : Menampilkan tampilan home
	for {
		fmt.Println("______________________________________________")
		fmt.Println("****          Aplikasi E-money            ****")
		fmt.Println("______________________________________________")
		fmt.Println()
		fmt.Print("Hai ")
		fmt.Println(A[idAkun].nama)
		fmt.Printf("Saldo Anda : Rp.%d,-\n", A[idAkun].saldo)
		fmt.Println()
		fmt.Println("Pilih menu: ")
		fmt.Println("1. Transfer")
		fmt.Println("2. Pembayaran")
		fmt.Println("3. Isi Saldo")
		fmt.Println("4. Riwayat")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih (1/2/3/4/5) : ")

		var pilih int
		fmt.Scanf("%d\n", &pilih)
		fmt.Println()

		if pilih == 1 {
			transfer(A, idAkun)
		} else if pilih == 2 {
			pembayaran(A, idAkun)
		} else if pilih == 3 {
			isiSaldo(A, idAkun)
		} else if pilih == 4 {
			riwayat(A, idAkun)
		} else if pilih == 5 {
			return
		} else {
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
func isiSaldo(A *tabAkun, idAkun int) {
	// I.S. Terdefinisi array A dan idAkun bilangan bulat
	// F.S. Menampilkan tampilan isi saldo
	var nominal int
	fmt.Println()
	fmt.Println("______________________________________________")
	fmt.Println("> > >          I S I   S A L D O         < < <")
	fmt.Println("______________________________________________")
	fmt.Println()
	fmt.Print("Masukan nominal : ")
	fmt.Scanf("%d\n", &nominal)
	A[idAkun].saldo = A[idAkun].saldo + nominal
	fmt.Println()
	fmt.Println("Saldo berhasil ditambahkan.")
	fmt.Println()
}
func transfer(A *tabAkun, idAkun int) {
	// I.S. Terdefinisi array A dan idAkun bilangan bulat
	// F.S. Menampilkan tampilan transfer
	var akunTujuan, password, tipePengirim, tipe string
	var nominal int
	var tujuanIndex int
	var waktuTransaksi time.Time

	tujuanIndex = -1

	fmt.Println()
	fmt.Println("______________________________________________")
	fmt.Println("> > >          T R A N S F E R           < < <")
	fmt.Println("______________________________________________")
	fmt.Println()

	fmt.Printf("Masukan akun tujuan  : ")
	fmt.Scanf("%s\n", &akunTujuan)
	
	fmt.Printf("Masukan nominal      : ")
	fmt.Scanf("%d\n", &nominal)
	
	fmt.Printf("Masukan password     : ")
	fmt.Scanf("%s\n", &password)
	

	if A[idAkun].password != password {
		fmt.Println("Password salah!")
		return
	}

	if A[idAkun].saldo < nominal {
		fmt.Println()
		fmt.Println("Saldo tidak mencukupi!")
		return
	}

	tujuanIndex = findAkun_SequentialSearch(*A, akunTujuan, idAkun)

	if tujuanIndex == -1 {
		fmt.Println()
		fmt.Println("Akun tujuan tidak ditemukan!")
		return
	}


	// Mendapatkan waktu saat ini
	waktuTransaksi = time.Now()

	// Update transaksi pengirim
	A[idAkun].transaksi[A[idAkun].countTransaksi].pengeluaran = nominal
	A[idAkun].transaksi[A[idAkun].countTransaksi].saldo = A[idAkun].saldo - nominal
	tipePengirim = "Transfer ke " + akunTujuan
	A[idAkun].transaksi[A[idAkun].countTransaksi].tipe = tipePengirim
	A[idAkun].transaksi[A[idAkun].countTransaksi].waktu = waktuTransaksi
	A[idAkun].countTransaksi++

	// Update saldo pengirim
	A[idAkun].saldo = A[idAkun].saldo - nominal

	// Update transaksi penerima
	A[tujuanIndex].transaksi[A[tujuanIndex].countTransaksi].pemasukan = nominal
	A[tujuanIndex].transaksi[A[tujuanIndex].countTransaksi].saldo = A[tujuanIndex].saldo + nominal
	tipe = "Transfer dari " + A[idAkun].nama
	A[tujuanIndex].transaksi[A[tujuanIndex].countTransaksi].tipe = tipe

	A[tujuanIndex].transaksi[A[tujuanIndex].countTransaksi].waktu = waktuTransaksi
	A[tujuanIndex].countTransaksi++

	// Update saldo penerima
	A[tujuanIndex].saldo = A[tujuanIndex].saldo + nominal

	fmt.Println()
	fmt.Println("Transfer berhasil!")
}
func pembayaran(A *tabAkun, idAkun int) {
	// I.S. Terdefinisi array A dan idAkun bilangan bulat
	// F.S. Menampilkan tampilan pembayaran
	var emailTujuan, password string
	var nominal, layanan int
	var tujuanIndex int
	var waktuTransaksi time.Time

	tujuanIndex = -1

	fmt.Println("______________________________________________")
	fmt.Println("> > >       P E M B A Y A R A N          < < <")
	fmt.Println("______________________________________________")
	fmt.Println()

	fmt.Println("Layanan :")
	fmt.Println("1. Listrik")
	fmt.Println("2. BPJS")
	fmt.Println("3. WIFI")
	fmt.Println("4. Makanan")
	fmt.Printf("Pilih layanan (1/2/3/4): ")
	fmt.Scanf("%d\n", &layanan)

	fmt.Println()
	fmt.Printf("Masukan akun tujuan : ")
	fmt.Scanf("%s\n", &emailTujuan)

	fmt.Printf("Masukan nominal     : ")
	fmt.Scanf("%d\n", &nominal)

	fmt.Printf("Masukan password    : ")
	fmt.Scanf("%s\n", &password)


	if A[idAkun].password != password {
		fmt.Println("Password salah!")
		return
	}

	if A[idAkun].saldo < nominal {
		fmt.Println("Saldo tidak mencukupi!")
		return
	}

	tujuanIndex = findAkun_SequentialSearch(*A, emailTujuan, idAkun)

	if tujuanIndex == -1 {
		fmt.Println("Akun tujuan tidak ditemukan!")
		return
	}

	// Mendapatkan waktu saat ini
	waktuTransaksi = time.Now()

	// Update transaksi pengirim
	A[idAkun].transaksi[A[idAkun].countTransaksi].pengeluaran = nominal
	A[idAkun].transaksi[A[idAkun].countTransaksi].saldo = A[idAkun].saldo - nominal
	if layanan == 1 {
		A[idAkun].transaksi[A[idAkun].countTransaksi].tipe = "Pembayaran Listrik"
	} else if layanan == 2 {
		A[idAkun].transaksi[A[idAkun].countTransaksi].tipe = "Pembayaran BPJS"
	} else if layanan == 3 {
		A[idAkun].transaksi[A[idAkun].countTransaksi].tipe = "Pembayaran WIFI"
	} else if layanan == 4 {
		A[idAkun].transaksi[A[idAkun].countTransaksi].tipe = "Pembayaran Makanan"
	}
	A[idAkun].transaksi[A[idAkun].countTransaksi].waktu = waktuTransaksi
	A[idAkun].countTransaksi++

	// Update saldo pengirim
	A[idAkun].saldo = A[idAkun].saldo - nominal

	// Update transaksi penerima
	A[tujuanIndex].transaksi[A[tujuanIndex].countTransaksi].pemasukan = nominal
	A[tujuanIndex].transaksi[A[tujuanIndex].countTransaksi].saldo = A[tujuanIndex].saldo + nominal
	A[tujuanIndex].transaksi[A[tujuanIndex].countTransaksi].tipe = "Penerima Pembayaran"
	A[tujuanIndex].transaksi[A[tujuanIndex].countTransaksi].waktu = waktuTransaksi
	A[tujuanIndex].countTransaksi++

	// Update saldo penerima
	A[tujuanIndex].saldo = A[tujuanIndex].saldo + nominal

	fmt.Println()
	fmt.Println("Transfer berhasil!")
}
func findAkun_SequentialSearch(A tabAkun, akunTujuan string, idAkun int) int {
	/* mengembalikan indek dari akun yang memiliki nama sama dengan akunTujuan,
	   dan jika tidak ditemukan, maka mengembalikan nilai index -1 */
	var i, idx int
	i = 0
	idx = -1
	for i < NMAX && idx == -1 {
		if A[i].nama == akunTujuan && i != idAkun {
			idx = i
		}
		i++ 
	}
	return idx
}
func riwayat(A *tabAkun, idAkun int) {
	// I.S. Terdefinisi array A dan idAkun bilangan bulat
	// F.S. Menampilkan tampilan riwayat transaksi
	fmt.Println()
	fmt.Println("______________________________________________")
	fmt.Println("> > >           R I W A Y A T            < < <")
	fmt.Println("______________________________________________")
	fmt.Println()

	if A[idAkun].countTransaksi == 0 {
		fmt.Println("Tidak ada riwayat transaksi.")
		return
	}

	fmt.Println("Riwayat Transaksi:")
	fmt.Println()
	fmt.Println("____________________________________________________________________________________________________")
	fmt.Printf("| %-2s | %-19s | %-22s | %-12s | %-12s | %-14s |\n", "No", "Waktu", "Tipe", "Pemasukan", "Pengeluaran", "Saldo")
	for i := 0; i < A[idAkun].countTransaksi; i++ {
		fmt.Printf("| %-2d | %-19s | %-22s | %12d | %12d | %14d |\n",
			i+1, A[idAkun].transaksi[i].waktu.Format("2006-01-02 15:04:05"), A[idAkun].transaksi[i].tipe, A[idAkun].transaksi[i].pemasukan, A[idAkun].transaksi[i].pengeluaran, A[idAkun].transaksi[i].saldo)
	}
	fmt.Println("____________________________________________________________________________________________________")

	var fitur string
	fmt.Println()
	fmt.Print("Pilih filter (y/n) ? ")
	fmt.Scanf("%s\n", &fitur)
	if fitur == "y" || fitur == "Y" {

		var pilihan int
		fmt.Println("Pilihan urutan:")
		fmt.Println("1. Pemasukan terbesar")
		fmt.Println("2. Pemasukan terkecil")
		fmt.Println("3. Pengeluaran terbesar")
		fmt.Println("4. Pengeluaran terkecil")
		fmt.Print("Pilih (1/2/3/4): ")
		fmt.Scanf("%d\n", &pilihan)

		if pilihan == 1 || pilihan == 3 {
			urutMembesar(A, idAkun, pilihan)
		}else if pilihan == 2 || pilihan == 4{
			urutMengecil(A, idAkun, pilihan)
		}
	} else if fitur == "n" || fitur == "N" {
		return
	}
}
func urutMembesar(A *tabAkun, idAkun int, pilihan int) {
	// I.S. Terdefinisi array A, idAkun, dan pilihan fitur
	// F.S. Menampilkan tampilan riwayat transaksi sesuai pilihan fitur menggunakan selection sort

	var i, j, idx_min, n int
	var t transaksi
	n = A[idAkun].countTransaksi
	i = 0
	for i < n-1 {
		idx_min = i
		j = i + 1 
		if pilihan == 1 {
			for j < n {
				if A[idAkun].transaksi[idx_min].pemasukan < A[idAkun].transaksi[j].pemasukan {
					idx_min = j
				}
				j++
			}
		} else if pilihan == 3 {
			for j < n {
				if A[idAkun].transaksi[idx_min].pengeluaran < A[idAkun].transaksi[j].pengeluaran {
					idx_min = j
				}
				j++
			}
		}
		t = A[idAkun].transaksi[idx_min]
		A[idAkun].transaksi[idx_min] = A[idAkun].transaksi[i]
		A[idAkun].transaksi[i] = t
		i++
	}
	fmt.Println()
	fmt.Println("Riwayat Transaksi:")
	fmt.Println()
	if pilihan == 1 {
		fmt.Println("____________________________________________________________________")
		fmt.Printf("| %-2s | %-19s | %-22s | %-12s |\n", "No", "Waktu", "Tipe", "Pemasukan")
		for idx := 0; idx < n; idx++ {
			fmt.Printf("| %-2d | %-19s | %-22s | %12d |\n", idx+1, A[idAkun].transaksi[idx].waktu.Format("2006-01-02 15:04:05"), A[idAkun].transaksi[idx].tipe, A[idAkun].transaksi[idx].pemasukan)
		}
		fmt.Println("____________________________________________________________________")
	} else if pilihan == 3 {
		fmt.Println("____________________________________________________________________")
		fmt.Printf("| %-2s | %-19s | %-22s | %-12s |\n", "No", "Waktu", "Tipe", "Pengeluaran")
		for idx := 0; idx < n; idx++ {
			fmt.Printf("| %-2d | %-19s | %-22s | %12d |\n", idx+1, A[idAkun].transaksi[idx].waktu.Format("2006-01-02 15:04:05"), A[idAkun].transaksi[idx].tipe, A[idAkun].transaksi[idx].pengeluaran)
		}
		fmt.Println("____________________________________________________________________")
	}
}
func urutMengecil(A *tabAkun, idAkun int, pilihan int) {
	// I.S. Terdefinisi array A, idAkun, dan pilihan fitur
	// F.S. Menampilkan tampilan riwayat transaksi sesuai pilihan fitur menggunakan insert sort

	var i, j, n int
	n = A[idAkun].countTransaksi
	if pilihan == 2 {
		for i = 1; i < n; i++ {
			temp := A[idAkun].transaksi[i]
			j = i - 1
			for j >= 0 && A[idAkun].transaksi[j].pemasukan > temp.pemasukan {
				A[idAkun].transaksi[j+1] = A[idAkun].transaksi[j]
				j--
			}
			A[idAkun].transaksi[j+1] = temp
		}
	} else if pilihan == 4 {
		for i = 1; i < n; i++ {
			temp := A[idAkun].transaksi[i]
			j = i - 1
			for j >= 0 && A[idAkun].transaksi[j].pengeluaran > temp.pengeluaran {
				A[idAkun].transaksi[j+1] = A[idAkun].transaksi[j]
				j--
			}
			A[idAkun].transaksi[j+1] = temp
		}
	}
	fmt.Println()
	fmt.Println("Riwayat Transaksi:")
	fmt.Println()
	
	if pilihan == 2 {
		fmt.Println("____________________________________________________________________")
		fmt.Printf("| %-2s | %-19s | %-22s | %-12s |\n", "No", "Waktu", "Tipe", "Pemasukan")
		for idx := 0; idx < n; idx++ {
			fmt.Printf("| %-2d | %-19s | %-22s | %12d |\n", idx+1, A[idAkun].transaksi[idx].waktu.Format("2006-01-02 15:04:05"), A[idAkun].transaksi[idx].tipe, A[idAkun].transaksi[idx].pemasukan)
		}
		fmt.Println("____________________________________________________________________")
	} else if pilihan == 4 {
		fmt.Println("____________________________________________________________________")
		fmt.Printf("| %-2s | %-19s | %-22s | %-12s |\n", "No", "Waktu", "Tipe", "Pengeluaran")
		for idx := 0; idx < n; idx++ {
			fmt.Printf("| %-2d | %-19s | %-22s | %12d |\n", idx+1, A[idAkun].transaksi[idx].waktu.Format("2006-01-02 15:04:05"), A[idAkun].transaksi[idx].tipe, A[idAkun].transaksi[idx].pengeluaran)
		}
		fmt.Println("____________________________________________________________________")
	}
}
