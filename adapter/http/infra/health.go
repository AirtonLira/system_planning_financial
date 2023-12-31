package infra

import (
	"encoding/json"
	"net/http"

	"github.com/AirtonLira/system_planning_financial/model/infra"
)

// Health - Function of the return status about health application
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	statusApp := infra.StatusApp{Status: "Alive"}

	_ = json.NewEncoder(w).Encode(statusApp)

}
