package infrastracture

import "github.com/tf63/go-tx-sample/internal/uow-pattern/domain"

// RepositoryManagerの実装
type RepositoryManager interface {
	accountRepository() domain.AccountRepository
}

type repositoryManager struct {
	ar domain.AccountRepository
}

func (r *repositoryManager) accountRepository() domain.AccountRepository {
	return r.ar
}

func NewRepositoryManager(ar domain.AccountRepository) RepositoryManager {
	return &repositoryManager{ar: ar}
}
