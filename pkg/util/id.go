package util

import (
	"errors"
	"strconv"
	"time"

	"github.com/sdming/gosnow"
)

var snow *gosnow.SnowFlake
var since = time.Date(2017, 1, 0, 0, 0, 0, 0, time.UTC)

const (
	mask uint = gosnow.WorkerIdBits + gosnow.SequenceBits
)

func init() {
	var err error
	if snow, err = gosnow.Default(); err != nil {
		panic(err)
	}
}

// OID is unique id
type OID uint64

// NewOID returns a OID instance
func NewOID() OID {
	id, err := snow.Next()
	if err != nil {
		panic(err)
	}
	return OID(id)
}

// ParseOID ...
func ParseOID(str string) (OID, error) {
	i, _ := strconv.ParseUint(str, 36, 64)
	id := OID(i)
	if id.IsValid() {
		return id, nil
	}
	return id, errors.New("invalid OID: " + str)
}

// Equal ...
func (s OID) Equal(a OID) bool {
	return uint64(s) == uint64(a)
}

// IsValid ...
func (s OID) IsValid() bool {
	return s.GetTime().After(since)
}

// GetTime ...
func (s OID) GetTime() time.Time {
	t := int64(uint64(s)>>mask) + gosnow.Since
	return time.Unix(t/1000, t%1000)
}

func (s OID) String() string {
	return strconv.FormatUint(uint64(s), 36)
}

// MarshalText ...
func (s *OID) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// UnmarshalText ...
func (s *OID) UnmarshalText(b []byte) error {
	id, err := ParseOID(string(b))
	if err == nil {
		*s = id
	}
	return err
}
