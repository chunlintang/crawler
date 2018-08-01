package api

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"log"
)

type Server struct {
	container *restful.Container
}

type Servers []Server

func (s *Server) init() {
	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)
	wsContainer.Router(restful.CurlyRouter{})
	s.container = wsContainer
}

func (s *Server) register() {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Path("/gocn")
	ws.Route(ws.GET("/api/v1/tenants/contents/{page}").To(GetContents))

	s.container.Add(ws)
}

func (s *Server) Start() {
	server := &http.Server{Addr: ":8081", Handler: s.container}
	log.Fatal(server.ListenAndServe())
}

func New() *Server {
	server := &Server{container: nil}
	server.init()
	server.register()
	return server
}
