package log

import (
	"net"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
)

// RestyTracer is a custom tracer for Resty HTTP client, designed to log detailed trace information
// about HTTP requests and responses.
//
// It leverages zerolog for structured logging and captures information such as DNS lookup time,
// connection time, TLS handshake duration, and more.
type RestyTracer struct {
	logger zerolog.Logger
}

// TraceInfo represents detailed timing and connection metrics for an HTTP request.
// It mirrors the trace information provided by Resty's TraceInfo, with fields formatted as strings
// for structured logging purposes.
type TraceInfo struct {
	DNSLookup      string
	ConnTime       string
	TCPConnTime    string
	TLSHandshake   string
	ServerTime     string
	ResponseTime   string
	TotalTime      string
	IsConnReused   bool
	IsConnWasIdle  bool
	ConnIdleTime   string
	RequestAttempt int
	RemoteAddr     net.Addr
}

// NewRestyTracer creates a new instance of RestyTracer.
//
// Parameters:
// - logger: A zerolog.Logger instance to be used for logging.
//
// Returns:
// - A pointer to the newly created RestyTracer.
func NewRestyTracer(logger zerolog.Logger) *RestyTracer {
	return &RestyTracer{logger}
}

// Trace logs detailed information about an HTTP request and its trace details using zerolog.
//
// This method captures the following details:
// - Request URL, method, and status code.
// - Protocol used for the request.
// - Request and response headers.
// - TraceInfo metrics such as DNS lookup time, connection time, and more.
//
// Parameters:
// - r: The Resty Response object containing the HTTP request and trace details.
func (t *RestyTracer) Trace(r *resty.Response) {
	ti := r.Request.TraceInfo()

	t.logger.Trace().
		Str("url", r.Request.URL).
		Str("method", r.Request.Method).
		Int("status", r.StatusCode()).
		Str("proto", r.Proto()).
		Interface("request_header", r.Request.Header).
		Interface("response_header", r.Header()).
		Interface("trace_info", TraceInfo{
			DNSLookup:      ti.DNSLookup.String(),
			ConnTime:       ti.ConnTime.String(),
			TCPConnTime:    ti.TCPConnTime.String(),
			TLSHandshake:   ti.TLSHandshake.String(),
			ServerTime:     ti.ServerTime.String(),
			ResponseTime:   ti.ResponseTime.String(),
			TotalTime:      ti.TotalTime.String(),
			IsConnReused:   ti.IsConnReused,
			IsConnWasIdle:  ti.IsConnWasIdle,
			ConnIdleTime:   ti.ConnIdleTime.String(),
			RequestAttempt: ti.RequestAttempt,
			RemoteAddr:     ti.RemoteAddr,
		}).
		Msg("HTTP request completed")
}
