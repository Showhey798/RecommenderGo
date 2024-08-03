package gateway

import "net/http"

func (g *Gateway) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	g.logger.Info("Health check")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}
