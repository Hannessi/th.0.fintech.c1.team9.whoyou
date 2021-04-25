package userApi

import (
	identifierApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/api"
	"github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/unmarshaller"
)

type User struct {
	Id                 string                           `json:"id" bson:"id"`
	Name               string                           `json:"name" bson:"name"`
	Surname            string                           `json:"surname" bson:"surname"`
	Address            string                           `json:"address" bson:"address"`
	Identifications    []identifierApi.Identification   `json:"identifications" bson:"-"`
	IdentificationsRaw []unmarshaller.IdentificationRaw `json:"-" bson:"identifications"`
	KycStatus          KycStatus                        `json:"kycStatus" bson:"kycStatus"`
}

type KycStatus string

const KYC_STATUS_PENDING KycStatus = "KYC_STATUS_PENDING"
const KYC_STATUS_CLEAR KycStatus = "KYC_STATUS_CLEAR"
const KYC_STATUS_INCONCLUSIVE KycStatus = "KYC_STATUS_INCONCLUSIVE"
const KYC_STATUS_FRAUDULENT KycStatus = "KYC_STATUS_FRAUDULENT"
