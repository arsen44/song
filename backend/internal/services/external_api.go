package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"song/internal/models"
)

func FetchSongDetails(group, song string) (*models.SongDetail, error) {
	// Создаем URL с параметрами
	apiURL := fmt.Sprintf("http://external-api-address/info?group=%s&song=%s", group, song)

	// Выполняем GET-запрос
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к API: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API вернул ошибку: %d", resp.StatusCode)
	}

	// Декодируем JSON-ответ
	var songDetail models.SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %v", err)
	}

	return &songDetail, nil
}
