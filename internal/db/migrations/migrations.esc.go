// Code generated by "esc -o=migrations.esc.go -pkg=migrations -ignore=\.go$ -private ."; DO NOT EDIT.

package migrations

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
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
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
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

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
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

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/20190504194418_initial_schema.down.sql": {
		name:    "20190504194418_initial_schema.down.sql",
		local:   "20190504194418_initial_schema.down.sql",
		size:    103,
		modtime: 1557163117,
		compressed: `
H4sIAAAAAAAC/3Jydff0s+bicgnyD1AIcXTycVXwdFNwjfAMDglWKM7MLchJjU/Oz81NzEspVnB2DHZ2
dHHFpTw5IzEvLzUHWZ2zv6+vZ4g1FyAAAP//rNXPS2cAAAA=
`,
	},

	"/20190504194418_initial_schema.up.sql": {
		name:    "20190504194418_initial_schema.up.sql",
		local:   "20190504194418_initial_schema.up.sql",
		size:    882,
		modtime: 1557296329,
		compressed: `
H4sIAAAAAAAC/9SRwW7CMBBE7/mKPYLUP+AUyIKiBqcKRm1Olom3YMl2othUqF9fESgBQiv12KvHbzQ7
M8VFyiZRNCsw5gg8nmYI1U46R8bDKAIA0Ao2eqtdgAUyLGKOCUxLSHAerzMO8QrSBBlPeQkvRbqMixKe
sXzq2KolGUgJGSBoSz5I24TPC8vy19EYWM6BrbPshOwb9RfkxHhqRZ/z1tBJSxDocP/etPSuD4+Uzd4Y
Cp1yevC7em+UsLWiVgaCTV0bku6CReO+w5Ql+HbpUJyjCa0OULurbs/Ckbyt32vbGBJVba106h+tcL7t
aogC51ggm+Hq6nCtBuRPE1nyXm4fSrKqyHth6IPM/e7DOe46FX3UbpicDVvvv4wnv5sd03/vO7A5il2c
fLlM+ST6CgAA//+TtYZVcgMAAA==
`,
	},

	"/": {
		name:  "/",
		local: `.`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	".": {
		_escData["/20190504194418_initial_schema.down.sql"],
		_escData["/20190504194418_initial_schema.up.sql"],
	},
}
