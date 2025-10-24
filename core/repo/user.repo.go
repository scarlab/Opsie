package repo

import (
	"opsie/core/models"
	"opsie/def"
	"opsie/pkg/errors"
	"opsie/pkg/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}


// CreateOwnerAccount - inserts a new owner user
// CreateOwnerAccount inserts a new owner user
func (r *UserRepository) CreateOwnerAccount( payload models.NewUserPayload) (models.User, *errors.Error) {
	user := models.User{
		BaseModel: models.BaseModel{
			ID: utils.GenerateID(),
		},
		DisplayName: payload.DisplayName,
		Email:       payload.Email,
		Password:    payload.Password,
		SystemRole:  def.SystemRoleOwner.ToString(),
		IsActive:    true,
	}

	if err := r.db.Create(&user).Error; err != nil {
		// Handle duplicate email (Postgres unique constraint)
		if errors.IsPgConflict(err) {
			return models.User{}, errors.Conflict("Email already in use")
		}
		return models.User{}, errors.Internal(err)
	}

	return user, nil
}


// GetOwnerCount - counts number of owners
func (r *UserRepository) GetOwnerCount() (int, *errors.Error) {
	var count int64
	if err := r.db.Model(&models.User{}).Where("system_role = ?", def.SystemRoleOwner).Count(&count).Error; err != nil {
		return 0, errors.Internal(err)
	}
	return int(count), nil
}

// Delete - deletes a user
func (r *UserRepository) Delete(userID int64) *errors.Error {
	result := r.db.Delete(&models.User{}, userID)
	if result.Error != nil {
		return errors.Internal(result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.NotFound("user not found")
	}
	return nil
}

// GetByEmail - fetch user by email
func (r *UserRepository) GetByEmail(email string) (models.User, *errors.Error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.NotFound("user not found")
		}
		return user, errors.Internal(err)
	}
	return user, nil
}

// GetByID - fetch user by ID
func (r *UserRepository) GetByID(ID int64) (models.User, *errors.Error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.NotFound("user not found")
		}
		return user, errors.Internal(err)
	}
	return user, nil
}

// UpdateAccountName - updates display name
func (r *UserRepository) UpdateAccountName(userID int64, name string) (models.User, *errors.Error) {
	var user models.User
	if err := r.db.Model(&models.User{}).Where("id = ?", userID).Update("display_name", name).First(&user, "id = ?", userID).Error; err != nil {
		return user, errors.Internal(err)
	}
	return user, nil
}

// UpdateAccountPassword - updates password
func (r *UserRepository) UpdateAccountPassword(userID int64, password string) (bool, *errors.Error) {
	res := r.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password)
	if res.Error != nil {
		return false, errors.Internal(res.Error)
	}
	return res.RowsAffected > 0, nil
}
