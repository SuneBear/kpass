package model

import (
	"io"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/service"
	"github.com/seccom/kpass/src/util"
	"github.com/tidwall/buntdb"
)

// File is database access oject for files
type File struct {
	db *service.DB
}

// Init ...
func (m *File) Init(db *service.DB) *File {
	m.db = db
	return m
}

// Create ...
func (m *File) Create(userID, key, name string, r io.Reader) (
	fileResult *schema.FileResult, err error) {
	FileID := util.NewOID()
	signed := ""
	var data []byte
	if data, err = ioutil.ReadAll(r); err != nil {
		return
	}
	size := len(data)
	if key != "" {
		key = auth.AESKey(FileID.String(), key)
		if data, err = auth.Encrypt([]byte(key), data); err != nil {
			return
		}
		if signed, err = auth.SignedFileKey(FileID, key); err != nil {
			return
		}
	}

	file := &schema.File{UserID: userID, Name: name, Size: size, Created: util.Time(time.Now())}
	file.Updated = file.Created
	fileResult = file.Result(FileID, url.QueryEscape(signed))
	err = m.db.DB.Update(func(tx *buntdb.Tx) error {
		_, _, e := tx.Set(schema.FileKey(FileID), file.String(), nil)
		if e == nil {
			blob := schema.FileBlob(data)
			_, _, e = tx.Set(schema.FileBlobKey(FileID), blob.String(), nil)
		}
		return e
	})

	if err != nil {
		return nil, dbError(err)
	}
	return
}

// FindFile ...
func (m *File) FindFile(FileID util.OID, key string) (
	file *schema.File, fileBlob schema.FileBlob, err error) {
	err = m.db.DB.View(func(tx *buntdb.Tx) error {
		res, e := tx.Get(schema.FileKey(FileID))
		if e != nil {
			return e
		}
		if file, e = schema.FileFrom(res); e != nil {
			return e
		}

		if res, e = tx.Get(schema.FileBlobKey(FileID)); e != nil {
			return e
		}
		if fileBlob, e = schema.FileBlobFromString(res); e != nil {
			return e
		}
		if key != "" {
			data, err := auth.Decrypt([]byte(key), []byte(fileBlob))
			if err != nil {
				return err
			}
			fileBlob = schema.FileBlob(data)
		}
		return nil
	})
	if err != nil {
		return nil, nil, dbError(err)
	}
	return
}

// Delete ...
func (m *File) Delete(FileID util.OID) error {
	err := m.db.DB.Update(func(tx *buntdb.Tx) error {
		_, e := tx.Delete(schema.FileKey(FileID))
		_, e = tx.Delete(schema.FileBlobKey(FileID))
		return e
	})

	return dbError(err)
}

// FindFiles ...
func (m *File) FindFiles(EntryID util.OID, key string, ids ...string) (
	files []*schema.FileResult, err error) {
	files = make([]*schema.FileResult, 0)
	entryID := EntryID.String()
	err = m.db.DB.View(func(tx *buntdb.Tx) error {
		for _, id := range ids {
			ID, e := util.ParseOID(id)
			if e != nil {
				return e
			}
			res, e := tx.Get(schema.FileKey(ID))
			if e != nil {
				return e
			}
			file, e := schema.FileFrom(res)
			if e != nil {
				return e
			}
			signed, err := auth.SignedFileKey(ID, auth.AESKey(ID.String(), key))
			if err != nil {
				return err
			}
			fileResult := file.Result(ID, url.QueryEscape(signed))
			fileResult.SetDownloadURL("entry", entryID)
			files = append(files, fileResult)
		}
		return nil
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}
