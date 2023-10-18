package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pix_key *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"account_id" gorm:"column:account_id;type:uuid;notnull" valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pix_key *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pix_key)

	if pix_key.Kind != "email" && pix_key.Kind != "cpf" {
		return errors.New("invalid type of key")
	}

	if pix_key.Status != "active" && pix_key.Kind != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(kind string, key string, account *Account) (*PixKey, error) {
	pix_key := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pix_key.ID = uuid.NewV4().String()
	pix_key.CreatedAt = time.Now()

	err := pix_key.isValid()

	if err != nil {
		return nil, err
	}

	return &pix_key, nil
}
