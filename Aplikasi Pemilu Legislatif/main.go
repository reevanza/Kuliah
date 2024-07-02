package main

import (
	"fmt"
	"strings"
	"time"
)

const NMAX int = 1000

type calon struct {
	namaDepancalon    string
	namaBelakangcalon string
	namapartai        string
	nomorurut         int
	totalsuara        int
}

type petugas struct {
	username string
	password string
}

type pemilih struct {
	nomorNik            string
	namaDepanPemilih    string
	namaBelakangPemilih string
	pilihanDPD          int
	pilihanDPRRI        int
	pilihanDPRDProv     int
	pilihanDPRDKab      int
	sudahMemilih        bool
}

type tabCalon [NMAX]calon
type tabPetugas [NMAX]petugas
type tabPemilih [NMAX]pemilih

var A_DPD, A_DPR_RI, A_DPRDProv, A_DPRDKab tabCalon
var B tabPemilih
var C tabPetugas
var n_DPD, n_DPR_RI, n_DPRDProv, n_DPRDKab, n_Pemilih, n_Petugas int

func main() {
	menu()
}

func menu() {
	var choice int
	maskot()
	gambarMenu()
	for choice != 3 {
		fmt.Println("╔═══╦═════════════════╗")
		fmt.Println("║   ║ Pilih Role Anda ║")
		fmt.Println("╠═══╬═════════════════╣")
		fmt.Println("║ 1 ║   Petugas KPU   ║")
		fmt.Println("║ 2 ║     Pemilih     ║")
		fmt.Println("║ 3 ║      Exit       ║")
		fmt.Println("╚═══╩═════════════════╝")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&choice)
		fmt.Println()
		switch choice {
		case 1:
			DaftarPetugas(&C, &n_Petugas)
		case 2:
			MenuPemilih()
		case 3:
			close()
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia")
		}
	}

}

func close() {
	fmt.Println("╔═══════════════════════════════════════════╗")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("║            Program dibuat oleh            ║")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("╠═══════════════════════════════════════════╣")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("║ Reevanza Abel Desta Arifin (103012330104) ║")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("║                                           ║")
	fmt.Println("╚═══════════════════════════════════════════╝")
	fmt.Println()
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Terimakasih")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Thank you")
	time.Sleep(6 * time.Second)
	fmt.Println()
}

