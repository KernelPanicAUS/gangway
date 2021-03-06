// Code generated by "esc -o cmd/gangway/bindata.go templates/"; DO NOT EDIT.

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/commandline.tmpl": {
		local:   "templates/commandline.tmpl",
		size:    4187,
		modtime: 1531428210,
		compressed: `
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
`,
	},

	"/templates/home.tmpl": {
		local:   "templates/home.tmpl",
		size:    2088,
		modtime: 1531422938,
		compressed: `
H4sIAAAAAAAC/4xWUW/jNgx+76/g6V42LLKbuxXYcnaGrd2AYnfrgHYH7JGWGZutLPkkOW5u2H8fJMdJ
mhuwBUgsUhT5kR9ppXh1c3f98OfvP0MbOr2+KOIDNJqmFGREVBDW6wuAoqOA0IbQS/o08LYU19YEMkE+
7HoSoCapFIGeQx7dvAPVovMUyj8efpHfifzoxmBHpdgyjb114eTwyHVoy5q2rEgmYQFsODBq6RVqKpcL
6PCZu6GbFdnl3nXgoGndoGlG3BX5JF7EnVdSwvX9PYCUyVKzeYLW0aYUMSO/yvONNcFnjbWNJuzZZ8p2
OStrfthgx3pXfsBAjlF/c6us8QIc6VL4sNPkW6Igjo7Pd84iqdo8+kxpO9QbjY5SJHzE51xz5fNuH4c/
U36ZLS8vsze58i/0WccmU95PMVOkuALoHWUV+hb+SmL8VKieGmcHU0tltXUrqDSqp3cHg7329bffb6q3
V0d9LIj0/JlW4DvU+mxnqsoKrq3xVqNffLAGlV28HxTXuFfT4j1X5DCwNfDBGru4oUf8OMA9Gj8pfuLg
gyPs4CM5PNm4toNjcvAbjQvorLG+R0VHFHZLbqPtuAIcgj3qR+tqOTrsV2Cs6/AE+NhyIJn8rGKtpp2/
L9Iji2TLSlv1tC9fj3XNplnBJSyv+ufZ+tw4m5mRUelPan9SQjYtOQ6zjyLfs1bk04AVla13iU6DW1Aa
vS+F5qYNstIDQfwR4KymUhjccpNqKibai5oPZwxuU/Y9uTRWyIacWBcIXJdC28bKo3rfma/FfLpyaGoZ
rcRxkHC9z6gY9GzoIjRouSZpjeyolvFgbceIaTbPB30UBp0QRHydrVjTIaiPXgxuz0/u13hEWWNAiSrw
FgP5f3VWDSFYE3tdY+9JrAuet17yJNYdmaHIeX3IsMhr3iYScoPT4qSynlTqY2Nlj7WsbBApIzY1PcsK
TSr0F4SckDBnVLl1/M5iu5xtYy9E3sgEclCjeyIj3x6ogF+HipyhQB5+HEJLJrBKjVDk7fLg7yS4s+Pe
2yF6DHh1HtBq8Ms3kBpOrB9a9jAE1hx2MLLW0JLuYWeHOGtzWIKRQxu17hSY0oOP6NMmGrjrydzexBeC
IRXgq7vbm+uvIU5uBvfcGGADwUJDAXxAF6jOiry9OmQzc/L/UpubJde2YSMmhmJbahs5S81x7JVgpEbX
EIy4JS9ps4kQJyHVYpq7dcJ5a05H4QWsmdDzLkqLqD3cQnCvHPfB7y8inyTwTp3cD7am7PHTQG6XroZp
Kd9ky2yZ3vyPXqyLfDr6hZfHlzfFf9nGy/XMKCGfXkdFPv0tuPgnAAD//yEDNLMoCAAA
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/templates": {
		isDir: true,
		local: "templates",
	},
}
