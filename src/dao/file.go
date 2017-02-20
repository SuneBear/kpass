package dao

import (
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

// NewFile return a File intance
func NewFile(db *service.DB) *File {
	return &File{db}
}

// Create ...
func (o *File) Create(userID, name, key string, data []byte) (fileResult *schema.FileResult, err error) {
	FileID := util.NewOID()
	signed := ""
	if key != "" {
		key = auth.AESKey(FileID.String(), key)
		if data, err = auth.Encrypt([]byte(key), data); err != nil {
			return
		}
		if signed, err = auth.SignedFileKey(FileID, key); err != nil {
			return
		}
	}

	file := &schema.File{UserID: userID, Name: name, Created: util.Time(time.Now())}
	file.Updated = file.Created
	fileResult = file.Result(FileID, signed)
	err = o.db.DB.Update(func(tx *buntdb.Tx) error {
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
func (o *File) FindFile(FileID util.OID, key string) (file *schema.File, fileBlob schema.FileBlob, err error) {
	err = o.db.DB.View(func(tx *buntdb.Tx) error {
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
			key = auth.AESKey(FileID.String(), key)
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
func (o *File) Delete(FileID util.OID) error {
	err := o.db.DB.Update(func(tx *buntdb.Tx) error {
		_, e := tx.Delete(schema.FileKey(FileID))
		return e
	})

	return dbError(err)
}

// FindFiles ...
func (o *File) FindFiles(key string, ids ...string) (files []*schema.FileResult, err error) {
	files = make([]*schema.FileResult, 0)
	err = o.db.DB.View(func(tx *buntdb.Tx) error {
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
			files = append(files, file.Result(ID, signed))
		}
		return nil
	})
	if err != nil {
		return nil, dbError(err)
	}
	return
}
