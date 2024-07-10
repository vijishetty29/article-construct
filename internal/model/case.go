package model

type Case struct {
	Country          string    `bson:"country,omitempty"`
	Ian              string    `bson:"ian,omitempty"`
	Nat              string    `bson:"nat,omitempty"`
	UniqueID         string    `bson:"uniqueId,omitempty"`
	Status           string    `bson:"status,omitempty"`
	Source           string    `bson:"source,omitempty"`
	ProductLine      string    `bson:"productLine,omitempty"`
	ItemID           string    `bson:"itemID,omitempty"`
	UniqueVariantIDs []string  `bson:"uniqueVariantIds,omitempty"`
	Variants         []Variant `bson:"variants,omitempty"`
}
