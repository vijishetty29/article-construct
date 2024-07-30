package model

type Item struct {
	Country       string   `bson:"country,omitempty"`
	Ian           string   `bson:"ian,omitempty"`
	Nat           string   `bson:"nat,omitempty"`
	UniqueID      string   `bson:"uniqueId,omitempty"`
	ItemStatus    string   `bson:"itemStatus,omitempty"`
	Source        string   `bson:"source,omitempty"`
	ProductLine   string   `bson:"productLine,omitempty"`
	UniqueCaseIDs []string `bson:"uniqueCaseIds,omitempty"`
	Cases         []Case   `bson:"cases,omitempty"`
}
