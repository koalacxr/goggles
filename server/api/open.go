package api

import (
	"fmt"
	"net/http"
)

// openFileExplorer opens the system File Explorer to the package provided.
func openFileExplorer(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")

	provider.OpenFileExplorer(name)
	fmt.Fprintf(w, "{}")
}