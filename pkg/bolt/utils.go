package bolt

import (
	"fmt"
	"net/http"
	"opsie/core/models"
	"opsie/def"
	"opsie/pkg/errors"
)


func GetSessionUser(r *http.Request) (models.AuthUser, *errors.Error) {
	userVal := r.Context().Value(def.ContextKeyUser)
	if userVal == nil {
		return models.AuthUser{}, errors.New(http.StatusUnauthorized, "no active session user found")
	}
	user, ok := userVal.(models.AuthUser)
	if !ok {
		return models.AuthUser{}, errors.Internal(fmt.Errorf("invalid user context"))
	}

	return  user, nil
}

func GetSession(r *http.Request) (models.Session, *errors.Error) {
	sessionVal := r.Context().Value(def.ContextKeySession)
	if sessionVal == nil {
		return models.Session{}, errors.New(http.StatusUnauthorized, "no active session found")
	}

	session, ok := sessionVal.(models.Session)
	if !ok {
		return models.Session{}, errors.Internal(fmt.Errorf("invalid session context"))
	}

	return  session, nil
}
