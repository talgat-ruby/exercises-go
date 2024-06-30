package api

import (
	"log/slog"

	todoController "github.com/talgat-ruby/exercises-go/exercise3/problem1/api/controllers/todo"
	todoControllerRoute "github.com/talgat-ruby/exercises-go/exercise3/problem1/api/controllers/todo"
	todoRepository "github.com/talgat-ruby/exercises-go/exercise3/problem1/api/repositories/todo"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/api/router"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/config"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/mlog"
	todoServices "github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/services/todo"
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/server"
)

func Run() {
	mlog.InitLogger()
	slog.Info("hello from Run")
	appConfig := config.NewConfig("./.env")
	repo := todoRepository.NewRepository(appConfig)
	service := todoServices.NewService(repo)
	controller := todoController.NewController(service)
	controllerRoute := todoControllerRoute.NewControllerRoute(controller)
	appRouter := router.NewRouter(controllerRoute)
	server := server.NewServer(appConfig, appRouter)
	server.Run()
}
