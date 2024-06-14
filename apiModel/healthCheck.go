package apiModel

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GeldNetworkMVP/GeldMVPBackend/dtos/responseDtos"
	"github.com/gorilla/context"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	defer context.Clear(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp := responseDtos.HealthCheckResponse{
		Note:    "Geld backend up and running",
		Time:    time.Now().Format("Mon Jan _2 15:04:05 2006"),
		Version: "0",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
