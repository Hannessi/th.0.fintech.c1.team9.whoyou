package userRegistrarDefaultImpl

import (
	"errors"
	identifierApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/api"
	userApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/user/api"
	"github.com/sirupsen/logrus"
)

type Registrar struct {
	Recordkeeper userApi.Recordkeeper
}

func (r *Registrar) validateRegisterNewUserRequest(request userApi.RegisterNewUserRequest) error {
	_, err := r.Recordkeeper.Retrieve(userApi.RetrieveRequest{
		Identification: identifierApi.Identification{
			IdentifierType: request.Identifier.GetType(),
			Identifier:     request.Identifier,
		},
	})
	if err == nil {
		return errors.New("user exists")
	}
	return nil
}

func (r *Registrar) RegisterNewUser(request userApi.RegisterNewUserRequest) (*userApi.RegisterNewUserResponse, error) {
	logrus.Info("register user request: ", request)
	if err := r.validateRegisterNewUserRequest(request); err != nil {
		return nil, err
	}

	id := request.Identifier

	request.User.Identifications = append(request.User.Identifications, identifierApi.Identification{
		IdentifierType: id.GetType(),
		Identifier:     id,
	})

	createResponse, err := r.Recordkeeper.Create(userApi.CreateRequest{
		User: request.User,
	})
	if err != nil {
		return nil, err
	}

	return &userApi.RegisterNewUserResponse{
		User: createResponse.User,
	}, nil
}
