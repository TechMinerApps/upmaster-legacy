package endpoints

import (
	"github.com/TechMinerApps/upmaster/models"
	"github.com/gin-gonic/gin"
)

// swagger:response IndexEndpointResponse
type IndexEndpointResponse struct {
	Endpoints []models.Endpoint `json:"endpoints"`
}

// Index list all endpoints
func Index(c *gin.Context) {
	// swagger:route GET /endpoints IndexEndpoint
	//		list all endpoints, FOR AGENT/ADMIN ONLY
	// responses:
	//		200: IndexEndpointResponse
}

// Store create a new endpoint
func Store(c *gin.Context) {
	// swagger:route POST /endpoints StoreEndpoit Endpoint
	//		create a new endpoint
	// responses:
	//		200: Endpoint
}

// Update update an existing endpoint
func Update(c *gin.Context) {
	// swagger:route PUT /endpoints/{id} UpdateEndpoint Endpoint
	//		update an existing endpoint
	// responses:
	//		200: Endpoint
	//		404:
}

// Destroy delete an existing endpoint
func Destroy(c *gin.Context) {
	// swagger:route DELETE /endpoints/{id} UpdateEndpoint Endpoint
	//		delete an existing endpoint
	// responses:
	//		204:
	//		404:
}
