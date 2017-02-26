package schema

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/seccom/kpass/src/util"
)

// FileBlob ...
type FileBlob []byte // binary data encoded with base64, size <= 200KB

// FileBlobFromString ...
func FileBlobFromString(text string) (FileBlob, error) {
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil, err
	}
	return FileBlob(data), nil
}

// String ...
func (b FileBlob) String() string {
	return base64.StdEncoding.EncodeToString(b)
}

// Reader returns a Reader instance
func (b FileBlob) Reader() *bytes.Reader {
	return bytes.NewReader(b)
}

// File represents file
type File struct {
	UserID  string    `json:"userID"`
	Name    string    `json:"name"`
	Size    int       `json:"size"`
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

// Result returns FileResult intance for entry
func (f *File) Result(ID util.OID, signed string) *FileResult {
	return &FileResult{
		ID:      ID,
		UserID:  f.UserID,
		Name:    f.Name,
		Size:    f.Size,
		Signed:  signed,
		Created: f.Created,
		Updated: f.Updated,
	}
}

// FileResult represents file info
type FileResult struct {
	ID          util.OID  `json:"id"`
	UserID      string    `json:"userID"`
	Name        string    `json:"name"`
	Size        int       `json:"size"`
	Signed      string    `json:"-"`
	DownloadURL string    `json:"DownloadURL"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

// SetDownloadURL ...
func (f *FileResult) SetDownloadURL(refType, refID string) {
	f.DownloadURL = DownloadURL(f.ID, refType, refID, f.Signed)
}

// DownloadURL ...
func DownloadURL(FileID util.OID, refType, refID, signed string) string {
	if !FileID.Valid() {
		return ""
	}
	return fmt.Sprintf(
		`/download/%s?refType=%s&refID=%s&signed=%s`, FileID, refType, refID, signed)
}
