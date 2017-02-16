package protocol

import (
	"sync"

	"github.com/golang/groupcache/lru"

	"chain/errors"
	"chain/protocol/bc"
	"chain/protocol/validation"
)

// ValidateTxCached checks a cache of prevalidated transactions
// before attempting to perform a context-free validation of the tx.
func (c *Chain) ValidateTxCached(hdrRef *bc.EntryRef) error {
	// Consult a cache of prevalidated transactions.
	hash, err := hdrRef.Hash()
	if err != nil {
		return err
	}
	err, ok := c.prevalidated.lookup(hash)
	if ok {
		return err
	}

	err = validation.CheckTxWellFormed(hdrRef)
	c.prevalidated.cache(hash, err)
	return err
}

type prevalidatedTxsCache struct {
	mu  sync.Mutex
	lru *lru.Cache
}

func (c *prevalidatedTxsCache) lookup(txID bc.Hash) (err error, ok bool) {
	c.mu.Lock()
	v, ok := c.lru.Get(txID)
	c.mu.Unlock()
	if !ok {
		return err, ok
	}
	if v == nil {
		return nil, ok
	}
	return v.(error), ok
}

func (c *prevalidatedTxsCache) cache(txID bc.Hash, err error) {
	c.mu.Lock()
	c.lru.Add(txID, err)
	c.mu.Unlock()
}

func (c *Chain) checkIssuanceWindow(tx *bc.EntryRef) error {
	hdr := tx.Entry.(*bc.Header)
	_, issuances := hdr.Inputs()
	for range issuances {
		// TODO(tessr): consider removing 0 check once we can configure this
		if c.MaxIssuanceWindow != 0 && hdr.MinTimeMS()+bc.DurationMillis(c.MaxIssuanceWindow) < hdr.MaxTimeMS() {
			// xxx should this be checking the iss->Anchor->TimeRange bounds instead?
			return errors.WithDetailf(validation.ErrBadTx, "issuance input's time window is larger than the network maximum (%s)", c.MaxIssuanceWindow)
		}
	}
	return nil
}
