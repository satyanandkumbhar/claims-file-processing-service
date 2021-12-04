package models

//Represents indivisual Claim
type Claim struct {
	Id       string `json:"id"`
	MemberId string `json:"memberId"`
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
}

//Holder for claims
type ClaimRequests struct {
	ClaimRequests []Claim `json:"claims"`
}
