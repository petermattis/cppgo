// Code generated by go-bindata.
// sources:
// cppgo.go
// DO NOT EDIT!

package main

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

var _cppgoGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\xff\x6f\xdb\xb6\xb6\xff\xd9\xfa\x2b\x4e\x5d\x24\x93\x12\xc5\x6e\xbb\xa1\x6f\xb0\x63\x3f\x64\x69\xfa\x6a\xb4\x75\x82\xda\x79\xbb\x43\x16\x78\x8c\x44\xd9\x5c\x65\x4a\x57\xa4\xf2\x65\x85\xff\xf7\x8b\x43\x52\x12\x25\xcb\x6d\xb3\x75\x17\x17\x77\x18\x10\x8b\x3c\x5f\x3e\x3c\xe7\xc3\x73\x48\xa9\xfd\x3e\x9c\x26\xe9\x43\xc6\x96\x2b\x09\x2f\x9e\x3d\xff\x1f\xb8\xa0\x92\x66\xf0\x9e\x48\xc9\x44\xcf\xe9\xf7\x9d\x7e\x1f\xde\xb1\x80\x72\x41\x43\xc8\x79\x48\x33\x90\x2b\x0a\x27\x29\x09\x56\xb4\x98\xf1\xe1\xff\x69\x26\x58\xc2\xe1\x45\xef\x19\xb8\x28\xd0\x35\x53\x5d\x6f\x88\x26\x1e\x92\x1c\xd6\xe4\x01\x78\x22\x21\x17\x14\xe4\x8a\x09\x88\x58\x4c\x81\xde\x07\x34\x95\xc0\x38\x04\xc9\x3a\x8d\x19\xe1\x01\x85\x3b\x26\x57\xca\x8f\xb1\x82\x48\xe0\x17\x63\x23\xb9\x91\x84\x71\x20\x10\x24\xe9\x03\x24\x91\x2d\x08\x44\x1a\xd0\xf8\xdf\x4a\xca\x74\xd0\xef\xdf\xdd\xdd\xf5\x88\x02\xdc\x4b\xb2\x65\x3f\xd6\xa2\xa2\xff\x6e\x72\x7a\x36\x9d\x9d\x1d\xbd\xe8\x3d\x33\x4a\x97\x3c\xa6\x42\x40\x46\xff\x99\xb3\x8c\x86\x70\xf3\x00\x24\x4d\x63\x16\x90\x9b\x98\x42\x4c\xee\x20\xc9\x80\x2c\x33\x4a\x43\x90\x09\x82\xbe\xcb\x98\x64\x7c\xe9\x83\x48\x22\x79\x47\x32\x8a\x66\x42\x26\x64\xc6\x6e\x72\x59\x8b\x59\x01\x91\x89\x9a\x40\xc2\x81\x70\xe8\x9e\xcc\x60\x32\xeb\xc2\x4f\x27\xb3\xc9\xcc\x47\x23\x3f\x4f\xe6\x6f\xce\x2f\xe7\xf0\xf3\xc9\x87\x0f\x27\xd3\xf9\xe4\x6c\x06\xe7\x1f\xe0\xf4\x7c\xfa\x6a\x32\x9f\x9c\x4f\x67\x70\xfe\x1a\x4e\xa6\xbf\xc0\xdb\xc9\xf4\x95\x0f\x94\xc9\x15\xcd\x80\xde\xa7\x19\xae\x20\xc9\xd0\x04\xc3\x80\xd2\xb0\x07\x33\x4a\x6b\x10\xa2\x44\x43\x12\x29\x0d\x58\xc4\x02\x88\x09\x5f\xe6\x64\x49\x61\x99\xdc\xd2\x8c\x33\xbe\x44\xfd\x94\x66\x6b\x26\x30\xb1\x02\x08\x0f\x21\x66\x6b\x26\x89\x54\xcf\x5b\xeb\xaa\xbc\x9c\x5c\xce\xdf\x9c\x7f\x98\xa9\xfc\xa2\x19\xf4\xc6\xc9\x9a\x0a\x4c\x56\x90\x70\xbd\xf4\x24\x13\x3d\x07\xa7\xa7\x89\x44\x3d\x22\x2d\x56\x30\x01\x74\x7d\x43\xc3\x90\x86\x18\x65\xb4\x1a\xa4\xe9\x32\x81\x1b\xc6\x49\xf6\x00\xb9\x60\x7c\x09\xcb\xe4\xe8\x86\xf1\x90\x48\xd2\x83\x4b\xa1\x7c\x75\x97\x09\x2c\x29\xa7\x19\x91\xb4\x0b\x77\x2b\xca\xe9\x2d\xcd\x80\x49\x34\xb9\x4e\x42\x16\x31\x9d\xbb\x8c\xde\xe4\x2c\x0e\x95\xe9\x42\x21\x84\xc2\xdc\x32\x51\xc0\x59\x4c\x35\xc6\x43\x2d\xcc\x96\x3c\xc9\xa8\xe3\xa4\x24\xf8\x88\xd1\x5a\x13\xc6\x1d\x87\xad\xd3\x24\x93\xe0\x3a\x9d\x6e\xb4\x96\x5d\xa7\xd3\x8d\x93\x25\xfe\x49\x89\x5c\xe1\xdf\x8c\x46\x31\x0d\xd4\x8c\x48\x32\xf5\x37\xe7\x2c\x48\x42\xda\x75\x3a\xfd\x3e\x4c\xde\x5f\x9c\x7f\x98\xcf\x1c\xcf\x71\x6e\x49\x06\xf2\x21\xa5\x02\x46\xb0\x26\x1f\xa9\xbb\x26\xe9\x95\xd1\xef\xcd\x1f\x52\x7a\x7d\x93\x24\xb1\xa7\xe4\x42\x1a\xc4\xf3\xaf\x94\x55\xf1\x4f\x49\x50\x17\x46\x16\xf2\xe5\xf5\x41\x39\xeb\x39\x0e\x7a\xaf\xc4\x41\xc8\x2c\x0f\x24\x7c\x72\x3a\x38\x06\xa0\x55\x9c\x4e\x48\x53\x01\x80\x46\x2a\x6d\xe5\xcf\xe9\x68\xfc\x57\xd7\x36\x14\xa7\x23\x28\xe5\x5a\x61\x0b\xa2\xb3\x71\x9c\x28\xe7\x81\x5a\xf9\x94\xac\xa9\x2b\xc1\x16\xf2\x8c\x57\x44\x21\xee\x98\x0c\x56\x20\x7b\x6f\x19\x0f\x5d\x0f\x87\x02\x22\x68\x29\xff\x53\x92\xc4\x03\xa7\xd3\xc9\xa8\xcc\x33\x0e\x5d\xb4\xdf\x6d\xc8\x4c\xb8\xfc\xd1\x96\x61\x5c\xfe\xb8\x90\x2d\x52\xcf\x5f\x36\xc4\x9e\xbf\x6c\x95\xfb\xfe\x45\x43\xee\xfb\x17\xad\x72\xbe\xfd\xf0\xf2\x87\x86\xd2\xcb\x1f\xb6\x95\x2e\x59\x03\x6b\xde\x0e\xf6\x92\x35\xd1\xe6\x3b\xe0\x5e\xb2\x26\xde\x7c\x07\x60\x94\xf4\x6b\x4f\x75\xc8\xf9\x67\x30\xa7\x32\x6b\x8a\xa6\x32\xdb\x96\x7d\x1d\x27\xa4\x01\x27\xc2\xa1\x56\xb9\xba\xfb\x30\xc9\x6f\x62\xda\x14\x9c\x29\xa6\xd8\x72\x7a\xa4\x29\x77\x92\x65\xe4\x01\xc5\x68\x4c\xd7\xc8\x39\x18\x8c\x2c\xfe\xf5\xce\x62\xba\x76\x3d\xcf\xe9\x74\x58\x04\xa5\xcc\x68\x04\xdd\x2e\x72\xae\x32\xdf\x75\x3a\x9d\x4d\xe5\x2e\x5a\xcb\xde\x2c\xcd\x18\x97\x91\xdb\x55\x4e\x8e\xf7\x84\xbf\x17\x8e\xbb\x7e\x69\xc6\x07\xd9\x7b\x47\xb9\x32\x5f\x07\x8f\xbd\xe9\xef\x06\xa5\x9c\x1c\xef\x09\x1b\x51\x13\xc8\x85\x4e\xdf\xdf\x09\xe3\x22\x61\x5c\xd2\xec\x0b\x40\x66\xaa\xfc\x20\x96\xf4\xe3\xf2\x82\xc8\x95\x82\xd2\xbb\xd0\x0f\xae\xc1\x50\xcc\x6d\x43\x90\x3d\x85\xd9\xab\x03\xc1\xc2\xdc\xfb\x89\x08\xea\x1a\x4d\x0f\x0e\xa1\x3b\x18\x74\xe1\xd0\xd2\x08\x69\x44\xf2\x58\x39\xef\xf7\x61\x7e\xfe\xea\xdc\x4d\xf1\x78\xe4\x0d\x20\xe7\x22\x4f\xb1\xea\xd3\x50\xcf\x42\x89\xf8\x14\x8f\x31\xf4\xfe\xe5\x0f\x3b\x26\x9e\xbf\xf8\x71\x6b\x66\x45\x78\x73\xec\x75\xce\x83\xe6\xd8\x04\x03\x16\x91\x80\x36\x27\xde\x93\xb4\x39\x74\xc9\x05\x89\xa8\x09\xb2\x53\x4b\xc9\x06\xeb\x2d\x36\x85\x20\x4d\xdf\xd2\x87\xbb\x24\x0b\x75\x57\x28\x1b\x02\x16\xcd\x4f\x4e\xa7\x1b\xd2\x98\x4a\xda\x1d\x80\xcc\x72\xea\x3b\x9d\x2e\xbd\xc7\x45\x5b\x03\x8c\xc7\x8c\x57\x12\x45\x21\x17\x84\x33\xc9\xfe\xd0\x84\x11\xa6\x7c\xdb\x65\x9c\x45\xb6\xf7\x2b\x71\xad\xb2\x66\x40\x0a\xcc\xc6\x42\x21\x2d\x87\x4a\xd3\x24\x08\xa8\x10\x49\xb6\xd3\x74\x92\x4b\x24\x89\x6a\x72\x57\xd7\x59\xce\xa9\x0f\xcf\x7c\x88\x29\x77\x05\x92\x36\xcd\xe8\xed\x65\x9a\xd2\x0c\xa5\x22\x12\x0b\xea\x74\xf0\x88\xc2\x7c\x50\x43\x19\xe1\x4b\x0a\x42\x01\x62\x11\x98\x4e\xdd\x9b\x08\xa5\xe4\x66\x9e\x26\x18\x8b\x80\xc1\x18\x9e\xc1\xfe\x3e\x3c\xa9\x6c\xaa\x39\x85\x61\x84\x67\x47\xca\x43\x37\xc9\xa5\x0f\xdf\x2d\xbe\x43\x12\x2a\x16\x76\x32\x18\x95\x76\xe7\xc9\xbb\xe4\x4e\xd9\xc5\x99\xca\xd0\x48\x85\x14\x79\x0b\x34\x16\x54\x1b\xb6\xa7\x0d\x74\x65\x71\xdb\x1f\x9a\xb3\xc2\x57\xcb\x87\x0a\x15\x8a\x79\x5e\x19\xd6\x3b\x12\x7f\xdc\x6a\xbb\x55\x6b\xff\x9a\xd6\x7b\x21\x33\xbf\x5e\x5f\xfd\xed\xca\x66\x00\x69\x77\x55\x25\xd9\x28\x4a\x94\xa6\x9f\x8c\x1a\x15\xc0\xa6\x07\x67\x31\x2a\x98\x13\xc9\x60\x64\x6d\x67\xab\x32\x78\x4e\x87\xe3\x64\x75\xf2\xb9\xc2\x9f\xd7\xca\x11\xc7\x52\xc1\x59\xac\xcc\x72\x18\xc1\x7e\x29\xa6\xe2\x8c\x4f\x03\xa5\xea\xe3\x23\x1e\x76\x06\xd5\xb1\xa9\x71\xe2\xf1\x94\x0c\x9e\x6f\x06\x9f\x3b\x87\xf9\x26\x57\x4d\x40\x30\x02\x5e\x04\x80\xf7\xd0\xcc\x95\xac\x6d\x07\x3d\x5b\x4d\x15\xcc\xe0\xbd\xe2\xa4\x68\x12\x6f\x06\x7c\x90\x9e\x63\x38\x8d\x11\x78\x36\x04\x06\xc7\x58\xd8\xf2\xf5\x6b\x46\xe3\xd0\xf5\x86\xc0\x0e\x0f\x95\x8f\x48\x17\x54\x3d\xce\x4c\x3d\x8d\x8a\x30\x62\x22\x8a\x8a\xaa\xee\x47\xba\x00\xd0\x10\x22\x54\xe8\xe1\x38\x9e\xe8\x19\xcf\x0b\x2a\x1a\x96\x44\x6a\xe5\x16\x55\xda\xea\xba\xff\x38\xee\x20\xb4\x10\xf1\x2a\xf6\x68\x0f\xde\x10\x42\x44\x89\xc9\xdc\xdf\x37\xbf\xcd\x2e\xe4\x3d\x4c\xdc\x55\x58\x85\x4c\x43\xdc\x68\xfa\x94\xe1\x2d\xf6\xc0\x12\x7b\x32\xea\x62\xa3\xba\xd0\x7d\x2a\xcd\x92\x38\x59\xe6\x78\x32\xee\x08\xbd\xf4\xaa\xb8\x54\x44\x28\x4b\x4c\x95\x5d\x64\x20\xe6\x60\xe1\x03\xaf\xea\x8a\x75\x10\x47\x90\xc6\x64\x99\x42\xfd\xec\x03\xd7\x7b\x02\x1f\xf5\xfa\xcb\x19\x44\xea\x32\x1f\x7e\x07\xc6\xa5\x07\x48\xad\x5a\xed\x54\x62\x57\xec\xba\xa7\x76\xc7\x71\x31\xf0\xbb\x1e\x70\x3a\x1b\xcf\xe9\xdc\x32\xc1\xec\x95\xb4\x91\xba\x0d\xbc\x41\xab\x36\x4d\x0f\xa3\x65\x0c\x69\xb0\x56\xd4\x68\xca\x4c\xd4\x8a\xd8\xba\xdc\xaa\x26\x1e\x58\xca\x6d\x97\x08\xcf\xf4\x08\x23\x72\xc5\xed\x0d\xa1\x9c\x59\x33\x26\xb9\xe6\x4e\xf2\xf9\xec\x28\x46\x14\x99\x09\xad\xb4\xa8\x09\xe5\x44\xfd\x28\x13\x82\x4f\x3e\x84\x5b\xe9\xd0\xe3\x5f\x48\x86\xa2\x5f\x95\x0a\xf5\x58\x4b\x84\x09\xb1\x05\xa4\x82\xb1\x15\xe0\x5a\x84\xbb\x55\x69\xde\x13\xf0\xe9\x57\xde\xf5\x81\x2b\xcb\x95\x59\x69\xaf\x4f\xd7\x8a\x32\x75\xb8\x77\x5c\xb9\x95\xb8\xee\xaf\x7c\x03\xfd\x3e\xd8\xd6\x7f\xe5\x35\xeb\xbb\x33\xaa\x6d\x36\x7a\xc8\xb7\x6a\x1c\x16\x6a\xeb\x00\x6a\x31\xe2\x6b\xfb\x47\x29\x8c\x01\x29\x4a\x2d\x8b\xe0\x49\x79\x9f\x2e\x06\x3b\xb5\x91\xaa\x86\xd4\xe3\xd5\xf5\x4c\x69\xaa\x62\xb6\xb3\x27\x0d\x4d\x14\x55\x95\xb2\x7b\x6b\xa7\xd3\x89\x93\x65\xef\x35\x91\x24\x8e\xdc\x6e\x8e\x85\x96\x06\xb8\x3b\x4a\xb1\x01\x26\xfa\xc9\x08\xf6\x44\x99\x0c\x1f\xac\x9b\x7b\x71\xb8\xb0\xd1\x05\x31\x11\x02\xf6\xc4\x50\x65\xb0\x38\xda\xd6\x4f\xc3\x2a\x1e\xb2\xb1\x4a\xe7\x2f\x2f\xe8\x2f\x2f\x67\xf3\xdf\xd9\xc3\x2a\x1e\x9b\x06\x56\x35\xa4\x36\xba\x35\x77\x27\xb2\xad\x2d\xc3\xa6\x04\x58\x29\xb6\xa5\x20\xcd\x6f\x62\x16\x0c\xb6\xf5\x01\x28\xcf\xd7\xf0\x09\x66\xec\x0f\x0a\x23\xd8\x0b\x61\x53\xb0\x05\x87\xb6\x4c\xb5\x98\xb0\x8d\x3b\x1d\x75\xb3\x48\xb8\x7e\x71\x94\x64\xa0\x5f\x0a\xfd\xbb\x32\x59\x5f\x1b\x4a\x4d\xcf\xfe\x71\x71\xfe\x61\x7e\xf6\x4a\x31\x0e\xff\x0f\x71\x7d\x6a\xcf\xd5\x2e\x13\x91\x8a\x9d\xe7\x9b\xbc\xe3\xdf\xf3\x28\x12\x54\x7a\x2d\x14\x91\x5b\xf7\xe2\x2a\x9d\x58\x5c\x9a\x57\xe2\x16\x5c\x93\xf7\x17\xef\xce\xde\x9f\x4d\xbf\x35\x34\xac\x65\x76\x02\x3e\x19\x99\x72\xa0\xbd\x92\x01\xec\x09\xd7\xdb\x2a\x14\x0d\x19\x80\x01\x84\x44\x92\x85\xcb\xf3\x38\x4e\x25\x5e\x87\x8a\x1a\x58\x97\xdc\xb4\x0f\xef\x09\x57\x61\xc1\xe5\xee\xdf\xd7\xfd\x7d\xa5\xe7\xfb\x9e\xfa\xfb\xa7\x3d\xdf\x26\x2c\x84\x03\xb4\xf1\x88\xe5\x66\x54\x5d\xa5\xd3\x8c\xca\x45\x40\x84\x3c\xd6\xb6\xcc\xfb\xb8\x83\xb1\xab\xec\x3d\x0e\x93\x19\xc3\xa4\x35\xa1\xaa\x64\xe8\x7d\x54\x6c\x6d\xfd\xea\xa8\x8d\x17\xde\x96\x01\x00\x73\x04\x41\xba\x1c\xef\x89\xf1\x60\xb0\xa4\x52\x61\x5c\xc0\x21\xec\x85\xde\xd0\x36\x6a\xd3\xa9\x15\xb7\x6e\x99\x5b\xb4\xfa\x72\x80\xb7\xeb\xd2\x0e\x07\xd5\x2d\xb7\x1e\x21\x48\x33\x76\x4b\x24\x35\xb5\xab\xa1\x5d\x4b\x82\x76\xb9\x18\x96\x90\x6d\xd9\x8d\x1e\x36\xaf\x3e\x8a\x63\x3d\x8c\xe0\x37\xf5\x15\x2c\xac\x7d\x06\x78\xd0\x9f\x1b\x9c\x7e\x1f\x5e\x9d\xc3\xf4\x7c\x0e\x67\xaf\x26\xf3\x27\x8e\xf3\x94\x45\x3c\xa4\x11\x9c\x5e\x5c\xfc\xdf\xf9\xe2\x8d\xf3\x34\xa4\x11\xe3\xb4\x7c\x76\x9e\x32\x1e\xc4\x79\x48\xe1\x58\x5f\xb3\xc7\xd6\x08\x96\x8a\x85\xcc\x08\x93\x62\x5c\x13\x0c\x19\x97\xbd\xd5\xd8\x71\xaa\x16\xba\x4c\xe0\x93\xe3\x48\xba\x4e\x63\x22\x8d\xae\x6a\xb4\x73\x1f\xca\xdf\x67\x30\x52\xb1\x1e\x3b\xe6\x35\x3d\x66\x7b\xd8\xae\x56\x93\x39\xb6\xad\x08\x19\x0e\x06\x94\x93\x9b\x98\x2e\x58\x74\xac\x1e\x99\x58\xa8\xce\x72\x3c\x1f\x0f\x06\xb7\x24\xce\xe9\x78\x30\x40\x8d\x31\x7c\x72\x9a\x3d\x63\x3e\x18\xa8\x5f\x9b\xa1\x03\x20\x24\x91\x2c\x80\x39\x20\xdf\x5a\xe8\xe0\xd8\xe4\xd4\x9b\x06\xd5\x36\xce\xe6\x5b\x01\x27\x19\x93\xab\x35\x95\x2c\xf8\x2a\xf4\x82\xfd\x41\x93\xc8\x9d\x7b\x8f\xc7\x7f\xb0\xa3\x28\xcc\x8b\x72\x50\xad\x4c\xf7\xe9\x59\xf1\x96\xaa\xf8\xae\xf2\x86\x92\x50\xbd\x35\x82\x82\xcb\xc1\x8a\x64\xda\x19\x2a\x83\x79\xc7\xae\x50\x0e\x1d\x84\xe8\x14\xfd\x16\xa7\x5b\x97\xa2\x8d\x7a\x4d\x61\xed\x7c\x7b\x4d\x8e\x2e\x73\xab\x30\xdb\x59\xe5\xb4\x45\xab\xc8\xa9\x75\xa9\x68\x85\x83\x81\x79\xf7\x46\xc4\xc2\xbc\x5a\x2a\x0b\x57\x11\x28\x4b\x4c\x99\xc0\x52\x2f\xf4\xe9\x62\x08\x1b\xc7\x5e\xfb\x01\x68\x81\xa6\x09\x84\x77\x34\x56\x71\x51\x1a\x76\x60\x76\x49\xab\xa0\x29\xa0\x45\x11\x29\x5d\x99\x15\x29\xc1\xe1\x67\xb8\xa7\xf3\x66\xde\xa3\xe2\xba\xab\x78\xb6\xc6\x1e\xe3\x7a\xb0\x15\x7a\xa3\xbf\x33\xf6\xa9\xcc\x16\xee\x2e\x36\x15\x2d\x46\x23\x2f\x93\xe0\x2b\x5d\x50\xaf\xce\x17\x2e\x5a\x50\x87\x70\xdd\x95\xe1\x7f\xcd\x76\x31\x85\x1f\xa7\x3d\x18\xc0\xdc\xb5\xb3\x57\xc1\x39\x50\x74\x2f\xc3\x68\x93\x1c\x55\x35\x8f\x01\x92\x14\x2b\xa4\x39\xd4\xb5\x8b\x6b\x3b\x15\x92\x42\xd5\xec\x8c\xd2\xc4\x41\xbb\xba\x5a\x4d\x43\xe7\xa0\x54\x3a\x1a\xb7\x6b\xed\x5b\x6a\x2d\xc9\xd6\x01\x2f\x16\x52\x20\x31\x3a\x3b\x72\x7f\x5a\xe4\x7e\x22\xcd\x9a\x6b\xc9\x47\x39\x6c\x04\x95\xbc\xa9\x34\x0b\xf5\x01\xb5\xfa\x39\xdc\x21\x9c\xd1\x88\x66\x94\x07\x46\xa1\xfe\x68\x2b\x15\xfe\x8f\x4f\xc7\x20\x68\x1c\x19\x01\x1b\x4c\x21\x62\xf8\x75\x0a\x07\x81\x8f\x1b\x04\x18\x0f\xe9\x7d\x41\x32\x3c\x26\x12\xc6\x69\xb6\x70\x83\x92\x3d\x4a\x62\xe1\x6a\x41\xc3\x8c\x2d\x83\xa5\xdb\x7d\x68\xb3\x76\xdf\xab\x1e\x9a\x86\xef\x7b\xfa\x87\xc5\x3a\x2b\x4c\x5f\x20\x83\x7b\x60\x19\xbe\xd2\x86\xae\x87\x65\xed\x29\x50\x95\x66\x0e\x0f\x5d\xf5\x1e\x46\x1b\xa9\xe6\x25\x8c\xe0\x40\xae\x98\x18\xda\xd6\x0f\x0f\x65\x41\x34\x6b\x81\x96\xad\xc2\xd0\xe1\xa1\x76\x5d\xd3\x2e\xed\x6d\x5a\xa1\x1c\x1d\x3d\x06\xca\xd1\xd1\xc1\xe7\xb0\x1c\x1d\x95\x58\x8e\x8e\x3e\x8f\xc5\x01\xfd\x16\xaa\x50\x1d\x8d\xb6\x73\x98\xc8\x15\xb6\x87\x96\x80\x57\xe1\xc6\xbb\x8b\x92\xb3\x72\x0b\xfb\xfb\x26\xad\xd5\x6c\x85\x66\xd3\x74\xfd\xe4\x51\xae\x9f\xb8\x6a\x15\xa5\x65\x6f\xe7\x6e\x46\x7a\x97\x98\x86\xba\x15\x40\x81\xe3\x4b\x95\x5c\x5d\xc4\x95\xe7\x96\x0e\xdc\x76\x9c\xd4\x81\x6e\xb4\x61\x6b\x24\x20\xa9\x42\x5a\xdf\x91\x5f\xee\xcb\xe5\xfe\x9e\xef\x28\x16\xd6\x36\xd9\xff\x8a\xea\xa0\x3f\xeb\xce\xc7\x63\x60\x66\xc8\x16\x2b\xc6\xf4\x0a\x17\x95\x48\xed\x80\xa0\xde\x6c\x7e\xd3\xf3\x41\xcb\x5e\xbf\xba\x76\x55\xc2\x5a\x59\x50\x6b\x5a\x55\xcb\x87\x43\x60\x70\x50\xcd\x62\x5c\x4b\x82\x3c\xf2\x28\xd0\x08\x01\xdc\xd0\x25\xe3\xdb\x6a\x75\x29\x17\xb9\xe9\xc3\x33\xfb\xb8\x52\x99\xa0\xea\x05\xd1\xd7\x18\xb0\x0e\x3d\x7f\xfa\x50\xa2\x4b\xfb\xb4\x60\xb4\x7a\xd3\xf4\xb9\x93\xc9\xb4\x19\xb9\xbf\x87\x80\xfa\x1f\x3b\xcc\x7d\x98\xfe\x15\x0e\x2a\x2b\x3b\x39\xf8\xc8\xab\xf8\x37\x61\x61\x71\x67\xfe\xd3\x04\x9c\xfe\xc7\xf3\xae\x71\x7f\x56\xcc\xfb\x4d\x5f\x94\x8b\x2f\x39\x78\x51\x6e\x7c\x19\x58\x26\x8e\xf3\x94\xf2\x90\x45\xea\x85\x56\x71\x05\xfe\xcd\x7c\x1f\x58\x13\x66\x3e\xa7\xf5\xfb\xf0\xf3\xc9\xbb\xb7\x4e\x47\x7d\x60\x73\x36\xce\xbf\x02\x00\x00\xff\xff\x7d\x69\xab\x82\x78\x2a\x00\x00")

func cppgoGoBytes() ([]byte, error) {
	return bindataRead(
		_cppgoGo,
		"cppgo.go",
	)
}

func cppgoGo() (*asset, error) {
	bytes, err := cppgoGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cppgo.go", size: 10872, mode: os.FileMode(420), modTime: time.Unix(1503017232, 0)}
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
	"cppgo.go": cppgoGo,
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
	"cppgo.go": &bintree{cppgoGo, map[string]*bintree{}},
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

