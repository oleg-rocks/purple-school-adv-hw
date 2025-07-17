package verify

type VerifyRequest struct {
	Email string `json:"email" validate:"required,email"`
}