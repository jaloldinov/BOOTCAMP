package storage

import "market/models"

type StorageI interface {
	Close()
	Branch() BranchRepoI
	Category() CategoryRepoI
}

type BranchRepoI interface {
	Create(*models.CreateBranch) (string, error)
	GetByID(*models.BranchPrimaryKey) (*models.Branch, error)
	GetList(*models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Update(*models.UpdateBranch) (string, error)
	Delete(*models.BranchPrimaryKey) error
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (string, error)
	GetByID(*models.CategoryPrimaryKey) (*models.Category, error)
	GetList(*models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(*models.UpdateCategory) (string, error)
	Delete(*models.CategoryPrimaryKey) error
}