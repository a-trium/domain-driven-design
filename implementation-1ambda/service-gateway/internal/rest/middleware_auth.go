package rest

import (
	"net/http"
	"github.com/gorilla/sessions"
		"fmt"
	"encoding/json"
		"github.com/pkg/errors"
	e "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/exception"
	"github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/internal/config"
	"strings"
)

func InjectAuthMiddleware(sessionStore *sessions.CookieStore, h http.Handler) http.Handler {
	env := config.Env
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if CORS
		if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
			h.ServeHTTP(w, r)
			return
		}

		// disable swagger-ui based on the flag
		if strings.HasPrefix(r.URL.Path, "/api/docs") && !env.EnableSwaggerUI {
			message := fmt.Sprintf("Not Found: (%s) %s", r.Method, r.URL.Path)
			err := errors.New(message)
			ex := e.NewNotFoundException(err)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(ex.StatusCode())
			json.NewEncoder(w).Encode(ex.ToSwaggerError())
			return
		}

		//// if auth request
		//if r.Method == http.MethodPost && strings.HasPrefix(r.URL.Path, "/api/auth/") {
		//	h.ServeHTTP(w, r)
		//	return
		//}
		//
		//session, _ := sessionStore.Get(r, user.SessionCookieName)
		//if authenticated, _ := user.IsAuthenticated(session); !authenticated {
		//	message := fmt.Sprintf("Not Authenticated: (%s) %s", r.Method, r.URL.Path)
		//	err := errors.New(message)
		//	ex := e.NewUnauthorizedException(err)
		//
		//	w.Header().Set("Content-Type", "application/json")
		//	w.WriteHeader(ex.StatusCode())
		//	json.NewEncoder(w).Encode(ex)
		//	return
		//}

		h.ServeHTTP(w, r)
	})
}
