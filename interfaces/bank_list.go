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
	CreatedAt   string   `json:"createdAt"`    
	UpdatedAt   string   `json:"updatedAt"`    
}