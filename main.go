package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type apiResponse struct {
	Message string		`json:"message"`
	TimeStamp time.Time `json:"time_stamp"`
}

func main() {
	r := mux.NewRouter()
	r.Path("/greetings").Methods(http.MethodGet).HandlerFunc(GreetingsHandler)
	r.Path("/greetings/{name}").Methods(http.MethodGet).HandlerFunc(GreetingsHandler)

	http.ListenAndServe(":8000", r)
}

func GreetingsHandler(writer http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	name, found := variables["name"]
	if !found {
		name = "Generic User"
	}

	resp := fmt.Sprintf("Hello, Dear %s", name)

	fmt.Println(resp)

	jsonResp := apiResponse{
		Message:   resp,
		TimeStamp: time.Now(),
	}

	jsonBuffer, err := json.Marshal(jsonResp)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(jsonBuffer)
	if err != nil {
		fmt.Println(err)
	}
}
