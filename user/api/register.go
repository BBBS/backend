package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/bbbs/backend/user/register"
	"github.com/bbbs/backend/variables"
)

// PostRegister registers a new user
func PostRegister(w rest.ResponseWriter, r *rest.Request) {
	var registerReq map[string]interface{}
	err := r.DecodeJsonPayload(&registerReq)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = register.Register(registerReq)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.(http.ResponseWriter).Write(variables.BlankResponse)
}
