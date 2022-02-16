package helpers

import (
	"gochicoba/models"
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

func BatchItemDate(l []*models.Item) []models.ItemResponse {
	response := []models.ItemResponse{}
	for _, li := range l {
		createdAt := ConvertMonth(li.CreatedAt)
		r := models.ItemResponse{
			ID:          li.ID,
			Name:        li.Name,
			Description: li.Description,
			CreatedAt:   createdAt,
		}
		response = append(response, r)
	}
	return response
}

func SingleItemDate(m *models.Item) models.ItemResponse {
	return models.ItemResponse{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   ConvertMonth(m.CreatedAt),
	}
}
