package helpers

import (
	"time"
)

// convert time.Time > Format 01 January 2022 > string
//

func ConvertMonth(t time.Time) string {
	// fmt.Println(t)
	m := t.Format("01")
	var b string
	switch m {
	case "01":
		b = " Januari "
	case "02":
		b = " Februari "
	case "03":
		b = " Maret"
	case "04":
		b = " April "
	case "05":
		b = " Mei "
	case "06":
		b = " Juni "
	case "07":
		b = " Juli "
	case "08":
		b = " Agustus "
	case "09":
		b = " September "
	case "10":
		b = " Oktober "
	case "11":
		b = " November"
	case "12":
		b = " Desember "
	}
	return t.Format("02") + b + t.Format("2006")
}
