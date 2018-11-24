// Code generated by go-bindata. DO NOT EDIT.
// sources:
// concat-not-needed.md (301B)
// extensive-for-loop.md (152B)
// itoa.md (159B)
// loop-map.md (497B)
// many-loops.md (169B)
// plus-equal.md (91B)
// remove-extra-bool.md (171B)
// strings-builder.md (256B)

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

var _concatNotNeededMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\x8e\x31\x72\xe4\x30\x0c\x04\xf3\x7b\xc5\x64\x17\x58\xde\x07\xb8\x6a\x13\x7f\xc2\xa9\xb0\xd4\xc8\xa4\x4d\x12\x34\x00\x69\x4b\xbf\x77\x99\x59\x07\xc0\x74\xbf\xe2\x23\xb3\xe3\xd2\x03\x87\x13\xeb\xcb\x7d\x45\x28\x64\x0c\xf6\x6d\x52\x07\xdb\x88\x0b\xa7\x58\x91\x47\xe5\x32\x8f\x93\x1e\x75\x83\x71\x54\x49\x44\x09\x3c\x4b\x64\x7c\x1d\x1e\x58\xef\x2b\x4a\xf7\xa0\x6c\x37\xbc\x1f\x81\x67\xa6\x71\x7e\x65\x39\x09\xe7\x49\x93\x0a\xa9\x4d\x3d\x50\x36\xf6\x28\x49\x2a\x92\x38\x7d\x81\x6b\x23\x86\xe9\xa7\x49\x6b\x34\xc7\x73\xba\x86\x71\xa7\xfd\x25\x7d\x93\x03\x91\x89\xa4\x1b\x27\xb8\x34\x62\x57\x03\x25\xe5\xb9\xb3\x80\x27\x3b\xca\x8e\xc8\x12\xff\x1d\x5d\x03\x3f\x47\x09\x42\x47\x94\x26\xf5\x0d\x46\xd9\xe4\x51\x6a\x89\x0b\x0f\x4a\x38\x06\x6d\x57\x6b\xd2\x13\x17\xcc\x3a\xdd\xa7\x20\x4a\xe3\xed\xdf\x6f\x00\x00\x00\xff\xff\x3b\x23\x0c\x37\x2d\x01\x00\x00")

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

	info := bindataFileInfo{name: "concat-not-needed.md", size: 301, mode: os.FileMode(420), modTime: time.Unix(1543061558, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x84, 0x34, 0x1d, 0xd7, 0x94, 0x48, 0x3d, 0xcb, 0xc8, 0xd5, 0x26, 0x58, 0x20, 0xf8, 0x89, 0x5b, 0x9e, 0xfd, 0x4, 0x90, 0x36, 0xaa, 0x31, 0x8f, 0xa4, 0xce, 0xb2, 0x4a, 0x74, 0x31, 0xc3, 0xb2}}
	return a, nil
}

var _extensiveForLoopMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcc\x31\x0e\xc2\x30\x0c\x85\xe1\x9d\x53\xbc\x03\x40\x17\x84\x18\x19\x98\xd8\x39\x80\xad\xc6\x90\x48\xa9\x5d\xd9\x29\x28\xb7\xaf\xd2\xf5\x3d\xfd\xdf\x05\x2f\x84\x08\xba\x6d\x60\x17\x6c\x51\xf4\x0b\x06\x7d\xcc\x09\xd5\x6c\x45\xcb\xdc\x50\x02\xc9\xc6\xb5\x98\xcb\x98\x14\x2e\x5c\x6b\x87\xca\x2c\x11\xec\x7d\xc2\x93\xf5\x80\xca\xb2\xba\xfd\x04\xa6\x47\xfc\xc0\xdb\x3b\x92\x70\x1d\x80\x69\xed\xf8\x97\x96\xd1\xb2\x60\xe6\x90\x00\x5d\xe9\x0c\xba\x11\x58\x13\xe8\x4e\xd3\x69\x0f\x00\x00\xff\xff\x42\x70\x9d\x7a\x98\x00\x00\x00")

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

	info := bindataFileInfo{name: "extensive-for-loop.md", size: 152, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x32, 0xd, 0x55, 0xa9, 0x7f, 0xa3, 0x9b, 0xad, 0x3a, 0xc7, 0xad, 0x71, 0x3e, 0x60, 0xed, 0x73, 0x16, 0x66, 0x4e, 0x28, 0xc3, 0x72, 0xb8, 0xcf, 0x62, 0xa0, 0x46, 0xeb, 0x7e, 0xa7, 0x83, 0x47}}
	return a, nil
}

