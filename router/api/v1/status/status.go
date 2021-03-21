package status

import (
	"net/http"
	"time"

	"github.com/TechMinerApps/upmaster/modules/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type State uint

const (
	Up State = iota
	Down
)

const layout = "2006-01-02T15:04:05.000Z"

type WriteEndpointResponse struct {
	WriteEndpointRequest
}

type WriteEndpointRequest struct {
	TimeStamp  string `json:"time_stamp"`
	NodeID     int    `json:"node_id"`
	EndpointID int    `json:"endpoint_id"`
	Up         State  `json:"up" binding:"required,numeric"`
}

func WriteEndpointStatus(c *gin.Context) {
	// swagger:route POST /status Status WriteEndpointStatus
	//    write endpoints status
	// responses:
	//   200: WriteEndpointStatus
	//   400: BadRequestError
	var req WriteEndpointRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	timestamp, err := time.Parse(layout, req.TimeStamp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	influx := c.MustGet("InfluxDB").(database.InfluxDB)
	logger := c.MustGet("Logger").(*logrus.Logger)
	p := database.StatusPoint{
		Time:       timestamp,
		Up:         int(req.Up),
		NodeID:     req.NodeID,
		EndpointID: req.EndpointID,
	}
	influx.Write(&p)
	logger.Info("StatusPoint write success")
}
