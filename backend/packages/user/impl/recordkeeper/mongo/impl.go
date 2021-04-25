package userRecordkeeperMongoImpl

import (
	identifierApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/api"
	"github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/unmarshaller"
	userApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/user/api"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Recordkeeper struct {
	MongoSession *mgo.Session
	Database     string
	Collection   string
}

func (r *Recordkeeper) validateCreateRequest(request userApi.CreateRequest) error {
	// todo validate request
	return nil
}

func (r *Recordkeeper) Create(request userApi.CreateRequest) (*userApi.CreateResponse, error) {
	logrus.Info("create user request: ", request)
	if err := r.validateCreateRequest(request); err != nil {
		return nil, err
	}

	s := r.MongoSession.Copy()
	defer s.Close()
	c := s.DB(r.Database).C(r.Collection)

	request.User.Id = uuid.New()
	request.User.KycStatus = userApi.KYC_STATUS_PENDING

	temp, err := convertIdentifications(request.User.Identifications)
	if err != nil {
		return nil, err
	}
	request.User.IdentificationsRaw = temp

	err = c.Insert(request.User)
	if err != nil {
		return nil, err
	}

	return &userApi.CreateResponse{
		User: request.User,
	}, nil
}

func (r *Recordkeeper) validateRetrieveRequest(request userApi.RetrieveRequest) error {
	// todo validation
	return nil
}

func (r *Recordkeeper) Retrieve(request userApi.RetrieveRequest) (*userApi.RetrieveResponse, error) {
	logrus.Info("retrieve user request: ", request)
	if err := r.validateRetrieveRequest(request); err != nil {
		return nil, err
	}

	identifierFilter := request.Identification.Identifier.GetRawJson()

	comboFilter := bson.M{"$and": []bson.M{
		{"identifications.identifierType": request.Identification.Identifier.GetType()},
		{"identifications.identifierRaw": identifierFilter},
	}}

	logrus.Info("bson map: ", comboFilter)

	s := r.MongoSession.Copy()
	defer s.Close()
	c := s.DB(r.Database).C(r.Collection)

	identifiedUser := userApi.User{}

	err := c.Find(comboFilter).One(&identifiedUser)
	if err != nil {
		logrus.Info("Could not find user: " + err.Error())
		return nil, err
	}

	return &userApi.RetrieveResponse{
		User: identifiedUser,
	}, nil
}

func convertIdentifications(list []identifierApi.Identification) ([]unmarshaller.IdentificationRaw, error) {
	newList := make([]unmarshaller.IdentificationRaw,0)

	for _, id := range list {
		raw, err := unmarshaller.Marshal(id.Identifier)
		if err != nil {
			return nil, err
		}
		newList = append(newList, raw)
	}

	return newList, nil
}