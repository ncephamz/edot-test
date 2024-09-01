package user

type (
	Token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	Claims struct {
		Id        string
		ExpiredIn int64
	}
)