func maskot() {
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▓▓▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▓████▓▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▓████████▓▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▓████████████▓▒░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▓████████████████▓▒░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▓████████████████▓▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▓████████████████▓▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▓▓███████████████▓▓▒▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒░░░░▓███████████████▓▒░▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒░░░░░░░░░░▒▓███████████▓▒░░░░░░░▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░▒▓▓██████▓▒░░░░░░░░░░░░▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░▒▒▓▓░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒▓▓▓▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒█████▓▓▓▒▒▒▒▒▒▒▒░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒██████████▓▓▒▒▒▒▒▒▒▒░░░░░▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒██████████████▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░░░░░░▒▒▒▓▓░░░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒███████████████████▓▒▒▒▒▒▒░░░░░░▒▒▓▒▓▓▓▓███░░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒████████████████████▓▒▒▒▒░░░░▒▒▓▓▓█████████▓░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒████████████████████▓▒▒▒▒░░░░▒▒▓▓▓█████████▓░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒████████████████████▓▒▒▒▒░░░██▓████████▓██▓▓░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒▓███████████████████▓▒▒▒▒░░▓██████▓▓██▓▒████░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒░░░▒▒▓██████████████▓▒▒▒▒░▒███▓███▓▓▓▒▒▒▓▓██░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒░░░░░░░░▒▒▓█████████▓▒▒▒▒░▓▓█▓▒▒▒▓█▓▒▒▒░░░▓▓░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒░░░░░░░░░░░░░▒▓▓████▓▒▒▒▒░████▒▒▓▓▓▓▓░░░░▒▒▒░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒░░░░░░░░░░░░░▒▓▓████▓▒▒▒▒░████▒▒▓▓▓▓▓░░░░▒▒▒░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒██░░░░▒▒▒▒░░░░▒▒▓░░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒░░░░░░░░░░░░░░░░░░░░░▒▒▒▒░▓▒▒░░░░░░░░░▒▒▓▓░░░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒░░░░░░░░░░░░░░░░░░░░░▒▒▒▒░▓▓▓░░░░░▒▒▒▓█▒▒░░░░▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░▒▒▒▒░░▓▒▓▓▓▓▓▓▓▒▒░░░░░░▒▒▒▒▒░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░▒▒▒▒░░░▓▒▓▒▒░░░░░░▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒░░░░░░░░░▒▒▒▒░░░░░░░░░░▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒░░░░░▒▒▒▒░░░░░▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒░▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░▒▓███▓▒░░░░░▒▓███▓▒░░░░░▓▓███▓▒░░░░░░░░░▒█▓░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░██▒░░░▒██░░░▓█▒░░░▒█▓░░░██░░░░▒██░░░░░░░▓██▓░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░▒░░░░░░▓█▒░▒█▓░░░░░▓█▒░░▒░░░░░░▓█▒░░░░▒█▓▒█▓░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░▓█▓░░▓█▒░░░░░▒█▓░░░░░░░░▓█▓░░░░▓█▒░░█▓░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░▓██▒░░░▓█▒░░░░░▒█▓░░░░░░▓█▓▒░░░░█▓░░░░█▓░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░▓█▓░░░░░░▒█▓░░░░░▓█▒░░░▒▓█▓░░░░░░██████████▓░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░▓█▓░░░░░░░░░██░░░░░██░░░▓█▓░░░░░░░░░░░░░░▒█▓░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░▓█████████▒░░░▓██▓██▓░░░▓█████████▒░░░░░░░░█▓░░░░░░░░░░░░░░░░░░░░░░░░")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
}

func loading() {
	fmt.Println("Loading...")
	time.Sleep(100 * time.Millisecond)
	fmt.Print("█")
	time.Sleep(150 * time.Millisecond)
	fmt.Print("██")
	time.Sleep(150 * time.Millisecond)
	fmt.Print("███")
	time.Sleep(200 * time.Millisecond)
	fmt.Print("████")
	time.Sleep(300 * time.Millisecond)
	fmt.Print("█████")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("██████")
	time.Sleep(100 * time.Millisecond)
}

func gambarMenu() {
	fmt.Println("    ╔═══════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("    ║                          Komisi Pemilihan Umum(KPU)                           ║")
	fmt.Println("    ╠═══════════════════════════════════════════════════════════════════════════════╣")
	fmt.Println("    ║                    Selamat datang di Pemilihan Umum 2024!                     ║")
	fmt.Println("    ║ Mari bersama-sama menggunakan hak suara kita untuk masa depan yang lebih baik ║")
	fmt.Println("    ║       Pemilu adalah kesempatan kita untuk memilih pemimpin yang terbaik       ║")
	fmt.Println("    ║      Jangan lupa datang ke TPS dan berpartisipasi dalam pemilu kali ini       ║")
	fmt.Println("    ║                    Suara Anda menentukan arah bangsa kita                     ║")
	fmt.Println("    ║                          Ayo sukseskan Pemilu 2024!                           ║")
	fmt.Println("    ╚═══════════════════════════════════════════════════════════════════════════════╝")
}

//PETUGAS

func DaftarPetugas(C *tabPetugas, n_Petugas *int) {
	var akun string

	fmt.Println("Apakah anda ingin membuat akun?(Ya/Tidak).")
	fmt.Scan(&akun)
	if strings.ToLower(akun) == "ya" {
		fmt.Println("╔═══════════════════════════════════╗")
		fmt.Println("║ Pendaftaran Akun Petugas KPU Baru ║")
		fmt.Println("╚═══════════════════════════════════╝")
		fmt.Print("Masukkan Username: ")
		fmt.Scan(&C[*n_Petugas].username)
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&C[*n_Petugas].password)
		*n_Petugas++
		fmt.Println()
		loading()
		fmt.Println()
		fmt.Println("Pendaftaran berhasil.")
		fmt.Println()
		LoginPetugas()
	} else {
		LoginPetugas()
	}
}

