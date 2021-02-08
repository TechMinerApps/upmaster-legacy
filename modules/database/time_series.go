package database

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/pkg/errors"
)

// InfluxDBConfig is a Config for starting InfluxDB Connection
type InfluxDBConfig struct {
	Token   string
	URL     string
	Org     string
	Bucket  string
	Options *influxdb2.Options
}

// InfluxDB is a InfluxDB connection instance
type InfluxDB struct {
	Config InfluxDBConfig
	Client influxdb2.Client
}

// WriteAPI returns the asynchronous, non-blocking, Write client from InfluxDB library
func (db *InfluxDB) WriteAPI() api.WriteAPI {
	return db.Client.WriteAPI(db.Config.Org, db.Config.Bucket)
}

// WriteAPIBlocking returns the synchronous, blocking, Write client from InfluxDB library
func (db *InfluxDB) WriteAPIBlocking() api.WriteAPIBlocking {
	return db.Client.WriteAPIBlocking(db.Config.Org, db.Config.Bucket)
}

// Close ensures all ongoing asynchronous write clients finish.
// Also closes all idle connections, in case of HTTP client was created internally.
func (db *InfluxDB) Close() {
	db.Client.Close()
}

// QueryAPI returns Query client
func (db *InfluxDB) QueryAPI(org string) api.QueryAPI {
	return db.Client.QueryAPI(db.Config.Org)
}

// NewInfluxDBConnection is used to start a InfluxDB connection
func NewInfluxDBConnection(c InfluxDBConfig) (*InfluxDB, error) {
	var DB InfluxDB
	if c.URL == "" || c.Bucket == "" || c.Org == "" {
		return nil, errors.New("InfluxDB Connection Info Incomplete")
	}
	DB.Config = c
	if c.Options != nil {
		DB.Client = influxdb2.NewClientWithOptions(c.URL, c.Token, c.Options)
	} else {
		DB.Client = influxdb2.NewClient(c.URL, c.Token)
	}
	return &DB, nil
}
