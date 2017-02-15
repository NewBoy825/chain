package bc

import (
	"fmt"
	"io"
	"reflect"

	"chain/crypto/sha3pool"
	"chain/encoding/blockchain"
	"chain/errors"
)

type (
	Entry interface {
		io.ReaderFrom
		io.WriterTo
		Type() string
		Body() interface{}
	}

	// EntryRef holds one or both of an entry and its id. If the entry
	// is present and the id is not, the id can be generated (and then
	// cached) on demand. Both may also be nil to represent a nil entry
	// pointer.
	EntryRef struct {
		Entry
		ID *Hash
	}

	hasher interface {
		Hash() (Hash, error)
	}
)

// Hash returns the EntryRef's cached entry ID, computing it first if
// necessary. Satisfies the hasher interface.
func (r *EntryRef) Hash() (Hash, error) {
	if r.ID == nil {
		if r.Entry == nil {
			return Hash{}, nil
		}
		h, err := entryID(r.Entry)
		if err != nil {
			return Hash{}, err
		}
		r.ID = &h
	}
	return *r.ID, nil
}

func (r EntryRef) IsNil() bool {
	return r.Entry == nil && r.ID == nil
}

func (ref *EntryRef) WriteEntry(w io.Writer) (int64, error) {
	n, err := blockchain.WriteVarstr31(w, []byte(ref.Type()))
	if err != nil {
		return int64(n), err
	}
	n2, err := ref.Entry.WriteTo(w)
	return int64(n) + n2, err
}

func (ref *EntryRef) ReadEntry(r io.Reader) (int64, error) {
	typ, n, err := blockchain.ReadVarstr31(r)
	if err != nil {
		return int64(n), err
	}
	switch string(typ) {
	case typeData:
		ref.Entry = new(data)
	case typeHeader:
		ref.Entry = new(Header)
	case typeIssuance:
		ref.Entry = new(Issuance)
	case typeMux:
		ref.Entry = new(mux)
	case typeNonce:
		ref.Entry = new(Nonce)
	case typeOutput:
		ref.Entry = new(Output)
	case typeRetirement:
		ref.Entry = new(Retirement)
	case typeSpend:
		ref.Entry = new(Spend)
	case typeTimeRange:
		ref.Entry = new(TimeRange)
	default:
		return int64(n), fmt.Errorf("unknown type %s", typ)
	}
	n2, err := ref.Entry.ReadFrom(r)
	return int64(n) + n2, err
}

type extHash Hash

var errInvalidValue = errors.New("invalid value")

func entryID(e Entry) (Hash, error) {
	h := sha3pool.Get256()
	defer sha3pool.Put256(h)

	h.Write([]byte("entryid:"))
	h.Write([]byte(e.Type()))
	h.Write([]byte{':'})

	bh := sha3pool.Get256()
	defer sha3pool.Put256(bh)
	err := writeForHash(bh, e.Body())
	if err != nil {
		return Hash{}, err
	}
	var innerHash Hash
	bh.Read(innerHash[:])
	h.Write(innerHash[:])

	var hash Hash
	h.Read(hash[:])

	return hash, nil
}

func writeForHash(w io.Writer, c interface{}) error {
	switch v := c.(type) {
	case hasher:
		h, err := v.Hash()
		if err != nil {
			return errors.Wrap(err, "computing hash")
		}
		_, err = w.Write(h[:])
		return errors.Wrap(err, "writing hash")
	case byte:
		_, err := w.Write([]byte{v})
		return errors.Wrap(err, "writing byte for hash")
	case uint64:
		_, err := blockchain.WriteVarint63(w, v)
		return errors.Wrapf(err, "writing uint64 (%d) for hash", v)
	case []byte:
		_, err := blockchain.WriteVarstr31(w, v)
		return errors.Wrapf(err, "writing []byte (len %d) for hash", len(v))
	case string:
		_, err := blockchain.WriteVarstr31(w, []byte(v))
		return errors.Wrapf(err, "writing string (len %d) for hash", len(v))

		// TODO: The rest of these are all aliases for [32]byte. Do we
		// really need them all?

	case Hash:
		_, err := w.Write(v[:])
		return errors.Wrap(err, "writing Hash for hash")
	case extHash:
		_, err := w.Write(v[:])
		return errors.Wrap(err, "writing extHash for hash")
	case AssetID:
		_, err := w.Write(v[:])
		return errors.Wrap(err, "writing AssetID for hash")
	}

	// The two container types in the spec (List and Struct)
	// correspond to slices and structs in Go. They can't be
	// handled with type assertions, so we must use reflect.
	switch v := reflect.ValueOf(c); v.Kind() {
	case reflect.Slice:
		l := v.Len()
		_, err := blockchain.WriteVarint31(w, uint64(l))
		if err != nil {
			return errors.Wrapf(err, "writing slice (len %d) for hash", l)
		}
		for i := 0; i < l; i++ {
			c := v.Index(i)
			if !c.CanInterface() {
				return errInvalidValue
			}
			err := writeForHash(w, c.Interface())
			if err != nil {
				return errors.Wrapf(err, "writing slice element %d for hash", i)
			}
		}
		return nil

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			c := v.Field(i)
			if !c.CanInterface() {
				return errInvalidValue
			}
			err := writeForHash(w, c.Interface())
			if err != nil {
				t := v.Type()
				f := t.Field(i)
				return errors.Wrapf(err, "writing struct field %d (%s.%s) for hash", i, t.Name(), f.Name)
			}
		}
		return nil
	}

	return errors.Wrap(fmt.Errorf("bad type %T", c))
}
