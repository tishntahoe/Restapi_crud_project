package database

import (
	"context"
	"errors"
	"github.com/tishntahoe/redsoft/models"
	"time"
)

var getDataBySecname = "select name,secname,COALESCE(thirdname,'') as thirdname,COALESCE(gender,'') as gender,COALESCE(age,0) as age,COALESCE(nation,'') " +
	"as nation from users where secname = $1"

func GetDataBySecnameQueryFunc(humanStruct *models.AddingHumanStruct, secname string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rows, err := DbConnect.QueryContext(ctx, getDataBySecname, secname)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&humanStruct.Name, &humanStruct.Secname, &humanStruct.Thirdname, &humanStruct.Gender,
			&humanStruct.Age, &humanStruct.Nation); err != nil {
			return err
		}
	}
	defer rows.Close()
	return nil
}

var getAllUsersQuery = "select name,secname,COALESCE(thirdname,''),COALESCE(gender,'')," +
	"COALESCE(age,0),COALESCE(nation,'') from users"

func GetAllUsersQueryFunc() ([]models.AddingHumanStruct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var users []models.AddingHumanStruct
	rows, err := DbConnect.QueryContext(ctx, getAllUsersQuery)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.AddingHumanStruct
		err := rows.Scan(
			&user.Name,
			&user.Secname,
			&user.Thirdname,
			&user.Gender,
			&user.Age,
			&user.Nation,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

var CheckFriendshipQuery = "select name,secname from " +
	"(select COALESCE (nullif(f.user_id1,u.id),user_id2) as iduser from friendship f " +
	"join users u on u.id = f.user_id1  or u.id = f.user_id2 " +
	"where u.name = $1 and u.secname = $2) b " +
	"join users u on b.iduser = u.id"

func CheckFriendshipQueryFunc(name, secname string) ([]models.AddingHumanStruct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rows, err := DbConnect.QueryContext(ctx, CheckFriendshipQuery, name, secname)
	var users []models.AddingHumanStruct
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.AddingHumanStruct
		err := rows.Scan(
			&user.Name,
			&user.Secname,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.New("Данного человека не существует, или у него нет друзей.")
	}
	return users, nil
}
