package models

type CreateSales struct {
	Client_name      string
	Branch_id        int
	Shop_asissent_id int
	Cashier_id       int
	Price            float64
	Payment_Type     Payment
	Status           Status
	Created_at       string
}

type Payment string
type Status string

const (
	Card    Payment = "card"
	Cash    Payment = "cash"
	Success Status  = "success"
	Cancel  Status  = "cancel"
)

type Sales struct {
	Id               string
	Client_name      string
	Branch_id        int
	Shop_asissent_id int
	Cashier_id       int
	Price            float64
	Payment_Type     Payment
	Status           Status
	Created_at       string
}

type GetAllSalesRequest struct {
	Page        int
	Limit       int
	Client_name string
}

type GetAllSalesResponse struct {
	Sales []Sales
	Count int
}
