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



/// ____________________________________________________________________________________________________________
/// User Account -----------------------------------------------------------------------------------------------
/// --- For all auth user


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



/// ____________________________________________________________________________________________________________
/// User Management ---------------------------------------------------------------------------------------------
/// --- For Admin/Owner usage



// Create the owner (onboarding)
func (r *UserRepository) CreateOwner( payload models.NewOwnerPayload) (models.User, *errors.Error) {
	resetPass := false
	user := models.User{
		BaseModel: models.BaseModel{
			ID: utils.GenerateID(),
		},
		DisplayName: payload.DisplayName,
		Email:       payload.Email,
		Password:    payload.Password,
		SystemRole:  def.SystemRoleOwner.ToString(),
		ResetPass:	 &resetPass,
		Avatar: "/_static/default/avatars/avatar_25.png",
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

// Create inserts a new user
func (r *UserRepository) Create( payload models.NewUserPayload) (models.User, *errors.Error) {
	user := models.User{
		BaseModel: models.BaseModel{
			ID: utils.GenerateID(),
		},
		DisplayName: payload.DisplayName,
		Email:       payload.Email,
		Password:    payload.Password,
		SystemRole:  payload.SystemRole,
		Avatar: payload.Avatar,
		
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



// Get all users
func (r *UserRepository) GetAll() ([]models.User, *errors.Error) {
	var user []models.User
	if err := r.db.Find(&user).Error; err != nil {
        return nil, errors.Internal(err)
    }
	return user, nil
}



// GetByEmail - fetch user by email
func (r *UserRepository) GetByEmail(email string) (models.User, *errors.Error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.NotFound("User not found")
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




// Update
func (r *UserRepository) Update(userID int64, payload models.UpdateUserPayload) (models.User, *errors.Error) {
	var user models.User

	// Find user by ID first
	if err := r.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, errors.NotFound("User not found")
		}
		return models.User{}, errors.Internal(err)
	}

	// Apply updates
	if err := r.db.Model(&user).Updates(payload).Error; err != nil {
		return models.User{}, errors.Internal(err)
	}

	// Return the updated record
	return user, nil
}



// Delete - deletes a user
func (r *UserRepository) Delete(userID int64) *errors.Error {
	result := r.db.Delete(&models.User{}, userID)
	if result.Error != nil {
		return errors.Internal(result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.NotFound("User not found")
	}
	return nil
}

