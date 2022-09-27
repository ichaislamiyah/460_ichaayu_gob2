package main

//import package
import (
	"fmt"
	"os"
	"strconv"
)

// struc untuk orang
type orang struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// function main
func main() {
	var absensi = []orang{
		{Absen: 1, Nama: ": Mark Lee", Alamat: ": Canada", Pekerjaan: ": Dancer", Alasan: ": Saya ingin belajar website untuk SM Ent."},
		{Absen: 2, Nama: ": Jaemin Na", Alamat: ": Seoul", Pekerjaan: ": OB", Alasan: ": Saya ingin upgrade skill saya"},
		{Absen: 3, Nama: ": Jaehyun Jeong", Alamat: ": Gresik", Pekerjaan: ": Artis Kabupaten", Alasan: ": Saya ingin update CV"},
		{Absen: 4, Nama: ": Taeyong Kim", Alamat: ": Malang", Pekerjaan: ": Sol sepatu", Alasan: ": Biar jago kaya Bjorka"},
		{Absen: 5, Nama: ": Yuta", Alamat: ": Hokkaido", Pekerjaan: ": Yakuza", Alasan: "Saya ingin menguasai Golang secara spontan uhuyyy"},
	}

	var argsRaw = os.Args
	var args = argsRaw[1]
	num, err := strconv.Atoi(args)
	_ = err
	generateParticipants(absensi, num)
}

func generateParticipants(o []orang, num int) {
	if num <= 5 {
		fmt.Println("Data Orang")
		fmt.Println("Absen", o[num-1].Absen)
		fmt.Println("Nama", o[num-1].Nama)
		fmt.Println("Alamat", o[num-1].Alamat)
		fmt.Println("Pekerjaan", o[num-1].Pekerjaan)
		fmt.Println("Alasan", o[num-1].Alasan)
	} else {
		fmt.Println("Tidak ada data untuk dia")
	}
}
