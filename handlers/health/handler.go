package health

import (
	"io"
	"net/http"
)

func Read(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
