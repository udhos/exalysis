// Code generated by go-bindata. DO NOT EDIT.
// sources:
// concat-not-needed.md (168B)
// extensive-for-loop.md (156B)
// itoa.md (88B)
// loop-map.md (484B)
// many-loops.md (146B)
// plus-equal.md (59B)
// remove-extra-bool.md (166B)
// strings-builder.md (238B)

package tpl

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _concatNotNeededMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcc\xc1\x71\x83\x30\x10\x85\xe1\x7b\xaa\x78\xf7\x4c\xe8\x80\x02\x52\x42\x6e\x5a\xd8\x87\x50\x2c\xb4\xcc\x6a\x65\x0f\xdd\x7b\xf0\xfd\xff\xbf\x1f\xfc\xd9\xc0\x6a\xa3\x2a\xd6\x9d\xeb\x03\x97\x0d\xc7\xe8\x92\x09\xdb\x90\xbe\xe7\x84\xd2\xe0\xcc\xe2\xda\x11\x86\xd7\x4e\x27\x4a\xa0\x74\x38\xa5\xd6\x0b\x8d\x54\xea\x84\xdf\x0d\xb1\x13\x21\x9e\x19\x78\x8a\x17\x59\x2a\xef\xf0\x2b\x0f\x71\x69\x41\xea\x6d\x2c\x04\x8f\x33\x2e\x08\x7a\x39\xce\x4a\xa4\x39\x41\x8d\xfd\x03\xfc\xdb\x02\x69\x7a\x8f\x9b\xf4\xa0\x4f\xef\x00\x00\x00\xff\xff\x96\xb6\x64\xbe\xa8\x00\x00\x00")

func concatNotNeededMdBytes() ([]byte, error) {
	return bindataRead(
		_concatNotNeededMd,
		"concat-not-needed.md",
	)
}

func concatNotNeededMd() (*asset, error) {
	bytes, err := concatNotNeededMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "concat-not-needed.md", size: 168, mode: os.FileMode(420), modTime: time.Unix(1541844199, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7e, 0xff, 0xdf, 0xb, 0x87, 0x21, 0x84, 0xc4, 0x5e, 0x22, 0xf3, 0xa1, 0x52, 0xfa, 0x73, 0x41, 0xa, 0xa0, 0x96, 0xe9, 0x66, 0x76, 0x49, 0xa8, 0xef, 0x81, 0xcc, 0x7d, 0x24, 0x4c, 0x62, 0xa4}}
	return a, nil
}

var _extensiveForLoopMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcc\x31\xae\xc2\x40\x0c\x84\xe1\xfe\x9d\x62\x0e\xf0\xa0\x41\x88\x32\x05\x15\x47\xa0\xb3\x95\x1d\x48\xa4\x8d\x1d\xd9\x59\xa4\xbd\x3d\x4a\x68\x67\xf4\x7f\x27\x3c\x90\x24\xba\x37\x68\x10\x2d\x67\x7b\x43\x21\x2f\x0f\x41\x75\x5f\xb1\x4d\xba\x61\x4e\x14\xdf\xaf\xc5\x83\xfb\x64\x08\x6a\xad\x1d\x46\x16\x96\x33\xee\x6a\x87\x32\x2f\x6b\xf8\x87\x70\x3b\xca\xe1\xef\xe9\x0d\xa3\xb7\x5a\x7e\x9c\x5b\xed\xd8\x26\xc2\x38\x32\x53\xa3\x63\xd4\x64\x42\x2e\xf2\x0f\xb9\x0a\xd4\x0a\xe4\x26\xc3\x37\x00\x00\xff\xff\xd4\x78\x2e\x62\x9c\x00\x00\x00")

func extensiveForLoopMdBytes() ([]byte, error) {
	return bindataRead(
		_extensiveForLoopMd,
		"extensive-for-loop.md",
	)
}

func extensiveForLoopMd() (*asset, error) {
	bytes, err := extensiveForLoopMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "extensive-for-loop.md", size: 156, mode: os.FileMode(420), modTime: time.Unix(1541844199, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1c, 0x23, 0x7, 0xfd, 0x4c, 0x31, 0x4f, 0xc1, 0x4e, 0xc2, 0xa0, 0xb4, 0x9e, 0x56, 0x59, 0xd9, 0xf, 0x15, 0x84, 0x37, 0x23, 0x54, 0xff, 0x59, 0x76, 0x4c, 0xb2, 0x31, 0xeb, 0x84, 0x97, 0x5b}}
	return a, nil
}

