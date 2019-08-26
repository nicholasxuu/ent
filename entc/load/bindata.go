// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package load

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x41\x6b\xdc\x30\x10\x85\xcf\x9a\x5f\x31\x15\x5b\x22\x81\xa3\xd0\x6b\x61\x4f\xcd\x1e\x52\x68\x52\xd8\x40\x0f\xdb\x25\xc8\xf6\x78\x23\x6a\xcb\xae\xa4\x94\x06\xa1\xff\x5e\x24\xd9\x0b\x3d\xd9\xf2\x7b\xfa\xe6\x3d\x4f\x8c\xd8\xd3\x60\x2c\x21\x9f\xb4\xb1\x1c\x53\x82\xbb\x3b\xfc\x32\xf7\x84\x17\xb2\xe4\x74\xa0\x1e\xdb\x77\xbc\x21\x1b\xba\xeb\xa7\x1b\x85\xf7\x4f\xf8\xf8\xf4\x8c\x87\xfb\x87\x67\x05\x8b\xee\x7e\xe9\x0b\x61\x66\x00\x98\x69\x99\x5d\x40\x01\x8c\xcf\x9e\x03\xe3\xed\x7b\xa0\xfc\x12\x23\x06\x9a\x96\x51\x07\x42\x5e\x5d\xbe\x8c\x2c\xd2\xe2\x8c\x0d\x03\xf2\x8f\xbf\x39\xaa\xef\x2b\x31\x25\x90\x00\x31\xe2\xae\xd5\x9e\xf0\xf3\x1e\xcb\x73\xd3\xf3\xdd\x3f\xda\xa1\xef\x5e\x69\xd2\x1e\xf7\x78\x3a\x93\x0d\xea\xc1\x06\x72\x83\xee\x28\x16\xb4\xd3\xf6\x42\xb8\x7b\x69\x70\x67\xf5\x54\x30\xea\x51\x4f\xe4\x33\x9f\xb1\x18\x6f\x57\x7e\x4a\x2a\x1f\xae\x51\x7c\x4c\x7c\xbd\x93\x52\x53\x58\x64\x7b\xbc\x4d\x09\x12\xc0\xf0\x66\xbb\xd2\x59\x48\x8c\xc0\x72\x90\xd1\x58\xf2\x78\x3a\x9f\xce\xb9\x34\xb0\x61\x76\xf8\xd2\xac\xf9\xf2\xdc\x1a\x65\xcb\x1b\x81\xb1\xb6\x41\x72\x2e\x6b\xdf\xb4\xf3\xaf\x7a\x3c\x16\x51\x54\x8f\x04\xc6\xcc\x50\x1c\x1f\xf6\x68\xcd\x58\xee\xb0\x41\x9b\x51\x90\x73\x59\xce\x15\xea\xdc\x3d\xea\x65\x21\xdb\x8b\x72\x6c\xb0\x95\x90\xd5\xd9\xab\x63\xe8\xe7\xb7\xa0\x7e\x38\x13\x48\x94\x7d\xa8\xaf\xb3\xb1\x9b\xb1\xc6\x15\xfc\xa7\xe5\x52\xca\x6b\xb7\x6d\x4a\x1e\x3f\xbb\x52\xb2\xb2\xc8\xb9\xca\x3a\x06\x67\xec\x25\x7b\xd4\x21\x7b\x84\x94\xc5\x73\xf8\x6b\x82\xf8\x54\x48\xff\x6d\xbd\x96\xaa\x4b\x5f\x7f\x66\x4a\xf0\x2f\x00\x00\xff\xff\xb5\xb1\x2f\xf6\x87\x02\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 647, mode: os.FileMode(420), modTime: time.Unix(1566131522, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xdd\x8e\xdb\xb6\x12\xbe\xb6\x9e\x62\x8e\x81\x2c\xa4\xc0\xd0\xde\x2b\xf0\xc5\xc1\x39\x5b\x74\x5b\x74\x53\x24\x6d\x6f\x82\x20\xa1\xa5\x91\xcd\x44\xa2\x1c\x92\xf6\x26\x59\xf8\xdd\x8b\x19\x0e\xf5\x63\x39\x1b\x20\xc1\x06\x08\x56\x9c\x3f\xce\x7c\xf3\x71\x48\xef\x55\xf9\x51\x6d\x11\x9a\x4e\x55\x49\xa2\xdb\x7d\x67\x3d\xa4\xc9\x62\x89\xa6\xec\x2a\x6d\xb6\xd7\x1f\x5c\x67\x96\xc9\x62\x59\xb7\x9e\xfe\x58\xac\x1b\x2c\xfd\x32\x49\x16\xcb\xad\xf6\xbb\xc3\x26\x2f\xbb\xf6\xba\x56\x25\x6e\xba\xee\xa3\x36\xe5\x61\xa3\x7c\x67\xaf\xd1\xb0\xfd\xf7\x6c\xae\x5d\xb9\xc3\x56\x5d\xd7\x1a\x9b\x6a\x99\x64\x49\x72\x7d\x0d\xaf\x59\x06\x16\xf7\x16\x1d\x1a\xef\x40\x19\x40\xe3\x73\x51\xf8\x9d\xf2\x70\xaf\x1c\xa7\x8d\x15\xd4\xb6\x6b\x41\x41\xd9\xb5\xfb\x46\x63\x05\x07\x87\x16\xa4\xb4\x3c\xf1\x5f\xf6\x18\x43\x3a\x6f\x0f\xa5\x87\x87\x64\x71\xa7\x5a\x04\x00\x92\x68\xb3\x05\x80\xf7\x54\x69\xb1\x34\xaa\xc5\x55\xd7\x6a\x8f\xed\xde\x7f\x59\xbe\x4f\x16\x37\xd5\x16\x1d\x00\xbc\x79\xfb\x9c\x3e\x7b\x4b\x24\xf9\xd4\xf4\x17\xaa\xc2\xb1\x29\x7f\x46\x53\xae\xee\xcc\xf6\xd6\x54\xf8\x19\x1d\xd9\xf2\x67\xb4\xd5\x41\x3e\x31\x3e\x31\x2c\x21\xe4\x1c\x95\x20\xff\x01\x50\x82\xe3\x1c\x93\x31\x2c\x8f\x00\xf3\x17\xc5\x08\xff\xb8\xc0\x9c\x05\x62\x4e\x1b\x9c\x99\xab\x2d\x7c\x33\xba\x57\xdb\xa9\xf5\x6b\xfd\x35\x06\x7f\xae\x8d\x97\x4f\xb1\x76\xfa\xeb\x59\xf0\xff\xed\x94\x75\xc8\x66\xcf\x87\xe8\x62\x5e\x06\xe5\xd4\xe3\x6f\xa3\x3f\x1d\xc2\x16\x9b\xae\x6b\xa6\x1b\x1c\x58\x39\x75\xb8\xd3\x4d\xa3\x36\x0d\x5e\x74\x30\xa2\x9c\xba\xbc\xdc\x7b\xdd\x19\xd5\x5c\x74\xe9\x44\x39\x75\xf9\x3f\xd6\xea\xd0\xf8\xcb\x69\x55\x41\x79\x46\xa4\xb6\x3d\xf8\x90\xd8\xdc\x43\x47\xe5\xd4\xe7\x1f\xd5\xe8\x8a\xce\xa0\x83\x01\xdc\xe8\x73\xec\x95\x17\x48\xc8\x47\x60\xce\x41\x16\xff\x00\x05\xd9\xef\x02\x03\xa5\x87\xdf\xe7\xde\xd4\xf0\x11\xd6\x9d\x19\x9e\xf3\xed\x15\xd6\x61\xf3\xa9\x9d\xc5\xfa\xdd\x7c\xf7\x57\x58\x0b\x35\x27\x13\xc1\x62\xfd\x0d\x8e\x49\x63\x1e\x61\xd7\xad\x39\xa2\x75\x78\x6e\xaa\x83\xf8\x7c\xfb\x4f\x07\x6d\xb1\x3a\xb3\xb5\x22\xbe\xd0\xb5\x30\x61\xe6\x6d\x0b\xf2\x1f\xe8\x5b\x70\x1c\x1a\x27\x95\xf6\x04\x7c\xa4\x52\x99\xa7\x6f\xde\x4e\x91\xfe\xf6\x38\x3d\xb7\xbc\x30\x4d\x43\x95\x77\x78\xcf\xfd\x28\x2d\x2a\x8f\x5c\xa4\x54\x44\xc1\x43\x59\xe1\xbe\x21\xd6\xa3\xa5\x1b\x29\x4f\xea\x83\x29\xa3\x6b\x8a\x3d\x9b\x33\x69\xee\x43\xb2\x30\x08\xc5\x1a\xae\x68\xf9\x90\x2c\x98\xa4\x05\x57\x89\x39\x7d\xa7\xd9\x2a\x59\x30\x23\xa3\x94\xbe\x45\xaa\xb6\x41\x48\x52\xb5\x0d\x42\xe1\x5a\x41\x42\xf9\x0e\x8a\x80\x62\xc1\xd6\xb7\x2e\xac\x82\x46\xe8\x51\x04\x8d\xac\x62\xb4\xd0\xf6\x82\x55\x71\xc5\xba\x53\xb2\xd0\x35\x70\xf6\x98\xff\xd7\xb9\xae\x4c\xb3\x17\x80\xf0\x9f\x35\x18\xdd\x50\x69\x0b\xc3\x29\xc0\x7a\x40\x20\x63\x3f\x8b\xfe\x60\x0d\x18\x14\x6c\xff\x50\xd6\xed\x54\x23\xf7\x28\x3f\x0f\xe8\xc0\xe3\xf8\x5e\xee\x41\xa5\xaf\x0e\x14\xfc\xf6\xfa\xe5\x1d\x39\x33\xc1\x4a\x65\x60\x83\x50\x21\xb9\x56\xc1\x84\x02\x88\x73\xb7\xf9\x80\xa5\x97\x3f\xd2\x95\xc9\xa6\xa9\x8b\x7b\x13\x6f\x65\xa7\x0c\xd2\x0d\xbc\x79\xbb\xf9\xe2\x71\x05\x68\x2d\xfd\xef\x6c\x46\xa5\x39\x6e\x5a\xf0\x7d\x08\x78\x6b\x53\x69\x8b\xa5\x4f\xe5\x19\xc3\x8d\x7a\x59\x4b\xe4\x2c\x93\x76\x9e\x92\x85\x70\x8c\x43\x16\x6b\x70\xaa\xc6\xc0\xc6\x68\x1b\x90\xb5\x76\x8c\x65\xc4\x4c\x37\x2b\xa8\x5b\x9f\xdf\x50\x2e\x75\xba\x94\xc4\x9f\x7d\x2a\xe0\xd9\x71\xb9\x02\xc7\xfb\x70\xf0\x00\x76\xdd\x59\x78\xb7\x82\x9a\xb6\xb2\xca\x10\x57\x03\xf5\x29\xaa\x63\xf1\x15\x6f\x4f\xeb\x11\xff\x00\xea\x11\x03\x47\x14\x24\xc5\x40\xc2\x11\x0b\x59\x11\x79\x38\xe2\x1b\xc9\xa7\x8c\x8b\x37\x51\x11\x94\xbf\x2a\x27\x02\x51\xc7\xeb\xb0\x10\xdf\xb8\x16\x75\xbc\xfa\xa2\x3a\xae\x45\xdd\x5f\x5a\x45\x50\xf7\x6b\xd1\x0f\x17\x54\x01\x0d\x9a\xb4\xce\x07\x49\x9a\xb1\xcd\x29\x59\x50\x0f\xdc\x0a\xba\x8f\x84\x50\x9d\xa7\xe1\x19\x42\x4f\x07\x9b\xbd\x20\x31\xe3\x45\xaf\x05\x6e\x22\x6b\xd2\x8c\x65\x35\x2f\x60\x0d\x57\xa4\x1e\xc2\x95\xb3\x70\xf2\xb4\x98\x84\x94\x17\x05\xd9\x95\xd1\xa0\x0f\x1c\xdf\x22\x6b\xb8\x12\x3b\x09\xef\x72\x99\x68\x6b\x50\xfb\x3d\x9a\x2a\x8d\x92\x15\xb8\x3a\x50\x21\x8c\xc1\x31\xef\x78\x5e\x3e\x25\xed\x70\xa0\x1d\xef\x1e\x58\x97\x87\x39\x3d\x4a\xf5\x26\xa4\x36\x0c\x8a\x10\x25\x3e\x58\xc7\x39\xcb\xe3\xf6\x29\xb3\xd6\xd5\xe7\x21\x6f\xc9\x41\x32\x8f\x4f\xeb\x51\xee\xb7\x31\xc9\x2b\xfe\xe2\x26\x72\x3d\x05\x50\x24\xae\x8d\xf6\x71\xc2\xbf\xd0\x96\x82\x75\xfc\x3d\x56\xc6\x63\x43\xca\xe9\xa1\x39\x4d\x86\x27\x5d\x57\xb9\xcc\xb0\xd4\x65\x32\x49\x87\x59\x02\xf7\x56\xed\x1d\x0f\xc1\x50\x75\xe4\x47\x8b\x7e\xd7\x55\x70\xaf\xfd\x0e\x2c\x96\xdd\x11\x2d\xf8\x0e\xd0\xb8\x83\x45\x30\x1d\xec\x95\xd1\x25\xbd\xdd\xa0\x0d\xe1\xb5\xd9\xca\xcc\x9c\x8d\xaa\xd9\xc0\xac\xe3\xb5\xda\xff\x78\x38\x1f\x9d\x15\xd6\x68\x81\xc2\xa5\xbc\xa6\xf6\x1d\x19\xed\x90\x0c\xdd\x22\xc7\x71\x33\x17\xe4\xbf\xbe\xd0\xc7\x58\x51\x48\x58\x5a\x7a\xe4\x93\x52\xc7\xb3\x60\x74\x13\x0e\xc8\x89\x8e\x90\x60\x37\x71\x4f\xb3\x15\x5b\x0d\x00\x06\x72\xce\xf0\x0b\xe2\x9f\x85\x6f\x7c\xe2\x66\xe8\x85\x23\x12\xc0\x23\xc3\x27\xc4\x2e\x54\x73\x01\x3a\x94\xa3\xf9\x18\x72\xa1\x88\x19\x70\xf1\x6c\xcc\xa0\x8b\x8a\x9f\x05\x6f\x7a\xf4\x67\xf0\xe9\xfe\x67\x6f\xff\xfe\x7c\x42\x04\x63\x51\x17\x30\xd4\xfd\x90\x78\x0c\xc5\x58\xcd\x80\x23\x17\xda\x3f\x24\x3c\x8c\x9f\x12\xd9\x64\x45\xb9\xd1\xb0\xf2\xf9\xef\xda\x54\x69\x06\xeb\x75\xaf\xff\xd3\x5b\x4e\x9d\x6e\x09\x9f\xdf\x34\xd8\xa6\x93\xd1\xe1\x93\x53\xf2\x6f\x00\x00\x00\xff\xff\x0b\xa3\x7d\x9f\x9d\x11\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 4509, mode: os.FileMode(420), modTime: time.Unix(1566736551, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
