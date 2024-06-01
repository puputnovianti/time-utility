package timeutilitypackage_test

import (
	"fmt"
	"testing"
	"time"

	timeutilitypackage "github.com/puputnovianti/time-utility"
	"github.com/stretchr/testify/assert"
)

func TestFormatDate(t *testing.T) {
	date := time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)
	layout := "2006-01-02"
	expected := "2024-06-01"

	if result := timeutilitypackage.FormatDate(date, layout); result != expected {
		t.Errorf("FormatDate(%v, %s) = %s; want %s", date, layout, result, expected)
	}
}

func TestNextWeekend(t *testing.T) {
	result, err := timeutilitypackage.NextWeekday(time.Now(), time.Monday)
	assert.Nil(t, err)
	fmt.Println(result)
}

func TestRecurringEvent(t *testing.T) {
	result, err := timeutilitypackage.RecurringEvent(time.Now(), "3 week")
	if err != nil {
		assert.Nil(t, err)
	}
	fmt.Println(result)
}
