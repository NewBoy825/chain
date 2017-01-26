package tx

import "chain/protocol/bc"

type issuance struct {
	body struct {
		anchor  entryRef
		value   bc.AssetAmount
		data    entryRef
		extHash extHash
	}
}

func (issuance) Type() string           { return "issuance1" }
func (iss *issuance) Body() interface{} { return iss.body }

func newIssuance(anchor entryRef, value bc.AssetAmount, data entryRef) *issuance {
	iss := new(issuance)
	iss.body.anchor = anchor
	iss.body.value = value
	iss.body.data = data
	return iss
}
