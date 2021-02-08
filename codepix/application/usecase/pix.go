package usecase

import (
	"github.com/thauska/imersao-fullstack-fullcycle/codepix/domain/model"
)

// PixUseCase represents a use case for pix keys
type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

// RegisterKey registers a new key
func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

// FindKey search a pix key by kind and key
func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
