package server

import (
	"context"
	"fmt"
	"github.com/AlkBur/web1c/server/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Debug  bool
	router *gin.Engine
}

func New() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) Run(addr string) error {
	s.router.LoadHTMLGlob("./server/templates/*")

	htnlVal := make(map[string]interface{})
	htnlVal["bases"] = make(map[string]string)
	for _, db := range dbs {
		htnlVal["bases"].(map[string]string)[db.Url] = db.Name
	}

	s.router.GET("/", handlers.Index(htnlVal))

	s.router.StaticFile("/favicon.ico", "./server/public/img/favicon.ico")
	s.router.StaticFS("/img/", http.Dir("./server/public/img"))
	s.router.StaticFS("/css/", http.Dir("./server/public/css"))
	s.router.StaticFS("/js/", http.Dir("./server/public/js"))

	if s.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("Server Shutdown: %v", err)
	}
	err := CloseAllDB()
	log.Println("Server exiting")
	return err
}
