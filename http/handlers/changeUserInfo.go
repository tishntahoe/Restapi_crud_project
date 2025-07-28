package handlers

import (
	"github.com/tishntahoe/redsoft/database"
	"net/http"
)

// ChangeUserInfo godoc
// @Summary Обновить информацию о пользователе
// @Description Обновляет информацию о человеке по имени, отчеству и фамилии
// @Tags people
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "Имя"
// @Param secname formData string true "Фамилия"
// @Param thirdname formData string false "Отчество"
// @Param gender formData string false "Пол"
// @Param age formData string false "Возраст"
// @Param nation formData string false "Национальность"
// @Success 200 {string} string "Информация о человеке ... была обновлена!"
// @Failure 400 {string} string "Ошибка при обновлении"
// @Router /changeUserInfo [post]
func ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	variablesVar := []string{"name", "secname", "thirdname", "gender", "age", "nation"}
	var mapVariables = make(map[string]string)
	FromValuesToMap(mapVariables, r, variablesVar)
	unionName := mapVariables["secname"] + " " + mapVariables["name"] + " " + mapVariables["thirdname"]

	err := database.UpdateHumanInfo(variablesVar, mapVariables)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = ResponseFunc(w, "Информация о человеке  "+unionName+" была обновлена!")
}
