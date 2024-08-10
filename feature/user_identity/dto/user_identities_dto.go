package dto

type CreateUserIdentityRequest struct {
	UserID      int    `json:"user_id" binding:"required"`
	Number      string `json:"number" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Status      int    `json:"status" binding:"required"`
	ExpiryDate  string `json:"expiry_date" binding:"required"`
	PlaceIssued string `json:"place_issued" binding:"required"`
}

type UpdateUserIdentityRequest struct {
	UserID      int    `json:"user_id"`
	Number      string `json:"number"`
	Type        string `json:"type"`
	Status      int    `json:"status"`
	ExpiryDate  string `json:"expiry_date"`
	PlaceIssued string `json:"place_issued"`
}

type UserIdentityResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Number      string `json:"number"`
	Type        string `json:"type"`
	Status      int    `json:"status"`
	ExpiryDate  string `json:"expiry_date"`
	PlaceIssued string `json:"place_issued"`
}