func LoginPetugas() {
	var username, password string
	var idx int

	idx = -1
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║   Portal Login Petugas KPU   ║")
	fmt.Println("╚══════════════════════════════╝")
	fmt.Println("Silakan masukkan username dan password Anda!")
	fmt.Print("Masukkan Username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan Password: ")
	fmt.Scan(&password)
	for i := 0; i < n_Petugas; i++ {
		if C[i].username == username && C[i].password == password {
			fmt.Println("Login Berhasil")
			idx = i
			fmt.Println()
			MenuPetugas()

		}
	}
	if idx == -1 {
		fmt.Println()
		fmt.Println("Username atau Password salah!")
		fmt.Println()
	}

}

func MenuPetugas() {
	var choicePetugas int
	fmt.Println("Selamat datang, Petugas!")
	for choicePetugas != 7 {
		fmt.Println("╔════╦═════════════════════════════╗")
		fmt.Println("║    ║        Menu Petugas         ║")
		fmt.Println("╠════╬═════════════════════════════╣")
		fmt.Println("║ 1  ║ Input data Partai dan Calon ║")
		fmt.Println("║ 2  ║ Cetak data Partai dan Calon ║")
		fmt.Println("║ 3  ║      Cari data Partai       ║")
		fmt.Println("║ 4  ║        Delete Calon         ║")
		fmt.Println("║ 5  ║         Edit Calon          ║")
		fmt.Println("║ 6  ║     Cek hasil Perolehan     ║")
		fmt.Println("║ 7  ║           Kembali           ║")
		fmt.Println("╚════╩═════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choicePetugas)
		fmt.Println()
		switch choicePetugas {
		case 1:
			PetugasInputData()
		case 2:
			PetugasCetakData()
		case 3:
			PetugasCariData()
		case 4:
			PetugasDeleteData()
		case 5:
			PetugasEditData()
		case 6:
			suaraPerolehan()
		case 7:
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia.")
		}
	}
}

func MenuPemilih() {
	var choicePemilih int
	fmt.Println("Selamat Datang dalam Pemilihan Pemilu!")
	fmt.Println("Silahkan pilih menu yang tersedua.")
	for choicePemilih != 5 {
		fmt.Println("╔════╦══════════════════════════════════════╗")
		fmt.Println("║    ║            Menu Pemilihan            ║")
		fmt.Println("╠════╬══════════════════════════════════════╣")
		fmt.Println("║ 1  ║         Pendaftaran Pemilih          ║")
		fmt.Println("║ 2  ║          Lakukan Pemilihan           ║")
		fmt.Println("║ 3  ║         Cek hasil perolehan          ║")
		fmt.Println("║ 4  ║  Cari Calon berdasarkan nama Partai  ║")
		fmt.Println("║ 5  ║               Kembali                ║")
		fmt.Println("╚════╩══════════════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choicePemilih)
		fmt.Println()

		switch choicePemilih {
		case 1:
			registPemilih(&B, &n_Pemilih)
		case 2:
			Pemilihan()
		case 3:
			suaraPerolehan()
		case 4:
			PemilihCariData()
		case 5:
			return
		default:
			fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia.")
		}
	}
}

func PetugasInputData() {
	var choice int
	fmt.Println("╔════╦════════════════════════════════════════╗")
	fmt.Println("║    ║ Pilih jenis pemilihan untuk input data ║")
	fmt.Println("╠════╬════════════════════════════════════════╣")
	fmt.Println("║ 1  ║                  DPD                   ║")
	fmt.Println("║ 2  ║                 DPR-RI                 ║")
	fmt.Println("║ 3  ║             DPRD Provinsi              ║")
	fmt.Println("║ 4  ║             DPRD Kab/Kota              ║")
	fmt.Println("╚════╩════════════════════════════════════════╝")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&choice)
	fmt.Println()
	switch choice {
	case 1:
		inputDataCalon(&A_DPD, &n_DPD)
	case 2:
		inputDataCalon(&A_DPR_RI, &n_DPR_RI)
	case 3:
		inputDataCalon(&A_DPRDProv, &n_DPRDProv)
	case 4:
		inputDataCalon(&A_DPRDKab, &n_DPRDKab)
	default:
		fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia.")
	}
}

func PetugasCetakData() {
	var choice int
	fmt.Println("╔════╦════════════════════════════════════════╗")
	fmt.Println("║    ║ Pilih jenis pemilihan untuk cetak data ║")
	fmt.Println("╠════╬════════════════════════════════════════╣")
	fmt.Println("║ 1  ║                  DPD                   ║")
	fmt.Println("║ 2  ║                 DPR-RI                 ║")
	fmt.Println("║ 3  ║             DPRD Provinsi              ║")
	fmt.Println("║ 4  ║             DPRD Kab/Kota              ║")
	fmt.Println("╚════╩════════════════════════════════════════╝")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&choice)
	fmt.Println()
	switch choice {
	case 1:
		cetakDataCalon(A_DPD, n_DPD)
	case 2:
		cetakDataCalon(A_DPR_RI, n_DPR_RI)
	case 3:
		cetakDataCalon(A_DPRDProv, n_DPRDProv)
	case 4:
		cetakDataCalon(A_DPRDKab, n_DPRDKab)
	default:
		fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia.")
	}
}

