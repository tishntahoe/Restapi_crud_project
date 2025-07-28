package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

var patternReg = "^([А-ЯЁA-Z][а-яёa-z]+(?:-[А-ЯЁA-Z][а-яёa-z]+)?)\\s([А-ЯЁA-Z][а-яёa-z]+)(?:\\s((?:[А-ЯЁA-Z][а-яёa-z]+)|(?:[А-ЯЁA-Z]\\.[А-ЯЁA-Z]\\.)))?$"
var patternEmail = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"

type ResponseData struct {
	Message string `json:"message"`
}

func GetFromValue(r *http.Request, key string) string {
	return r.FormValue(key)
}

func FromValuesToMap(mapValues map[string]string, r *http.Request, variables []string) {
	for _, key := range variables {
		mapValues[key] = GetFromValue(r, key)
	}
}

func ResponseFunc(w http.ResponseWriter, msg string) (err error) {
	w.Header().Set("Content-Type", "application/json")
	resp := ResponseData{Message: msg}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusAccepted)
	return
}

func PushJsonResponse(w http.ResponseWriter, res []byte) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
	return nil
}

func RegexCheckNameEmail(mapVariables map[string]string) error {
	if mapVariables["email"] != "" {
		matched, err := regexp.Match(patternEmail, []byte(mapVariables["email"]))
		if err != nil || !matched {
			return errors.New("Емейл не прошел проверку")
		}
	}
	if mapVariables["name"] != "" || mapVariables["secname"] != "" {
		matched, err := regexp.Match(patternReg, []byte(mapVariables["name"]+" "+mapVariables["secname"]))
		fmt.Println("|" + mapVariables["name"] + " " + mapVariables["secname"] + "|")
		if err != nil || !matched {
			return errors.New("Произошла ошибка с именем или же, это имя - не имя вовсе. Попробуй другое...")
		}
	} else {
		return errors.New("Имя или фамилия не прошли проверку, таких имен не существует!")
	}
	return nil
}
