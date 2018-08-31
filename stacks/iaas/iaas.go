// Code generated by go-bindata. DO NOT EDIT.
// sources:
// iaas/cloudformation/fullSite.yml
// iaas/cloudformation/s3site.yml

package iaas

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

var _fullsiteYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x61\x6f\xdb\x36\x10\xfd\xae\x5f\x71\x49\x07\xa8\x0d\x2c\xcd\x9d\xdb\x01\x21\x06\x14\x9a\xe3\xa4\xd9\x52\xc7\xb0\x9c\x06\x88\xe1\x0f\xb4\x74\x72\x98\x51\xa4\x47\x52\x76\x92\xa2\xff\x7d\x10\x25\xcb\x94\xbd\x64\xc6\x30\x2c\x9f\x22\xde\xf1\xee\xdd\xbb\xf7\x24\x47\xb7\xf1\x04\xf3\x25\xa7\x06\xcf\xa5\xca\xa9\xf9\x8a\x4a\x33\x29\x08\xf8\x3f\x75\xdf\x77\x83\xee\x69\xd0\x3d\xf5\xbd\x33\xd4\x89\x62\x4b\x63\x23\x7d\x85\xd4\xa0\x06\x2a\x20\xca\xe9\xb3\x14\x30\x96\x85\x41\xf8\xd8\x83\x7b\xa9\x0d\xa6\xf0\x2c\x05\x7a\x23\xaa\x68\x8e\x06\x95\x26\x1e\xc0\x67\x1b\xb9\x93\x02\xcb\x27\x80\xc9\xd3\x12\x09\xc4\x46\x31\xb1\xb0\x07\xad\x16\x93\x7b\x84\xb3\x61\x0c\x82\xe6\x08\x32\xfb\x87\x56\x80\xe1\x22\x84\x07\x5c\x69\x7c\xc0\x55\xc8\xa4\x2d\x18\x71\x2e\xd7\x98\x8e\xa8\x31\xa8\x04\x81\xb7\x9f\x8e\x82\x77\x53\x1a\x3c\x47\xc1\x5d\x37\x38\x0d\xc2\xd9\xb7\xf7\x9d\x9f\x7b\xdf\xdf\x7e\xfa\xe5\x28\x78\x67\xaf\xf4\xa5\xd0\x46\x51\x26\x4c\x0b\x4d\x5e\x68\x03\x73\x04\x0a\x2b\xca\x59\x6a\x91\xd9\xbe\x25\xbc\xb0\x35\xdc\xe0\x91\x69\xa3\x5f\x18\xb1\x46\xf4\x95\xf2\x02\xeb\x1c\x80\x00\x8c\x2a\xb0\x79\xc8\x28\xd7\x58\x13\x92\xd1\x82\x1b\xd2\x3a\x72\x50\x35\xfc\x82\x91\x90\x96\xff\xe4\x4c\x20\xb0\xcc\x81\x03\x02\x31\xd5\x65\xc2\x1c\x21\xb1\x7b\x4b\x3d\x80\x81\x58\x31\x25\x45\x8e\xc2\x1c\xb2\x0d\x27\xbd\x5c\x86\xbb\x64\x80\x3e\x67\x28\x4c\xcc\x0c\x0e\x69\x7e\xd0\x72\x87\xf5\x52\x13\x7b\x13\x3c\x80\xf3\x82\xf3\x33\x99\x53\x26\x0e\xad\x51\x0a\x24\x2b\x38\x87\xd4\x5e\xab\x84\x62\x65\xb0\x5e\xaf\xc3\x95\xd6\x29\xae\x90\xcb\x65\x89\x39\x4c\x64\xfe\xff\x2a\xc2\xeb\x4b\x91\xb2\x32\xd5\xae\xb9\x32\x8c\xe3\x00\x38\x1a\xfc\x59\x50\xae\x61\x7a\x34\xc6\x6c\x4f\x3e\x9d\x6a\xe5\x33\x6f\x8c\x5a\x16\x2a\x29\xc5\xf2\xe6\xb0\x3f\xef\xcd\x9b\xca\x22\x1f\x7b\xae\x0c\x14\x26\x52\xa5\xc0\x04\x44\xb7\xb1\x67\xd3\x0e\x2b\x07\xe5\x68\xee\x42\xa2\xdb\x98\x90\xba\x03\x21\xdb\x16\x1b\xb6\xaa\xb1\xc9\xde\xcc\x36\x3e\x52\x72\x89\xca\xb0\xad\xfa\xb7\x09\x7d\x29\x32\xb6\xd8\x9c\x97\xb5\x72\xab\x4f\x38\xfa\x4d\x32\x01\x53\xdf\xef\xc0\xd4\xff\xec\xb8\x3e\x93\x0a\xfc\x0e\x58\x0a\xfd\x6d\x21\x7f\x36\xab\x8b\x58\x31\xed\xc7\xf7\x5a\x4f\xe8\xc2\xb1\xe3\xef\xf8\x44\x60\x72\x39\x18\x37\x50\xac\x61\xeb\x42\x8e\x17\xda\x37\xa2\xd1\x68\xf7\x42\xf4\x25\xba\xbb\x1e\x5e\x46\x43\xcf\x03\xb8\xc5\xb9\x66\x06\xfb\xe5\xfc\x19\x4b\xa8\xc1\x3d\x5a\x9d\xd8\x17\x2a\xe8\x02\x55\xeb\xec\x05\x0a\x1d\xe3\x40\x83\xa0\x26\xed\x38\x3c\xee\xc0\xf4\xf8\xe4\xb8\xa6\x69\x3b\x73\x43\x12\xc0\x16\xdc\xaf\x45\xf2\x07\x9a\x3d\x5c\x71\x8f\x90\x2a\xf4\x02\x84\x2a\xe8\xd2\xdd\x76\xf4\x86\xf2\x28\x49\x50\xeb\xbe\x14\x46\x49\x4e\x60\x54\xcc\x39\x4b\xc6\x48\xd3\x3a\xbe\xe1\xc8\x4a\xa1\x50\xd4\x4a\xa9\xc1\x79\x29\x52\x7c\x3c\x93\x49\x51\xe9\x82\x95\x8f\xe1\xbd\xc9\x79\x93\x31\x50\x4a\xaa\x6d\xc6\x87\xee\x87\x76\xdc\xdb\x40\xb8\x92\x8b\xff\x64\xd6\xb8\x98\x03\x97\x8b\xf0\x87\x6f\xed\x81\xbf\x7b\xbb\xa4\x8e\x24\x67\xc9\xd3\x2b\xed\xaa\x84\x57\x9b\x6e\xc8\x6d\x15\xde\x70\x5b\xdd\x6f\x86\x6f\x66\x8e\x0d\x35\xd8\x3e\x0a\x20\x66\xa9\x4b\xff\xb9\x54\x17\x68\xaa\x7a\xd7\xf3\x07\x4c\x8c\x6e\x92\x01\x06\x59\x86\x89\x21\xd5\xeb\xd3\x39\x1f\x29\x26\x12\xb6\xa4\x9c\x80\x7f\xe2\x3b\x81\x28\xa9\xde\x01\xba\x47\x2e\x36\x05\x9d\xf0\xe6\xad\xb6\x63\x6d\xaa\x04\xa1\x6b\x4d\x74\x8f\x10\xd2\xf8\xba\x3d\x6b\x07\x7e\x3c\x99\xcd\xbc\xeb\xc2\x2c\x8b\xea\x33\xeb\xd8\xca\xba\x84\x49\x11\xa9\x5a\x34\xae\x6f\xfd\x7d\xff\x55\x90\x07\x8f\x4b\xa9\x1a\x72\xea\xbd\xd6\xee\x09\x4a\xf7\xec\xda\xbe\x63\x0f\xda\x9f\xbd\x8e\x8b\x23\x52\xa2\x76\x57\xfb\x67\xd3\xb6\x33\x50\x25\x76\x05\x72\x33\xbe\x6a\xa3\xbe\x40\x13\x19\xd3\x4e\x0a\xeb\xa7\x9b\xf1\xd5\xbf\x04\x0f\xaf\xa1\xbf\x19\x5f\xfd\x1d\xf2\x5d\x9c\x7f\x05\x00\x00\xff\xff\xcf\x85\xe2\x9c\x35\x0a\x00\x00")

