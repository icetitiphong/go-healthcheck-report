package health

import (
	"encoding/csv"
	"go-healthcheck/pkg/datamodel"
	"io"
	"net/http"
	"os"
	"time"
)

type GetHealthCheckReportFunc func() (resp *datamodel.HealthResponse, err error)

type urlStatus struct {
	url    string
	status bool
}

func (s *service) GetHealthCheckReport() (resp *datamodel.HealthResponse, err error) {
	start := time.Now()
	csvPath := "../../test.csv"
	resp = &datamodel.HealthResponse{}
	ch := make(chan urlStatus)
	urlList, err := ReadCsvFile(csvPath)

	if err != nil {
		return resp, err
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

	return resp, nil
}

func ReadCsvFile(filePath string) (urls []string, err error) {

	// Load a csv file.
	f, _ := os.Open(filePath)

	defer f.Close()
	// Create a new reader.
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			return urls, err
		}

		for value := range record {
			urls = append(urls, record[value])
		}
	}

	return urls, nil
}

func checkUrl(url string, c chan urlStatus) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	_, err := client.Get(url)
	if err != nil {
		// The website is down
		c <- urlStatus{url, false}
	} else {
		// The website is up
		c <- urlStatus{url, true}
	}
}
