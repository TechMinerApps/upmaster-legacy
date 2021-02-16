package swagger

import "github.com/TechMinerApps/upmaster/router/api/v1/status"

// swagger:response WriteEndpointStatus
type swaggerResponseWriteStatus struct {
	Body status.WriteEndpointResponse
}
