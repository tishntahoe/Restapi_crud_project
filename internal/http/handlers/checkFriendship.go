package handlers

import (
	"github.com/tishntahoe/redsoft/database"
	"net/http"
)

// CheckFriendship godoc
// @Summary Проверить друзей пользователя
// @Description Возвращает список друзей пользователя по имени и фамилии
// @Tags friendship
// @Accept application/x-www-form-urlencoded
// @Produce plain
// @Param name formData string true "Имя пользователя"
// @Param secname formData string true "Фамилия пользователя"
// @Success 202 {string} string "Список друзей пользователя"
// @Failure 400 {string} string "Ошибка при получении друзей"
// @Router /checkfriendship [post]
func CheckFriendship(w http.ResponseWriter, r *http.Request) {
	variablesVar := []string{"name", "secname"}
	var mapVariable = make(map[string]string)
	FromValuesToMap(mapVariable, r, variablesVar)

	users, err := database.CheckFriendshipQueryFunc(mapVariable["name"], mapVariable["secname"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	for _, user := range users {
		w.Write([]byte(mapVariable["name"] + " " + mapVariable["secname"] + " дружит с " + user.Name + " " + user.Secname))
	}
	w.WriteHeader(http.StatusAccepted)
	return
}
