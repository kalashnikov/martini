package martini

import (
	"github.com/fatih/color"
	"log"
	"net/http"
	"time"
)

// Logger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func Logger() Handler {
	return func(res http.ResponseWriter, req *http.Request, c Context, log *log.Logger) {
		const layout = "Jan 2, 2006 at 3:04pm (MST)"
		start := time.Now()

		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
			}
		}

		rw := res.(ResponseWriter)
		c.Next()

		red := color.New(color.FgRed).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		magenta := color.New(color.FgMagenta).SprintFunc()
		log.Printf("[%v, %v] %v | %s %s from %s\n", blue(start.Format(layout)), cyan(time.Since(start)),
			red(rw.Status()), yellow(req.Method), green(req.URL.RequestURI()), magenta(addr))
	}
}
