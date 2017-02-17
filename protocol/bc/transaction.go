package bc

type Transaction struct {
	Header      *EntryRef
	Issuances   []*EntryRef
	Spends      []*EntryRef
	Outputs     []*EntryRef
	Retirements []*EntryRef
}

func NewTransaction(hdrRef *EntryRef) *Transaction {
	hdr := hdrRef.Entry.(*Header)
	spends, issuances := hdr.Inputs()
	tx := &Transaction{
		Header:    hdrRef,
		Issuances: issuances,
		Spends:    spends,
	}
	for _, r := range hdr.Results() {
		switch r.Entry.(type) {
		case *Output:
			tx.Outputs = append(tx.Outputs, r)
		case *Retirement:
			tx.Retirements = append(tx.Retirements, r)
		}
	}
	return tx
}

func (tx *Transaction) ID() Hash {
	return tx.Header.Hash()
}

func (tx *Transaction) Version() uint64 {
	return tx.Header.Entry.(*Header).Version()
}

func (tx *Transaction) Data() *EntryRef {
	return tx.Header.Entry.(*Header).Data()
}

func (tx *Transaction) MinTimeMS() uint64 {
	return tx.Header.Entry.(*Header).MinTimeMS()
}

func (tx *Transaction) MaxTimeMS() uint64 {
	return tx.Header.Entry.(*Header).MaxTimeMS()
}

func (tx *Transaction) RefDataHash() Hash {
	return tx.Header.Entry.(*Header).RefDataHash()
}
