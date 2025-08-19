package internal

import "fmt"

type dataPendidikan struct {
	namaSekolah string
	jurusan     string
}

type personalData struct {
	nama         string
	foto         string
	umur         uint32
	nomorTelepon string
	status       string
	pendidikan   dataPendidikan
}

func Task1() {
	userData := personalData{
		nama:         "Zack Vengenz",
		foto:         "avatar.jpg",
		umur:         30,
		nomorTelepon: "08123456789",
		status:       "Lajang",
		pendidikan: dataPendidikan{
			namaSekolah: "Koda Tech Academy",
			jurusan:     "Fullstack Developer",
		},
	}

	fmt.Println("Nama : ", userData.nama)
	fmt.Println("Foto : ", userData.foto)
	fmt.Println("Umur : ", userData.umur)
	fmt.Println("Nomor Telepon : ", userData.nomorTelepon)
	fmt.Println("Status : ", userData.status)
	fmt.Println("Sekolah : ", userData.pendidikan.namaSekolah)
	fmt.Println("Jurusan : ", userData.pendidikan.jurusan)
}
