package listener

import (
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/tishntahoe/redsoft/docs"
	"github.com/tishntahoe/redsoft/http/handlers"
	"net/http"
)

func StartListenHttp() {
	mux := http.NewServeMux()

	mux.HandleFunc("/add", handlers.AddNewHuman)
	mux.HandleFunc("/addwithparams", handlers.AddNewHumanWithParams)
	mux.HandleFunc("/addemail", handlers.EmailAddHandler)
	mux.HandleFunc("/getbysecname", handlers.GetDataBySecname)
	mux.HandleFunc("/getallusers", handlers.GetAllUsers)
	mux.HandleFunc("/changeUserInfo", handlers.ChangeUserInfo)
	mux.HandleFunc("/createfriendship", handlers.CreateFriendshipHandler)
	mux.HandleFunc("/checkfriendship", handlers.CheckFriendship)

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", mux)
}
