package main

import (
	"fmt"
	"time"
)

func main() {
	var tanggal, film, tiket, menu, jumlahpesanan, harga1, hargatotal, hargadisk, harga, hargaakhir int
	var kode, camilan, vocer string
	camilan = "Ya"
	tanggal = 32
	film = 99
	jumlahpesanan = 0

	fmt.Println(" Loading…")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("                                                                ▄▀▄     ▄▀▄")
	fmt.Println("                                                               ▄█░░▀▀▀▀▀░░█▄")
	fmt.Println("                                                           ▄▄  █░░░░░░░░░░░█  ▄▄")
	fmt.Println("                                                          █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
	fmt.Println("                                                          █▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█")
	fmt.Println("                                                          █░░╦─╦╔╗╦─╔╗╔╗╔╦╗╔╗░░█")
	fmt.Println("                                                          █░░║║║╠─║─║─║║║║║╠─░░█")
	fmt.Println("                                                          █░░╚╩╝╚╝╚╝╚╝╚╝╩─╩╚╝░░█")
	fmt.Println("                                                          █▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█")
	fmt.Println("")
	fmt.Println("█▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀█")
	fmt.Println("█ Selamat datang di tempat di mana keseruan menonton film dimulai! Pesan tiket Anda dengan mudah dan nikmati pengalaman tak terlupakan di Indo XXV █")
	fmt.Println("█▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄█")

	for tanggal != 0 && (camilan == "Ya" || camilan == "ya") {
		fmt.Println(" ")
		fmt.Println("Mungkin ini kabar baik: tiket bioskop di sini hanya Rp 30000! ")
		fmt.Println("Namun, jika hari libur nasional dan weekend, harga menjadi Rp 45000")
		fmt.Println(" ")
		fmt.Println("Pilihlah tanggal kunjungan Anda di bulan Januari 2024 untuk menikmati serangkaian film terbaik :")
		fmt.Scan(&tanggal)
		for tanggal < 1 || tanggal > 31 {
			fmt.Println("Mohon masukkan tanggal kunjungan Anda dengan format yang benar")
			fmt.Scan(&tanggal)
		}

		fmt.Println("")
		fmt.Println("                                           ██▄▄")
		fmt.Println("                                           ██▀▀")
		fmt.Println("        LOADING....                      ▄███▄")
		fmt.Println("                                       ▄█████")
		fmt.Println("                drap     drap     ▀▄▄▀▀  █▄ █▄")
		fmt.Println("")
		fmt.Println(" ")
		fmt.Println("Film yang tersedia pada tanggal", tanggal, ", Desember 2023 sebagai berikut ")
		fmt.Println(" ")
		if tanggal > 0 && tanggal < 8 {

			fmt.Println("                        ▄▀▄     ▄▀▄")
			fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
			fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
			fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  No  ||                 FILM                 ||  Jam  ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  1   || Dilan 1945                           || 08.30 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  2   || Aku Tahu Kapan Kamu Lulus            || 11.00 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  3   || Ada Apa dengan Ilham                 || 15.25 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  4   || Cinta di Kampus                      || 18.30 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                               ||")
			fmt.Println("||  0   ||                   Kembali                     ||")
			fmt.Println("||      ||                                               ||")
			fmt.Println("===========================================================")
			fmt.Println("Pilih nomor film yang akan ditonton")
			fmt.Scan(&film)
			for film < 0 || film > 4 {
				fmt.Scan(&film)
				if film < 0 || film > 4 {
					fmt.Println("Masukkan input yang valid")
				}
			}

			if film == 1 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||               Dilan 1945             ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			} else if film == 2 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||       Aku Tahu Kapan Kamu Lulus      ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")

			} else if film == 3 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||          Ada Apa dengan Ilham        ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")

			} else if film == 4 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||            Cinta di Kampus           ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")

			}

		} else if tanggal >= 8 && tanggal < 15 {
			fmt.Println("                        ▄▀▄     ▄▀▄")
			fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
			fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
			fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  No  ||                 FILM                 ||  Jam  ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  1   || Faishal The Explorer                 || 08.30 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  2   || Rizki: Suara dari Hati               || 11.00 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  3   || Pengabdi Teh Poci                    || 15.25 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  4   || Seni Berpedang Daring: Pedang Hitam  || 18.30 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                               ||")
			fmt.Println("||  0   ||                   Kembali                     ||")
			fmt.Println("||      ||                                               ||")
			fmt.Println("===========================================================")
			fmt.Println("Pilih nomor film yang akan ditonton")
			fmt.Scan(&film)
			for film < 0 || film > 4 {
				fmt.Scan(&film)
				if film < 0 || film > 4 {
					fmt.Println("Masukkan input yang valid")
				}
			}

			if film == 1 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||          Faishal The Explorer        ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")

			} else if film == 2 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||         Rizki: Suara dari hati       ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			} else if film == 3 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||           Pengabdi Teh Poci          ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			} else if film == 4 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||  Seni Berpedang Daring: Pedang Hitam ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			}

		} else if tanggal >= 15 && tanggal < 22 {
			fmt.Println("                        ▄▀▄     ▄▀▄")
			fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
			fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
			fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  No  ||                 FILM                 ||  Jam  ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  1   || Danus: Mencari Modal                 || 08.30 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  2   || Dear Hansen                          || 11.00 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  3   || Koboy Senior                         || 15.25 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  4   || Ancoko: Dia Yang Bersamaku Tahun 1945|| 18.30 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                               ||")
			fmt.Println("||  0   ||                   Kembali                     ||")
			fmt.Println("||      ||                                               ||")
			fmt.Println("===========================================================")
			fmt.Println("Pilih nomor film yang akan ditonton")
			fmt.Scan(&film)
			for film < 0 || film > 4 {
				fmt.Scan(&film)
				if film < 0 || film > 4 {
					fmt.Println("Masukkan input yang valid")
				}
			}

			if film == 1 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||          Danus: Mencari Modal        ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			} else if film == 2 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||              Dear Hansen             ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			} else if film == 3 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||              Koboy Senior            ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			} else if film == 4 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      || Ancoko: Dia yang Bersamaku Tahun 1945||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			}

		} else if tanggal >= 22 && tanggal < 32 {
			fmt.Println("                        ▄▀▄     ▄▀▄")
			fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
			fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
			fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  No  ||                 FILM                 ||  Jam  ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  1   || The Nasi Bungkus                     || 08.30 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  2   || Etet Marketet                        || 11.00 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  3   || My Name                              || 15.25 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("||  4   || Satu Potong: Merah                   || 18.30 ||")
			fmt.Println("||      ||                                      ||       ||")
			fmt.Println("===========================================================")
			fmt.Println("||      ||                                               ||")
			fmt.Println("||  0   ||                   Kembali                     ||")
			fmt.Println("||      ||                                               ||")
			fmt.Println("===========================================================")
			fmt.Println("Pilih nomor film yang akan ditonton")
			fmt.Scan(&film)
			for film < 0 || film > 4 {
				fmt.Scan(&film)
				if film < 0 || film > 4 {
					fmt.Println("Masukkan input yang valid")
				}
			}

			if film == 1 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||            The Nasi Bungkus          ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			} else if film == 2 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||             Etet Marketet            ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			} else if film == 3 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||                My Name               ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			} else if film == 4 {
				fmt.Println(" ")
				fmt.Println("                        ▄▀▄     ▄▀▄")
				fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("===========================================================")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("||      ||           Satu Potong: Merah         ||       ||")
				fmt.Println("||      ||                                      ||       ||")
				fmt.Println("===========================================================")
			}

		}
		if film >= 1 && film <= 4 {
			fmt.Println("")
			fmt.Println("Masukkan jumlah pembelian")
			fmt.Scan(&tiket)
			for !(tiket >= 1 && tiket <= 99999) {
				fmt.Scan(&tiket)
				if tiket < 0 || !(tiket >= 1 && tiket <= 99999) {
					fmt.Println("Masukkan input yang valid")
				}
			}
			harga1 = 30000
			if tanggal == 1 || tanggal == 6 || tanggal == 7 || tanggal == 13 || tanggal == 14 || tanggal == 20 || tanggal == 21 || tanggal == 27 || tanggal == 28 {
				harga1 += 15000
				harga1 *= tiket
			} else if tanggal == 12 {
				hargadisk = harga1 * 30 / 100
				harga1 = (harga1 - hargadisk) * tiket
			} else {
				harga1 *= tiket
			}
			fmt.Println("")
			fmt.Println("Apakah Anda ingin mendapatkan diskon spesial? Yuk, cek apakah Anda memiliki kode voucher yang dapat digunakan!(ya/tidak)")
			fmt.Scan(&vocer)
			for vocer != "ya" && vocer != "tidak" {
				fmt.Println("masukkan input yang valid")
				fmt.Scan(&vocer)

			}

			if vocer == "ya" {
				fmt.Println("Masukkan kode voucher")
				fmt.Scan(&kode)
				for kode != "PASTIHEMAT" && kode != "AWALTAHUN" && kode != "QWERTYUIOP" && kode != "tidak" {
					fmt.Println("Kode salah, Masukkan kode yang valid! Jika tidak memiliki voucher ketik 'tidak'")
					fmt.Scan(&kode)
				}
				if kode == "PASTIHEMAT" {
					harga1 = harga1 * 80 / 100
				} else if kode == "AWALTAHUN" {
					harga1 = harga1 * 90 / 100
				} else if kode == "QWERTYUIOP" {
					harga1 = harga1 * 70 / 100
				}
			}
			fmt.Println("")
			fmt.Println("Total yang harus dibayar : Rp", harga1, ", Apakah kamu ingin menambah camilan?(ya/tidak)")
			fmt.Scan(&camilan)
			for camilan != "ya" && camilan != "Ya " && camilan != "tidak" && camilan != "Tidak" {
				fmt.Println("Input invalid!")
				fmt.Println("Masukkan input yang sesuai")
				fmt.Scan(&camilan)
			}
			for camilan == "ya" || camilan == "Ya" {
				fmt.Println("")
				fmt.Println("                                           ██▄▄")
				fmt.Println("                                           ██▀▀")
				fmt.Println("        LOADING....                      ▄███▄")
				fmt.Println("                                       ▄█████")
				fmt.Println("                drap     drap     ▀▄▄▀▀  █▄ █▄")
				fmt.Println("")
				fmt.Println("")
				fmt.Println("                                                          ▄▀▄     ▄▀▄")
				fmt.Println("                                                         ▄█░░▀▀▀▀▀░░█▄")
				fmt.Println("                                                     ▄▄  █░░░░░░░░░░░█  ▄▄")
				fmt.Println("                                                    █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
				fmt.Println("==========================================================================================================================================")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("||  No  ||      Menu Makanan        ||  No  ||        Menu Minuman       ||  No  ||                      Menu Kombo                     ||")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("==========================================================================================================================================")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("||  1   || Salted Popcorn     20k   ||  6   ||  Milo Ice           10K   ||  11  ||  SAPOPMI (Salted Popcorn + Mineral Water)      22K  ||")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("==========================================================================================================================================")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("||  2   || Caramel Popcorn    22k   ||  7   ||  Lemon Tea          12K   ||  12  ||  Combo 1 (Fish and Chips + Milo)               30K  ||")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("==========================================================================================================================================")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("||  3   || Fish and Chips     25k   ||  8   ||  Hot Coffee         15K   ||  13  ||  Combo 2 (Chicken Popcorn + Lemon Tea)         30K  ||")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("==========================================================================================================================================")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("||  4   || French Fries       15k   ||  9   ||  Mineral Water       5K   ||  14  ||  Combo 3 (Caramel Popcorn + Milo Ice)          28K  ||")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("==========================================================================================================================================")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("||  5   || Chicken Popcorn    23k   ||  10  ||  Milk Tea           16K   ||  15  ||  Combo 4 (Fish and Chips + Lemon Tea)          35K  ||")
				fmt.Println("||      ||                          ||      ||                           ||      ||                                                     ||")
				fmt.Println("==========================================================================================================================================")
				fmt.Println("pilih momor menu")

				fmt.Scan(&menu)

				if menu == 1 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||        Salted Popcorn      20k       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 20000
				} else if menu == 2 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||        Caramel Popcorn     22k       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 22000
				} else if menu == 3 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||        Fish and Chips      25k       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 25000
				} else if menu == 4 {
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||        French Fries        15k       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					jumlahpesanan = 0

					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 15000
				} else if menu == 5 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||        Chicken Popcorn     23k       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 23000
				} else if menu == 6 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||         Milo Ice           10K       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 10000
				} else if menu == 7 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||         Lemon Tea          12K       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 12000
				} else if menu == 8 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||         Hot Coffee         15K       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 15000
				} else if menu == 9 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||         Mineral Water       5K       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 5000
				} else if menu == 10 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                        ▄▀▄     ▄▀▄")
					fmt.Println("                       ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                   ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                  █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("===========================================================")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("||      ||         Milk Tea           16K       ||       ||")
					fmt.Println("||      ||                                      ||       ||")
					fmt.Println("===========================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 16000
				} else if menu == 11 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                               ▄▀▄     ▄▀▄")
					fmt.Println("                              ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                          ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                         █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("=============================================================================")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("||      ||   SAPOPMI (Salted Popcorn + Mineral Water)      22K    ||       ||")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("=============================================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 22000
				} else if menu == 12 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                               ▄▀▄     ▄▀▄")
					fmt.Println("                              ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                          ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                         █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("=============================================================================")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("||      ||   Combo 1 (Fish and Chips + Milo)               30K    ||       ||")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("=============================================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 30000
				} else if menu == 13 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                               ▄▀▄     ▄▀▄")
					fmt.Println("                              ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                          ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                         █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("=============================================================================")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("||      ||   Combo 2 (Chicken Popcorn + Lemon Tea)         30K    ||       ||")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("=============================================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 30000
				} else if menu == 14 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                               ▄▀▄     ▄▀▄")
					fmt.Println("                              ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                          ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                         █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("=============================================================================")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("||      ||   Combo 3 (Caramel Popcorn + Milo Ice)          28K    ||       ||")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("=============================================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 28000
				} else if menu == 15 {

					jumlahpesanan = 0
					fmt.Println(" ")
					fmt.Println("                               ▄▀▄     ▄▀▄")
					fmt.Println("                              ▄█░░▀▀▀▀▀░░█▄")
					fmt.Println("                          ▄▄  █░░░░░░░░░░░█  ▄▄")
					fmt.Println("                         █▄▄█ █░░▀░░┬░░▀░░█ █▄▄█")
					fmt.Println("=============================================================================")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("||      ||   Combo 4 (Fish and Chips + Lemon Tea)          35K    ||       ||")
					fmt.Println("||      ||                                                        ||       ||")
					fmt.Println("=============================================================================")
					fmt.Println(" ")
					fmt.Println("Masukkan jumlah pesanan anda")
					fmt.Scan(&jumlahpesanan)

					harga = jumlahpesanan * 35000
				} else if menu == 0 {
					harga = jumlahpesanan * 0
				}
				for !(jumlahpesanan >= 0 && jumlahpesanan <= 999) {
					harga = jumlahpesanan * -1
					fmt.Println("Input invalid!")
					fmt.Println("Masukkan input yang sesuai")
					fmt.Scan(&jumlahpesanan)
				}
				if jumlahpesanan >= 1 && jumlahpesanan <= 999 {

					hargatotal += harga
					fmt.Println("Harga camilan yang harus dibayar sejumlah ", hargatotal)
					fmt.Println("Apakah anda ingin menambah pesanan?(ya/tidak)")
					fmt.Scan(&camilan)
					for camilan != "ya" && camilan != "tidak" {
						fmt.Println("Input invalid!")
						fmt.Println("Masukkan input yang sesuai")
						fmt.Scan(&camilan)
					}
				} else {
					for !(jumlahpesanan >= 0 && jumlahpesanan <= 999) {
						fmt.Println("Input invalid!")
						fmt.Println("Masukkan input yang sesuai")
						fmt.Scan(&jumlahpesanan)
					}

				}
			}
		}
	}
	hargaakhir = hargatotal + harga1
	fmt.Println("")
	fmt.Println("                                           ██▄▄")
	fmt.Println("                                           ██▀▀")
	fmt.Println("        LOADING....                      ▄███▄")
	fmt.Println("                                       ▄█████")
	fmt.Println("                drap..   drap...  ▀▄▄▀▀  █▄ █▄")
	fmt.Println("")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Harga yang harus dibayar adalah Rp", hargaakhir)
	fmt.Println("")
	fmt.Println("Selamat menikmati film Anda! Terima kasih telah mempercayai kami dalam perjalanan sinematik Anda. Sampai jumpa di film-film berikutnya!")
	fmt.Println("")
	fmt.Println("                                                                ▀▄▀     ▄     ▄")
	fmt.Println("                                                             ▄███████▄  ▀██▄██▀")
	fmt.Println("                                                           ▄█████▀█████▄  ▄█")
	fmt.Println("                                                           ███████▀████████▀")
	fmt.Println("                                                            ▄▄▄▄▄▄███████▀")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("Created by : ")
	fmt.Println("Reevanza Abel Desta Arifin : 103012330104")
	time.Sleep(6 * time.Second)
}