var _itoaMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xcb\xd1\x09\xc2\x40\x10\x04\xd0\x56\xa6\x01\x53\x87\x76\xb1\x43\x4c\x74\xd1\xec\xc0\x66\x38\xb1\xfb\x70\xbf\x0f\xde\x0d\x77\x8e\x0d\xc4\x57\xfa\x80\x46\x9c\xee\x55\x35\x96\x87\xc5\xc0\xae\x86\xdf\x1b\xf2\x99\x3a\xe8\x5c\xf1\xe3\x1f\x16\xdc\xac\x73\x57\x1f\x60\x21\xb2\x1c\x53\xe7\xce\x7a\xc5\x72\x05\x00\x00\xff\xff\x6a\xd0\x65\x71\x58\x00\x00\x00")

func itoaMdBytes() ([]byte, error) {
	return bindataRead(
		_itoaMd,
		"itoa.md",
	)
}

func itoaMd() (*asset, error) {
	bytes, err := itoaMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "itoa.md", size: 88, mode: os.FileMode(420), modTime: time.Unix(1541844199, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x89, 0xf7, 0xe0, 0x3, 0xe6, 0xf8, 0xe1, 0xef, 0x9c, 0xe4, 0x48, 0x3f, 0x87, 0x61, 0xb3, 0xba, 0x9a, 0x5e, 0x6f, 0x40, 0xfb, 0x29, 0xd8, 0x6e, 0x94, 0x77, 0x26, 0xd7, 0x25, 0xca, 0x27, 0xa6}}
	return a, nil
}

var _loopMapMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\xbb\x8e\x1b\x31\x0c\xec\xfd\x15\xd3\x25\x29\x76\xdd\xdf\x07\x04\x48\xef\x2e\x08\x70\xb4\xc5\x95\x08\x48\xe4\x82\xe2\x5d\xbc\xfe\xfa\x40\xeb\x38\x76\xda\xd1\x70\x5e\x9a\x70\x2a\x0c\x09\x76\x0a\x31\x85\x79\x62\x87\x2d\x68\xb4\x76\x88\x22\x1b\xa4\x43\x2d\xb0\xc8\x75\xc6\xa9\x48\x1f\x40\xfa\x60\x84\x21\x0a\x43\x29\x3e\x9c\xc7\x4d\xa1\x5e\xf6\xc3\x37\x88\x76\xf6\x10\xcd\x43\xbb\x75\x1c\x2e\xce\x14\xdc\x41\x0a\xf2\xb3\x84\x93\x6f\x7f\xdd\xce\x1b\x12\x2f\xa2\x32\x12\xcc\x38\x19\x0a\xd7\x15\x89\x3f\xb9\xda\xca\xde\xe1\x4c\x55\x6e\x8c\x18\xee\xbf\x25\xca\x30\x16\x47\x70\x0f\x24\x0a\x02\x7f\xb2\x42\x96\x3d\xd0\xe1\x09\x4b\xdf\x91\x4e\x8d\x07\xc5\x37\x84\x34\x46\xb6\x2f\xfd\x55\x5f\x34\x58\x87\x3b\xd5\xba\x81\x52\xe2\x84\x6e\x8d\xe1\xa4\xc9\x9a\xdc\xf6\x68\x8f\xc2\x2f\x6b\x2d\x38\x3c\x87\x9a\xf1\xdd\x1c\xcd\x9c\x21\xba\x98\xb7\x3b\x89\x34\x8d\xd6\x7c\xa5\xb6\x56\x1e\x5d\xd2\x2e\xf3\xfe\xe3\xff\xd5\xdf\xd1\xf9\xf2\x90\xdd\x7b\xfc\xcc\xf6\xef\x1f\x68\x7f\xfa\xf5\xb5\x44\xac\xfd\xed\x78\x3c\x57\xcb\x73\xb6\x4a\x9a\x67\xf3\x7c\xcc\x36\x0d\xea\x24\x3a\xdd\xa9\xdf\x40\x1e\x72\xa9\x3c\xff\x09\x00\x00\xff\xff\x52\x12\xf6\x07\xe4\x01\x00\x00")

func loopMapMdBytes() ([]byte, error) {
	return bindataRead(
		_loopMapMd,
		"loop-map.md",
	)
}

func loopMapMd() (*asset, error) {
	bytes, err := loopMapMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "loop-map.md", size: 484, mode: os.FileMode(420), modTime: time.Unix(1541844199, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xec, 0x70, 0xf, 0xed, 0xf7, 0x1a, 0xcc, 0xa2, 0xbd, 0x8c, 0x8e, 0x29, 0x71, 0xc7, 0x60, 0x8c, 0x58, 0x15, 0x8f, 0x9e, 0xa1, 0x76, 0x37, 0x85, 0xb8, 0xa3, 0xd4, 0x82, 0xd7, 0x52, 0xe2, 0x6d}}
	return a, nil
}

