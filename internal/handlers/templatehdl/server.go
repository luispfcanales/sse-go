package templatehdl

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/luispfcanales/sse-go/internal/core/ports"
	"github.com/luispfcanales/sse-go/pkg/middleware"
	"github.com/luispfcanales/sse-go/pkg/rendertpl"
)

type dataTemplate struct {
	ActiveSession bool
	Data          interface{}
}

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
	m.HandleFunc("/", middleware.StateUser(hdl.home))
	m.HandleFunc("/paquetes", hdl.paquetes).Methods("GET")
	m.HandleFunc("/paquetes/{idcard}/{infocard}", hdl.infoCard).Methods("GET")
	m.HandleFunc("/post", hdl.post).Methods(http.MethodGet)
	m.HandleFunc("/login", hdl.login).Methods(http.MethodGet, http.MethodPost)
	m.HandleFunc("/registrar", hdl.register).Methods(http.MethodGet, http.MethodPost)
	m.NotFoundHandler = http.HandlerFunc(notfound)
}
func notfound(w http.ResponseWriter, r *http.Request) {
	rendertpl.RenderPage(w, "notfound", nil)
}

func (hdl *HTTPHandlerTemplate) paquetes(w http.ResponseWriter, r *http.Request) {
	exits, emailSession := hdl.serviceUser.ExistSession(r)
	tours := hdl.serviceTour.GetAll()
	dataTpl := dataTemplate{
		ActiveSession: exits,
	}
	if exits {
		valid := hdl.serviceUser.GmailIsValid(emailSession)
		dataTpl.ActiveSession = valid
	}
	dataTpl.Data = tours
	log.Println(dataTpl)
	rendertpl.RenderPage(w, "reserva", dataTpl)
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
func (hdl *HTTPHandlerTemplate) home(w http.ResponseWriter, r *http.Request) {
	exits, emailSession := hdl.serviceUser.ExistSession(r)
	dataTpl := dataTemplate{
		ActiveSession: exits,
	}
	if exits {
		valid := hdl.serviceUser.GmailIsValid(emailSession)
		dataTpl.ActiveSession = valid
	}
	rendertpl.RenderPage(w, "home", dataTpl)
}
func (hdl *HTTPHandlerTemplate) post(w http.ResponseWriter, r *http.Request) {
	rendertpl.RenderPage(w, "post", nil)
}
func (hdl *HTTPHandlerTemplate) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rendertpl.RenderPage(w, "login", nil)
		return
	}
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		user := hdl.serviceUser.GetUserWithCredentials(email, password)
		if user.ID == "" {
			log.Println(user)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   user.Email,
			Expires: time.Now().Add(3 * time.Minute),
			Path:    "/",
		})
		http.Redirect(w, r, "/paquetes", http.StatusSeeOther)
		return
	}
	log.Println("aqui llego")
	w.WriteHeader(http.StatusNotFound)
	http.NotFound(w, r)
}
func (hdl *HTTPHandlerTemplate) register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println("render a template")
		return
	}
}
