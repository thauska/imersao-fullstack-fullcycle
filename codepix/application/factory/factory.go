package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/thauska/imersao-fullstack-fullcycle/codepix/application/usecase"
	"github.com/thauska/imersao-fullstack-fullcycle/codepix/infrastructure/repository"
)

// TransactionUseCaseFactory returns a transaction use case
func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase
}
