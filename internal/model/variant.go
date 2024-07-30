package model

type Variant struct {
	Country     string   `bson:"country,omitempty"`
	Ian         string   `bson:"ian,omitempty"`
	Nat         string   `bson:"nat,omitempty"`
	UniqueID    string   `bson:"uniqueId,omitempty"`
	ItemStatus  string   `bson:"itemStatus,omitempty"`
	Source      string   `bson:"source,omitempty"`
	ProductLine string   `bson:"productLine,omitempty"`
	ItemID      string   `bson:"itemID,omitempty"`
	CaseIDs     []string `bson:"caseIDs,omitempty"`
	Item        Item     `bson:"item,omitempty"`
	Cases       []Case   `bson:"cases,omitempty"`
}
