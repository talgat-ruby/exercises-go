package main

import (
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/config"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/constants"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/game"
	"log/slog"
	"os"
)

func main() {
	apiConfig := config.New().GetApi()
	preInitFun := preInitFunc(apiConfig.Stop)
	apiConfig.Start(preInitFun)
}

func preInitFunc(internalModifier func()) func() {
	return func() {
		err := game.JoinGame()
		if err != nil {
			slog.Error(fmt.Sprintf("could not join game with host %s: %s", constants.Host, err))
			internalModifier()
			os.Exit(1)
		}
	}
}
