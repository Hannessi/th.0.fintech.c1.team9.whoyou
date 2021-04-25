package userApi

import identifierApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/api"

type Registrar interface {
	RegisterNewUser(request RegisterNewUserRequest) (*RegisterNewUserResponse, error)
}

type RegisterNewUserRequest struct {
	User       User
	Identifier identifierApi.Identifier
}

type RegisterNewUserResponse struct {
	User User
}
