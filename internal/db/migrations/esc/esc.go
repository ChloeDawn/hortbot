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
		size:    7593,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/+xYW2/bthd/96c4/6c0+CdF8lx0gJOordFE6WIHXTAMBCUe2Vx4UUkqsTfsuw/UxZEs
ykqHFkOR9anOuZDn/H7nQh0fw2LFLWRcIHALtkh+x9SB05CuqFoiFMpxAW6FkHFjHRikAhjmQm8kKgc6
K4WJdq8nx8dwoUFpryU2oJUXWQSbrlBSkHxpqONaWTAoKVdcLSHVynLrvKvmKG69J4nW0iXCilpIEJW3
0Q/IXk8mZ9H7WfxmMpmc30TTRQSLu08R0DRFa4nABxQwnUMU317BqwkAwAE+oNlohQdH1W9bJDY1PEHT
/EVqhoY6vf1DYjRlKbXuSYcyydXB5LB98vTsMipTpVDY+jjOIOFLrhy8j+LoZrqILuDsDi6id9Pby4W/
2+wiihezxR18upldTW/u4GN0Vx2SGqQOGaEOHJdoHZW5+2NrG19/fnUI8fUC4tvLy8qkyNnXmFQ2Fg15
umcjhdt49vNtVPlVVCI4XA+IGbe5oBvSV6vkiXaDMpo6/uBJowVStSPMDWZ8HXRZCIGulNRxeNi6yO8E
ypdKG2Slza+/bZNy8OdfBzvO08I6LYl+VGjs8/WlZl+hbXBZCDrqvzLRWjD9qICrJlxBrSOZDOUmp8Yi
2ejCFclQZnHtDCWCZ+ixbwNfyQ3NMoEEFU0EsgEn1iGV3j5wicIkVI3Yu0esMdwN2GghahQHIN1qtXMT
kjPMaCHcjri6/0oXgpG64IdS1bD7kRrfpuxQRstQx7R8WerCEVZU/S9w69pRxoXz9Ov7KbUqMRFc3Q+d
laOR3Pl+UGmNU9M3QyLppjQYOTml+dDBLQ0iuSLpyvO8H2lbL0eTonK+ze9XLB36swOA1nootcORy1U6
RNL18Hm1juVqKXAkHXYjEy1GDq2VnhVso+vjbZwPhyxx/8mSrolAtXSrPU4SP7sYyVeG2rH8dXVJTp1D
o57VzWyRkGakBzpHSzzSPwyOuOoo7HFWKp9/iM4/wqt65vzvLRwcHB61JQN0PYsWn6MohhOgisHpyUnQ
KgD9iGGvVfz0Fk4Oy52jXjlm8UX0y3blIPUkJ5yt/cL1tIrUgsC2Uo2iVEtJFfuBlpY6ttbechO9i26i
+DyatwLnrGcZZEsgqd3MkKcDy/Rex/3cPan0E/2lKDvSC8ivKmSgp5XxByd9GZA2oepFxgOSUlRtn+2c
H/mTA9VRZb6Nnypkg2EDy66bHn60cNpgLvhLBtEZvlxiECpt+LKZAfVSXvffXCsb7M2pLlRwK/vejGhh
GSjrDtKdkh500iJUx9pncTo/b12gItMDNdxPopdBpYF33wMVBT4fSSoxAOU2k536phIbPFqp7rnrFbl7
5C5dEafvUf1A2NTXHnnCd9/gjaR+gpcvrDLuEFSlgLhNPrBjZQbtatgc1zk3m048w3O3gwHZhtbguQPR
Vt4HMxE6vUdWLkX/ApjPQmbSaw31NkEEt+5FtAfuUO5/NIQ+stVZ4irTL7qJdr6N7PlQUs3aHRIe/bNp
u/34VFhk7eCbKDt78UCwu7uzj7ndrdp1MOCiWyltB/unSP3Guo7ni5vpLF5AuronheJfivJTWP34KpX8
v6f/lb/Op/MIPn+I4kCcs3lV3IsP5aMuupxHcAqo2GHHx/87XnYCfYaPQ3gLp+Wv0GOwXRpDU3GnfsYn
o8G8qpAX8lJsU9AnaT8F6zS2XXS5uP8rBkNBN4F1u/lswXiWNac3oZ6GCnJPlXPFHdnmLFS1X9cG+sTr
USSwWwdoNLxhBxx2EdnntavZJ7RNV8gK8R+jvwujU6MVwXVu0Fqug0vhtyL3tyFvnw4B9oY4M0zfkMsw
f4N+QwS+vrqaLd5M/g4AAP//2KibD6kdAAA=
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
