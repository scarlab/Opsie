package repo

import (
	"opsie/config"
	"opsie/core/models"
	"opsie/pkg/errors"
	"opsie/pkg/utils"
	"time"

	"gorm.io/gorm"
)

// AuthRepository - Handles DB operations for auth.
// Talks only to the database (or other data sources).
type AuthRepository struct {
	db *gorm.DB
}

// NewAuthRepository - Constructor for Repository
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}


func (r *AuthRepository) CreateSession(userId int64, key string, expiry time.Time) (models.Session, *errors.Error) {
    session := models.Session{
        UserID:   userId,
        Key:      key,
        Expiry:   expiry,
        IsValid:  true,
    }

    if err := r.db.Create(&session).Error; err != nil {
        // Handle Postgres conflict if needed
        return models.Session{}, errors.Internal(err)
    }

    return session, nil
}



func (r *AuthRepository) GetValidSessionByKey(key string) (models.Session, *errors.Error) {
    var session models.Session
    if err := r.db.Where("key = ? AND is_valid = ? AND expiry > NOW()", key, true).First(&session).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return models.Session{}, errors.Unauthorized("invalid session")
        }
        return models.Session{}, errors.Internal(err)
    }
    return session, nil
}





func (r *AuthRepository) GetValidSessionWithAuthUser(key string) (models.SessionWithUser, *errors.Error) {
    var session models.Session
    if err := r.db.Preload("User").
        Where("key = ? AND is_valid = ? AND expiry > NOW()", key, true).
        Joins("JOIN users ON users.id = sessions.user_id AND users.is_active = true").
        First(&session).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return models.SessionWithUser{}, errors.Unauthorized("invalid session")
        }
        return models.SessionWithUser{}, errors.Internal(err)
    }

	authUser := models.AuthUser{
		ID: session.User.ID,
		DisplayName: session.User.DisplayName,
		Email: session.User.Email,
		Avatar: session.User.Avatar,
		SystemRole: session.User.SystemRole,
		IsActive: session.User.IsActive,
		Preference: session.User.Preference,
	}

    return models.SessionWithUser{
        Session:  session,
        AuthUser: authUser,
    }, nil
}


func (r *AuthRepository) ExpireSession(key string) *errors.Error {
    res := r.db.Model(&models.Session{}).
        Where("key = ? AND is_valid = ?", key, true).
        Updates(map[string]interface{}{
            "is_valid": false,
            "expiry":   time.Now(),
        })
    if err := res.Error; err != nil {
        return errors.Internal(err)
    }
    if res.RowsAffected == 0 {
        return errors.NotFound("No active session found")
    }
    return nil
}



func (r *AuthRepository) RegenerateSessionKey(key string) (models.Session, *errors.Error) {
    var session models.Session
    if err := r.db.Where("key = ? AND is_valid = ?", key, true).First(&session).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return models.Session{}, errors.NotFound("session not found")
        }
        return models.Session{}, errors.Internal(err)
    }

    newKey, err := utils.GenerateSessionKey()
    if err != nil {
        return models.Session{}, errors.Internal(err)
    }

    session.Key = newKey
    session.Expiry = time.Now().Add(time.Duration(config.App.SessionDays) * 24 * time.Hour)

    if err := r.db.Save(&session).Error; err != nil {
        return models.Session{}, errors.Internal(err)
    }

    return session, nil
}
