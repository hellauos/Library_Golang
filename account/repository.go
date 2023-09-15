package account

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(account Account) (Account, error)
	GetAllAccount() ([]Account, error)
	GetAccountByEmail(email string) (Account, error)
	GetAccountByRole(roleName string) ([]Account, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(account Account) (Account, error) {
	err := r.db.Create(&account).Error

	return account, err
}


func (r *repository) GetAllAccount() ([]Account, error) {
	accounts, err := r.GetAllAccount()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *repository) GetAccountByRole(roleName string) ([]Account, error){
	var account []Account
	err := r.db.Where("role.name = ?", roleName).
	Joins("JOIN roles ON account.role_id = roles.id").
	Find(&account).Error

	return account, err
}

func (r *repository) GetAccountByEmail(email string) (Account, error){
	var account Account
	err := r.db.First(&account, "email = ?", email).Error
	return account,err
}

