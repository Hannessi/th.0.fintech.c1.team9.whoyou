package fingerprintPatternImpl

import (
	"encoding/json"
	"errors"
	identifierApi "github.com/hannessi/th.0.fintech.c1.team9.whoyou/backend/packages/identifier/api"
	"gopkg.in/mgo.v2/bson"
)

type Identifier struct {
	SerializedFingerprintPattern string `json:"serializedFingerprintPattern" bson:"serializedFingerprintPattern"`
}

func (i Identifier) GetType() identifierApi.IdentifierType {
	return identifierApi.FINGERPRINT_PATTERN
}

func (i Identifier) GetRawJson() string {
	a , err := Marshal(i)
	if err != nil {
		return "error occurred"
	}

	return a
}

func (i Identifier) GetBsonFilter() bson.M {
	return bson.M{"serializedFingerprintPattern": i.SerializedFingerprintPattern}
}

func Unmarshal(raw string) (Identifier, error) {
	id := Identifier{}

	err := json.Unmarshal([]byte(raw), &id)
	if err != nil {
		return id, errors.New("could not unmarshal facial pattern identifier: "+err.Error())
	}
	return id, nil
}

func Marshal(identifier Identifier) (string, error) {
	b, err := json.Marshal(identifier)
	if err != nil {
		return "", err
	}
	return string(b), nil
}