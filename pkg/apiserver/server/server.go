package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/exiaohao/firmwarelibrary/pkg/apiserver/models"
	"github.com/exiaohao/firmwarelibrary/pkg/apiserver/router"

	"github.com/gin-gonic/gin"
)

type InitialOpts struct {
	ListenAddr string
	MySQLDSN   string
}

type Server struct {
	Context    context.Context
	ListenAddr string
	MySQLDSN   string
}

func (s *Server) Initialize(opts InitialOpts) {
	s.ListenAddr = opts.ListenAddr
	s.MySQLDSN = opts.MySQLDSN
	// todo: initialize other fields
	models.ConnectDatabase(s.MySQLDSN)
}

func (s *Server) Run(stopCh <-chan struct{}) {
	app := gin.New()
	router.RegisterRouter(app)

	serv := http.Server{
		Addr:    s.ListenAddr,
		Handler: app,
	}

	go func() {
		if err := serv.ListenAndServe(); err != nil {
			log.Fatal("Server shutdown", err)
		}
	}()

	<-stopCh
	log.Printf("Shutting down server...")
	if err := serv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server shutdown", err)
	}
	time.Sleep(1 * time.Second)
	s.Context.Done()
	log.Fatal("Server shutdown gracefully")
}
