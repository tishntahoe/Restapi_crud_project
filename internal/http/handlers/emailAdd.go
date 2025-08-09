package handlers

import (
	"github.com/tishntahoe/redsoft/database"
	"net/http"
)

// EmailAddHandler godoc
// @Summary Привязать email к человеку
// @Description Добавляет email к человеку по имени и отчеству
// @Tags people
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Email"
// @Param name formData string true "Имя"
// @Param secname formData string true "Фамилия"
// @Success 200 {string} string "Емейл: ... присоединен к человеку: ..."
// @Failure 400 {string} string "Ошибка"
// @Router /addemail [post]
func EmailAddHandler(w http.ResponseWriter, r *http.Request) {
	variablesVar := []string{"email", "name", "secname"}
	var mapVariable = make(map[string]string)
	FromValuesToMap(mapVariable, r, variablesVar)

	err := database.InsertEmail(mapVariable["email"], mapVariable["name"], mapVariable["secname"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = ResponseFunc(w, "Емейл: "+mapVariable["email"]+" присоединен к человеку: "+mapVariable["name"]+" "+mapVariable["secname"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	return
}