func PetugasCariData() {
	var choice int
	fmt.Println("╔════╦═══════════════════════════════════════╗")
	fmt.Println("║    ║ Pilih jenis pemilihan untuk cari data ║")
	fmt.Println("╠════╬═══════════════════════════════════════╣")
	fmt.Println("║ 1  ║                  DPD                  ║")
	fmt.Println("║ 2  ║                DPR-RI                 ║")
	fmt.Println("║ 3  ║             DPRD Provinsi             ║")
	fmt.Println("║ 4  ║             DPRD Kab/Kota             ║")
	fmt.Println("╚════╩═══════════════════════════════════════╝")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&choice)
	fmt.Println()
	switch choice {
	case 1:
		cariPartai(A_DPD, n_DPD)
	case 2:
		cariPartai(A_DPR_RI, n_DPR_RI)
	case 3:
		cariPartai(A_DPRDProv, n_DPRDProv)
	case 4:
		cariPartai(A_DPRDKab, n_DPRDKab)
	default:
		fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia.")
	}
}

func PetugasDeleteData() {
	var choice int
	fmt.Println("╔════╦═════════════════════════════════════════╗")
	fmt.Println("║    ║ Pilih jenis pemilihan untuk delete data ║")
	fmt.Println("╠════╬═════════════════════════════════════════╣")
	fmt.Println("║ 1  ║                  DPD                    ║")
	fmt.Println("║ 2  ║                 DPR-RI                  ║")
	fmt.Println("║ 3  ║             DPRD Provinsi               ║")
	fmt.Println("║ 4  ║             DPRD Kab/Kota               ║")
	fmt.Println("╚════╩═════════════════════════════════════════╝")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&choice)
	fmt.Println()
	switch choice {
	case 1:
		deleteCalon(&A_DPD, &n_DPD)
	case 2:
		deleteCalon(&A_DPR_RI, &n_DPR_RI)
	case 3:
		deleteCalon(&A_DPRDProv, &n_DPRDProv)
	case 4:
		deleteCalon(&A_DPRDKab, &n_DPRDKab)
	default:
		fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia.")
	}
}

