package main

import "fmt"

// data
type tagihan struct {
	index        int
	name         string
	nominal      int
	date         int
	status       bool
	statusString string
	kategori     string
}

type list []tagihan

//main function
func main() {
	// data sample
	var data = list{
		{index: 1, name: "sample1", nominal: 500, date: 20},
		{index: 2, name: "sample2", nominal: 600, date: 21},
		{index: 3, name: "sample3", nominal: 400, date: 22},
		{index: 4, name: "sample4", nominal: 700, date: 23},
		{index: 5, name: "sample5", nominal: 50, date: 10},
		{index: 6, name: "sample6", nominal: 130, date: 32},
		{index: 7, name: "sample7", nominal: 230, date: 10},
		{index: 8, name: "sample8", nominal: 34, date: 11},
		{index: 9, name: "sample9", nominal: 100, date: 10},
		{index: 10, name: "sample10", nominal: 40, date: 7},
	}
	note(&data)
	// menu
	fmt.Println("==================")
	fmt.Println("||    SIMTAB    ||")
	fmt.Println("==================")
	var pilihan int = -1
	for pilihan != 7 {
		fmt.Println("=======================================================")
		fmt.Println("1.view 2.edit 3.cari 4.urut 5.Lunas 6.statistik 7.exit")
		fmt.Println("=======================================================")
		fmt.Print("pilihan: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Println("==================")
			fmt.Println("||     LIHAT    ||")
			fmt.Println("==================")
			var pilih int
			fmt.Print("1.semua 2.lunas 3.hutang")
			fmt.Scan(&pilih)
			dataLunas, dataUtang := statistik(data)
			switch pilih {
			case 1:
				read(data)
			case 2:
				read(dataLunas)
				fmt.Println("total lunas:", total(dataLunas))
			case 3:
				read(dataUtang)
				fmt.Println("total lunas:", total(dataUtang))
			}
		case 2:
			fmt.Println("==================")
			fmt.Println("||     EDIT     ||")
			fmt.Println("==================")
			crud(&data)
		case 3:
			var SEARCH int
			fmt.Println("==================")
			fmt.Println("||    SEARCH    ||")
			fmt.Println("==================")
			fmt.Println("pilih cari 1.nama 2.kategori")
			fmt.Scan(&SEARCH)
			switch SEARCH {
			case 1:
				searchName(data)
			case 2:
				searchKategori(data)
			}
		case 4:
			var urutan int
			fmt.Println("==================")
			fmt.Println("||    URUTKAN   ||")
			fmt.Println("==================")
			fmt.Println("pilih urutkan 1.nominal 2.tanggal")
			fmt.Scan(&urutan)
			switch urutan {
			case 1:
				insertionSort(data)
			case 2:
				selectionSort(data)
			}
		case 5:
			fmt.Println("===================")
			fmt.Println("||   PELUNASAN   ||")
			fmt.Println("===================")
			validate(&data)
		case 6:
			fmt.Println("===================")
			fmt.Println("||   STATISTIK   ||")
			fmt.Println("===================")
			lunas,utang := statistik(data)
			fmt.Println("persentase Lunas:", (total(lunas)/total(data))*100, "%")
			fmt.Println("persentase Utang:", (total(utang)/total(data))*100, "%")
		}

	}
}

// crud function
func crud(data *list) {
	// read
	read(*data)
	var pilihan int
	fmt.Println("Menu :")
	fmt.Println("1.tambah 2.edit 3.hapus 4.quit")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		// create
		var newList tagihan
		var found bool
		newList.index = len(*data) + 1
		fmt.Println("masukan nama nominal tanggal:")
		fmt.Scan(&newList.name, &newList.nominal, &newList.date)
		for _, vel := range *data {
			if newList.name == vel.name {
				found = true
			}
		}
		if found {
			fmt.Println("nama sudah ada")
		} else {
			*data = append(*data, newList)
		}
		note(data)
		crud(data)
	case 2:
		// edit
		var found bool
		var edit int
		fmt.Println("masukan index data yang ingin di edit:")
		fmt.Scan(&edit)
		for _, vel := range *data {
			if edit == vel.index {
				found = true
			}
		}
		if found {
			fmt.Println("masukan nama nominal tanggal:")
			fmt.Scan(&(*data)[edit-1].name, &(*data)[edit-1].nominal, &(*data)[edit-1].date)
		} else {
			fmt.Println("data tidak ditemukan")
		}
		note(data)
		crud(data)
	case 3:
		// delete
		var delete int
		var found bool
		fmt.Println("masukan index data yang ingin di hapus:")
		fmt.Scan(&delete)
		for _, vel := range *data {
			if delete == vel.index {
				found = true
				*data = append((*data)[:delete-1], (*data)[delete:]...)
			}
		}
		if found {
			fmt.Println("data berhasil dihapus")
		} else {
			fmt.Println("data tidak ditemukan")
		}
		note(data)
		crud(data)
	}
}

