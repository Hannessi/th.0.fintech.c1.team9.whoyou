package userRpc

import (
	"errors"
	"github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/unmarshaller"
	userApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/user/api"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Adaptor struct {
	Registrar userApi.Registrar
}

type RegisterNewUserRequest struct {
	User              userApi.User                   `json:"user"`
	RawIdentification unmarshaller.IdentificationRaw `json:"rawIdentification"`
}

type RegisterNewUserResponse struct {
	User userApi.User `json:"user"`
}

func (a *Adaptor) RegisterNewUser(r *http.Request, request *RegisterNewUserRequest, response *RegisterNewUserResponse) error {
	log.Info("RegisterNewUser: ", request)

	identifier, err := request.RawIdentification.Unmarshal()
	if err != nil {
		return errors.New("Could not unmarshal identifier: " + err.Error())
	}

	implResponse, err := a.Registrar.RegisterNewUser(userApi.RegisterNewUserRequest{
		User:       request.User,
		Identifier: identifier.Identifier,
	})
	if err != nil {
		return err
	}

	response.User = implResponse.User
	return nil
}