var _itoaMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcc\x31\xae\xc2\x40\x0c\x45\xd1\xfe\xaf\xe2\x75\x1f\x24\xc8\x16\xa8\xa9\x29\x68\x6d\x4d\x3c\x89\x8b\x78\xa4\xf1\x4b\xd0\xec\x1e\x45\xb4\x57\xba\xe7\x0e\x49\xf6\xd2\xe2\x98\x9e\x6c\x2a\xf0\x04\x57\x83\xcf\xde\x36\xa5\x17\x7c\x74\x80\x0d\xec\x1a\x59\x5b\xdf\xa0\x01\xf1\xa0\x9c\xf5\x9c\x3d\x16\xb9\x41\x63\x86\xf3\x3f\x71\x58\x1f\xb0\x5a\xbd\xb8\x05\x27\x5c\xde\xeb\xf8\xa9\x9e\xa8\x9a\xb4\x0e\xae\x1a\xd8\xd3\x63\x81\xd4\x8d\xf2\xc0\xcb\x0c\x5e\x31\xda\x8e\xa2\x81\x65\xb7\xcc\xe9\xfa\xf7\x0d\x00\x00\xff\xff\x8e\x14\x43\x0d\x9f\x00\x00\x00")

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

	info := bindataFileInfo{name: "itoa.md", size: 159, mode: os.FileMode(420), modTime: time.Unix(1543061558, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfa, 0x63, 0x11, 0x5, 0x6f, 0xcb, 0xb3, 0x64, 0xd5, 0x7f, 0xcd, 0x8c, 0x0, 0xf3, 0xc9, 0xb0, 0x2e, 0x6a, 0x8c, 0x7b, 0x40, 0x73, 0x79, 0x2d, 0x99, 0xa2, 0xd8, 0x0, 0x36, 0x63, 0xe, 0x37}}
	return a, nil
}

var _loopMapMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x50\x4b\x6e\xdb\x40\x0c\xdd\xf7\x14\x6f\x99\x00\xb2\xbc\xf7\x05\xda\xac\x83\xae\x8a\x02\xa1\xa5\x37\x1a\x22\xa3\xa1\xc1\xa1\x62\xf8\xf6\xc5\x48\x85\xb3\x24\xf8\xfe\x27\xfc\x6e\x5a\x17\x08\x56\xb9\x41\x1b\x04\x37\x7a\xe2\x14\xe5\x81\x2a\xb1\xb9\x14\xe8\x4c\x81\x56\x44\xd6\x86\x4f\xad\x33\x2c\xa1\x69\x6c\x12\x6a\x75\xc4\x2f\xbb\xf3\x8b\x3e\xe0\x4a\x4c\xe2\x4c\x5b\xb9\xe0\xca\x49\xb6\xc6\xae\xdb\x20\x4e\xe8\x7a\x2b\x5c\x59\x83\x33\xa4\x21\x4b\xcb\x08\xb9\x16\xb6\x01\x91\x09\xf3\x99\xde\x95\x23\x53\x1d\x9f\x7c\xb4\x1e\xa8\x5a\x60\xd9\xc4\xa5\x06\x39\x8f\x78\xab\x48\x32\xc5\x80\x7b\x66\xc5\xc3\x36\x68\xd0\x25\x08\xfb\xa2\xef\x4a\x3b\xd5\xd2\x51\x6a\xd8\x31\x77\x2d\x05\x0b\xa3\xff\xd7\xde\x45\x30\x6b\x4a\x74\xd6\x18\xe0\x52\x67\x5b\xff\x27\xa0\x4c\x19\xa1\x2b\x47\xbc\xdb\xb3\xf6\xb3\xee\xf0\xbd\x55\x8f\xd6\xfd\x5c\x97\x1c\x98\xb2\xe9\x44\xbc\x84\x3f\x20\x68\xa5\x1f\x5a\x5b\x50\xe6\xd7\x11\xef\xe4\x8e\xfd\x78\xdb\xd3\xaa\xd5\xc3\xee\x03\x8d\xd3\x71\xa6\xc3\x49\x3c\x74\x2a\x44\x32\xc7\x6a\x7d\xb8\x9a\xcc\xd7\x9d\x73\xc1\x9f\x9f\x76\x4c\xda\x3b\xec\xc4\xbf\x2f\x39\xe2\xd6\x2e\xe7\xf3\xb5\xd8\x32\x2e\x56\xa4\x2e\xa3\xf9\x72\x5e\xec\xd4\xa1\x27\xad\xa7\x03\xfa\x3a\xfe\xf8\x17\x00\x00\xff\xff\x24\xef\x1a\x06\xf1\x01\x00\x00")

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

	info := bindataFileInfo{name: "loop-map.md", size: 497, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf2, 0x87, 0x94, 0x54, 0x8d, 0xa8, 0xe7, 0xe2, 0x3d, 0xa2, 0x66, 0x67, 0x26, 0xac, 0x48, 0x5c, 0xa2, 0xca, 0xa2, 0x2, 0x2e, 0x81, 0x21, 0xbb, 0x6d, 0xe5, 0x60, 0xe, 0x79, 0xc2, 0x5a, 0xf6}}
	return a, nil
}

