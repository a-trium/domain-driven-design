package user

import (
	"net/http"
	"github.com/go-openapi/runtime"
	dto "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagserver/swagapi/auth"
	"github.com/gorilla/sessions"
)

type LoginSessionResponder struct {
	dto.LoginOK
	request *http.Request
	sessionStore sessions.Store
	uid string
}

func NewLoginSessionResponder(responder *dto.LoginOK, r *http.Request, sessionStore sessions.Store, uid string) *LoginSessionResponder {
	return &LoginSessionResponder{
		*responder, r, sessionStore, uid,
	}
}

func (responder *LoginSessionResponder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	r := responder.request

	session, _ := responder.sessionStore.Get(responder.request, "SESSION")
	session.Values["UID"] = responder.uid

	session.Save(r, w)
	responder.LoginOK.WriteResponse(w, p)
}
