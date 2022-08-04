package templatehdl

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luispfcanales/sse-go/internal/core/ports"
	"github.com/luispfcanales/sse-go/pkg/rendertpl"
)

//HTTPHandlerTemplate is struct to server
type HTTPHandlerTemplate struct {
	serviceUser ports.UserService
	serviceTour ports.TourService
}

//New return HTTPHandlerTemplate
func New(srvUser ports.UserService, srvTour ports.TourService) *HTTPHandlerTemplate {
	return &HTTPHandlerTemplate{
		serviceUser: srvUser,
		serviceTour: srvTour,
	}
}

//SetupRoutes initials routes template
func (hdl *HTTPHandlerTemplate) SetupRoutes(m *mux.Router) {
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hola"))
	})
	m.HandleFunc("/paquetes", hdl.paquetes).Methods("GET")
	m.HandleFunc("/paquetes/{idcard}/{infocard}", hdl.infoCard).Methods("GET")
	m.HandleFunc("/post", hdl.post).Methods("GET")
	m.HandleFunc("/login", hdl.login).Methods("GET")
	m.HandleFunc("/registrar", hdl.register).Methods(http.MethodGet, http.MethodPost)
	m.NotFoundHandler = http.HandlerFunc(notfound)
}
func notfound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not found"))
}

func (hdl *HTTPHandlerTemplate) paquetes(w http.ResponseWriter, r *http.Request) {
	tours := hdl.serviceTour.GetAll()
	rendertpl.RenderPage(w, "reserva", tours)
}
func (hdl *HTTPHandlerTemplate) infoCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idTour := vars["idcard"]
	tour, _ := hdl.serviceTour.GetOneTour(idTour)
	log.Println(idTour)
	switch vars["infocard"] {
	case "information":
		tour.ChangeRender("information")
		rendertpl.RenderPage(w, "paquetes-information", tour)
	case "pictures":
		tour.ChangeRender("pictures")
		rendertpl.RenderPage(w, "paquetes-pictures", tour)
	case "videos":
		tour.ChangeRender("videos")
		rendertpl.RenderPage(w, "paquetes-videos", tour)
	}
}
func (hdl *HTTPHandlerTemplate) post(w http.ResponseWriter, r *http.Request) {
	rendertpl.RenderPage(w, "post", nil)
}
func (hdl *HTTPHandlerTemplate) login(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("no hay cookie")
			rendertpl.RenderPage(w, "login", nil)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error in cookie")
		rendertpl.RenderPage(w, "login", nil)
		return
	}
	log.Println("exits:", c)
	rendertpl.RenderPage(w, "login", nil)
}
func (hdl *HTTPHandlerTemplate) register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println("render a template")
		return
	}
}
