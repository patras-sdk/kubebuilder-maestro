// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package convert generated by go-bindata.// sources:
// config/json-schema/full.json
// config/json-schema/limited.json
package convert

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configJsonSchemaFullJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x4b\x73\xdb\x36\x10\xbe\xfb\x57\x60\x60\x9f\x32\xa6\x25\xf7\x66\x5d\x3a\x9a\xe6\x92\x3e\xa4\x4e\xd2\x5e\xaa\x51\x3a\x10\xb9\xa4\x90\xe2\xc1\x80\x80\x2a\x55\xa3\xff\xde\xe1\x1b\x20\x41\x8b\xb4\xda\xdc\x2c\x60\x1f\xdf\x2e\x76\xbf\x5d\xfa\x7c\x87\x10\x7e\xc8\xc2\x3d\x70\x82\x17\x08\xef\xb5\x4e\xb3\xc5\x6c\xf6\x25\x93\x22\x28\x8f\x9f\xa4\x4a\x66\x91\x22\xb1\x9e\x7d\x37\x7f\x7e\x09\xe6\x2f\xb3\x4a\xfe\xb1\x50\x56\x10\x1a\x95\xd1\x03\x2c\x45\xb8\x97\x0a\x2f\x90\x56\x06\xca\x3b\x1a\xd9\x46\xff\x32\x91\x7c\x8a\xe0\x60\x5b\x0f\x62\xc3\x58\x69\x49\x53\xcd\x20\x97\xff\xe9\xf7\xf7\x6b\xf4\xe3\xa7\xf5\x2a\xf8\x64\x39\xd2\xa7\x34\xbf\xdd\x60\xb9\xfb\x02\xa1\xc6\x8f\x08\xef\xa4\x64\x40\x04\xde\x16\x02\x24\x8a\xa8\xa6\x52\x10\xf6\xab\x92\x29\x28\x4d\x21\xc3\x0b\x14\x13\x96\x95\x70\x52\xfb\x38\x8f\xbc\x81\x58\xfe\x68\x9d\xe0\x4c\x2b\x2a\x92\xc2\x71\x71\x1e\x4b\xc5\x89\xce\x6f\x8c\xa2\x81\x82\x18\x14\x88\x10\x5a\x81\x87\x50\x72\x0e\xa2\x10\x59\x49\x11\x00\x4f\xf5\x09\xc5\x8a\x24\xf9\x69\x86\x84\xd4\x88\x30\x26\xff\x86\xe8\xa9\xd5\x4a\x89\xd6\xa0\x44\xae\xf4\x79\xf3\xf9\x7e\xfb\xee\xfe\xfb\x07\x5c\x5c\x5e\x1e\x2b\x7c\xcd\xdb\x4c\xc3\xd8\xb1\x42\xea\xb7\xb9\x6a\xc5\x81\xb4\x0c\xfe\x20\xc1\x3f\xdb\x4d\x50\xfe\x31\x0f\x5e\x9e\x16\x7f\x6e\xdf\x75\x31\x2a\x88\x6f\x49\x62\xd7\x58\x55\x4f\x1f\xff\x17\xab\xcb\xc1\x4c\xd4\xd5\xd4\x58\x8e\x20\x26\x86\xe9\xba\x86\x5c\x83\xed\x7b\x0f\x21\x74\xe5\x23\x88\x33\x9f\x70\x5d\xcd\xf5\xf9\x40\x15\x9f\xfb\x99\xc1\xf7\xb8\x36\xef\x80\x3d\x5f\x1c\xcf\x75\x5b\x8d\x82\x19\x41\x16\x2a\x9a\xe6\x00\xc6\xaa\xc0\x91\xf0\x94\x81\x37\x38\xa2\x14\x39\xb5\xb1\x51\x0d\x3c\xab\x08\xa2\xe3\xb6\x06\x5f\x73\x07\x42\x98\x1b\xa6\x69\xca\x60\xed\x2d\x03\x61\xf8\x0e\x54\x6b\x1b\x8e\x21\x33\x79\x72\x7e\xa1\x82\x72\xc3\xf1\x02\xcd\x1d\x1f\x9c\x1c\xab\xf3\x21\x5b\x9d\xb0\x6a\x7b\xd3\xf4\x78\xe3\x7e\xa2\x9b\x69\x7a\x9c\x1c\x7f\x06\x91\xe8\x7d\x53\x1b\x65\x49\xcc\x8a\x42\x9b\x09\x29\x56\x90\x10\x4d\x0f\xf0\x41\x68\x48\x40\x61\x1b\xe1\x34\xd5\xf7\xe5\xeb\xcc\x5b\x13\x2d\x49\x4c\x68\x4f\x05\x09\x1c\x7b\x51\x7c\xa8\x6a\x62\x7a\x10\x93\x34\xfb\x31\x18\x41\xbf\x1a\x68\x8c\xdc\xc4\x06\x9c\x1c\x7f\x90\x42\x13\x2a\xde\x16\x8a\xad\xdc\x4c\x94\xeb\x26\x3c\xd8\x9e\xbb\xb8\x3c\x2c\x32\x05\xd9\x74\xf5\x7e\xa6\x15\x7c\x35\x54\x41\xe4\xb7\x50\x16\xcc\xb2\x60\x0a\x8b\x10\x52\x10\x11\x08\xfd\xd1\xd2\x7d\x2b\x75\x56\xf7\x9e\x9c\xda\xbe\x2b\x29\x97\x3c\x43\x29\xb2\x0e\x2d\x81\xf0\x77\xe9\x58\xae\xab\xe4\x1b\x7d\x22\x4e\x05\xc3\x6d\x1a\x98\xbe\x24\xd1\x9c\x63\x7f\x3b\xa5\x90\x59\xa4\x8f\xac\xd8\x06\x91\x38\x68\x26\x98\x76\x9b\xec\xd9\xb9\x70\x7b\xa7\x09\xb1\xcd\x1f\x42\x5b\x27\xe8\xf6\x69\xdc\xae\x1d\x1a\x68\xd8\x08\x38\x10\x66\x88\x86\x68\xa4\x06\xed\x76\xf2\x50\x62\x87\x87\xa8\x3f\xf5\xc5\x02\x56\xd7\xa7\x37\xba\xb0\xdb\xfb\x63\xc2\x9a\x30\xde\x7d\x9b\xeb\xb7\x5c\x21\x2a\xba\xf7\x76\xd5\x7f\x8e\xa2\x8a\xf5\xb4\x22\xbc\x56\xe8\xce\x90\x31\x98\x1b\x06\x29\xbf\x21\x6e\x82\xec\x10\x48\x17\xba\x97\x37\x3c\x41\xbc\x5a\xbb\xf1\x75\x19\xbd\x07\x71\x5d\x0a\x58\x06\xd7\xa5\x08\x63\xeb\x78\x80\x0e\x9c\x72\xaf\xe5\xab\x56\x1a\x2b\x2f\x05\x4c\x92\x17\x52\x8f\x00\x1d\x1d\x88\x08\xbd\xa3\x60\xe2\xb8\xde\xd3\xf1\x8b\x3b\xa3\x99\xce\x5f\x71\xac\xbc\x56\x34\xc9\x87\xe8\x48\x71\xca\xb9\xd1\x64\xe7\xdf\xd0\x27\x05\x95\x2a\x2a\x15\xd5\x27\x9f\x21\x5a\x4d\xf6\xbb\xba\x52\x2f\xe5\xf7\xb9\xf3\x51\x82\xed\xa7\xb9\x3e\xe0\x7c\xc3\xa1\x33\x67\xfa\x4f\xe9\x20\xf6\xac\x1e\xaf\x60\xb7\x1d\x7b\x17\xfc\x57\x56\x91\x9b\xd7\x2a\xd7\x93\x3d\x32\x2d\xd3\xd5\x66\xd0\xce\x9b\xde\x24\xee\xbd\x28\xea\xc7\x97\x47\x52\xff\x4f\xa4\xf9\xed\x7c\xeb\xa0\x1e\x73\x21\xb7\xb8\xba\xf3\xc9\x5e\x73\xc6\x6f\x2e\xe7\x5e\xe5\x5a\xbc\xdb\xdf\x00\x3c\x59\xdb\x6c\xdb\x82\xbb\xbb\xdc\xfd\x1b\x00\x00\xff\xff\x45\x47\xff\xdc\x6d\x12\x00\x00")