var _manyLoopsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcb\xc1\xad\x02\x31\x0c\x84\xe1\xfb\xab\x62\x0a\x78\x6c\x0b\x1c\x68\x80\x03\x0d\x18\x70\x76\x2d\x1c\x7b\x15\x3b\x44\x74\x8f\x02\x12\xb7\x99\x5f\xfa\x0e\x38\x37\x5f\x1b\xd5\x2a\xb6\x42\x02\x5e\x92\x0d\x74\xf5\x9e\x28\x62\xf7\x99\x73\x63\x84\xd4\x5d\x39\x12\xc3\xdb\x63\xc6\x70\xed\x29\x6e\x0b\x2e\x9b\xc4\xef\xa2\x07\x07\x6a\xd7\x94\x5d\x19\xea\xbe\xc7\x82\x13\x19\x5e\xde\x11\xcc\x90\xf2\x99\x37\xb2\x89\x9e\x0c\x49\x0c\xc9\x0d\x85\x07\xb7\xaf\xf8\x87\x37\x98\x1b\x83\x12\xa4\x7a\xfc\x7b\x07\x00\x00\xff\xff\x9e\x90\x92\x12\xa9\x00\x00\x00")

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

	info := bindataFileInfo{name: "many-loops.md", size: 169, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4, 0x10, 0x8a, 0x1d, 0x48, 0xb3, 0xce, 0xb3, 0xc1, 0x8f, 0x9b, 0xed, 0x8e, 0x3, 0x75, 0x69, 0x58, 0x13, 0xac, 0x82, 0x96, 0x62, 0xeb, 0xa2, 0x35, 0x19, 0x5d, 0x9f, 0xc8, 0xf6, 0x95, 0x13}}
	return a, nil
}

var _plusEqualMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xc6\x41\x0a\x82\x21\x10\x05\xe0\x7d\xa7\x78\xfc\xdb\xe8\x02\x41\xb4\x09\xda\x76\x04\x07\x1d\x74\x50\x9e\xd2\x28\xe5\xed\x83\x76\xdf\x05\x0f\x4b\xd8\x7d\xa1\xb2\x7f\xfe\x88\x42\xc8\x18\xca\x04\x9f\x6f\x63\x76\x18\xf1\xec\x58\x6e\xcc\x10\xf8\xe6\x94\x2f\x9a\x55\xc5\x2c\xe6\x57\x04\x1f\xaa\xb1\xe0\x7c\xc3\xf1\x6a\xc2\x7c\x84\xfb\xe9\x17\x00\x00\xff\xff\x47\xe2\xcd\x98\x5b\x00\x00\x00")

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

	info := bindataFileInfo{name: "plus-equal.md", size: 91, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xff, 0x6e, 0xa5, 0xab, 0x70, 0x97, 0x6b, 0xc5, 0xd2, 0x61, 0x97, 0x1a, 0x5d, 0x5d, 0x50, 0x84, 0x27, 0x76, 0x9f, 0x3b, 0x34, 0xc2, 0x74, 0xd0, 0x46, 0x5a, 0x6c, 0xc7, 0xba, 0x97, 0xf6, 0xd8}}
	return a, nil
}

