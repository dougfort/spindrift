package rpterr

import (
	"fmt"
	"os"
)

// ReportError reports an internal progam error
func ReportError(err error) {
	// I'm using Fprintf here insted of log to make it really clear that
	// this isn't a log
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
}
