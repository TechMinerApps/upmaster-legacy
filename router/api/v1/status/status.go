package status

import (
	"net/http"

	"github.com/TechMinerApps/upmaster/modules/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WriteEndpointResponse struct {
	WriteEndpointRequest
}

type WriteEndpointRequest struct {
	NodeID     int `json:"node_id"`
	EndpointID int `json:"endpoint_id"`
	Up         int `json:"up" binding:"required,numeric"`
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
	influx := c.MustGet("InfluxDB").(database.InfluxDB)
	logger := c.MustGet("Logger").(*logrus.Logger)
	p := database.StatusPoint{
		Up:         req.Up,
		NodeID:     1,
		EndpointID: 1,
	}
	influx.Write(&p)
	logger.Info("StatusPoint write success")
}
