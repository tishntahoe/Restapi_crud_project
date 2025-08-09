package main

import (
	"fmt"
	"github.com/tishntahoe/redsoft/database"
	_ "github.com/tishntahoe/redsoft/docs"
	"github.com/tishntahoe/redsoft/internal/http/listener"
)

// @title redsoft_TEST
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func main() {
	database.ConnectToDatabase()
	defer database.DbConnect.Close()
	fmt.Println("Подключение к БД завершено")
	fmt.Println("Служба listener запущена")
	listener.StartListenHttp()
}
