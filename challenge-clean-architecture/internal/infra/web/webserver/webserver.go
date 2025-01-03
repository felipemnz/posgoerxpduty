package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpHandler struct {
	Method     string
	Path       string
	HandleFunc http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      []HttpHandler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []HttpHandler{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) NewHttpHandler(method string, path string, handleFunc http.HandlerFunc) HttpHandler {
	return HttpHandler{Method: method, Path: path, HandleFunc: handleFunc}
}

func (s *WebServer) AddHandler(handler HttpHandler) {
	s.Handlers = append(s.Handlers, handler)
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		s.Router.Method(handler.Method, handler.Path, handler.HandleFunc)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
