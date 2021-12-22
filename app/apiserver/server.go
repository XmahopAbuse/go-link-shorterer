package apiserver

import (
	"github.com/gorilla/mux"
	"go-link-shorterer/app/store"
	"io"
	"net/http"
)

type APIserver struct {
	bindAddr string
	router *mux.Router
	store *store.Store
}

func New(bindAddr string) *APIserver {
	return &APIserver{
		bindAddr: bindAddr,
		router: mux.NewRouter(),
		store: store.New("host=localhost dbname=link-shorterer sslmode=disable password=yvxn6akk user=xmahop"),
	}
}

func (s *APIserver) Start() error {
	s.configureRouter()
	err := s.store.Open()
	if err != nil {
		return err
	}
	return http.ListenAndServe(s.bindAddr, s.router)
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/test", s.indexHandler())
}

func (s *APIserver) indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "its test")
	}
}

