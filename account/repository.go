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


func (s *service) GetAllAccount() ([]Account, error) {
	accounts, err := s.repository.GetAllAccount()
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

func (r *repository) GetAccountByEmail(email string) (Account, error) {
	var account Account
	err := r.db.Where("email = ?", email).First(&account).Error
	if err != nil {
		return Account{}, err
	}
	return account, nil
}	
