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
		size:    318,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/5SQvQrDIBRGd5/C93AyxhahaUJ1aKcgeiGB60+j9vk7l1Ka7Od8fJxOntWVEdLfxoka
3l0kVScq70obTW2raYOMKxQquBa8l7/Q4hbwDcHPLoVgo/9rbJDB1gPCs6W648gaMsLuUbfYGAG/ucf0
UcI5KGVGeAEyQsQ4DMow8g4AAP//d8ml+j4BAAA=
`,
	},

	"/20190504194418_initial_schema.up.sql": {
		name:    "20190504194418_initial_schema.up.sql",
		local:   "20190504194418_initial_schema.up.sql",
		size:    5153,
		modtime: 0,
		compressed: `
H4sIAAAAAAAC/+xY3W7cNhO936eYO9tAHDjXxvcBa5tJF7HldLNGahQFQYkjLVv+qCRl77bouxfUj1eW
qF0HSIAW7uVqzpzhcA55aJ+ewmotHORCIggHrkp/xcyDN5CtmS4QKu2FBL9GyIV1HiwyCRxLabYKtQeT
18HU+Lez01O4MqBNQMktGB1CDsFla1QMlCgs88JoBxYVE1roAjKjnXA+UHWlhAtMCp1jBcKaOUgRdcgx
D8jfzmYX5MMiOZ/NLpdkviKwuv9EgGUZOkclPqCE+Wcgyd0NHM8AAI7wAe3WaDx60/x2VeoyK1K03Rdl
OFrmzdOH1BrGM+b8DsO4EvpodtIrPL+4JvVGaZSurSY4pKIQ2sMHkpDlfEWu4OIersj7+d31KixtcUWS
1WJ1D5+Wi5v58h4+kvumRmaReeSUefBCofNMlf6Pp9zk9svxCSS3K0jurq+blKrkX5PS5Di0dLfOLgp3
yeLHO9LwaqYQPG78oF5qPJ2KscyLh6AFI5HpQbC0mItNlLKSEn0daRcoCm0s8vrTz788NXP0519Hg9ys
ct4oah41WvdyvDL8K9AWi0qyg/xNijGSm0cNQvuGRDLnaWaUYno0pkFyDc1VbJNKZh3Sral8lU5tMW68
ZVSKHMN0+6Nt4pbluUSKmqUSeYSkRrm1qSSn7ZGYqsWFKyXb0kdmwzl2U0uqax1ChS0xlae8ai6IyNpb
olxIHyY9sfYmTKXQv03VKtEq4cORaVAvmGnLmrFyirSHoEpomq2DXMZd9HEl2gy1D3fcfmBNGGo/h/Vx
qIzHA4trMFSxzXS9FuOELiQe2GS3VamRB4q2oBc122FDvx35dMsK91dWbEMl6sKv95Ck4ermtFxb5g7t
33MsLZn3aPWLBHT5A7n8CMcT078gqy+EJHAGTHN4d3Z28iaSFdnJA4mjU/X//8HZSd/AFskV+enJwGjr
C1TwTTDvnbG1gZH1OaFKid3l9i9ywLa1ngkuyXuyJMkl+dzrW/BR5pT3dc+VqC32nibPfgzcxlTaD025
W3DYF2Nj9MhFJFKHGkOH4127b+oGIiIYjJLuUmo53CbjYe8gJ+f7yULNTlUjmnpBQ2X9XtUX2msQVKUi
V2Ldf3Sm31AJlYoIodn5yPy7kUyPvU3VleqG3eWEUsMRWyyb0byS66N+Wz0Tfzx/dED6NE9v9N6jiE89
0VCybURb7VVFucjzrn7X7rsBtn28frOLaay3kQwi0otIZVqFY8LRru/hHWHHppetkVfyP+F+N+Fm1miK
m9Kic+HdMm23/xANjyUR88+IbvZY6JhyUsYx5sM6ZpU3FkspXrPPeiuKAqOjN1YU3Su//edE/Uc0utJo
F33pNSL7HgrrDSsirWejnNZUn6Qz6WF259S3NzeL1fns7wAAAP//E5JSxCEUAAA=
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
