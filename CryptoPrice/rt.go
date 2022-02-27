package CryptoPrice

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type logRT struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l logRT) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "%s | %s | %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	return l.next.RoundTrip(r)
}
