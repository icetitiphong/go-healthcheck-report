package healthhandler

import (
	"encoding/json"
	"fmt"
	"go-healthcheck/pkg/core/health"
	"net/http"
)

type getHealthCheckReportHandler struct {
	getHealthCheckReportFunc health.GetHealthCheckReportFunc
}

func NewGetHealthCheckReportHandler(getHealthCheckReportFunc health.GetHealthCheckReportFunc) *getHealthCheckReportHandler {
	return &getHealthCheckReportHandler{
		getHealthCheckReportFunc: getHealthCheckReportFunc,
	}
}

func (h *getHealthCheckReportHandler) HealthCheckReport(w http.ResponseWriter, r *http.Request) {
	resp, err := h.getHealthCheckReportFunc()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Print operation result to console
	fmt.Println("Checked websites: ", resp.TotalWebsites)
	fmt.Println("Successful websites: ", resp.Success)
	fmt.Println("Failure websites: ", resp.Failure)
	fmt.Println("Total times to finished checking website: ",
		resp.TotalTime)

	//convert struct to JSON
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
