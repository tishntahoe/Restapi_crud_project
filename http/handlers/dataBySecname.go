package handlers

import (
	"encoding/json"
	"github.com/tishntahoe/redsoft/database"
	"github.com/tishntahoe/redsoft/models"
	"net/http"
)

// GetDataBySecname godoc
// @Summary Получить данные о человеке по фамилии
// @Description Возвращает структуру человека по фамилии (secname), если он найден
// @Tags people
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param secname formData string true "Фамилия человека"
// @Success 202 {object} models.AddingHumanStruct
// @Failure 400 {string} string "Ошибка или человек не найден"
// @Router /getbysecname [post]
func GetDataBySecname(w http.ResponseWriter, r *http.Request) {
	var (
		HumanStruct models.AddingHumanStruct
	)
	secname := GetFromValue(r, "secname")
	database.GetDataBySecnameQueryFunc(&HumanStruct, secname)
	if HumanStruct.Secname == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Человека с такой фамилией не найдено"))
		return
	}

	result, err := json.Marshal(HumanStruct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = PushJsonResponse(w, result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
