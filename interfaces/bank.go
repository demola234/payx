package interfaces

type Bank struct {
	Status  bool   `json:"status"` 
	Message string `json:"message"`
	Data    Data   `json:"data"`   
}

type Data struct {
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`  
	BankID        int64  `json:"bank_id"`       
}


