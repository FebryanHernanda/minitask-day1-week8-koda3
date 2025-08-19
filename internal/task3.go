package internal

import "fmt"

func Task3() {
	var arrPositive [5]uint16 = [5]uint16{10, 50, 25, 3, 17}

	/* Slice Array */
	slice1 := arrPositive[1:3]
	slice2 := arrPositive[2:4]
	slice3 := arrPositive[4:5]

	/* search maxNumber */
	maxNumber := arrPositive[0]
	for i := 0; i < len(arrPositive); i++ {
		if arrPositive[i] > maxNumber {
			maxNumber = arrPositive[i]
		}
	}

	fmt.Println(arrPositive)
	fmt.Println("Nilai indeks ke 1 - 3 : ", slice1)
	fmt.Println("Nilai indeks ke 2 - 4 : ", slice2)
	fmt.Println("Nilai indeks terakhir : ", slice3)
	fmt.Println("Nilai Terbesar adalah : ", maxNumber)
}
