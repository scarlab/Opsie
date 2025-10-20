package bolt

import (
	"fmt"
	"net/http"
	"opsie/def"
	"opsie/pkg/errors"
	"opsie/types"
)


func GetSessionUser(r *http.Request) (types.AuthUser, *errors.Error) {
	userVal := r.Context().Value(def.ContextKeyUser)
	if userVal == nil {
		return types.AuthUser{}, errors.New(http.StatusUnauthorized, "no active session user found")
	}
	user, ok := userVal.(types.AuthUser)
	if !ok {
		return types.AuthUser{}, errors.Internal(fmt.Errorf("invalid user context"))
	}

	return  user, nil
}

func GetSession(r *http.Request) (types.Session, *errors.Error) {
	sessionVal := r.Context().Value(def.ContextKeySession)
	if sessionVal == nil {
		return types.Session{}, errors.New(http.StatusUnauthorized, "no active session found")
	}

	session, ok := sessionVal.(types.Session)
	if !ok {
		return types.Session{}, errors.Internal(fmt.Errorf("invalid session context"))
	}

	return  session, nil
}
