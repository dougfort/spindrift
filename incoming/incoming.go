package incoming

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dougfort/spindrift/types"
)

// Listen to an HTTP socket, feeding the
// data to various channels
func Listen(cfg types.Config) error {
	http.HandleFunc("/status", statusHandler)
	return http.ListenAndServe(cfg.Address, nil)
}

func statusHandler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, fmt.Sprintf("%s: status\n", "xxx"))
}
