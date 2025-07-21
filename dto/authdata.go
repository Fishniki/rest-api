package dto 

type AuthRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

 type AuthResponse struct {
	Token string `json:"token"`
 }

 type RegisterRequest struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
 }