func PetugasEditData() {
	var choice int
	fmt.Println("╔════╦═══════════════════════════════════════╗")
	fmt.Println("║    ║ Pilih jenis pemilihan untuk edit data ║")
	fmt.Println("╠════╬═══════════════════════════════════════╣")
	fmt.Println("║ 1  ║                  DPD                  ║")
	fmt.Println("║ 2  ║                DPR-RI                 ║")
	fmt.Println("║ 3  ║             DPRD Provinsi             ║")
	fmt.Println("║ 4  ║             DPRD Kab/Kota             ║")
	fmt.Println("╚════╩═══════════════════════════════════════╝")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&choice)
	fmt.Println()
	switch choice {
	case 1:
		editCalon(&A_DPD, n_DPD)
	case 2:
		editCalon(&A_DPR_RI, n_DPR_RI)
	case 3:
		editCalon(&A_DPRDProv, n_DPRDProv)
	case 4:
		editCalon(&A_DPRDKab, n_DPRDKab)
	default:
		fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia.")
	}
}

func inputDataCalon(A *tabCalon, n *int) {
	var cekWNI string
	fmt.Println("Apakah Calon merupakan Warga Negara Indonesia(WNI)? Ya atau Tidak.")
	fmt.Println("Are the Candidates Indonesian Citizens(WNI)? Yes or No.")
	fmt.Scan(&cekWNI)
	if strings.ToLower(cekWNI) == "yes" || strings.ToLower(cekWNI) == "ya" {
		fmt.Print("Masukkan nama depan calon: ")
		fmt.Scan(&A[*n].namaDepancalon)
		fmt.Print("Nama belakang calon (ketik '-' jika tidak ada): ")
		fmt.Scan(&A[*n].namaBelakangcalon)
		if A[*n].namaBelakangcalon == "-" {
			A[*n].namaBelakangcalon = ""
		}
		fmt.Print("Masukkan nama partai: ")
		fmt.Scan(&A[*n].namapartai)
		A[*n].nomorurut = *n + 1
		*n++
		fmt.Println()
		loading()
		fmt.Println()
		fmt.Println("Pendaftaran berhasil.")
		fmt.Println()
	} else {
		fmt.Println("candidate is not qualified.")
		fmt.Println("Exit.")
		fmt.Println()
		return
	}
}

func cetakDataCalon(A tabCalon, n int) {
	maxNoUrut := len("No Urut")
	maxNamaCalon := len("Nama Calon")
	maxPartai := len("Partai")

	for i := 0; i < n; i++ {
		noUrutWidth := len(fmt.Sprintf("%d", i+1))
		namaCalonWidth := len(A[i].namaDepancalon + " " + A[i].namaBelakangcalon)
		partaiWidth := len(A[i].namapartai)

		if noUrutWidth > maxNoUrut {
			maxNoUrut = noUrutWidth
		}
		if namaCalonWidth > maxNamaCalon {
			maxNamaCalon = namaCalonWidth
		}
		if partaiWidth > maxPartai {
			maxPartai = partaiWidth
		}
	}
	fmt.Println()
	loading()
	fmt.Println()
	fmt.Println(strings.Repeat("━", maxNoUrut+maxNamaCalon+maxPartai+10))
	fmt.Printf("| %-*s | %-*s | %-*s |\n", maxNoUrut, "No Urut", maxNamaCalon, "Nama Calon", maxPartai, "Partai")
	fmt.Println(strings.Repeat("━", maxNoUrut+maxNamaCalon+maxPartai+10))

	for i := 0; i < n; i++ {
		namaCalon := A[i].namaDepancalon + " " + A[i].namaBelakangcalon
		fmt.Printf("| %-*d | %-*s | %-*s |\n", maxNoUrut, i+1, maxNamaCalon, namaCalon, maxPartai, A[i].namapartai)
		fmt.Println(strings.Repeat("─", maxNoUrut+maxNamaCalon+maxPartai+10))
	}
	fmt.Println()
}

