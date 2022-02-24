package server

import (
	"log"
	"net/http"
	"os"

	websocket "github.com/gorilla/websocket"
)

var (
	logger *log.Logger
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	logger = log.New(os.Stdout, "web ", log.LstdFlags)
	server := http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}
	server.ListenAndServe()
}
func routes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/foo", foo)
	return r
}
func foo(w http.ResponseWriter, r *http.Request) {
	logger.Println("request to foo")
}
func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
