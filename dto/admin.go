package dto 

type (
	AdminCreateDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AdminLoginDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)