func cariPartai(A tabCalon, n int) {
	var found bool
	var keyword string
	fmt.Print("Masukkan nama partai yang ingin dicari: ")
	fmt.Scan(&keyword)

	fmt.Println("Daftar calon pada partai:", keyword)

	maxNoUrut := len("No Urut")
	maxNamaCalon := len("Nama Calon")

	for i := 0; i < n; i++ {
		if A[i].namapartai == keyword {
			noUrutWidth := len(fmt.Sprintf("%d", i+1))
			namaCalonWidth := len(A[i].namaDepancalon + " " + A[i].namaBelakangcalon)
			if noUrutWidth > maxNoUrut {
				maxNoUrut = noUrutWidth
			}
			if namaCalonWidth > maxNamaCalon {
				maxNamaCalon = namaCalonWidth
			}
			found = true
		}
	}

	if !found {
		fmt.Println("Partai tidak ditemukan.")
		fmt.Println()
		return
	}
	fmt.Println()
	loading()
	fmt.Println()
	fmt.Println(strings.Repeat("━", maxNoUrut+maxNamaCalon+8))
	fmt.Printf(" | %-*s | %-*s |\n", maxNoUrut, "No Urut", maxNamaCalon, "Nama Calon")
	fmt.Println(strings.Repeat("━", maxNoUrut+maxNamaCalon+8))

	for i := 0; i < n; i++ {
		if A[i].namapartai == keyword {
			namaCalon := A[i].namaDepancalon + " " + A[i].namaBelakangcalon
			fmt.Printf(" | %-*d | %-*s |\n", maxNoUrut, i+1, maxNamaCalon, namaCalon)
			fmt.Println(strings.Repeat("─", maxNoUrut+maxNamaCalon+8))
		}
	}
}

func editCalon(A *tabCalon, n int) {
	var idx int
	var newNamaDepan, newNamaBelakang, newPartai string

	cetakDataCalon(*A, n)
	fmt.Print("Masukkan nomor calon yang akan diubah: ")
	fmt.Scan(&idx)

	if idx < 1 || idx > n {
		fmt.Println("Nomor calon tidak valid.")
		return
	}
	idx--

	fmt.Println("Data calon yang akan diubah:")
	fmt.Printf("Nama Depan: %s\n", A[idx].namaDepancalon)
	fmt.Printf("Nama Belakang: %s\n", A[idx].namaBelakangcalon)
	fmt.Printf("Partai: %s\n", A[idx].namapartai)
	fmt.Println()
	fmt.Println("Masukkan data baru (ketik '-', jika tidak ingin mengubah): ")
	fmt.Print("Nama Depan baru: ")
	fmt.Scan(&newNamaDepan)

	if newNamaDepan != "-" {
		A[idx].namaDepancalon = newNamaDepan
	}

	fmt.Print("Nama Belakang baru: ")
	fmt.Scan(&newNamaBelakang)
	if newNamaBelakang != "-" {
		A[idx].namaBelakangcalon = newNamaBelakang
	}

	fmt.Print("Partai baru: ")
	fmt.Scan(&newPartai)
	if newPartai != "-" {
		A[idx].namapartai = newPartai
	}
	fmt.Println()
	loading()
	fmt.Println()
	fmt.Println("Data calon berhasil diubah.")
	fmt.Println()
}

func deleteCalon(A *tabCalon, n *int) {
	var idx int
	cetakDataCalon(*A, *n)
	fmt.Print("Masukkan nomor calon yang akan dihapus: ")
	fmt.Scan(&idx)
	if idx < 1 || idx > *n {
		fmt.Println("Nomor calon tidak valid.")
		return
	}

	for i := idx - 1; i <= *n-2; i++ {
		A[i] = A[i+1]
	}
	*n--
	fmt.Println()
	loading()
	fmt.Println()
	fmt.Println("Calon berhasil dihapus.")
	fmt.Println()
}