func configJsonSchemaFullJsonBytes() ([]byte, error) {
	return bindataRead(
		_configJsonSchemaFullJson,
		"config/json-schema/full.json",
	)
}

func configJsonSchemaFullJson() (*asset, error) {
	bytes, err := configJsonSchemaFullJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/json-schema/full.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configJsonSchemaLimitedJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x3d\x6f\xdb\x30\x10\xdd\xf3\x2b\x08\x36\x63\x1c\x25\xdd\xe2\x2d\x40\x97\xb4\x45\x5c\x34\xed\x64\x78\xa0\xa5\xb3\x7c\xa9\x44\x2a\x24\x65\xc0\x30\xfc\xdf\x0b\x8a\xfa\xa0\x28\x32\xb6\x83\x02\x5d\x0c\xeb\x78\xef\xbe\xde\x3b\x4a\x87\x2b\x42\xe8\xb5\x4a\xb7\x50\x32\x3a\x27\x74\xab\x75\xa5\xe6\x49\xf2\xaa\x04\x9f\x59\xf3\xad\x90\x79\x92\x49\xb6\xd1\xc9\xe7\xbb\xfb\x87\xd9\xdd\x43\xd2\xfa\xdf\x34\x60\x09\x69\x2d\x15\xee\xe0\x91\xa7\x5b\x21\xe9\x9c\x68\x59\x83\x3d\xc3\xcc\x0d\xfa\xa7\xce\xc4\x6d\x06\x3b\x37\xfa\xac\xc0\x12\x35\x64\x36\x98\x46\x5d\x80\x81\x7c\xb7\x56\xf2\xed\xf7\x97\x05\xf9\xfa\xb2\x78\x9e\xbd\x38\x39\xf5\xbe\x32\x5e\x4b\x2a\xd6\xaf\x90\x6a\x7a\x43\xe8\x5a\x88\x02\x18\xa7\xab\xc6\x81\x65\x19\x6a\x14\x9c\x15\x3f\xa4\xa8\x40\x6a\x04\x45\xe7\x64\xc3\x0a\x65\x2b\xab\x5c\xb3\x19\xc2\x10\xd5\x3e\xd9\xce\x36\xa6\x96\x4f\xc9\x75\x06\x1b\x95\x28\x2c\xab\x02\x7e\xed\x2b\x50\xb4\xf1\x39\xde\xb4\xc0\xb6\xea\x1e\xd9\x46\xa2\x4a\x4b\xe4\xf9\xd8\x39\x03\x95\x4a\xac\x4c\x75\xe7\x42\x02\xc5\xc6\x7b\xec\xce\x8d\x07\xdf\x2f\x4c\x07\xcb\xde\x44\xc8\x61\xda\x57\x2e\x45\x5d\xd1\x2e\x5b\xd4\xad\x62\x92\x95\xa0\x41\x52\x72\xec\x3d\x57\xed\x3f\x6b\x31\xbf\x47\xcb\xbc\x81\x0c\xa3\xb5\x29\x02\xed\x76\x04\x9e\x68\xaa\x27\x2e\x3e\x8f\x29\x81\x8d\x2d\x15\x5c\x69\x27\x53\x7f\xe6\xb4\x3b\x21\x30\xca\x88\x0f\x0c\x93\x79\x36\xbc\x92\x28\x24\xea\x7d\x0c\x8b\x5c\x43\x0e\x32\x0c\x96\xf0\x56\xa3\x04\xb3\x61\x01\xb6\x6c\xda\x47\x29\xd9\x9e\x7a\x39\x83\xc3\x8b\x73\xf2\x1e\x2f\x2e\x3e\x22\x38\x72\xbe\xe8\x22\xae\x41\xe1\x11\x47\x7c\xde\x64\x1a\x5e\x36\xac\x2e\x0c\xef\x87\x01\x72\x1c\x6b\xb5\x5b\xae\x3e\xfa\xbf\x92\xa7\x43\xcc\xd2\x13\xa7\xc3\x43\x81\x4a\x3f\xb3\x12\x3a\x72\x57\xa7\xe5\xfd\x9f\x54\x3a\x0c\xb3\xbb\xd8\xa7\x8d\x5e\xa2\x40\xe0\x75\x19\xab\x83\x35\xee\x23\x2e\x51\x43\xa9\xda\xe4\xe1\xed\x0d\xac\x7d\x58\x89\xa1\x2a\x9d\x4b\xdd\x57\xe3\xc1\xd3\xe6\x3b\x55\x8e\x2a\xbd\x30\x0d\x21\xb4\x44\xfe\xd4\x82\xef\x27\x87\x35\xc7\xb7\x1a\x9e\x42\x63\x20\x8e\xac\xc9\x68\x23\xdc\x01\xb1\x6c\xc7\x78\x6a\x79\x0a\x0c\x7d\x78\x79\x46\x76\xa8\x91\x77\x30\xf2\x16\xb9\xfe\x88\xa4\x7a\xf1\x7f\x00\xab\x25\xe6\xf9\x68\x5d\xcf\x86\x62\x59\xd6\x9a\xad\xe3\x5b\x74\xd9\x20\x82\xf7\x89\xfd\xac\xb1\xb2\x0f\xdc\x28\x63\xe9\x04\x89\xf7\x74\xd4\x7e\x5d\xfd\x6c\x05\x45\xfd\x7c\x8e\xb2\x9c\x7c\xed\x92\x39\x17\x90\x2f\xda\x40\xb3\xfd\x0b\xc7\x31\xf1\xba\x28\xc6\xcf\xe5\x7a\xec\x31\x79\x5f\x78\x1c\xac\xc6\xf5\x3a\xd7\xc2\xe9\xf9\x38\xa3\xf0\x08\x1e\x88\x0d\x2c\x48\x7f\x34\x70\xb7\x5c\x0d\x9f\x28\x57\xc7\xab\xbf\x01\x00\x00\xff\xff\x1d\xf9\xe7\x5f\xf9\x0a\x00\x00")

func configJsonSchemaLimitedJsonBytes() ([]byte, error) {
	return bindataRead(
		_configJsonSchemaLimitedJson,
		"config/json-schema/limited.json",
	)
}

func configJsonSchemaLimitedJson() (*asset, error) {
	bytes, err := configJsonSchemaLimitedJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/json-schema/limited.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"config/json-schema/full.json":    configJsonSchemaFullJson,
	"config/json-schema/limited.json": configJsonSchemaLimitedJson,
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
	"config": &bintree{nil, map[string]*bintree{
		"json-schema": &bintree{nil, map[string]*bintree{
			"full.json":    &bintree{configJsonSchemaFullJson, map[string]*bintree{}},
			"limited.json": &bintree{configJsonSchemaLimitedJson, map[string]*bintree{}},
		}},
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