var _manyLoopsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcc\xc1\xad\x83\x40\x0c\x84\xe1\xfb\xab\x62\x1a\x78\xb4\x90\x03\x2d\xa4\x01\x47\x2c\x60\xc5\x78\xd0\xda\x16\xa2\xfb\x28\x7b\xc8\x71\x46\x9f\xfe\x7f\xcc\x5c\xd4\x37\x68\x40\x60\x4c\xc8\x8b\x95\x58\xd5\xc7\x2d\x08\x3d\x4e\x6b\x08\x5a\xa5\xd2\x27\x3c\x77\x8d\xdf\x44\x45\x0b\x1c\x65\xa9\x5f\x65\xe4\x19\x13\x66\x71\xdc\x2c\x64\xbf\x21\xbe\xe0\x62\x7f\xe3\xef\xd2\xdc\x47\xd0\x37\x6b\x58\xd9\x07\x07\x3b\x9c\xde\x20\x09\x31\x7b\x7c\x02\x00\x00\xff\xff\x31\x8f\x37\xf5\x92\x00\x00\x00")

func manyLoopsMdBytes() ([]byte, error) {
	return bindataRead(
		_manyLoopsMd,
		"many-loops.md",
	)
}

func manyLoopsMd() (*asset, error) {
	bytes, err := manyLoopsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "many-loops.md", size: 146, mode: os.FileMode(420), modTime: time.Unix(1539184320, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5, 0x1f, 0x5b, 0xd7, 0xf4, 0xb2, 0xe, 0x23, 0x8, 0x59, 0x56, 0xba, 0x80, 0xba, 0xc8, 0xfc, 0x4a, 0xda, 0x52, 0xad, 0xee, 0x92, 0x2d, 0xea, 0x3a, 0x6, 0xe0, 0x4b, 0xa8, 0x3c, 0x4d, 0x6e}}
	return a, nil
}

var _plusEqualMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x55\x70\xc9\x4c\x51\xa8\xcc\x2f\x55\xc8\xce\xcb\x2f\xb7\x52\x48\xcf\x57\x28\x2e\x2d\x28\xc8\x2f\x2a\x29\x56\x28\xc9\x48\x55\x48\x28\x2e\x48\x4d\x4d\xce\x50\xd0\xb6\x55\x50\x0a\xc8\x49\xcc\x4b\x57\x4a\x50\x28\xae\xcc\x2b\x49\xac\xd0\x03\x04\x00\x00\xff\xff\x9e\x43\x35\xee\x3b\x00\x00\x00")

func plusEqualMdBytes() ([]byte, error) {
	return bindataRead(
		_plusEqualMd,
		"plus-equal.md",
	)
}

func plusEqualMd() (*asset, error) {
	bytes, err := plusEqualMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plus-equal.md", size: 59, mode: os.FileMode(420), modTime: time.Unix(1541844199, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x41, 0xff, 0x90, 0xf5, 0x14, 0xc3, 0x56, 0xc5, 0x52, 0x26, 0x40, 0xf6, 0x6b, 0xf4, 0x9a, 0xd0, 0x49, 0x11, 0xdd, 0xf, 0x3a, 0xff, 0xee, 0x91, 0x4, 0xf6, 0xd5, 0xa4, 0x33, 0x69, 0x63, 0xc0}}
	return a, nil
}

var _removeExtraBoolMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\x3d\xb2\xc2\x30\x0c\x04\xe0\xfe\x9d\x62\x9b\x57\x92\xfb\xd0\x45\xd8\x0b\xd6\xe0\xd8\x19\x49\x0e\xe4\xf6\x0c\x49\x45\xff\xed\xcf\x05\xd7\x3e\x20\x46\x0c\xd7\xf6\x80\x34\xf0\x1d\x26\xd8\xc4\x54\x6e\x95\x98\xff\x7d\x46\x74\xa4\xc2\xf4\xc4\xab\x30\x0a\x0d\xd2\x76\xf4\x3b\xa2\x10\x49\x9c\x0e\x59\xd7\xaa\xcc\xd3\xd1\x97\xfa\xa8\x19\xda\x3c\x28\x19\x7f\x67\x54\x4f\xee\x61\xdf\x21\x75\x78\x68\xad\xe0\xb2\xc6\x0e\x69\x19\xc6\xa5\x6f\x3c\xd0\xef\x87\xe9\x13\x00\x00\xff\xff\xbb\xd1\x4d\xdd\xa6\x00\x00\x00")