// PEMILIH

func PemilihCariData() {
	var choice int
	fmt.Println("╔════╦═══════════════════════════════════════╗")
	fmt.Println("║    ║ Pilih jenis pemilihan untuk cari data ║")
	fmt.Println("╠════╬═══════════════════════════════════════╣")
	fmt.Println("║ 1  ║                  DPD                  ║")
	fmt.Println("║ 2  ║                DPR-RI                 ║")
	fmt.Println("║ 3  ║             DPRD Provinsi             ║")
	fmt.Println("║ 4  ║             DPRD Kab/Kota             ║")
	fmt.Println("╚════╩═══════════════════════════════════════╝")
	fmt.Print("Pilih menu : ")
	fmt.Scan(&choice)
	fmt.Println()
	switch choice {
	case 1:
		cariPartai(A_DPD, n_DPD)
	case 2:
		cariPartai(A_DPR_RI, n_DPR_RI)
	case 3:
		cariPartai(A_DPRDProv, n_DPRDProv)
	case 4:
		cariPartai(A_DPRDKab, n_DPRDKab)
	default:
		fmt.Println("Menu tidak tersedia, silahkan pilih menu yang tersedia.")
	}
}

func registPemilih(B *tabPemilih, n_Pemilih *int) {
	var cekWNI string
	fmt.Println("Apakah anda Warga Negara Indonesia(WNI)? Ya atau Tidak.")
	fmt.Println("Are you an Indonesian citizen (WNI)? Yes or No")
	fmt.Scan(&cekWNI)
	if strings.ToLower(cekWNI) == "ya" || strings.ToLower(cekWNI) == "yes" {
		fmt.Print("Masukkan nama depan pemilih: ")
		fmt.Scan(&B[*n_Pemilih].namaDepanPemilih)
		fmt.Print("Nama belakang pemilih (ketik '-' jika tidak ada): ")
		fmt.Scan(&B[*n_Pemilih].namaBelakangPemilih)
		if B[*n_Pemilih].namaBelakangPemilih == "-" {
			B[*n_Pemilih].namaBelakangPemilih = ""
		}
		fmt.Print("Masukkan nomor NIK: ")
		fmt.Scan(&B[*n_Pemilih].nomorNik)
		*n_Pemilih++
		fmt.Println()
		loading()
		fmt.Println()
		fmt.Println("Pendaftaran berhasil.")
		fmt.Println()
	} else {
		fmt.Println("You don't qualify for enrollment.")
		fmt.Println("Exit.")
		fmt.Println()
		return
	}
}

func Pemilihan() {
	var nomorNIK string
	var found bool

	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║ Silahkan login menggunakan nomor NIK ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Println("")

	fmt.Print("Masukkan nomor NIK: ")
	fmt.Scan(&nomorNIK)
	for i := 0; i < n_Pemilih; i++ {
		if B[i].nomorNik == nomorNIK {
			found = true
			fmt.Print("Nama Pemilih: ")
			fmt.Println(B[i].namaDepanPemilih, B[i].namaBelakangPemilih)
			fmt.Println()
			if B[i].sudahMemilih {
				fmt.Println("Anda sudah memberikan suara. Terima kasih.")
				return
			} else {
				fmt.Println("╔═════════════════════════════════════════════╗")
				fmt.Println("║ Silakan memilih calon untuk setiap kategori ║")
				fmt.Println("╚═════════════════════════════════════════════╝")
				fmt.Println()
				fmt.Println("Kategori DPD")
				pilihCalon(&A_DPD, &B[i].pilihanDPD, n_DPD)
				fmt.Println("Kategori DPR-RI")
				pilihCalon(&A_DPR_RI, &B[i].pilihanDPRRI, n_DPR_RI)
				fmt.Println("Kategori DPRD Provinsi")
				pilihCalon(&A_DPRDProv, &B[i].pilihanDPRDProv, n_DPRDProv)
				fmt.Println("Kategori DPRD Kab/Kota")
				pilihCalon(&A_DPRDKab, &B[i].pilihanDPRDKab, n_DPRDKab)
				B[i].sudahMemilih = true
				return
			}
		}
	}
	if !found {
		fmt.Println("NIK tidak ditemukan.")
	}
}

