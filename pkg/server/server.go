package server

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"pkg/config"
)

func NewServer(router *chi.Mux) {
	conf := config.Conf
	log.Fatal(http.ListenAndServe(":"+conf.APP_PORT, router))
}

func NewTLSServer(router *chi.Mux) {
	conf := config.Conf
	log.Fatal(http.ListenAndServeTLS(":"+conf.APP_PORT, conf.TLS_CERT, conf.TLS_CERT_KEY, router))
}
