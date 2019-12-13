package main

import (
	"fmt"
	"strings"
)

const N = 10

type employee struct {
	kode     int
	nama     string
	golongan int
	umur     int
	alamat   string
	reward   string
}
type history struct {
	kode  int
	jmasuk  float64
	jkeluar float64
	lembur int
	pulcep int
}
type Arr [N]employee
type Arr1 [N]history
var (
	imax  int
	empl   employee
	hist   history
	tengah int
	ar     Arr
	ar1    Arr1
	x, y, z   int
	rewardmasuk, rewardlembur string
	ctr int = 0
	ctr1 int = 0
	n int = N-1
)

func identity(data *Arr) {
	found := false
	for ctr <= n && !found{
		fmt.Print("Masukan kode pegawai: ")
		fmt.Scanln(&data[ctr].kode)
		if data[ctr].kode==00 {
			break
		}
		fmt.Print("Masukan nama pegawai(Pisahkan antar kata dengan _ ): ")
		fmt.Scanln(&data[ctr].nama)
		fmt.Print("Masukan golongan pegawai: ")
		fmt.Scanln(&data[ctr].golongan)
		fmt.Print("Masukan umur pegawai: ")
		fmt.Scanln(&data[ctr].umur)
		fmt.Print("Masukan alamat pegawai(Pisahkan antar kata dengan _ ): ")
		fmt.Scanln(&data[ctr].alamat)
		fmt.Println(" ")
		ctr++
	}
	fmt.Println(" ")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&x)
	if x==1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x == 3 {
		History(&ar1)
	} else if x == 4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x == 5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 0 {
		menu()
	} else if x == 000 {
		fmt.Println("Close Program")
	}
}

func searchI(data Arr, key int) {

	for pass := 0; pass < ctr-1; pass++ {
		imax = pass
		for i := pass + 1; i < ctr; i++ {
			if data[imax].kode > data[i].kode {
				imax = i
			}
		}
		empl = data[imax]
		data[imax] = data[pass]
		data[pass] = empl
	}
	found := false
	bawah := 0
	atas := ctr
	for bawah < atas && !found {
		tengah = (bawah + atas) / 2
		if key == data[tengah].kode {
			found = true
		} else if key < data[tengah].kode {
			atas = tengah
		} else if key > data[tengah].kode {
			bawah = tengah + 1
		}
	}
	if found == true {
		fmt.Println("Kode pegawai: ", data[tengah].kode, "Nama pegawai: ", data[tengah].nama, "Golongan pegawai: ", data[tengah].golongan, "Umur pegawai: ", data[tengah].umur, "Alamat pegawai: ", data[tengah].alamat, "Reward pegawai: ", data[tengah].reward)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
	fmt.Println(" ")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&x)
	if x == 1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x == 3 {
		History(&ar1)
	} else if x == 4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x == 5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 0 {
		menu()
	} else if x == 000 {
		fmt.Println("Close Program")
	}
}

func History(data1 *Arr1) {
	found:=false
	for  ctr1 <= n && !found {
		fmt.Print("Masukan kode pegawai: ")
		fmt.Scanln(&data1[ctr1].kode)
		if data1[ctr1].kode==00 {
			break
		}
		fmt.Print("Masukan jam masuk pegawai: ")
		fmt.Scanln(&data1[ctr1].jmasuk)
		fmt.Print("Masukan jam keluar pegawai: ")
		fmt.Scanln(&data1[ctr1].jkeluar)
		fmt.Println(" ")

		jam := data1[ctr1].jkeluar - data1[ctr1].jmasuk
		if int(jam) > 8 {
			data1[ctr1].lembur = int(jam) - 8
		} else if int(jam) < 8 {
			data1[ctr1].pulcep = 8 - int(jam)
		}
		ctr1++
	}
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&x)
	if x == 1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x==3 {
		History(&ar1)
	} else if x == 4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x == 5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 0 {
		menu()
	} else if x == 000 {
		fmt.Println("Close Program")
	}
}

func searchH(data1 Arr1, key int) {
	found := false
	pass:=0
	for pass < ctr1 && !found{
		if key == data1[pass].kode {
			found = true
		}
		pass++
	}
	pass--
	if found == true {
		fmt.Println("Kode Pegawai: ", data1[pass].kode, "Jam masuk pegawai: ", data1[pass].jmasuk, "Jam keluar pegawai: ", data1[pass].jkeluar, "Lembur: ", data1[pass].lembur, "Pulang cepat: ", data1[pass].pulcep)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
	fmt.Println(" ")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&x)
	if x == 1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x == 3 {
		History(&ar1)
	} else if x==4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x == 5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 0 {
		menu()
	} else if x == 000 {
		fmt.Println("Close Program")
	}
}

