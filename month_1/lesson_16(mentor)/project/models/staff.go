package models

type CreateStaff struct {
	BranchId string
	TariffId string
	TypeId   StaffType
	Name     string
	Balance  float64
}

type Staff struct {
	Id       string
	BranchId string
	TariffId string
	TypeId   StaffType
	Name     string
	Balance  float64
}

type ChangeBalance struct {
	Id      string
	Balance float64
}

type StaffType string

const (
	Cashier       StaffType = "cashier"
	ShopAssistant StaffType = "shop_assistant"
)

type GetAllStaffRequest struct {
	Page        int
	Limit       int
	Type        StaffType
	Name        string
	BalanceFrom float64
	BalanceTo   float64
}

type GetAllStaff struct {
	Staffes []Staff
	Count   int
}
