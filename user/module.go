package user

import "net/http"

func Module() http.Handler {
	return Handle(findById())
}
