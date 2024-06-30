package cli

import "github.com/talgat-ruby/exercises-go/exercise3/problem2/internal/mlog"

func Run() int {
	mlog.InitLogger()
	rd, code := parseFlag()
	if code != 0 || rd == nil {
		return code
	}
	return do(rd)
}
