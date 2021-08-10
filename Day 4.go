package main

import "fmt"

type riwayatPendidikan struct {
	school     string
	jurusan    string
	tahunMulai int
	tahunLulus int
}

type person struct {
	firstName string
	lastName  string
	hobi      string
	riwayatPendidikan
}

func main() {
	dataSiswa := person{
		firstName: "Benaya",
		lastName:  "Lisias",
		hobi:      "Makan",
		riwayatPendidikan: riwayatPendidikan{
			school:     "Binus",
			jurusan:    "Computer science",
			tahunMulai: 2013,
			tahunLulus: 2017,
		},
	}

	cobaPointer := &dataSiswa
	cobaPointer.updateName("Beng")
	dataSiswa.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
