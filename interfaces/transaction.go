package interfaces

import(
	
)

type Deposit struct{
	Amount	int		`json:"amount"`
	Email	string	`json:"email"`
	Metadata Metadata	`json:"metadata"`
}

type Metadata struct{
	Amount int		`json:"amount"`
	Message	string	`json:"message"`
	DebitorAccount	string	`json:"debitorAccount"`
	CreditorAccount string	`json:"creditorAccount"`
}

type DepositPayload struct{
	Amount int		`json:"amount"`
	Message	string	`json:"message"`
	DebitorAccount	string	`json:"debitorAccount"`
	CreditorAccount string	`json:"creditorAccount"`
}

type DepositResponse struct{
	Status bool 	`json:"status"`
	Message	string	`json:"message"`
	Data	DepositResponseData	
}

type DepositResponseData struct{
	AuthorizationUrl string 		`json:"authorization_url"`
	AccessCode	string	`json:"access_code"`
	Reference	string	`json:"reference"`
}

type VerificationResponse struct{
	Status	bool    `json:"status"`
	Data	VerificationResponseData	`json:"data"`
}

type VerificationResponseData struct{
	Metadata VerificationResponseDataMetadata	`json:"metadata"`
}

type VerificationResponseDataMetadata struct{
	Amount string	`json:"amount"`
	Message	string		`json:"message"`
	DebitorAccount	string		`json:"debitorAccount"`
	CreditorAccount string	`json:"creditorAccount"`
}