func pilihCalon(A *tabCalon, pilihan *int, n int) {
	cetakDataCalon(*A, n)
	fmt.Print("Masukkan nomor urut calon yang dipilih: ")
	fmt.Scan(pilihan)
	for i := 0; i < n; i++ {
		if A[i].nomorurut == *pilihan {
			A[i].totalsuara++
			fmt.Println()
			loading()
			fmt.Println()
			fmt.Println("Pilihan berhasil disimpan.")
			return
		}
	}
	fmt.Println("Nomor urut tidak valid.")
	fmt.Println()
}

func suaraPerolehan() {
	fmt.Println("╔═══════════════════════╗")
	fmt.Println("║ Hasil Perolehan Suara ║")
	fmt.Println("╚═══════════════════════╝")
	fmt.Println()
	fmt.Println("Kategori DPD:")
	cetakHasil(A_DPD, n_DPD)
	fmt.Println("Kategori DPR-RI:")
	cetakHasil(A_DPR_RI, n_DPR_RI)
	fmt.Println("Kategori DPRD Provinsi:")
	cetakHasil(A_DPRDProv, n_DPRDProv)
	fmt.Println("Kategori DPRD Kab/Kota:")
	cetakHasil(A_DPRDKab, n_DPRDKab)
}

func cetakHasil(A tabCalon, n int) {
	insertionSorting(&A, n)
	maxNoUrut := len("No Urut")
	maxNamaCalon := len("Nama Calon")
	maxPartai := len("Partai")
	maxTotalSuara := len("Total Suara")

	fmt.Println()
	loading()
	fmt.Println()

	for i := 0; i < n; i++ {
		noUrutWidth := len(fmt.Sprintf("%d", A[i].nomorurut))
		namaCalonWidth := len(A[i].namaDepancalon + " " + A[i].namaBelakangcalon)
		partaiWidth := len(A[i].namapartai)
		totalSuaraWidth := len(fmt.Sprintf("%d", A[i].totalsuara))

		if noUrutWidth > maxNoUrut {
			maxNoUrut = noUrutWidth
		}
		if namaCalonWidth > maxNamaCalon {
			maxNamaCalon = namaCalonWidth
		}
		if partaiWidth > maxPartai {
			maxPartai = partaiWidth
		}
		if totalSuaraWidth > maxTotalSuara {
			maxTotalSuara = totalSuaraWidth
		}
	}

	fmt.Println(strings.Repeat("━", maxNoUrut+maxNamaCalon+maxPartai+maxTotalSuara+13))
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s |\n", maxNoUrut, "No Urut", maxNamaCalon, "Nama Calon", maxPartai, "Partai", maxTotalSuara, "Total Suara")
	fmt.Println(strings.Repeat("━", maxNoUrut+maxNamaCalon+maxPartai+maxTotalSuara+13))

	for i := 0; i < n; i++ {
		namaCalon := A[i].namaDepancalon + " " + A[i].namaBelakangcalon
		fmt.Printf("| %-*d | %-*s | %-*s | %-*d |\n", maxNoUrut, A[i].nomorurut, maxNamaCalon, namaCalon, maxPartai, A[i].namapartai, maxTotalSuara, A[i].totalsuara)
		fmt.Println(strings.Repeat("─", maxNoUrut+maxNamaCalon+maxPartai+maxTotalSuara+13))
	}
}

func insertionSorting(A *tabCalon, n int) {
	var pass, i int
	var temp calon

	pass = 1
	for pass <= n-1 {
		temp = A[pass]
		i = pass
		for i > 0 && temp.totalsuara > A[i-1].totalsuara {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}
