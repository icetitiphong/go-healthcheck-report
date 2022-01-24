package health

import (
	"fmt"
	"go-healthcheck/pkg/datamodel"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetHealthCheckReportResponseWhenFoundCsv(t *testing.T) {
	mockCSVPath := "../../../test.csv"

	resp := &datamodel.HealthResponse{}
	start := time.Now()
	ch := make(chan urlStatus)
	urlList, err := ReadCsvFile(mockCSVPath)

	if err != nil {
		fmt.Println(err)
	}

	for _, s := range urlList {
		go checkUrl(s, ch)
	}

	result := make([]urlStatus, len(urlList))
	for i := range result {
		result[i] = <-ch
		if result[i].status {
			resp.Success++
		} else {
			resp.Failure++
		}
	}

	resp.TotalWebsites = len(urlList)
	resp.TotalTime = time.Since(start)

	assert.Equal(t, resp.TotalWebsites, len(urlList))
	assert.Nil(t, err)
}

func TestGetErrorWhenNotFoundCsv(t *testing.T) {
	mockCSVPath := "test.csv"

	_, err := ReadCsvFile(mockCSVPath)

	if err != nil {
		fmt.Println(err)
	}

	assert.NotNil(t, err)
}