func rwd(data *Arr, data1 *Arr1) {
	for i := 0; i < ctr; i++ {
		if data1[i].jmasuk <= 8.00 && data1[i].jkeluar >= 16.00 {
			rewardmasuk = "menghargai waktu"
		} else {
			rewardmasuk = "null"
		}
		if data1[i].lembur >= 10 {
			rewardlembur = "pekerja keras"
		} else {
			rewardlembur = "null"
		}
		if rewardmasuk !="null" && rewardlembur !="null" {
			data[i].reward = "pekerja keras dan menghargai waktu"
		} else if rewardmasuk !="null" && rewardlembur =="null" {
			data[i].reward = "menghargai waktu"
		} else if rewardmasuk =="null" && rewardlembur !="null" {
			data[i].reward = "pekerja keras"
		} else if rewardmasuk =="null" && rewardlembur =="null" {
			data[i].reward = "No reward"
		}
	}
}

func kerjas(data Arr) {
	rwd(&data, &ar1)
	for i := 0; i < ctr; i++ {
		if strings.Contains(data[i].reward, "pekerja keras") {
			fmt.Println(data[i].nama)
		}
	}
	fmt.Println(" ")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&x)
	if x == 1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x == 3 {
		History(&ar1)
	} else if x == 4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x==5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 0 {
		menu()
	} else if x == 000 {
		fmt.Println("Close Program")
	}
}

func updateI(data *Arr, key int) {
	for pass := 0; pass < ctr-1; pass++ {
		imax = pass
		for i := pass + 1; i < ctr; i++ {
			if data[imax].kode < data[i].kode {
				imax = i
			}
		}
		empl = data[imax]
		data[imax] = data[pass]
		data[pass] = empl
	}
	found := false
	bawah := 0
	atas := ctr
	for bawah < atas && !found {
		tengah = (bawah + atas) / 2
		if key == data[tengah].kode {
			found = true
		} else if key < data[tengah].kode {
			atas = tengah
		} else if key > data[tengah].kode {
			bawah = tengah + 1
		}
	}
	if found == true {
		fmt.Print("Masukan kode pegawai: ")
		fmt.Scanln(&data[tengah].kode)
		fmt.Print("Masukan nama pegawai(Pisahkan antar kata dengan _ ): ")
		fmt.Scanln(&data[tengah].nama)
		fmt.Print("Masukan golongan pegawai: ")
		fmt.Scanln(&data[tengah].golongan)
		fmt.Print("Masukan umur pegawai: ")
		fmt.Scanln(&data[tengah].umur)
		fmt.Print("Masukan alamat pegawai(Pisahkan antar kata dengan _ ): ")
		fmt.Scanln(&data[tengah].alamat)
		fmt.Println(" ")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
	fmt.Println(" ")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&x)
	if x == 1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x == 3 {
		History(&ar1)
	} else if x==4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x == 5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 0 {
		menu()
	} else if x == 000 {
		fmt.Println("Close Program")
	}
	if x == 1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x == 3 {
		History(&ar1)
	} else if x==4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x == 5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 0 {
		menu()
	} else if x == 000 {
		fmt.Println("Close Program")
	}
}

func updateH(data1 *Arr1, key int) {
	found := false
	pass:=0
	for pass < ctr1 && !found{
		if key == data1[pass].kode {
			found = true
		}
		pass++
	}
	pass--
	if found {
		fmt.Print("Masukan kode pegawai: ")
		fmt.Scanln(&data1[pass].kode)
		fmt.Print("Masukan jam masuk pegawai: ")
		fmt.Scanln(&data1[pass].jmasuk)
		fmt.Print("Masukan jam keluar pegawai: ")
		fmt.Scanln(&data1[pass].jkeluar)
		fmt.Println(" ")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
	fmt.Println(" ")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&x)
	if x == 1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x == 3 {
		History(&ar1)
	} else if x==4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x == 5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 0 {
		menu()
	} else if x == 000 {
		fmt.Println("Close Program")
	}
}

func delete(data *Arr, data1 *Arr1, key int) {
	for pass := 0; pass < ctr-1; pass++ {
		imax = pass
		for i := pass + 1; i < ctr; i++ {
			if data[imax].kode > data[i].kode {
				imax = i
			}
		}
		empl = data[imax]
		data[imax] = data[pass]
		data[pass] = empl
	}
	found := false
	bawah := 0
	atas := ctr
	for bawah < atas && !found {
		tengah = (bawah + atas) / 2
		if key == data[tengah].kode {
			found = true
		} else if key < data[tengah].kode {
			atas = tengah
		} else if key > data[tengah].kode {
			bawah = tengah + 1
		}
	}
	if found {
		data[tengah].kode = 0
		data[tengah].nama = "0"
		data[tengah].golongan = 0
		data[tengah].umur = 0
		data[tengah].alamat = "0"
	} else {
		fmt.Println("Data tidak ditemukan")
	}
	found = false
	pass:=0
	for pass < ctr1 && !found{
		if key == data1[pass].kode {
			found = true
		}
		pass++
	}
	pass--
	if found {
		data1[pass].kode = 0
		data1[pass].jmasuk = 0
		data1[pass].jkeluar = 0
		fmt.Println(" ")
	} else{
		fmt.Println("Data tidak ditemukan")
	}
	fmt.Println(" ")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&x)
	if x == 1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x == 3 {
		History(&ar1)
	} else if x==4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x == 5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 0 {
		menu()
	} else if x == 000 {
		fmt.Println("Close Program")
	}
}

