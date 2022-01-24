package main

import (
	"go-healthcheck/pkg/core/health"
	healthhandler "go-healthcheck/pkg/health"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainMustPassWhenHeaderHaveAuthorization(t *testing.T) {
	req, err := http.NewRequest("POST", "/healthcheck/report", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.yzGZDsqV-EtH9QppKFV26pHGLjTehoXP7LUXexiKcotWH-7M7tu0vShL4i6wqawOzni9tHWvysOZnOZWkHw8aKIhyZEqj68FJ2asX0E87idttyODVN3GGjy_a4KZz7s7VZ34THkzwuKiDlZ0d6P0AYd-LNKijkk8wQN_o3IknyQ.DPK1_9VCRXJwqFk-qgjDXPZtmfcvsgPfF-I4KdDSpPg")

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
	}))

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	healthService := health.NewService()
	handler = http.HandlerFunc(healthhandler.NewGetHealthCheckReportHandler(healthService.GetHealthCheckReport).HealthCheckReport)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestValidateAuthMustPassWhenAuthIsFollowFormat(t *testing.T) {
	mockAuth := "Bearer eyJhbGciOiJIUzI1NiJ9.yzGZDsqV-EtH9QppKFV26pHGLjTehoXP7LUXexiKcotWH-7M7tu0vShL4i6wqawOzni9tHWvysOZnOZWkHw8aKIhyZEqj68FJ2asX0E87idttyODVN3GGjy_a4KZz7s7VZ34THkzwuKiDlZ0d6P0AYd-LNKijkk8wQN_o3IknyQ.DPK1_9VCRXJwqFk-qgjDXPZtmfcvsgPfF-I4KdDSpPg"
	expect := "eyJhbGciOiJIUzI1NiJ9.yzGZDsqV-EtH9QppKFV26pHGLjTehoXP7LUXexiKcotWH-7M7tu0vShL4i6wqawOzni9tHWvysOZnOZWkHw8aKIhyZEqj68FJ2asX0E87idttyODVN3GGjy_a4KZz7s7VZ34THkzwuKiDlZ0d6P0AYd-LNKijkk8wQN_o3IknyQ.DPK1_9VCRXJwqFk-qgjDXPZtmfcvsgPfF-I4KdDSpPg"

	accessToken, err := ValidateAuth(mockAuth)

	assert.Equal(t, expect, accessToken)
	assert.Nil(t, err)
}

func TestValidateAuthMustFailWhenAuthIsNotFollowFormat(t *testing.T) {
	mockAuth := "sclkasncklasnclkansclkasc"
	accessToken, err := ValidateAuth(mockAuth)

	assert.NotNil(t, err)
	assert.Empty(t, accessToken)
}

func TestVerifyAccessTokenMustBePassWhenGetClientID(t *testing.T) {
	mockAccessToken := "eyJhbGciOiJIUzI1NiJ9.yzGZDsqV-EtH9QppKFV26pHGLjTehoXP7LUXexiKcotWH-7M7tu0vShL4i6wqawOzni9tHWvysOZnOZWkHw8aKIhyZEqj68FJ2asX0E87idttyODVN3GGjy_a4KZz7s7VZ34THkzwuKiDlZ0d6P0AYd-LNKijkk8wQN_o3IknyQ.DPK1_9VCRXJwqFk-qgjDXPZtmfcvsgPfF-I4KdDSpPg"

	verifyResp, err := VerifyAccessToken(mockAccessToken)

	assert.NotNil(t, verifyResp.ClientID)
	assert.NotNil(t, verifyResp.ExpiresIn)
	assert.NotNil(t, verifyResp.Scope)
	assert.Empty(t, verifyResp.Error)
	assert.Empty(t, verifyResp.ErrorDescription)
	assert.Nil(t, err)
}

func TestVerifyAccessTokenMustBeFailWhenNotGetClientID(t *testing.T) {
	mockAccessToken := "testfail"

	verifyResp, err := VerifyAccessToken(mockAccessToken)

	assert.NotNil(t, verifyResp.Error)
	assert.NotNil(t, verifyResp.ErrorDescription)
	assert.Empty(t, verifyResp.ClientID)
	assert.Empty(t, verifyResp.ExpiresIn)
	assert.Empty(t, verifyResp.Scope)
	assert.Nil(t, err)
}
