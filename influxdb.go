package main

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/influxdata/influxdb/client/v2"
)

func newInfluxDBClient(config *config) (*client.Client, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:%d", config.InfluxDBHost, config.InfluxDBPort),
		Username: config.InfluxDBUserName,
		Password: config.InfluxDBPassword,
	})
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func createDatabase(clnt *client.Client, DBName string) error {
	query := client.Query{
		Command:  fmt.Sprintf("CREATE DATABASE %s", DBName),
		Database: DBName,
	}
	resp, err := (*clnt).Query(query)
	if err != nil {
		return err
	}
	if resp.Error() != nil {
		return resp.Error()
	}
	return nil
}

func report(result *Result) error {
	dbClient, err := newInfluxDBClient(defaultConfig)
	if err != nil {
		return err
	}

	logrus.Infof("%#v", result)

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  defaultConfig.InfluxDBName,
		Precision: "h",
	})
	if err != nil {
		return err
	}

	// Create a point and add to batch
	tags := map[string]string{"repo": result.Repo.Name}
	fields := map[string]interface{}{
		"ErrBreakDependInfo": result.ErrBreakDependInfo,
		"ErrBreakDepends":    result.ErrBreakDepends,
		"ErrInvalidBinary":   result.ErrInvalidBinary,
		"ErrMismatchBinary":  result.ErrMismatchBinary,
		"ErrNoSrc":           result.ErrNoSrc,
		"ErrNoSrc2":          result.ErrNoSrc2,
	}

	pt, err := client.NewPoint("repo_errors", tags, fields, time.Now())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	// Write the batch
	return (*dbClient).Write(bp)
}