func removeExtraBoolMdBytes() ([]byte, error) {
	return bindataRead(
		_removeExtraBoolMd,
		"remove-extra-bool.md",
	)
}

func removeExtraBoolMd() (*asset, error) {
	bytes, err := removeExtraBoolMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "remove-extra-bool.md", size: 166, mode: os.FileMode(420), modTime: time.Unix(1541844199, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd7, 0x3, 0x4a, 0x24, 0x28, 0xf2, 0x42, 0xa3, 0xbf, 0xd8, 0x3f, 0xb5, 0x1d, 0x6, 0xcf, 0xb9, 0xc0, 0x15, 0x3c, 0x57, 0x70, 0xe3, 0x9a, 0x33, 0xaf, 0x97, 0x2b, 0xab, 0x8c, 0xc5, 0xff, 0x0}}
	return a, nil
}

var _stringsBuilderMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8f\x3b\x4e\x04\x41\x0c\x44\xf3\x3d\x45\x25\x44\x88\x09\x90\x08\x39\xc0\x1e\x81\xac\x3d\x3d\x5e\xc6\xda\xfe\x48\x2e\x37\x62\x12\xce\x8e\x7a\x10\x9b\x39\x78\x7e\x55\xf5\x82\x2b\xa8\x8a\xa3\x0f\x88\x2b\x06\xad\x7d\x42\x90\x9e\x98\x16\x5c\x1b\x62\x37\x42\xbf\xd5\xb3\x51\x11\xbb\xc4\xe4\x2b\x21\x58\x2d\xd0\xbf\xd4\xef\x56\x0a\xa4\x6d\x48\xcf\xef\x09\x92\x63\x48\x29\x07\x2e\xb9\x37\x8e\xaa\xc4\xcf\xeb\xf2\x86\xb0\x79\x16\x25\x51\xb5\x76\x3f\xce\x17\xfb\x17\xdd\x84\xa1\xbe\xe0\xa3\x0f\x64\x69\xf0\x31\xb3\x15\x69\xd5\x96\xf7\x2a\x7e\x67\x82\x10\x9b\x32\xbb\xad\xba\xc1\xfe\x80\xcb\xa3\x9c\x35\x86\x8f\x1c\xd6\x1b\x4f\x79\xf8\x01\x8b\xb9\xcd\xa9\xe5\xb6\xe0\x37\x00\x00\xff\xff\x67\x96\x41\xf1\xee\x00\x00\x00")

func stringsBuilderMdBytes() ([]byte, error) {
	return bindataRead(
		_stringsBuilderMd,
		"strings-builder.md",
	)
}

func stringsBuilderMd() (*asset, error) {
	bytes, err := stringsBuilderMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "strings-builder.md", size: 238, mode: os.FileMode(420), modTime: time.Unix(1541844199, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3b, 0x3c, 0xa9, 0x70, 0x7b, 0x7c, 0xf9, 0x4f, 0x22, 0x37, 0xdc, 0x73, 0x1f, 0x54, 0xf2, 0x85, 0xbc, 0x96, 0xdc, 0xa3, 0x5e, 0xdf, 0x19, 0xe9, 0x5d, 0x64, 0x64, 0xee, 0x32, 0xd1, 0x8f, 0xff}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"concat-not-needed.md": concatNotNeededMd,

	"extensive-for-loop.md": extensiveForLoopMd,

	"itoa.md": itoaMd,

	"loop-map.md": loopMapMd,

	"many-loops.md": manyLoopsMd,

	"plus-equal.md": plusEqualMd,

	"remove-extra-bool.md": removeExtraBoolMd,

	"strings-builder.md": stringsBuilderMd,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"concat-not-needed.md":  &bintree{concatNotNeededMd, map[string]*bintree{}},
	"extensive-for-loop.md": &bintree{extensiveForLoopMd, map[string]*bintree{}},
	"itoa.md":               &bintree{itoaMd, map[string]*bintree{}},
	"loop-map.md":           &bintree{loopMapMd, map[string]*bintree{}},
	"many-loops.md":         &bintree{manyLoopsMd, map[string]*bintree{}},
	"plus-equal.md":         &bintree{plusEqualMd, map[string]*bintree{}},
	"remove-extra-bool.md":  &bintree{removeExtraBoolMd, map[string]*bintree{}},
	"strings-builder.md":    &bintree{stringsBuilderMd, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
