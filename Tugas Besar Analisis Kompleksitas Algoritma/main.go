package main

import (
	"fmt"
	"math/rand"
	"time"
)

type calon struct {
	no    int
	nama  string
	suara int
}

var calonData []calon

func main() {
	// Input jumlah calon
	var jumlahCalon int
	fmt.Print("Masukkan jumlah calon: ")
	fmt.Scan(&jumlahCalon)

	// Generate data calon secara otomatis
	calonData = make([]calon, jumlahCalon)
	for i := 0; i < jumlahCalon; i++ {
		calonData[i] = calon{no: i + 1, nama: fmt.Sprintf("Calon %d", i+1), suara: 0}
	}

	// Proses pemilihan suara secara acak
	var jumlahPemilih = 10000
	fmt.Printf("Dataset awal diinisialisasi dengan %d calon dan %d pemilih.\n", jumlahCalon, jumlahPemilih)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < jumlahPemilih; i++ {
		pilihan := rand.Intn(len(calonData))
		calonData[pilihan].suara++
	}

	// Urutkan data menggunakan Insertion Sort (Iteratif)
	fmt.Println("\nHasil perhitungan suara menggunakan Insertion Sort:")
	startInsertion := time.Now()
	insertionSort()
	durationInsertion := time.Since(startInsertion)
	for i := 0; i < 10 && i < len(calonData); i++ { // Cetak 10 besar
		fmt.Printf("%d. %s - Suara: %d\n", calonData[i].no, calonData[i].nama, calonData[i].suara)
	}
	fmt.Printf("Waktu eksekusi Insertion Sort: %.6f detik\n", durationInsertion.Seconds())

	// Urutkan data menggunakan Bubble Sort (Rekursif)
	fmt.Println("\nHasil perhitungan suara menggunakan Bubble Sort:")
	startBubble := time.Now()
	recursiveBubbleSort(calonData, len(calonData))
	durationBubble := time.Since(startBubble)
	for i := 0; i < 10 && i < len(calonData); i++ { // Cetak 10 besar
		fmt.Printf("%d. %s - Suara: %d\n", calonData[i].no, calonData[i].nama, calonData[i].suara)
	}
	fmt.Printf("Waktu eksekusi Bubble Sort: %.6f detik\n", durationBubble.Seconds())
}

func insertionSort() {
	for i := 1; i < len(calonData); i++ {
		key := calonData[i]
		j := i - 1
		for j >= 0 && calonData[j].suara < key.suara {
			calonData[j+1] = calonData[j]
			j--
		}
		calonData[j+1] = key
	}
}

func recursiveBubbleSort(data []calon, n int) {
	if n == 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		if data[i].suara < data[i+1].suara {
			data[i], data[i+1] = data[i+1], data[i]
		}
	}
	recursiveBubbleSort(data, n-1)
}
