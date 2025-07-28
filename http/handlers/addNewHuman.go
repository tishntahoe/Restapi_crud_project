package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/tishntahoe/redsoft/database"
	"github.com/tishntahoe/redsoft/models"
	"io"
	"net/http"
	"net/url"
	"sync"
)

var sites = [3]string{"agify", "genderize", "nationalize"}
var patternSitesApi = [2]string{"https://api.", ".io?name="}

// AddNewHuman godoc
// @Summary Добавить нового человека
// @Description Добавляет человека по имени, отчеству, фамилии и email(не обязательно, идет как дополнение), получая возраст, пол и нацию из внешних API
// @Tags people
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "Имя"
// @Param secname formData string true "Фамилия"
// @Param thirdname formData string false "Отчество"
// @Param email formData string false "Email"
// @Success 200 {string} string "Человек с именем ... был добавлен"
// @Failure 400 {string} string "Ошибка"
// @Router /add [post]
func AddNewHuman(w http.ResponseWriter, r *http.Request) {
	var variablesVar = []string{"name", "secname", "thirdname", "email"}
	errChan := make(chan error, 3)
	var mapVariables = make(map[string]string)
	var (
		ageStruct      models.Agify
		genderStruct   models.Genderize
		nationalStruct models.Nationalize
		wg             sync.WaitGroup
		err            error
	)
	FromValuesToMap(mapVariables, r, variablesVar)
	unionName := mapVariables["secname"] + " " + mapVariables["name"] + " " + mapVariables["thirdname"]

	err = RegexCheckNameEmail(mapVariables)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	for _, site := range sites {
		wg.Add(1)
		go readingUnmarshalling(&ageStruct, &genderStruct, &nationalStruct, site, unionName, &wg, errChan)
	}
	wg.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			w.Write([]byte("Ошибка при разборе readingUnmarshalling(): " + err.Error()))
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	data := formingData(&ageStruct, &genderStruct, &nationalStruct, mapVariables["name"], mapVariables["secname"], mapVariables["thirdname"])

	err = database.InsertAddHuman(data, mapVariables["email"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = ResponseFunc(w, "Человек с именем "+unionName+" был добавлен!")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	return
}

func formingData(ag *models.Agify, gen *models.Genderize, nat *models.Nationalize, name, secname, thirdname string) models.AddingHumanStruct {
	var data models.AddingHumanStruct
	data.Name = name
	data.Secname = secname
	data.Thirdname = thirdname
	data.Age = ag.Age
	data.Gender = gen.Gender

	if len(nat.Country) != 0 {
		data.Nation = nat.Country[0].Country_id
	}
	return data
}

func readingUnmarshalling(ageStruct *models.Agify, genderStruct *models.Genderize, nationalStruct *models.Nationalize,
	site, unionName string, wg *sync.WaitGroup, ch chan error) {

	postresp, err := http.Get(patternSitesApi[0] + site + patternSitesApi[1] + url.QueryEscape(unionName))
	if err != nil {
		ch <- err
	}
	readedBody, err := io.ReadAll(postresp.Body)
	if err != nil {
		ch <- err
	}
	fmt.Println(postresp)
	fmt.Println(patternSitesApi[0] + site + patternSitesApi[1] + string(unionName))
	switch site {
	case "agify":
		json.Unmarshal(readedBody, &ageStruct)
	case "genderize":
		json.Unmarshal(readedBody, &genderStruct)
	case "nationalize":
		json.Unmarshal(readedBody, &nationalStruct)
	}
	wg.Done()
	ch <- nil
}
