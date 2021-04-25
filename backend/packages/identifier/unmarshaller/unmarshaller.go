package unmarshaller

import (
	"errors"
	identifierApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/api"
	facialPatternImpl "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/impl/facialPattern"
	fingerprintPatternImpl "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/impl/fingerprintPattern"
)

type IdentificationRaw struct {
	IdentifierType identifierApi.IdentifierType `json:"identifierType" bson:"identifierType"`
	IdentifierRaw  string                       `json:"identifierRaw" bson:"identifierRaw"`
}

func (i IdentificationRaw) Unmarshal() (identifierApi.Identification, error) {
	switch i.IdentifierType {
	case identifierApi.FACIAL_PATTERN:
		id, err := facialPatternImpl.Unmarshal(i.IdentifierRaw)
		return identifierApi.Identification{
			IdentifierType: i.IdentifierType,
			Identifier:     id,
		}, err

	case identifierApi.FINGERPRINT_PATTERN:
		id, err := fingerprintPatternImpl.Unmarshal(i.IdentifierRaw)
		return identifierApi.Identification{
			IdentifierType: i.IdentifierType,
			Identifier:     id,
		}, err

	default:
		return identifierApi.Identification{}, errors.New("invalid identifier provided")
	}
}

func Marshal(identifier identifierApi.Identifier) (IdentificationRaw, error) {
	switch v := identifier.(type) {
	case fingerprintPatternImpl.Identifier:
		s, err := fingerprintPatternImpl.Marshal(v)
		return IdentificationRaw{
			IdentifierType: identifierApi.FINGERPRINT_PATTERN,
			IdentifierRaw:  s,
		}, err
	case facialPatternImpl.Identifier:
		s, err := facialPatternImpl.Marshal(v)
		return IdentificationRaw{
			IdentifierType: identifierApi.FACIAL_PATTERN,
			IdentifierRaw:  s,
		}, err

	default:
		return IdentificationRaw{}, errors.New("invalid type")
	}
}
