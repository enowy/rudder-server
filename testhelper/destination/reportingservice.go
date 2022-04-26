package destination

import (
	_ "encoding/json"
	"fmt"
	_ "github.com/Shopify/sarama"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"log"
	"database/sql"
)

type  ReportingServiceResource struct {
	ReportingserviceURL string
}

func  SetupReportingService(pool *dockertest.Pool, d deferer) (*ReportingServiceResource, error) {
	// Set  timescale DB
	// pulls an image, creates a container based on it and runs it
	database := "temo"
	timescaleContainer, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "timescale/timescaledb",
		Tag:        "latest-pg13",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=" + database,
			"POSTGRES_PASSWORD=password",
		},
	})
	if err != nil {
		log.Println("Could not start resource Timescale DB: %w", err)
	}
	timescaleDB_DSN_Internal := fmt.Sprintf("postgresql://postgres:password@host.docker.internal:%s/%s?sslmode=disable", timescaleContainer.GetPort("5432/tcp"), database)
	fmt.Println("timesacaleDB_DSN", timescaleDB_DSN_Internal)
	timescaleDB_DSN_viaHost := fmt.Sprintf("postgresql://postgres:password@localhost:%s/%s?sslmode=disable", timescaleContainer.GetPort("5432/tcp"), database)
	if err := pool.Retry(func() error {
		var err error
		rs_db, err := sql.Open("postgres", timescaleDB_DSN_viaHost)
		if err != nil {
			return err
		}
		return rs_db.Ping()
	}); err != nil {
		log.Println("Could not connect to postgres %w", err)
	}
	log.Println("timescaleDB_DSN_viaHost", timescaleDB_DSN_viaHost)

	// Set  reporting service
	// pulls an image, creates a container based on it and runs it
	reportingContainer, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "rudderstack/rudderstack-reporting",
		Tag:          "v1.0.12",
		ExposedPorts: []string{"5000"},
		Env: []string{
			"DATABASE_URL=" + timescaleDB_DSN_Internal,
		},
	})
	if err != nil {
		log.Println("Could not start resource Reporting Service: %w", err)
	}

	reportingserviceURL := fmt.Sprintf("http://localhost:%s", reportingContainer.GetPort("5000/tcp"))
	fmt.Println("reportingserviceURL", reportingserviceURL)

	if err != nil {
		return nil, err
	}

	d.Defer(func() error {
		if err := pool.Purge(reportingContainer); err != nil {
			log.Printf("Could not purge resource: %s \n", err)
		}
		return nil
	})
	return &ReportingServiceResource{
		ReportingserviceURL: reportingserviceURL,
	}, nil
}