// untuk menampilkan data
func read(data list) {
	fmt.Println("===================================================")
	fmt.Println("no   name   nominal   tanggal bayar  katergori  status")
	fmt.Println("---------------------------------------------------")
	for _, vel := range data {
		fmt.Println(vel.index, " ", vel.name, " ", vel.nominal, " ", vel.date, " ", vel.kategori, " ", vel.statusString)
		fmt.Println("---------------------------------------------------")
	}
	fmt.Println("===================================================")
}

// mencatat status dan kategori
func note(data *list) {
	for i, vel := range *data {
		if vel.status == true {
			(*data)[i].statusString = "Lunas"
		} else {
			(*data)[i].statusString = "Utang"
		}
		if vel.nominal >= 500 {
			(*data)[i].kategori = "Besar"
		} else if vel.nominal >= 100 {
			(*data)[i].kategori = "Menengah"
		} else {
			(*data)[i].kategori = "Kecil"
		}
	}
}

// Sequential
func searchKategori(data list) {
	var kategori string
	var found bool
	fmt.Print("masukan kategori (Ringan/Menengah/Berat):")
	fmt.Scan(&kategori)
	for _, vel := range data {
		if kategori == vel.kategori {
			found = true
			fmt.Println(vel.index, " ", vel.name, " ", vel.nominal, " ", vel.date, " ", vel.kategori, " ", vel.statusString)
		}
	}
	if found == false {
		fmt.Println("data tidak ditemukan")
	} else {
		fmt.Println("data ditemukan")
	}
}

// Binary Search
func searchName(data list) {
	var found bool = false
	var X string
	fmt.Print("search name:")
	fmt.Scan(&X)
	var med int
	var kr int = 0
	var kn int = len(data) - 1

	// selection data name untuk mengurutkan alfabet
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j].name < data[min].name {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}

	for kr <= kn && !found {
		med = (kr + kn) / 2
		if X == data[med].name {
			fmt.Println(data[med].index, " ", data[med].name, " ", data[med].nominal, " ", data[med].date, " ", data[med].kategori, " ", data[med].statusString)
			found = true
		}
		if X > data[med].name {
			kr = med + 1
		} else if X < data[med].name {
			kn = med - 1
		}
	}
	if found == true {
		fmt.Println("data ditemukan")
	} else {
		fmt.Println("data tidak ditemukan")
	}
}

// selectionSort
func selectionSort(data list) {
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j].date < data[min].date {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
	read(data)
}

// insertionSort
func insertionSort(data list) {
	for i := 0; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j].nominal > key.nominal {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
	read(data)
}

// validate lunas
func validate(data *list) {
	var found bool
	var index int
	fmt.Print("masukan index data yang dilunasi:")
	fmt.Scan(&index)
	for i, vel := range *data {
		if index == vel.index {
			(*data)[i].status = true
			found = true
		}
	}
	if found {
		fmt.Println("Berhasil")
	} else {
		fmt.Println("data tidak di temukan")
	}
	note(data)
}

// manghitung statistik
func statistik(data list) (list, list) {
	var dataLunas, dataUtang list
	for _, vel := range data {
		if vel.status {
			dataLunas = append(dataLunas, vel)
		} else {
			dataUtang = append(dataUtang, vel)
		}
	}
	return dataLunas, dataUtang
}

// menghitung total
func total(data list) float32 {
	var total float32
	for _, vel := range data {
		total += float32(vel.nominal)
	}
	return total
}
