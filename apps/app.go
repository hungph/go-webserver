package apps

import (
	"../controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
}

func (curApp *App) Initialize(runType string) {
	os.Setenv("runType", runType)

	curApp.Router = mux.NewRouter()
	curApp.initializeRoutes()
}

func (curApp *App) Run(addr string) {
	err := http.ListenAndServe(addr, curApp.Router)

	if err != nil {
		fmt.Println(err)
	}
}

func (curApp *App) initializeRoutes() {
	curApp.Router.HandleFunc("/v1/user/sign-up", controllers.SignUp).Methods("POST")
	curApp.Router.HandleFunc("/v1/user/sign-in", controllers.SignIn).Methods("POST")
}
