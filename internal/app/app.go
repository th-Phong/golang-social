package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"phongtran/go-social/golang-social/internal/config"
	"phongtran/go-social/golang-social/internal/db"
	"phongtran/go-social/golang-social/internal/db/sqlc"
	"phongtran/go-social/golang-social/internal/routes"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Application struct {
	config *config.Config
	router *gin.Engine
	modude []Module
}

type Module interface {
	Routes() routes.Route
}

type ModuleContext struct {
	DB sqlc.Querier
}

func NewApplication(config *config.Config) (*Application, error) {
	r := gin.Default()

	if err := db.InitDB(); err != nil {
		log.Fatal("init db fail", err)
		return nil, err
	}

	ctx := &ModuleContext{
		DB: db.DB,
	}

	modules := []Module{
		NewItemModule(ctx),
	}

	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config: config,
		router: r,
		modude: modules,
	}, nil
}

func (app *Application) Run() error {
	srv := &http.Server{
		Addr:    app.config.ServerAddress,
		Handler: app.router,
	}

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	return nil
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}
	return routeList
}
