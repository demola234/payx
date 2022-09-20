package interfaces

type BankList struct {
	Status  bool    `json:"status"` 
	Message string  `json:"message"`
	Data    []Datum `json:"data"`   
}

type Datum struct {
	ID          int64    `json:"id"`           
	Name        string   `json:"name"`         
	Slug        string   `json:"slug"`         
	Code        string   `json:"code"`         
	PayWithBank bool     `json:"pay_with_bank"`
	Active      bool     `json:"active"`              
	CreatedAt   string   `json:"createdAt"`    
	UpdatedAt   string   `json:"updatedAt"`    
}