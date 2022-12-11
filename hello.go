package hello

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

type HelloWorldMsg struct {
	Msg string `json:"msg"`
}

func init() {
	functions.HTTP("HelloGET", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	msgJson := HelloWorldMsg{
		Msg: "Olá, está é a minha primeira função sem servidor",
	}

	if err := json.NewEncoder(w).Encode(&msgJson); err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Erro ao executar a função: %v", err.Error())
		return
	}

}
