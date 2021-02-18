package auth

import (
	"github.com/gin-gonic/gin"
)

// swagger:parameters LoginCredencials
type AuthLoginRequest struct {
	// in:body
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// swagger:response AuthLoginResponse
type AuthLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login handles login request, receiving username/email and password from request body and verify it
func Login(c *gin.Context) {
	// swagger:route POST /auth/login Auth LoginCredencials
	//    login with username/email and password
	// responses:
	//		200: AuthLoginResponse
	//		401:
}

// Logout handles logout request, revoking user's AccessToken and RefreshToken
func Logout(c *gin.Context) {
	// swagger:route DELETE /auth/logout Auth Logout
	//		logout user
	//	responses:
	//		204:
}
