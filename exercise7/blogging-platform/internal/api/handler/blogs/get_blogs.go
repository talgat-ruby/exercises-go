package blogs

import (
	"fmt"
	"net/http"
)

func (b *Blogs) GetBlogs(w http.ResponseWriter, r *http.Request) {
	_ = r.Context()
	b.logger.Debug("in GetBlogs")
	fmt.Println("in GetBlogs")
}
