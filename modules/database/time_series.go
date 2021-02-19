package database

import (
	"context"
	"strconv"

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

// InfluxDB is a interface provided basic APIs
type InfluxDB interface {
	Config() InfluxDBConfig
	Client() influxdb2.Client
	Close()
	WriteAPI() api.WriteAPI
	WriteAPIBlocking() api.WriteAPIBlocking
	QueryAPI(org string) api.QueryAPI
	Write(s *StatusPoint)
	WriteBlocking(s *StatusPoint) error
}

// influxDB is a InfluxDB connection instance
type influxDB struct {
	config InfluxDBConfig
	client influxdb2.Client
}

func (db *influxDB) Config() InfluxDBConfig {
	return db.config
}

func (db *influxDB) Client() influxdb2.Client {
	return db.client
}

// StatusPoint is a struct to write into InfluxDB
type StatusPoint struct {
	Up         int
	NodeID     int
	EndpointID int
}

// WriteAPI returns the asynchronous, non-blocking, Write client from InfluxDB library
func (db *influxDB) WriteAPI() api.WriteAPI {
	return db.Client().WriteAPI(db.Config().Org, db.Config().Bucket)
}

// WriteAPIBlocking returns the synchronous, blocking, Write client from InfluxDB library
func (db *influxDB) WriteAPIBlocking() api.WriteAPIBlocking {
	return db.Client().WriteAPIBlocking(db.Config().Org, db.Config().Bucket)
}

// Close ensures all ongoing asynchronous write clients finish.
// Also closes all idle connections, in case of HTTP client was created internally.
func (db *influxDB) Close() {
	db.Client().Close()
}

// QueryAPI returns Query client
func (db *influxDB) QueryAPI(org string) api.QueryAPI {
	return db.Client().QueryAPI(db.Config().Org)
}

// Write write a StatusPoint to InflxuDB asyncronously
func (db *influxDB) Write(s *StatusPoint) {
	p := influxdb2.NewPointWithMeasurement("status").
		AddField("up", s.Up).
		AddTag("node", strconv.Itoa(s.NodeID)).
		AddTag("endpoint", strconv.Itoa(s.EndpointID))

	api := db.WriteAPI()
	api.WritePoint(p)
}

// WriteBlocking write a StatusPoint to InfluxDB immediately
func (db *influxDB) WriteBlocking(s *StatusPoint) error {
	p := influxdb2.NewPointWithMeasurement("status").
		AddField("up", s.Up).
		AddTag("node", strconv.Itoa(s.NodeID)).
		AddTag("endpoint", strconv.Itoa(s.EndpointID))

	api := db.WriteAPIBlocking()
	if err := api.WritePoint(context.Background(), p); err != nil {
		return err
	}
	return nil
}

// NewInfluxDBConnection is used to start a InfluxDB connection
func NewInfluxDBConnection(c InfluxDBConfig) (InfluxDB, error) {
	if c.URL == "" || c.Bucket == "" || c.Org == "" {
		return nil, errors.New("InfluxDB Connection Info Incomplete")
	}
	if c.Options != nil {
		return &influxDB{
			config: c,
			client: influxdb2.NewClientWithOptions(c.URL, c.Token, c.Options),
		}, nil

	}
	return &influxDB{
		config: c,
		client: influxdb2.NewClient(c.URL, c.Token),
	}, nil

}
