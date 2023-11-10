package postgres

import (
	"context"
	"errors"
	"fmt"
	"market/models"
	"market/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type personRepo struct {
	db *pgxpool.Pool
}

func NewPerson(db *pgxpool.Pool) *personRepo {
	return &personRepo{
		db: db,
	}
}
func (p *personRepo) Create(newPerson models.CreatePerson) (string, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO 
		persons(id,name,job,age) 
	VALUES($1,$2,$3,$4)`
	_, err := p.db.Exec(context.Background(), query,
		id,
		newPerson.Name,
		newPerson.Job,
		newPerson.Age,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}
	return id, nil
}

func (p *personRepo) Update(newPerson models.Person) (string, error) {
	query := `
	UPDATE persons
	SET name=$2,job=$3,age=$4
	WHERE id=$1`
	resp, err := p.db.Exec(context.Background(), query,
		newPerson.Id,
		newPerson.Name,
		newPerson.Job,
		newPerson.Age,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (p *personRepo) GetAll(req models.GetAllPersonsRequest) (*models.GetAllPersonsResponse, error) {
	var (
		params  = make(map[string]interface{})
		filter  = "WHERE true"
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
		count   = 0
	)
	s := `
	SELECT
	id,
	name,
	job,
	age,
	created_at::text
	FROM persons
	`
	countQ := `
	SELECT
	count(*)
	FROM persons
	`
	if req.Search != "" {
		filter += ` AND name ILIKE '%' || @search || '%' `
		params["search"] = req.Search
	}
	if req.Job != "" {
		filter += ` AND job=@job `
		params["job"] = req.Job
	}
	if req.Age > 0 {
		filter += ` AND age=@age `
		params["age"] = req.Age
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf(" OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ
	countQ = countQ + filter
	c, pArr := helper.ReplaceQueryParams(countQ, params)
	err := p.db.QueryRow(context.Background(), c, pArr...).Scan(&count)
	if err != nil {
		return nil, err
	}
	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := p.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	resp := []models.Person{}
	for rows.Next() {
		var person models.Person
		err := rows.Scan(
			&person.Id,
			&person.Name,
			&person.Job,
			&person.Age,
			&person.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		resp = append(resp, person)
	}
	return &models.GetAllPersonsResponse{Persons: resp, Count: int32(count)}, nil
}

func (p *personRepo) Get(req models.RequestByID) (*models.Person, error) {
	s := `
	SELECT
	id,
	name,
	job,
	age
	FROM persons
	WHERE id=$1
	`
	person := models.Person{}
	err := p.db.QueryRow(context.Background(), s, req.ID).Scan(
		&person.Id,
		&person.Name,
		&person.Job,
		&person.Age,
	)
	if err != nil {
		return nil, err
	}
	return &person, err
}

func (p *personRepo) GetByUsername(req models.RequestByUsername) (*models.Person, error) {
	s := `
	SELECT
	id,
	name,
	job,
	age
	FROM persons
	WHERE id=$1
	`
	person := models.Person{}
	err := p.db.QueryRow(context.Background(), s, req.Username).Scan(
		&person.Id,
		&person.Name,
		&person.Job,
		&person.Age,
	)
	if err != nil {
		return &person, err
	}
	return &person, err
}
func (p *personRepo) Delete(req models.RequestByID) (string, error) {

	return "", errors.New("not found")
}
