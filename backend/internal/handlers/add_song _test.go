package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"song/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Mock для внешнего сервиса
type MockService struct {
	mock.Mock
}

func (m *MockService) FetchSongDetails(group, song string) (*models.SongDetail, error) {
	args := m.Called(group, song)
	return args.Get(0).(*models.SongDetail), args.Error(1)
}

// Mock для репозитория
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) SaveSong(db *gorm.DB, song models.Song) error {
	args := m.Called(db, song)
	return args.Error(0)
}

// Тест для ошибки при парсинге JSON
func TestAddSong_InvalidJSON(t *testing.T) {
	router := gin.Default()
	router.POST("/add-song", AddSong(nil))

	// Передаем некорректный JSON (например, без обязательных полей)
	body := []byte(`{"group": "The Beatles"}`) // отсутствует поле "song"
	req, _ := http.NewRequest("POST", "/add-song", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

// Тест для ошибки при получении деталей песни из внешнего API
func TestAddSong_FetchSongDetailsError(t *testing.T) {
	mockService := new(MockService)
	mockService.On("FetchSongDetails", "The Beatles", "Hey Jude").Return(nil, assert.AnError)

	router := gin.Default()
	router.POST("/add-song", AddSong(nil))

	body := models.SongInput{
		Group: "The Beatles",
		Song:  "Hey Jude",
	}
	reqBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/add-song", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Failed to fetch song details")
	mockService.AssertExpectations(t)
}

// Тест для ошибки при неверном формате даты релиза
func TestAddSong_InvalidReleaseDate(t *testing.T) {
	mockService := new(MockService)
	mockService.On("FetchSongDetails", "The Beatles", "Hey Jude").Return(&models.SongDetail{
		ReleaseDate: "invalid-date",
	}, nil)

	router := gin.Default()
	router.POST("/add-song", AddSong(nil))

	body := models.SongInput{
		Group: "The Beatles",
		Song:  "Hey Jude",
	}
	reqBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/add-song", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid release date format")
	mockService.AssertExpectations(t)
}

// Тест для ошибки при сохранении песни в базу данных
func TestAddSong_SaveSongError(t *testing.T) {
	mockService := new(MockService)
	mockRepository := new(MockRepository)

	// Мокаем успешный ответ от внешнего API
	mockService.On("FetchSongDetails", "The Beatles", "Hey Jude").Return(&models.SongDetail{
		ReleaseDate: "2024-01-01",
		Text:        "Some text",
		Link:        "some-link",
	}, nil)

	// Мокаем ошибку при сохранении в базу
	mockRepository.On("SaveSong", mock.Anything, mock.Anything).Return(assert.AnError)

	router := gin.Default()
	router.POST("/add-song", AddSong(nil))

	body := models.SongInput{
		Group: "The Beatles",
		Song:  "Hey Jude",
	}
	reqBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/add-song", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Failed to save song to database")
	mockService.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}

// Тест для успешного добавления песни
func TestAddSong_Success(t *testing.T) {
	mockService := new(MockService)
	mockRepository := new(MockRepository)

	// Мокаем успешный ответ от внешнего API
	mockService.On("FetchSongDetails", "The Beatles", "Hey Jude").Return(&models.SongDetail{
		ReleaseDate: "2024-01-01",
		Text:        "Some text",
		Link:        "some-link",
	}, nil)

	// Мокаем успешное сохранение в базу данных
	mockRepository.On("SaveSong", mock.Anything, mock.Anything).Return(nil)

	router := gin.Default()
	router.POST("/add-song", AddSong(nil))

	body := models.SongInput{
		Group: "The Beatles",
		Song:  "Hey Jude",
	}
	reqBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/add-song", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Song added successfully")
	mockService.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}
