package auth

type authResponse struct {
	ID 					string `json:"id"`
	Token 				string `json:"token"`
	UserName		 	string `json:"user_name"`
	Role				string `json:"role"`
}