func sort() {
	rwd(&ar, &ar1)
	fmt.Println("Menu Sort(Ascending): ")
	fmt.Println("1. Nama", " ", "2. Kode", " ", "3. Umur")
	fmt.Println("0. Menu Program", " ", "000. End Program")
	
	fmt.Println(" ")
	fmt.Print("Pilih sort/menu: ")
	fmt.Scanln(&z)
	if z==1{
		nama(ar)
	} else if z==2 {
		kode(ar)
	} else if z==3 {
		umur(ar)
	} else if z == 0 {
		menu()
	} else if z == 000 {
		fmt.Println("Close Program")
	}
}

func nama(data Arr) {
	for pass := 0; pass < ctr-1; pass++ {
		imax = pass
		for i := pass + 1; i < ctr; i++ {
			if data[imax].nama > data[i].nama {
				imax = i
			}
		}
		empl = data[imax]
		data[imax] = data[pass]
		data[pass] = empl
	}
	for i := 0; i < ctr; i++ {
		fmt.Println("Kode pegawai: ", data[i].kode, "Nama pegawai: ", data[i].nama, "Golongan pegawai: ", data[i].golongan, "Umur pegawai: ", data[i].umur, "Alamat pegawai: ", data[i].alamat, "Reward pegawai: ", data[i].reward)
	}
	fmt.Println(" ")
	fmt.Print("Pilih Sort/menu: ")
	fmt.Scanln(&z)
	if z == 1 {
		nama(ar)
	} else if z == 2 {
		kode(ar)
	} else if z == 3 {
		nama(ar)
	} else if z == 0 {
		menu()
	} else if z == 000 {
		fmt.Println("Close Program")
	}
}

func kode(data Arr) {
	for pass := 0; pass < ctr-1; pass++ {
		imax = pass
		for i := pass + 1; i < ctr; i++ {
			if data[imax].kode > data[i].kode {
				imax = i
			}
		}
		empl = data[imax]
		data[imax] = data[pass]
		data[pass] = empl
	}
	for i := 0; i < ctr; i++ {
		fmt.Println("Kode pegawai: ", data[i].kode, "Nama pegawai: ", data[i].nama, "Golongan pegawai: ", data[i].golongan, "Umur pegawai: ", data[i].umur, "Alamat pegawai: ", data[i].alamat, "Reward pegawai: ", data[i].reward)
	}
	fmt.Println(" ")
	fmt.Print("Pilih Sort/menu: ")
	fmt.Scanln(&z)
	if z == 1 {
		nama(ar)
	} else if z == 2 {
		kode(ar)
	} else if z == 3 {
		umur(ar)
	} else if z == 0 {
		menu()
	} else if z == 000 {
		fmt.Println("Close Program")
	}
}

func umur(data Arr) {
	for pass := 0; pass < ctr-1; pass++ {
		imax = pass
		for i := pass + 1; i < ctr; i++ {
			if data[imax].umur > data[i].umur {
				imax = i
			}
		}
		empl = data[imax]
		data[imax] = data[pass]
		data[pass] = empl
	}
	for i := 0; i < ctr; i++ {
		fmt.Println("Kode pegawai: ", data[i].kode, "Nama pegawai: ", data[i].nama, "Golongan pegawai: ", data[i].golongan, "Umur pegawai: ", data[i].umur, "Alamat pegawai: ", data[i].alamat, "Reward pegawai: ", data[i].reward)
	}
	fmt.Println(" ")
	fmt.Print("Pilih Sort/menu: ")
	fmt.Scanln(&z)
	if z == 1 {
		nama(ar)
	} else if z == 2 {
		kode(ar)
	} else if z == 3 {
		umur(ar)
	} else if z == 0 {
		menu()
	} else if z == 000 {
		fmt.Println("Close Program")
	}
}

func menu() {
	fmt.Println(" ")
	fmt.Println("Menu Program: ")
	fmt.Println("1. Create Identity	 ", " ", "5. Show Hardworker")
	fmt.Println("2. Search Identity 	 ", " ", "6. Show Identity Sorted by name(Ascending)")
	fmt.Println("3. Create History 	 ", " ", "7. Update Identity")
	fmt.Println("4. Show History 	 ", " ", "8. Update History")
	fmt.Println("9. Delete Specific Data", " ", "000. Close Program")
	
	fmt.Println(" ")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&x)
	if x == 1 {
		identity(&ar)
	} else if x == 2 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchI(ar, y)
	} else if x == 3 {
		History(&ar1)
	} else if x == 4 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		searchH(ar1, y)
	} else if x == 5 {
		kerjas(ar)
	} else if x == 6 {
		sort()
	} else if x == 7 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateI(&ar, y)
	} else if x == 8 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		updateH(&ar1, y)
	} else if x == 9 {
		fmt.Print("Masukan keycode: ")
		fmt.Scanln(&y)
		delete(&ar, &ar1, y)
	} else if x == 000 {
		fmt.Println("End Program")
	}
}
func main() {
	menu()
}
