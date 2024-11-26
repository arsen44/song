package utils

import (
	"fmt"
	"time"
)

// parseReleaseDate преобразует строку даты в формате "dd.MM.yyyy" в time.Time
func ParseReleaseDate(dateStr string) (time.Time, error) {
	// Ожидаемый формат даты
	const layout = "02.01.2006"

	// Разбор строки в time.Time
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %s, expected format is dd.MM.yyyy", dateStr)
	}

	return parsedDate, nil
}
