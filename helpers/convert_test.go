package helpers

import (
	"gochicoba/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConvertMonth(t *testing.T) {
	assert.Equal(t, ConvertMonth(time.Now()), "25 Februari 2022", "Sesuaikan tanggal, format Indonesia.")
}

func TestSingleItemDate(t *testing.T) {
	ir := models.Item{
		ID:          1,
		Name:        "TEST1",
		Description: "TEST1",
		CreatedAt:   time.Now(),
		Price:       3000,
	}

	actual := SingleItemDate(&ir)
	expectation := ConvertMonth(time.Now())

	assert.Equal(t, actual.CreatedAt, expectation, "Output must same.")
}

func TestBatchItemDate(t *testing.T) {
	var mI []*models.Item = []*models.Item{
		{ID: 1,
			Name:        "TEST1",
			Description: "TEST1",
			CreatedAt:   time.Now(),
			Price:       3000,
		}, {
			ID:          2,
			Name:        "TEST2",
			Description: "TEST2",
			CreatedAt:   time.Now(),
			Price:       3000,
		},
	}

	moI := BatchItemDate(mI)

	for _, l := range moI {
		assert.Equal(t, l.CreatedAt, ConvertMonth(time.Now()), "Tanggal harus sesuai.")
	}
}