var _removeExtraBoolMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xce\xb1\x0e\xc2\x40\x0c\x03\xd0\x9d\xaf\xf0\xc2\x06\xfd\x1f\xb6\x86\x6b\xca\x45\x84\x4b\x95\xe4\x0a\xfd\x7b\x54\x40\x62\xb5\xac\x67\x9f\x71\xb1\x0e\x72\x46\x0f\x69\x37\x50\x03\xbf\xd2\x09\x2b\xb9\xd0\x55\x19\xe3\x31\x46\xa4\xa1\x54\x2e\x77\x3c\x2b\x67\x65\x87\x39\x9a\x25\xa8\x6d\xb0\x19\x59\x19\x85\x82\x03\xb4\x2c\x2a\x3c\x0d\x1f\xb6\x58\xd7\x09\xd2\x22\x99\xa6\x1f\x20\xdf\x76\xa4\xef\x73\x12\x88\x14\x55\xf0\x63\xc9\xed\x84\xa0\x75\x8f\x37\xeb\x1e\xac\x33\xfe\x3f\x86\xc3\x3b\x00\x00\xff\xff\x5b\x6c\xbc\x27\xab\x00\x00\x00")

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

	info := bindataFileInfo{name: "remove-extra-bool.md", size: 171, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe8, 0x1c, 0x71, 0x70, 0xc8, 0x67, 0xc2, 0xd2, 0x2e, 0x79, 0x8e, 0x5a, 0x35, 0x95, 0xe8, 0x2b, 0x9b, 0xe2, 0x3f, 0xa4, 0x1c, 0xed, 0x23, 0xdc, 0x6e, 0x77, 0x9b, 0xc, 0xf6, 0x14, 0xa6, 0xce}}
	return a, nil
}

var _stringsBuilderMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8f\x31\x52\xc5\x30\x10\x43\x7b\x4e\xa1\x86\xa1\xe0\x93\x1b\x70\x00\x7a\x1a\xba\x38\x8e\x82\x3d\x49\xbc\x99\xd5\x1a\xc8\xed\x99\x1f\x86\x5a\xa3\x27\xbd\x17\xbc\x41\x24\x4e\xeb\x48\x4e\x74\xd5\xf6\x89\x84\xf1\x51\x23\x0a\x9d\x03\xde\x4b\x0a\x7c\x9b\xaf\xba\x61\xea\x81\x1a\x4f\xc2\xe1\x36\xa5\x69\x3b\x61\x5f\xf4\xb5\x6e\x1b\x16\x73\x44\xa9\x02\x7f\xe8\xb9\xea\x6a\x12\x0a\xbf\x88\xc7\xc1\x36\xc3\x0e\x7a\x0a\x73\x8c\xcf\xaf\x23\xaa\xb0\x24\x05\x1d\xa9\xcd\xe8\xa2\xb0\x51\xc2\xce\xdd\xfc\x1c\xf0\x61\x1d\x39\x35\x78\x6f\x88\x42\x4c\x6c\xb9\xec\xc9\x57\x21\x09\x33\x95\xbd\x4e\x9c\x51\xff\xe2\xff\x61\xd4\xa6\xf0\x9e\xa3\x5a\xd3\xed\x62\xe7\xc2\xbc\x22\xee\x26\xf7\x9f\xa7\x75\x17\xb7\x65\x78\xf8\x0d\x00\x00\xff\xff\x28\x78\x0f\x5a\x00\x01\x00\x00")

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

	info := bindataFileInfo{name: "strings-builder.md", size: 256, mode: os.FileMode(420), modTime: time.Unix(1542198423, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd, 0xe9, 0x7c, 0xc6, 0xa0, 0xd4, 0xa4, 0x64, 0x66, 0x63, 0xd8, 0xa5, 0xa1, 0x63, 0x38, 0x6e, 0xc, 0xee, 0x15, 0xc9, 0x73, 0x77, 0xc9, 0xca, 0x1, 0xfd, 0x8a, 0xb3, 0x6, 0xf5, 0xfc, 0x2d}}
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
