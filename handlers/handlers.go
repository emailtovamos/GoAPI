package handlers

import (
	"encoding/json"
	"github.com/emailtovamos/GoAPI/accounts"
	u "github.com/emailtovamos/GoAPI/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &accounts.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create accounts
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	givenAccount := &accounts.Account{}
	err := json.NewDecoder(r.Body).Decode(givenAccount) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := accounts.Login(givenAccount.Email, givenAccount.Password)
	u.Respond(w, resp)
}

var GetRoles = func(w http.ResponseWriter, r *http.Request) {
	input := &accounts.Input{}
	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request for Getting Roles"))
		return
	}

	resp := getRoles(input)
	u.Respond(w, resp)
}

func getRoles(i *accounts.Input) map[string]interface{} {
	resp := u.Message(true, "Getting roles")
	resp["accounts"] = accounts.Role{
		Subject: "testSubject",
	} // TODO Get roles here by calling Kubernetes API
	return resp
}
