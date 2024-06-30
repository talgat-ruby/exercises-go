package cli

import (
	"github.com/talgat-ruby/exercises-go/exercise3/problem2/internal/models"
	"github.com/talgat-ruby/exercises-go/exercise3/problem2/run"
)

func do(requestData *models.RequestData) int {
	if requestData == nil {
		return 1
	}
	return run.Do(requestData)
}
