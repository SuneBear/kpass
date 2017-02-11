package schema

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/seccom/kpass/pkg/util"
)

// FileBlob ...
type FileBlob []byte // binary data encoded with base64, size <= 100KB

// FileBlobFromString ...
func FileBlobFromString(text string) (FileBlob, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(text)))
	if _, err := base64.StdEncoding.Decode(dst, []byte(text)); err != nil {
		return nil, err
	}
	return FileBlob(dst), nil
}

// MarshalText ...
func (b FileBlob) String() string {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(dst, b)
	return string(dst)
}

// Reader returns a Reader instance
func (b FileBlob) Reader() *bytes.Reader {
	return bytes.NewReader(b)
}

// File represents file
type File struct {
	Name    string    `json:"name"`
	UserID  string    `json:"userID"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// FileFrom parse JSON string and returns a File intance.
func FileFrom(str string) (file *File, err error) {
	file = new(File)
	if err = json.Unmarshal([]byte(str), file); err != nil {
		return nil, err
	}
	return file, nil
}

// String returns JSON string with full file info
func (f *File) String() string {
	return jsonMarshal(f)
}

// Result returns FileResult intance
func (f *File) Result(ID util.OID, signed string) *FileResult {
	return &FileResult{
		ID:      ID,
		Signed:  signed,
		UserID:  f.UserID,
		Name:    f.Name,
		Created: f.Created,
		Updated: f.Updated,
	}
}

// FileResult represents file info
type FileResult struct {
	ID      util.OID  `json:"id"`
	UserID  string    `json:"userID"`
	Name    string    `json:"name"`
	Signed  string    `json:"signed"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
