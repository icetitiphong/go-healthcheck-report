package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-healthcheck/pkg/core/health"
	"go-healthcheck/pkg/datamodel"
	healthhandler "go-healthcheck/pkg/health"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	healthService := health.NewService()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/healthcheck/report", healthhandler.NewGetHealthCheckReportHandler(healthService.GetHealthCheckReport).HealthCheckReport).Methods("POST")
	r.Use(LineAccessMiddleware)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))

}

func LineAccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		accessToken, err := ValidateAuth(auth)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		verifyResp, err := VerifyAccessToken(accessToken)

		if err != nil {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			return
		}

		if verifyResp.ClientID == "" {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			return
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func ValidateAuth(auth string) (string, error) {

	splitToken := strings.Split(auth, "Bearer")
	if len(splitToken) != 2 {
		// Error: Bearer token not in proper format
		return "", Error("Error: Bearer token not in proper format")
	}

	auth = strings.TrimSpace(splitToken[1])

	return auth, nil
}

func VerifyAccessToken(accessToken string) (verifyAccessToken *datamodel.VerifyAccessTokenReponse, err error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	response, err := client.Get("https://api.line.me/oauth2/v2.1/verify?access_token=" + accessToken)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(responseData), &verifyAccessToken)
	if err != nil {
		return nil, err
	}

	return verifyAccessToken, nil
}

func Error(text string) error {
	return errors.New(text)
}
