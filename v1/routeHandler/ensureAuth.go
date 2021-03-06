package routeHandler

import (
	"github.com/goincremental/negroni-sessions"
	"go.iondynamics.net/passPad/v1/passpad"
	"go.iondynamics.net/passPad/v1/passpad/account"
	"net/http"
	"time"
)

func ensureAuth(w http.ResponseWriter, req *http.Request) *account.Account {
	acc := ensureAuthNoRedirect(w, req)
	if acc == nil {
		http.Redirect(w, req, "/v1/login", http.StatusFound)
	}
	return acc
}

func ensureAuthNoRedirect(w http.ResponseWriter, req *http.Request) *account.Account {
	session := sessions.GetSession(req)
	user := session.Get("user")
	pass := session.Get("pass")
	if user != nil && pass != nil {
		acc := passpad.AuthAccount(user.(string), pass.(string))
		if acc != nil {
			session.Set("last-access", int32(time.Now().Unix()))
			return acc
		}
	}
	return nil
}
