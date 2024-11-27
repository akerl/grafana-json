package server

import (
	"fmt"
	"net/http"

	"github.com/akerl/timber/v2/log"
)

var logger = log.NewLogger("grafana-json.server")

type JSONServer struct {
	Host string
	Port int
}

func buildBindAddr(j *JSONServer) string {
	port := 8080
	if j.Port != 0 {
		port = j.Port
	}
	return fmt.Sprintf("%s:%d", j.Host, port)
}

func buildRoutes(j *JSONServer) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", j.getHealth)
	mux.HandleFunc("POST /metrics", j.postMetrics)
	mux.HandleFunc("POST /metric-payload-options", j.postMetricPayloadOptions)
	mux.HandleFunc("POST /query", j.postQuery)
	mux.HandleFunc("POST /variable", j.postVariable)
	mux.HandleFunc("POST /tag-keys", j.postTagKeys)
	mux.HandleFunc("POST /tag-values", j.postTagValues)
	return mux
}

func (j *JSONServer) Start() {
	bind := buildBindAddr(j)
	mux := buildRoutes(j)
	http.ListenAndServe(bind, mux)
}

func (j *JSONServer) getHealth(w http.ResponseWriter, r *http.Request) {
	logger.DebugMsgf("received request to %s", r.URL.Path)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("PONG"))
}

func (j *JSONServer) postMetrics(w http.ResponseWriter, r *http.Request) {
	logger.DebugMsgf("received request to %s", r.URL.Path)
}

func (j *JSONServer) postMetricPayloadOptions(w http.ResponseWriter, r *http.Request) {
	logger.DebugMsgf("received request to %s", r.URL.Path)
}

func (j *JSONServer) postQuery(w http.ResponseWriter, r *http.Request) {
	logger.DebugMsgf("received request to %s", r.URL.Path)
}

func (j *JSONServer) postVariable(w http.ResponseWriter, r *http.Request) {
	logger.DebugMsgf("received request to %s", r.URL.Path)
}

func (j *JSONServer) postTagKeys(w http.ResponseWriter, r *http.Request) {
	logger.DebugMsgf("received request to %s", r.URL.Path)
}

func (j *JSONServer) postTagValues(w http.ResponseWriter, r *http.Request) {
	logger.DebugMsgf("received request to %s", r.URL.Path)
}
