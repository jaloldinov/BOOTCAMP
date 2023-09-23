package postgres

import (
	"context"
	"fmt"
	"market/config"
	"market/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db         *pgxpool.Pool
	branches   *branchRepo
	categories *categoryRepo
	products   *productRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnections

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Branch() storage.BranchRepoI {
	if s.branches == nil {
		s.branches = NewBranchRepo(s.db)
	}
	return s.branches
}

func (s *store) Category() storage.CategoryRepoI {
	if s.categories == nil {
		s.categories = NewCategoryRepo(s.db)
	}
	return s.categories
}

func (s *store) Product() storage.ProductRepoI {
	if s.products == nil {
		s.products = NewProductRepo(s.db)
	}
	return s.products
}

func (s *store) Close() {
	s.db.Close()
}
