package user

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func Handle(getUser GetUserById) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		user := getUser(id)
		writeResponse(w, user)
	}
}

func writeResponse(w http.ResponseWriter, user *User) {
	bytes, _ := json.Marshal(user)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
