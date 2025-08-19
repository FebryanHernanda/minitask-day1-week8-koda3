package internal

import (
	"errors"
	"fmt"
	"strings"
)

func Task4() {
	var nilaiSuhu float64 = 100
	asalSatuan := "C"
	tujuanSatuan := "R"

	results, error := convertTemp(nilaiSuhu, asalSatuan, tujuanSatuan)

	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Printf("%.2f %s = %.2f %s\n", nilaiSuhu, asalSatuan, results, tujuanSatuan)
	}
}

func convertTemp(nilaiSuhu float64, asalSatuan, tujuanSatuan string) (float64, error) {
	asal := strings.ToLower(asalSatuan)
	tujuan := strings.ToLower(tujuanSatuan)

	var celcius float64

	switch asal {
	case "c":
		celcius = nilaiSuhu
	case "f":
		celcius = (nilaiSuhu - 32) * 5.0 / 9.0
	case "r":
		celcius = nilaiSuhu * 5.0 / 4.0
	case "k":
		celcius = nilaiSuhu - 273.5
	default:
		return 0, errors.New("satuan asal tidak dikenal")
	}

	switch tujuan {
	case "c":
		return celcius, nil
	case "r":
		return celcius * 4.0 / 5.0, nil
	case "f":
		return (celcius * 9.0 / 5.0) + 32.0, nil
	case "k":
		return celcius + 273.15, nil
	default:
		return 0, errors.New("satuan tujuan tidak dikenal")
	}
}
