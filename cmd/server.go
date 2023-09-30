package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

const CtxCorrelationIdName = "x-correlationId"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	r := mux.NewRouter()
	r.Use(correlationMiddleware)
	r.PathPrefix("/").HandlerFunc(s.handleRequest)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Err(err).Msgf("Failed to listen and serve on %s", ":8080")
	}
}

func correlationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationId := r.Header.Get("X-Correlation-ID")
		if correlationId == "" {
			correlationId = "very_unique"
		}

		ctx := context.WithValue(r.Context(), CtxCorrelationIdName, correlationId)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) handleRequest(_ http.ResponseWriter, r *http.Request) {
	log.Info().Ctx(r.Context()).Msg("Heeeeey, you")
}