func fullsiteYmlBytes() ([]byte, error) {
	return bindataRead(
		_fullsiteYml,
		"fullSite.yml",
	)
}

func fullsiteYml() (*asset, error) {
	bytes, err := fullsiteYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fullSite.yml", size: 2613, mode: os.FileMode(420), modTime: time.Unix(1535751607, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1a, 0xae, 0x2b, 0x7e, 0x28, 0xac, 0x13, 0xc8, 0xfb, 0x79, 0xe2, 0x54, 0x4e, 0x5a, 0x1d, 0x1f, 0x93, 0x8c, 0x6c, 0x4e, 0xee, 0x8f, 0x8f, 0x7d, 0x6f, 0xf1, 0xe9, 0xa7, 0x54, 0xfc, 0xf0, 0x99}}
	return a, nil
}

var _s3siteYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x6d\x4f\xdb\x48\x10\xfe\x9e\x5f\x31\xa0\x93\x02\xa7\x38\x4d\x1c\xa0\x87\x75\x52\xe5\x38\x09\x2f\x25\x21\x8d\x0d\x14\x10\x3a\x6d\xec\x89\xb3\x57\x7b\xd7\xda\x5d\x13\x68\xd5\xff\x7e\xf2\xda\x89\x6d\x52\x0a\xdf\x8e\x4f\xce\xce\xdb\x33\xb3\xcf\xcc\x2c\xf6\x8d\xeb\x61\x9c\x44\x44\xe1\x88\x8b\x98\xa8\x6b\x14\x92\x72\x66\x41\xd3\xec\x74\x3b\x46\xe7\xd8\xe8\x1c\x37\x1b\x03\x94\xbe\xa0\x89\xd2\x12\x47\x20\x51\x28\x81\x30\x70\x7b\x30\x4f\xfd\x6f\xa8\xc0\xe7\x6c\x41\xc3\x54\x60\x00\x0b\x2e\x60\xc9\xa5\xa2\x2c\x04\x02\x52\x11\x45\x7d\x58\xe1\x5c\x52\x85\x2d\x20\x2c\x00\x02\x33\x9e\x2a\x6c\x00\x1c\xf6\x60\x30\x71\x41\xa0\xcf\x45\x00\x09\xa7\x4c\x9b\x29\x0e\x6a\x89\x85\xef\xc6\x94\x08\x12\xa3\x42\x21\xad\x06\xc0\x29\x97\x0a\x83\x3b\xce\x30\xfb\x05\xe0\x3d\x27\x68\x81\xab\x04\x65\xa1\x3e\xa8\x81\xf5\x96\xa8\x23\x30\x12\x23\xf0\x45\x06\x1a\x9f\x68\x0e\xce\x8e\xc9\x77\xce\x72\x2c\x19\x92\xa5\xf6\x0c\xdf\x39\x43\xc0\x76\xd8\x86\x47\x29\x03\x7c\xc4\x88\x27\x31\x32\xd5\xf6\x79\xac\x03\xd8\x51\xc4\x57\x18\x4c\x89\x52\x28\x98\x05\x7b\x9f\x76\x8c\xfd\x7b\x62\x7c\xb7\x8d\xbb\x8e\x71\x6c\xb4\x1f\x7e\x74\x5b\x47\xbd\x9f\x7b\x9f\xfe\xde\x31\xf6\xb5\x89\xc3\x99\x54\x82\x50\xa6\x6a\xe8\xe2\x54\x2a\x98\x23\x10\x78\x24\x11\x0d\x34\x52\x1d\x3e\x83\xdb\x6e\x00\x8c\xd2\x28\x1a\xf0\x98\x50\x36\x21\xf1\xbb\x13\x5e\xa4\x51\x04\x81\x36\xcb\x13\xd7\xd9\xac\x56\xab\xf6\xff\x9f\x91\xed\xc7\x0e\x0a\x45\x17\xd4\x27\x0a\x6d\xc1\xde\x93\x54\x46\x86\xf5\x6d\xa1\xe4\xa9\xf0\x11\xb2\x82\xc0\x9e\x3d\x9b\xec\x17\xf7\x6a\xdf\xb8\x50\x71\x0d\x63\xc2\x48\x88\x02\xf6\x6c\x67\xbc\x0f\x7e\x29\x69\xff\x32\xe7\x5d\x22\x98\x45\x56\xd2\x22\x7e\x6c\xb5\xff\xdc\x6d\x00\xdc\xe4\xa4\xed\x6b\x1a\x5e\xcd\x2e\xde\x03\xb5\xb0\x59\xf7\x45\x2a\xa2\x06\x80\x43\xfc\x25\x5e\x93\x28\x45\xcf\xab\x79\x99\xa4\xf1\x1c\x45\xe1\x65\x41\xd2\x48\x59\xd0\xeb\x74\xaa\x00\xb5\x99\xb4\x40\x9f\x01\x18\x70\xd4\xd9\x7c\x76\xcd\xf2\x7b\x6d\x96\xab\x94\xdf\xc7\xc5\x77\x9d\x25\xde\x05\x70\xfd\x2d\xa1\xdb\x02\xb3\x05\x87\x2d\xe8\x76\x5a\xd0\x3d\x84\x98\x32\x99\x47\x6b\x8c\x49\x92\x50\x16\xea\xbe\x9b\x61\x48\x39\x1b\x93\x24\xc7\x9f\x4a\x03\x89\x54\x46\xd7\x2a\x22\xb9\xbd\xb2\x33\xcf\x02\x0b\xee\x7a\xf6\x97\xbe\xeb\x9d\x8c\x6e\xcf\x5d\x6f\xb4\xb6\x59\xe1\x6f\x6d\xcc\xd1\xe1\xd1\xd5\xdd\x85\x39\xee\xda\xce\xa0\x66\x63\xbe\x1e\xa7\x7f\x7e\xf4\xf9\x68\x76\x76\x76\x39\xf9\x38\xd6\x4a\x98\xbe\x15\xa7\xdb\xff\xec\x78\x5f\x07\x1f\x0f\x86\x77\xd3\xa1\x56\x22\x89\x21\x79\xaa\x96\x6f\x25\x75\xd9\x39\x37\x07\x5f\xfb\xc3\xee\xc8\xeb\x6f\x1b\xbe\x8e\xb2\x7b\xe3\x9c\x9d\xdc\x9e\x39\x13\xb3\x7f\x3b\x58\x1b\x32\x2e\xde\x8c\x68\x8e\x0f\x86\xa7\x57\x33\xf3\x68\xfa\xf1\xee\x46\x2b\x49\xf2\x66\xe9\xbb\x27\x23\xaf\x73\x65\x77\xcf\xcc\xd3\xeb\xc6\xba\x65\xa4\xd5\x28\x49\xed\x44\x3c\x0d\x16\x82\x33\x55\xe5\xa3\x7d\xe3\x5a\x96\x16\x8d\xb4\xc8\x1a\x50\xa9\x04\x9d\xa7\x19\x55\xb4\xde\x54\xf0\x24\x6b\x25\x94\xeb\xe8\x55\x15\x47\xaf\x81\xb5\x24\x9b\x12\x71\x36\x68\x2c\x28\xc3\xd5\xf4\x6b\x23\x7f\xb3\x4d\x4a\xf3\x54\x2a\x1e\x0f\x85\xe0\x62\x86\x32\xe1\x4c\x96\x61\x33\x6e\x6b\x49\xd6\x5c\x94\x85\x63\xca\xb2\xe6\xaa\xf4\x41\xf6\x97\x6b\xf0\x00\x2d\x38\xe8\x1c\x54\x04\x6b\x7f\xb9\xcc\xac\x19\xad\x65\x53\x12\xe2\x94\xa8\xa5\x05\xbb\x1f\x28\x0b\xf0\xa9\xbd\x54\x71\xb4\xbb\xd1\xbc\x14\x34\xa4\xac\x86\xa8\x32\xac\x61\xc7\xc5\x08\x7d\x05\xf7\x66\x0b\x76\xdc\x24\xa2\x0a\xee\x77\x3f\xec\xb6\x60\x67\x86\x8b\xad\xe1\xf2\xf0\x50\x41\x90\xdd\xa2\xdb\xcb\xfd\x57\x8e\xf3\x7a\xe4\xc7\x2f\x4b\x9d\xfd\x9d\x7a\xde\x74\xca\x85\xb2\xa0\xf9\x57\xa7\xb9\x25\x72\x0b\xd9\xc1\x41\xaf\x2e\xcc\x3d\x4e\x05\x57\xdc\xe7\xd1\x94\x47\xd4\x7f\xb6\x60\xa9\x54\x62\x70\x16\x3d\x6f\x74\x87\x8c\xcc\x23\x0c\x2c\x50\x22\xc5\xcd\xe9\xa9\x52\x49\xf9\x6e\xc8\xac\xcc\xd2\x7d\x31\xd6\x66\x9c\xab\xcb\xf9\xbf\xe8\x2b\x0b\xca\x52\x6e\xb4\xec\x88\x92\x17\x77\xab\x8b\xd4\xac\x2f\xc0\xd2\xed\x05\x0f\x43\xca\x6a\xe9\xe7\x95\xcc\xca\x9e\xce\x21\xe2\x61\xfb\x8f\x1f\x75\xeb\x9f\x6d\xd9\x6b\x13\xbd\x46\xc8\x4a\x66\xcb\xaf\x5d\x31\x9f\x0a\x5c\xd0\x27\x0b\x88\xef\xa3\x94\x55\x0a\xfd\x32\x69\x80\x33\xe6\x47\x69\x80\x0e\xe7\xdf\xb2\x6e\x80\x05\x89\x24\xbe\xcc\x5b\x4f\xfe\x3e\x2e\xc9\x23\xe5\xa2\x8a\xb6\x98\xee\x63\x54\x4b\x1e\xc8\xaa\xc4\x80\x93\xa1\x57\xfb\x7d\x3a\xb4\x07\x95\x83\xc2\xb5\x66\xbb\xae\x52\x6d\xbd\x54\x14\xc7\xe4\xe9\x1d\x4a\x45\xdf\xfc\x56\xc9\xe1\x71\x22\x50\xca\xad\x22\x78\x44\x84\xa8\x72\xfe\xbc\x42\xda\x11\x17\x2b\x22\x82\xcd\x26\xab\x31\xef\x4b\x8a\xe2\x39\xdf\xa5\x5b\xbe\xb3\xb0\x79\x6d\x6b\x87\x1b\x8f\x16\x30\xce\xaa\x06\xd7\x14\x57\x28\x5e\xd2\x58\x60\x40\x05\xfa\xca\x50\xdc\xc8\xc8\x59\xde\xed\x54\x50\x1f\x9d\x88\x64\x79\x95\xdf\xff\xd8\x51\x49\xcc\xdc\x65\xe5\x59\x51\xbb\xc3\xad\xb7\x4c\x5e\xc5\xad\xf3\x8a\x8d\x2b\x23\x37\x4d\x12\x2e\x54\x7e\xf5\x16\x48\x46\xf3\x2e\x2b\x87\xf3\x60\xe2\xbe\x7c\xef\xe9\xc9\xac\x9f\xaa\x87\x3d\xcb\x9a\xe9\x07\xb3\x8b\xea\x44\xf0\x34\x79\x65\x30\x97\x4b\xa1\x98\x47\xe7\x9c\x32\xb8\x6f\x36\x5b\x70\x9f\x77\x57\xa9\xd1\x6c\x41\x7b\x33\x81\x36\xde\x37\xae\x0c\x28\x5c\xfc\xb6\x29\x0b\xa4\xf5\xae\xce\x19\x52\x2d\xdb\xd6\xce\x1f\x78\x93\x81\xed\xd9\x5f\x6e\x6f\xcc\x2a\xcd\x8b\x22\xc0\xce\x09\x2a\x5b\x29\xb8\xdf\x5a\x5d\xad\xca\xbc\x7d\x68\x5c\xa6\x2a\x49\x73\xcc\xa5\xca\x90\x05\x7a\xc7\xe4\x00\x34\x05\xdf\xef\x11\x5e\xbe\x9b\xd6\xde\xf4\x3f\x39\x95\x9d\x16\xd4\xd7\x64\x59\xa0\x7a\xd8\x57\xab\x57\x0b\x92\x29\x54\x60\x34\x1a\xff\x05\x00\x00\xff\xff\xd1\x3c\x79\xaf\xa6\x0d\x00\x00")

func s3siteYmlBytes() ([]byte, error) {
	return bindataRead(
		_s3siteYml,
		"s3site.yml",
	)
}

func s3siteYml() (*asset, error) {
	bytes, err := s3siteYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "s3site.yml", size: 3494, mode: os.FileMode(420), modTime: time.Unix(1535751656, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdf, 0x54, 0x3b, 0x67, 0xc4, 0xbe, 0x71, 0xbf, 0xa7, 0xfd, 0x26, 0xcf, 0xa9, 0xbd, 0xd2, 0x15, 0x4e, 0xa4, 0xf2, 0x4c, 0x27, 0x86, 0x63, 0x94, 0x86, 0xeb, 0x26, 0x49, 0x30, 0x88, 0x9b, 0x16}}
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
	"fullSite.yml": fullsiteYml,

	"s3site.yml": s3siteYml,
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
	"fullSite.yml": &bintree{fullsiteYml, map[string]*bintree{}},
	"s3site.yml":   &bintree{s3siteYml, map[string]*bintree{}},
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
