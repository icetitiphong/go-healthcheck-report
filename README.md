# go-healthcheck

## About the project

The API is used to check website and response time when given the CSV list, calls the Healthcheck Report API to send the statistic of each website




## Sample
POST http://localhost:8080/healthcheck/report

Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.yzGZDsqV-EtH9QppKFV26pHGLjTehoXP7LUXexiKcotWH-7M7tu0vShL4i6wqawOzni9tHWvysOZnOZWkHw8aKIhyZEqj68FJ2asX0E87idttyODVN3GGjy_a4KZz7s7VZ34THkzwuKiDlZ0d6P0AYd-LNKijkk8wQN_o3IknyQ.DPK1_9VCRXJwqFk-qgjDXPZtmfcvsgPfF-I4KdDSpPg

Content-Type: application/json


## Installation

```bash
cd go-healthcheck/cmd/health/
go run main.go
```


### Layout

```tree
├── README.md
├── cmd
│   ├── health
│       └── healthcheck_report.go
│       └── usecase_get_heath_check_report_test.go
│       └── usecase_get_heath_check_report.go
├── pkg
│   ├── core
│   │   └── health
│   │       └── README.md
│   ├── datamodel
│   │   └── model.go
│   └── health
│       └── health_check_report_handler.go
├── go.mod
├── go.sum
├── README.md
└── test.csv
```
