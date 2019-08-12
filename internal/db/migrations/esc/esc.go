// Code generated by "esc -o=esc/esc.go -pkg=esc -ignore=esc -include=\.sql$ -private -modtime=0 ."; DO NOT EDIT.

package esc

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
		size:    283,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/1zPwW7DIAwG4DtP4QfYG+SUpmxCWtdqyWE7IQqegupABiZ7/R1Cqig3/9/B9n+Sb+qj
EeL8eb3B0J7eJahXkF+qH3oBAGBHEwJSfllTyRwnbeM0meAq/pbIWGdTOCacyW+wmOTNnbbIf57tqDk+
MFS6U7QPdLpkTNuZdb8mn/lAPvzESglnNIzu8E62I7pCOxdd23ftWT57ft92NcFYizlrwgWpEaK7Xi5q
aMR/AAAA//955cCYGwEAAA==
`,
	},

	"/20190504194418_initial_schema.up.sql": {
		name:    "20190504194418_initial_schema.up.sql",
		local:   "20190504194418_initial_schema.up.sql",
		size:    7232,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/+xY32/bthN/919x36c0+CZF8lx0gJOordFE6WwHXTAMBC2eLC78oZJUYm/Y/z5QP2zJ
ouy0QLEVWZ4i3/GO97nP8Y48PYV5xi2kXCBwC7ZY/I6JA6chyahaIhTKcQEuQ0i5sQ4MUgEMc6HXEpUD
nZbChXavR6encKVBaa8l1qCVF1kEm2QoKUi+NNRxrSwYlJQrrpaQaGW5dd5U44pbb0mitXSJkFELC0Tl
1+hHZK9Ho4vo/SR+MxqNLqfReB7B/P5TBDRJ0Foi8BEFjGcQxXc38GoEAHCEj2jWWuHRSfVti4VNDF+g
aX6RmqGhTm9+WBhNWUKt2+pQJrk6Gh23PY8vrqMSKoXC1u44gwVfcuXgfRRH0/E8uoKLe7iK3o3vrud+
b5OrKJ5P5vfwaTq5GU/v4WN0XzlJDFKHjFAHjku0jsrc/bFZG99+fnUM8e0c4rvr62pJkbOvWVKtsWjI
dp+NFO7iyc93UWVXUYngcOV2/C20I0Mymjj+6NmgBVK1I8wNpnwVNFkIga6U1Bv0+eimdCcCvlTaICvX
/PrbJtqjP/862jGeFNZpSfSTQmOfry81+wptg8tC0IP2qyVaC6afFHDVhCuodSSVIWxyaiyStS5csRhC
FlfOUCJ4ij6p7YxWckPTVCBBRRcCWcBIpaWFqMEeQH6j1Q4hJGeY0kK4HXEpt5kuBCN1wQ1FxLjNBV2T
J2r8MWGHAi8jOqTly0IXjrCiOn8Cu64NpVw4z5IBhCoxEVw9DPnK0UjufD1WWocZ5A8jIum6XHDAc0Lz
IcctDSK5Iknm6diPtK2Xo0lQOX/M7lcsDXrfgYTWeii1wwObq3SIpKthf7WO5Wop8AAcdi0XWhxwWis9
K9hG18fbGB8OWeJ+z5KuiEC1dNkeIwvfOxjJM0PtIfy6uiSnzqFRzzp0Lj9Elx/hVX0G/+8tHB0dn7Ql
A7y4iOafoyiGM6CKwfnZWXBVAOMDC3s1+dNbODsum2vdWyfxVfTLpreSumURzlZ+stj23FoQaMvV0Zxo
KaliP1B3rmNrNehp9C6aRvFlNGsFzllvZTM0dfpIANQuMmTrsIT3Nu5jt1XpA/2lKEv/BeCrChk4PMr4
d3v3NiBtQn0dGQ9ISlE1hbUxP/GeA9VRId/Onypkk8MmLbtmevmjhdMGc8FfchKd4cslBlOlDV82h209
pJaTDtpcKxschBNdqOD4870Z0cploKw7me6U9KCRFqE6qz2K49llawMVmR6p4X6WehlUGrgHPVJR4PMz
SSUGUrlBslPfVGKTjxbUPXO9IndP3CUZcfoB1Q+Um3rbe+6q9cWzvLCU0YUSUgqIW+fBdBlMDdpseDmu
cm7WnV0Pd9cO0mQTQJO1nURs5P2ULYROHpCVo88/kLJn4T/qHQD1zEAEt+5FHALcodw/g4fejGqUuEr1
Cz0qO69NnReHPc8PVWPd4eLJt7XWzctLYZG1MWiC7QzBAzHvDso+9HZg7XIYMNEtmLaB6sZ0G8/m0/Ek
nkOSPZBC8S9F+dBTX6VKJf+3/a/8uhzPIvj8IYoDgUxmVQrmH8orWnQ9i+AcULHjjo3/d6zsRPIMG8fw
Fs7Lr9DVrl0CQz1up04O9zmDeVUJL+Te1+aYB2k/x2oY2ya6ZBt+Iyyf5VDQdWB4ru+dhPE0bbw3oZ6H
Km6gjL+ljvvE6lEgMAkHaDI8DwcMdhHfZ7Wr2SesTTJkhfiPsd+FsYnRiuAqN2gt18Hh7t9F3j4dAuwN
cWaYviGTYf4G7YYIfHtzM5m/Gf0dAAD//9y5APdAHAAA
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
