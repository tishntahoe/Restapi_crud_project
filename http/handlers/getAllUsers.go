package handlers

import (
	"encoding/json"
	"github.com/tishntahoe/redsoft/database"
	"github.com/tishntahoe/redsoft/models"
	"net/http"
)

// GetAllUsers godoc
// @Summary Получить всех пользователей
// @Description Возвращает массив всех людей из базы
// @Tags people
// @Accept json
// @Produce json
// @Success 202 {array} models.AddingHumanStruct
// @Failure 400 {string} string "Ошибка при получении пользователей"
// @Router /getallusers [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var (
		allUsersStruct []models.AddingHumanStruct
		err            error
	)
	allUsersStruct, err = database.GetAllUsersQueryFunc()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	result, err := json.Marshal(allUsersStruct)
	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	return
}
