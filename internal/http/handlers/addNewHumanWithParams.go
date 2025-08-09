package handlers

import (
	"github.com/tishntahoe/redsoft/database"
	"github.com/tishntahoe/redsoft/pkg/models"
	"net/http"
	"strconv"
)

// AddNewHumanWithParams godoc
// @Summary Добавить нового человека с ручными параметрами
// @Description Добавляет человека, используя параметры, переданные напрямую (без внешних API)
// @Tags people
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "Имя"
// @Param secname formData string true "Фамилия"
// @Param thirdname formData string false "Отчество"
// @Param age formData int false "Возраст"
// @Param nation formData string false "Национальность (напр. US, RU)"
// @Param gender formData string false "Пол (male или female, но другой тоже подойдет:) )"
// @Param email formData string false "Email"
// @Success 200 {string} string "Человек с именем ... был добавлен с собственными параметрами"
// @Failure 400 {string} string "Ошибка"
// @Router /addwithparams [post]
func AddNewHumanWithParams(w http.ResponseWriter, r *http.Request) {

	variablesVar := []string{"name", "secname", "thirdname", "gender", "age", "nation", "email"}
	var mapVariables = make(map[string]string)
	FromValuesToMap(mapVariables, r, variablesVar)
	unionName := mapVariables["secname"] + " " + mapVariables["name"] + " " + mapVariables["thirdname"]

	err := RegexCheckNameEmail(mapVariables)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	age, _ := strconv.Atoi(mapVariables["age"])

	err = database.InsertAddHuman(
		models.AddingHumanStruct{mapVariables["name"], mapVariables["secname"], mapVariables["thirdname"],
			mapVariables["gender"], mapVariables["nation"], age}, mapVariables["email"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = ResponseFunc(w, "Человек с именем "+unionName+" был добавлен с собственными параметрами")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	return
}
