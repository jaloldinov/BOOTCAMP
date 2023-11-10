package memory

import (
	"encoding/json"
	"errors"
	"lesson_15/models"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type saleRepo struct {
	fileName string
}

func NewSaleRepo(fileName string) *saleRepo {
	return &saleRepo{fileName: fileName}
}

func (c *saleRepo) CreateSale(req models.CreateSales) (string, error) {
	id := uuid.NewString()
	sales, err := c.read()
	if err != nil {
		return "", err
	}
	sales = append(sales, models.Sales{
		Id:               id,
		Client_name:      req.Client_name,
		Price:            req.Price,
		Payment_Type:     req.Payment_Type,
		Status:           req.Status,
		Branch_id:        req.Branch_id,
		Shop_asissent_id: req.Shop_asissent_id,
		Cashier_id:       req.Cashier_id,
		Created_at:       time.Now().Format("2006-01-02 15:04:05"),
	})
	err = c.write(sales)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *saleRepo) UpdateSale(req models.Sales) (string, error) {
	sales, err := c.read()
	if err != nil {
		return "", err
	}
	for i, v := range sales {
		if v.Id == req.Id {
			sales[i].Client_name = req.Client_name
			sales[i].Price = req.Price
			sales[i].Payment_Type = req.Payment_Type
			sales[i].Status = req.Status
			sales[i].Branch_id = req.Branch_id
			sales[i].Shop_asissent_id = req.Shop_asissent_id
			sales[i].Cashier_id = req.Cashier_id
			err = c.write(sales)
			if err != nil {
				return "", err
			}
			return "updated", nil
		}
	}
	return "", errors.New("not sale found with ID")
}

func (c *saleRepo) GetSale(req models.IdRequest) (resp models.Sales, err error) {
	sales, err := c.read()
	if err != nil {
		return models.Sales{}, err
	}
	for _, v := range sales {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Sales{}, errors.New("not found")
}

func (c *saleRepo) GetAllSale(req models.GetAllSalesRequest) (resp models.GetAllSalesResponse, err error) {
	sales, err := c.read()
	if err != nil {
		return models.GetAllSalesResponse{}, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	var filtered []models.Sales

	for _, v := range sales {
		if strings.Contains(v.Client_name, req.Client_name) {
			filtered = append(filtered, v)
		}
	}

	if start > len(filtered) {
		resp.Sales = []models.Sales{}
		resp.Count = len(filtered)
		return
	} else if end > len(filtered) {
		return models.GetAllSalesResponse{
			Sales: filtered[start:],
			Count: len(filtered),
		}, nil
	}

	return models.GetAllSalesResponse{
		Sales: filtered[start:end],
		Count: len(filtered),
	}, nil

}

func (c *saleRepo) DeleteSale(req models.IdRequest) (resp string, err error) {
	sales, err := c.read()
	if err != nil {
		return "", err
	}
	for i, v := range sales {
		if v.Id == req.Id {
			sales = append(sales[:i], sales[i+1:]...)
			err = c.write(sales)
			if err != nil {
				return "", err
			}
			return "deleted", nil
		}
	}
	return "", errors.New("not found")
}

func (u *saleRepo) read() ([]models.Sales, error) {
	var branches []models.Sales

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return branches, nil
}

func (u *saleRepo) write(sales []models.Sales) error {
	body, err := json.Marshal(sales)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
