package usecase

import (
	"github.com/codeedu/imersao/codepix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (pix *PixUseCase) Register(key string, kind string, account_id string) (*model.PixKey, error) {
	account, err := pix.PixKeyRepository.FindAccount(account_id)

	if err != nil {
		return nil, err
	}

	new_pix_key, err := model.NewPixKey(kind, key, account)

	if err != nil {
		return nil, err
	}

	pix_key, err := pix.PixKeyRepository.RegisterKey(new_pix_key)

	if pix_key.ID == "" {
		return nil, err
	}

	return pix_key, nil
}
