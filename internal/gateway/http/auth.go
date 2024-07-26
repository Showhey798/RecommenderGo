package gateway

type SignupRequest struct {
	Email    string
	Password string
}

type SignupResponse struct {
	Success bool
}
