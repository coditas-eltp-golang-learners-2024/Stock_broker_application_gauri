package constants

const (
	// ErrorMessageBadRequest represents the error message for HTTP status code 400
	ErrorMessageBadRequest = "Bad request"
	// ErrorMessageUnauthorized represents the error message for HTTP status code 401
	ErrorMessageUnauthorized = "OTP is expired or invalid"
	// SuccessMessageOTPValidated represents the success message for OTP validation
	SuccessMessageOTPValidated = "OTP validated successfully"
	SuccessMessage             = "Password changed successfully"
)

//secret key
var JwtKey = []byte("iqwetrhkjmnffhgh")
