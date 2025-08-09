package database

import (
	"context"
	"errors"
	"github.com/tishntahoe/redsoft/pkg/models"
	"strconv"
	"strings"
	"time"
)

var addHumanQuery = "insert into users (name,secname,thirdname,age,nation,gender) " +
	"values ($1,$2,NULLIF ($3,\"\"),NULLIF ($4,\"\"),NULLIF ($5,\"\"),NULLIF ($6,\"\")) " +
	"ON CONFLICT (name,secname) " +
	"DO UPDATE SET " +
	"thirdname = NULLIF (EXCLUDED.thirdname,\"\")," +
	"age = NULLIF (EXCLUDED.age,\"\")," +
	"nation = NULLIF (EXCLUDED.nation,\"\")," +
	"gender = NULLIF (EXCLUDED.gender,\"\")"

var addEmail_idx = "insert into emails (email,user_id) values($1,$2) ON CONFLICT(email) DO NOTHING"

func InsertAddHuman(data models.AddingHumanStruct, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	result, err := DbConnect.ExecContext(ctx, addHumanQuery, data.Name, data.Secname, data.Thirdname, data.Age, data.Nation, data.Gender)
	if err != nil {
		return err
	}

	if email != "" {
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		_, err = DbConnect.ExecContext(ctx, addEmail_idx, email, id)
		if err != nil {
			return err
		}
	}
	return nil
}

var addEmail_name = "insert into emails (email,user_id) values($1, (select id from users where name = $2 and secname = $3)) ON CONFLICT(email) DO NOTHING"

func InsertEmail(email, name, secname string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	result, err := DbConnect.ExecContext(ctx, addEmail_name, email, name, secname)
	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Такого Человека не существует! Сначала создайте его!")
	}
	if err != nil {
		return err
	}
	return nil
}

func UpdateHumanInfo(data []string, mapVariables map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var (
		updateHumanInfo   = "UPDATE users SET "
		buffInterfaceArgs []interface{}
		setClauses        []string
	)
	for _, field := range data {
		val := mapVariables[field]
		if val != "" {
			buffInterfaceArgs = append(buffInterfaceArgs, val)
			placeholder := "$" + strconv.Itoa(len(buffInterfaceArgs))
			setClauses = append(setClauses, field+" = "+placeholder)
		}
	}
	if len(buffInterfaceArgs) == 0 {
		return errors.New("нет данных для обновления")
	}
	buffInterfaceArgs = append(buffInterfaceArgs, mapVariables["name"], mapVariables["secname"])
	whereClause := " WHERE name = $" + strconv.Itoa(len(buffInterfaceArgs)-1) +
		" AND secname = $" + strconv.Itoa(len(buffInterfaceArgs))
	updateHumanInfo += strings.Join(setClauses, ", ") + whereClause
	result, err := DbConnect.ExecContext(ctx, updateHumanInfo, buffInterfaceArgs...)
	if err != nil {
		return err
	}
	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("такого пользователя не существует (name, secname)")
	}
	return nil
}

var createFriendshipQuery = "insert into friendship (user_id1,user_id2) values " +
	"(MIN((select id from users where name = $1 and secname = $2),(select id from users where name = $3 and secname = $4))," +
	"MAX((select id from users where name = $1 and secname = $2),(select id from users where name = $3 and secname = $4))) ON conflict (user_id1,user_id2) DO NOTHING"

func CreateFriendshipQueryFunc(mapVariables map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := DbConnect.ExecContext(ctx, createFriendshipQuery, mapVariables["name"], mapVariables["secname"], mapVariables["nameFr"], mapVariables["secnameFr"])
	if err != nil {
		return errors.New("Одного или обоих людей не существует! Сначала добавьте его!")
	}
	count, err := result.RowsAffected()
	if count == 0 || err != nil {
		return errors.New("Дружба уже создана!")
	}
	if err != nil {
		return err
	}
	return nil
}
