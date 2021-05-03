package main
import (
	"github.com/gorilla/mux"
	"github.com/noguchidaisuke/clean-architecture/adapter/controller"
	"log"
	"net/http"
)

func main() {
	Start()
}

func Start() {
	router := mux.NewRouter()
	userController := controller.NewUser()
	router.HandleFunc("/", userController.GetUserByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
