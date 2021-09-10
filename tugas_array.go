package main

import (
	"fmt"
	"strconv"
)

var id int = 0
var menu int
var next string
var games = []map[string]string{}

func main() {
	Looping()
}

func Looping() {
	Menu()

	if menu == 1 {
		input()
	} else if menu == 2 {
		delete()
	} else if menu == 3 {
		tampil()
	} else if menu == 4 {
		Cari()
	} else if menu == 5 {
		top3()
	} else if menu == 6 {
		menu6()
	}
	fmt.Scan(&next)
	if next == "m" {
		main()
	} else {
		fmt.Println("Selesai")
	}
}

func Menu() {
	fmt.Printf("\t\t\t Daftar Pilihan\n")
	fmt.Println("1. Input data game baru")
	fmt.Println("2. Hapus data game berdasarkan id game")
	fmt.Println("3. Tampilkan seluruh data game beserta jumlah data yang tersimpan dalam list")
	fmt.Println("4. Cari data game berdasarkan nama")
	fmt.Println("5. Menampilkan top 3 game terfavorit")
	fmt.Println("6. Menampilkan seluruh data game dengan rating diatas 4.0")
	fmt.Println("0. keluar")
	fmt.Printf("\n")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scan(&menu)
	fmt.Println("========================")
}

func input() {
	var judul, ketRat, ratStr, idStr string
	var rat float64
	id++

	fmt.Print("Masukkan Judul Game : ")
	fmt.Scan(&judul)
	fmt.Print("Masukkan Rating 1 sampai 5: ")
	fmt.Scan(&rat)

	idStr = strconv.Itoa(id)
	ratStr = strconv.Itoa(int(rat))

	if rat >= 4.0 {
		ketRat = "Good"
	} else if rat <= 4.0 && rat >= 2.0 {
		ketRat = "Average"
	} else if rat < 2.0 {
		ketRat = "Poor"
	}

	var oldGame []map[string]string = games

	var newGame = []map[string]string{
		{"ID": idStr, "Judul Game": judul, "Rating": ratStr, "Keterangan Rating": ketRat},
	}

	games = append(oldGame, newGame...)
	fmt.Println(games)
	fmt.Println("==========================")
	fmt.Printf("ketik m lalu enter untuk kembali ke menu \n atau \n")
	fmt.Printf("ketik 0 lalu enter untuk keluar\n")
}

func tampil() {
	fmt.Printf("Tampilan Seluruh Data Games : \n")

	for _, game := range games {
		fmt.Println(game["ID"], game["Judul Game"], game["Rating"], game["Keterangan Rating"])

	}

	fmt.Println("Jumlah Data Dalam List : ", len(games), "Item")
	fmt.Println("==========================")
	fmt.Printf("ketik m lalu enter untuk kembali ke menu \n atau \n")
	fmt.Printf("ketik 0 lalu enter untuk keluar\n")
}

func delete() {
	var ID int
	if len(games) == 0 {
		fmt.Println("Tidak ada yang bisa dihapus")
	}
	fmt.Print("Masukkan ID untuk dihapus : ")
	fmt.Scan(&ID)

	var before = games[:ID-1]
	var after = games[ID:]

	games = append(before, after...)
	fmt.Println(games)
	fmt.Println("==========================")
	fmt.Printf("ketik m lalu enter untuk kembali ke menu \n atau \n")
	fmt.Printf("ketik 0 lalu enter untuk keluar\n")
}

func Cari() {
	var judul string
	var cari int = 0

	fmt.Print("Cari Judul Game : ")
	fmt.Scan(&judul)

	for _, game := range games {
		if game["Judul Game"] == judul {
			cari++
			fmt.Println(game["ID"], game["Judul Game"], game["Rating"], game["Keterangan Rating"])
		}

	}
	if cari == 0 {
		fmt.Println("Data Tidak Ditemukan")
	}
	fmt.Println("==========================")
	fmt.Printf("ketik m lalu enter untuk kembali ke menu \n atau \n")
	fmt.Printf("ketik 0 lalu enter untuk keluar\n")
}

func top3() {
	if len(games) > 0 {

		var dataTop3 = []map[string]string{}

		fmt.Printf("Top 3 Games : \n")

		for _, game := range games {
			if game["Rating"] == "5" {
				var oldGame []map[string]string = dataTop3

				var newGame = []map[string]string{
					{"ID": game["ID"], "Judul Game": game["Judul Game"], "Rating": game["Rating"], "Keterangan Rating": game["Keterangan Rating"]},
				}

				dataTop3 = append(oldGame, newGame...)
			}
		}

		for _, game := range games {
			if game["Rating"] == "4" {
				var oldGame []map[string]string = dataTop3

				var newGame = []map[string]string{
					{"ID": game["ID"], "Judul Game": game["Judul Game"], "Rating": game["Rating"], "Keterangan Rating": game["Keterangan Rating"]},
				}

				dataTop3 = append(oldGame, newGame...)
			}
		}

		if len(dataTop3) < 3 {
			for _, game := range games {
				if game["Rating"] == "3" {
					var oldGame []map[string]string = dataTop3

					var newGame = []map[string]string{
						{"ID": game["ID"], "Judul Game": game["Judul Game"], "Rating": game["Rating"], "Keterangan Rating": game["Keterangan Rating"]},
					}

					dataTop3 = append(oldGame, newGame...)
				}
			}
		}

		if len(dataTop3) < 3 {
			for _, game := range games {
				if game["Rating"] == "2" {
					var oldGame []map[string]string = dataTop3

					var newGame = []map[string]string{
						{"ID": game["ID"], "Judul Game": game["Judul Game"], "Rating": game["Rating"], "Keterangan Rating": game["Keterangan Rating"]},
					}

					dataTop3 = append(oldGame, newGame...)
				}
			}
		}

		if len(dataTop3) < 3 {
			for _, game := range games {
				if game["Rating"] == "1" {
					var oldGame []map[string]string = dataTop3

					var newGame = []map[string]string{
						{"ID": game["ID"], "Judul Game": game["Judul Game"], "Rating": game["Rating"], "Keterangan Rating": game["Keterangan Rating"]},
					}

					dataTop3 = append(oldGame, newGame...)
				}
			}
		}
		var sort3 int = 1
		for _, top := range dataTop3 {
			if sort3 <= 3 {
				sort3++
				fmt.Println(top["ID"], top["Judul Game"], top["Rating"], top["Keterangan Rating"])
			}
		}

	} else {
		fmt.Println("data game 0, silahkan masukkan data game")
	}
	fmt.Println("==========================")
	fmt.Printf("ketik m lalu enter untuk kembali ke menu \n atau \n")
	fmt.Printf("ketik 0 lalu enter untuk keluar\n")
}

func menu6() {
	fmt.Printf("Data Games Yang Memiliki Rating Diatas 4 : \n")
	var up4rating int = 0

	for _, games := range games {
		if games["Rating"] >= "4" {
			up4rating++
			fmt.Println(games["ID"], games["Judul Game"], games["Rating"], games["Keterangan Rating"])
		}
	}

	if up4rating == 0 {

		fmt.Println("tidak ada data game dengan rating diatas 4")

	}

	fmt.Println("==========================")
	fmt.Printf("ketik m lalu enter untuk kembali ke menu \n atau \n")
	fmt.Printf("ketik 0 lalu enter untuk keluar\n")
}
