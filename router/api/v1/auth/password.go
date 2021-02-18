package auth

import "github.com/gin-gonic/gin"

// swagger:parameters ResetTokenRequest
type AuthResetTokenRequest struct {
	Email string `json:"email"`
}

// swagger:parameters ResetPasswordRequest
type AuthResetPasswordRequest struct {
	Token string `json:"token"`
}

// SendResetToken sends reset password token to user's email
func SendResetToken(c *gin.Context) {
	// swagger:route POST /auth/reset Auth ResetTokenRequest
	//		send reset token to email
	// responses:
	//		204:
	//		500:
}

// ResetPassword reset password of user
func ResetPassword(c *gin.Context) {
	// swagger:route PUT /auth/reset Auth ResetPasswordRequest
	//		reset password of curresponding user of the token
	// responses:
	//		204:
	//		404:
}
