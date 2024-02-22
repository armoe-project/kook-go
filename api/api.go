package api

type Api struct {
	token string
}

func NewApi(token string) *Api {
	return &Api{token: token}
}
