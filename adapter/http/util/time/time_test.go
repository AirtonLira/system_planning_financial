package time

import "testing"

func TestConvertStringToTime(t *testing.T) {
	var convertedToTime = ConvertStringToTime("2023-11-19T10:00:00")

	if convertedToTime.Year() != 2023 || convertedToTime.Month() != 11 || convertedToTime.Day() != 19 {
		t.Errorf("A data informada não é uma data valida!")
	}
}
