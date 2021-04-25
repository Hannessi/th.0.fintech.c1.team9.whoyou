package identifierApi

type Identifier interface {
	GetType() IdentifierType
	GetRawJson() string
	//GetBsonFilter() (bson.M, error)
	// todo IsValid() bool
}

type IdentifierType string

const FACIAL_PATTERN IdentifierType = "FACIAL_PATTERN"
const FINGERPRINT_PATTERN IdentifierType = "FINGERPRINT_PATTERN"

type Identification struct {
	IdentifierType IdentifierType `json:"identifierType" bson:"identifierType"`
	Identifier     Identifier     `json:"identifier" bson:"identifier"`
}

