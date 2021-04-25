package userApi

import identifierApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/api"

type Recordkeeper interface {
	Create(request CreateRequest) (*CreateResponse, error)
	Retrieve(request RetrieveRequest) (*RetrieveResponse, error)
}

type CreateRequest struct {
	User User
}

type CreateResponse struct {
	User User
}

type RetrieveRequest struct {
	Identification identifierApi.Identification
}

type RetrieveResponse struct {
	User User
}
