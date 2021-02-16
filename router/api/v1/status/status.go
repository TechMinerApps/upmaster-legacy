package status

import "github.com/gin-gonic/gin"

type WriteEndpointResponse struct {
	WriteEndpointRequest
}

type WriteEndpointRequest struct {
	NodeID     int
	EndpointID int
	Up         int
}

func WriteEndpointStatus(c *gin.Context) {
	// swagger:route POST /status/{endpoint_id} Status WriteEndpointStatus
	//    abc
	// responses:
	//   200: WriteEndpointStatus

}
