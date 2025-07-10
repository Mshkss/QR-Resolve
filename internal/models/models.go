package models

type ApiEntry struct {
	Mac         string `json:"mac" bson:"mac"`
	RedirectUrl string `json:"redirect_url" bson:"redirect_url"`
}

// LoginRequest defines model for LoginRequest.
type LoginRequest struct {
	Password string `json:"password" bson:"password"`
	Username string `json:"username" bson:"username"`
}

type LoginResponse struct {
	Token *string `json:"token,omitempty" bson:"token,omitempty"`
}
