package g

import (
	"github.com/toolkits/file"
	"net/http"
	"path/filepath"
	"strings"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			if !file.IsExist(filepath.Join(Root, "/public", r.URL.Path, "index.html")) {
				http.NotFound(w, r)
				return
			}
		}
		http.FileServer(http.Dir(filepath.Join(Root, "/public"))).ServeHTTP(w, r)
	})
}
