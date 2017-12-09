package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/v1/users", getAllUsersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/users/{name}", getUserByNameHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/users", userRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/users/{name}", deleteUserByNameHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/v1/users/{name}", updateUserByNameHandler(formatter)).Methods("PATCH")

	mx.HandleFunc("/v1/meetings/{title}", getMeetingByTitleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/meetings", createMeetingHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/meetings/{title}", deleteMeetingByTitleHandler(formatter)).Methods("DELETE")
	// mx.HandleFunc("/service/userinfo", getUserInfoHandler(formatter)).Methods("GET")

}
