package handlers

import (
	"github.com/tishntahoe/redsoft/database"
	"net/http"
)

// CreateFriendshipHandler godoc
// @Summary Создать дружбу между двумя пользователями
// @Description Создаёт дружбу между пользователями по имени и фамилии
// @Tags friendship
// @Accept application/x-www-form-urlencoded
// @Produce plain
// @Param name formData string true "Имя первого пользователя"
// @Param secname formData string true "Фамилия первого пользователя"
// @Param nameFr formData string true "Имя второго пользователя"
// @Param secnameFr formData string true "Фамилия второго пользователя"
// @Success 202 {string} string "Дружба между ... создана!"
// @Failure 400 {string} string "Ошибка при создании дружбы"
// @Router /createfriendship [post]
func CreateFriendshipHandler(w http.ResponseWriter, r *http.Request) {
	variablesVar := []string{"name", "secname", "nameFr", "secnameFr"}
	var mapVariable = make(map[string]string)
	FromValuesToMap(mapVariable, r, variablesVar)

	err := database.CreateFriendshipQueryFunc(mapVariable)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Дружба между " + mapVariable["name"] + " " + mapVariable["secname"] + " и " + mapVariable["nameFr"] + " " + mapVariable["secnameFr"] + " создана!"))
	return
}
