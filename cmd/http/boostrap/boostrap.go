package boostrap

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/luispfcanales/sse-go/internal/core/services/loginsrv"
	"github.com/luispfcanales/sse-go/internal/core/services/toursrv"
	"github.com/luispfcanales/sse-go/internal/core/services/usersrv"
	"github.com/luispfcanales/sse-go/internal/handlers/templatehdl"
	"github.com/luispfcanales/sse-go/internal/handlers/userhdl"
	"github.com/luispfcanales/sse-go/internal/repository/memory"
	"github.com/luispfcanales/sse-go/internal/repository/memtour"
)

//Run execute webserver
func Run() error {
	m := mux.NewRouter()
	m.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	repoUser := memory.New()
	repoTour := memtour.New()

	serviceUser := usersrv.New(repoUser)
	serviceLogin := loginsrv.New(repoUser)
	serviceTour := toursrv.New(repoTour)
	//handle user http
	hdluser := userhdl.New(serviceUser, serviceLogin)
	hdluser.SetupRoutes(m)
	//handle template page
	hdltempl := templatehdl.New(serviceUser, serviceTour)
	hdltempl.SetupRoutes(m)

	port := fmt.Sprintf(":%s", getPort())
	log.Println("run server to port", port)
	return http.ListenAndServe(port, m)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
