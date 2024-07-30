package dto

type VariantDto struct {
	Country     string    `json:"country,omitempty"`
	Ian         string    `json:"ian,omitempty"`
	Nat         string    `json:"nat,omitempty"`
	UniqueID    string    `json:"uniqueId,omitempty"`
	ItemStatus  string    `json:"itemStatus,omitempty"`
	Source      string    `json:"source,omitempty"`
	ProductLine string    `json:"productLine,omitempty"`
	ItemID      string    `json:"itemID,omitempty"`
	CaseIDs     []string  `json:"caseIDs,omitempty"`
	Cases       []CaseDto `json:"cases,omitempty"`
}
