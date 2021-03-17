// Code generated by vfsgen; DO NOT EDIT.

// +build vfs

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2021, 3, 15, 6, 9, 35, 278427387, time.UTC),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2021, 3, 17, 10, 29, 32, 132603751, time.UTC),
			uncompressedSize: 465,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x51\xcd\x4e\xf3\x30\x10\xbc\xf7\x29\xf6\xf3\xf9\x33\xa6\x37\x0e\x6b\x4b\x51\x01\xc1\x89\x0a\xa5\x50\x8e\x6e\xb2\xad\x57\xf2\x4f\x49\xb7\x8d\x78\x7b\x94\xa4\x88\x4a\x70\xe0\x34\xe3\xdd\x99\xd1\x58\x8b\xff\x6e\x9f\x16\xf5\xdb\xf2\x0e\x82\xa4\xe8\x66\x38\x00\x44\x9f\x77\x56\x51\x56\xc3\x80\x7c\xeb\x66\x00\x98\x48\x3c\x34\xc1\x77\x07\x12\xab\x56\xf5\xbd\xbe\x51\xdf\x8b\xec\x13\x59\x75\x62\xea\xf7\xa5\x13\x05\x4d\xc9\x42\x59\xac\xea\xb9\x95\x60\x5b\x3a\x71\x43\x7a\x7c\xfc\x07\xce\x2c\xec\xa3\x3e\x34\x3e\x92\x9d\x5f\x5d\x5f\x04\x05\x91\xbd\xa6\xf7\x23\x9f\xac\x5a\xeb\x55\xa5\x17\x25\xed\xbd\xf0\x26\xd2\x45\x2a\x93\xa5\x76\x47\x93\x4f\x58\x22\xb9\xaa\x5e\xa3\x99\xe8\x0c\xcd\x54\x1b\x37\xa5\xfd\x18\x35\x61\xee\x1e\x28\xc6\x02\x55\xbd\xd6\xd5\x8e\xb2\xa0\x09\xf3\x71\x75\x8c\x03\x00\x60\xe4\x89\x00\xa0\x87\xd0\xd1\xd6\x2a\x23\xd4\x25\xe5\x5e\x69\x53\x53\x97\x38\xfb\x88\xc6\x9f\xe5\xe6\x4b\xff\xab\x91\xf3\xb6\x28\x57\x2d\x1f\x61\xa4\x7f\xb5\x75\x94\x8a\x90\x72\xcf\x23\xbe\x30\xf5\x70\x9e\xfd\x48\x40\x33\x34\x47\x33\xfd\x11\xcd\x78\xc1\xcf\x00\x00\x00\xff\xff\x34\x0b\xd7\x3b\xd1\x01\x00\x00"),
		},
		"/remote.html": &vfsgen۰CompressedFileInfo{
			name:             "remote.html",
			modTime:          time.Date(2021, 3, 17, 10, 27, 48, 366795977, time.UTC),
			uncompressedSize: 8406,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x5a\xcd\x72\xdc\xc6\xf1\xbf\xf3\x29\xfa\x8f\xbf\xaa\x16\x6b\xef\x02\xa4\x12\xc5\xf2\x8a\x60\xbc\xa2\x68\x4b\xb1\x28\x29\x22\x19\x5a\x95\x8a\xe3\x59\xa0\x77\x31\xe2\xec\x0c\x3c\x33\x58\xec\x5a\x64\x55\x72\x73\xca\x55\xc9\x25\x95\x93\x93\x43\x2e\x49\xe5\x90\xe4\x94\x54\xc5\x87\x3c\x8d\x65\xe7\x2d\x52\x33\xf8\x58\x00\xfb\x61\x2a\x72\x95\xf7\x40\x00\x3d\xdd\xbf\x9e\xfe\x9a\xe9\x01\xb8\xff\x7f\xf7\x1e\x1f\x9e\x3e\x7b\x72\x04\xb1\x9e\xb2\x83\x9d\x7d\x73\x01\x46\xf8\x24\x70\x90\x3b\x07\x3b\x3b\xfb\x31\x92\xe8\x60\x07\x60\x7f\x8a\x9a\x40\x18\x13\xa9\x50\x07\xce\xd9\xe9\xbb\xfd\xdb\xce\x72\x80\x93\x29\x06\xce\x8c\x62\x96\x08\xa9\x1d\x08\x05\xd7\xc8\x75\xe0\x64\x34\xd2\x71\x10\xe1\x8c\x86\xd8\xb7\x0f\x3d\xa0\x9c\x6a\x4a\x58\x5f\x85\x84\x61\xb0\xe7\xed\xd6\x80\x62\xad\x93\x3e\x7e\x9c\xd2\x59\xe0\x7c\xd0\x3f\x1b\xf6\x0f\xc5\x34\x21\x9a\x8e\x18\xd6\x50\x29\x06\x18\x4d\x30\x97\xd3\x54\x33\x3c\x78\xf9\x87\x5f\x7e\xf5\xd7\x5f\xbc\xfc\xd5\x67\x2f\x3f\xff\xd7\xd7\xff\xfe\xfc\xab\x3f\x7f\xf6\xf2\xd7\x7f\xfa\xf2\xd3\x7f\xec\xfb\xf9\xf8\xce\xbe\x9f\xdb\xb2\xb3\x3f\x12\xd1\xc2\x4a\x46\x74\x06\x34\x0a\x1c\x92\x24\x16\xaa\x20\x29\xbd\x60\x18\x38\x53\x22\x27\x94\x0f\xe0\x56\x32\x87\xdd\x64\x5e\x70\x00\xec\xab\x84\xf0\x92\x69\x44\xc2\x8b\x89\x14\x29\x8f\xfa\xa1\x60\x42\x0e\x46\x2c\xc5\x19\x15\x0c\xf5\x1d\xc8\x29\x59\x4c\x99\x46\x07\x66\x7d\x15\x8b\x2c\x70\x24\x92\x48\x70\xb6\x70\x0e\xbe\xfc\xcd\x5f\xbe\xfe\xdb\x17\xfb\xbe\x01\xac\xd0\x47\xa9\xd6\x82\xc3\x3b\x21\xa3\xe1\x85\xe1\x1e\x4b\x54\xf1\x83\xe9\xc4\x39\xf8\xf2\xd3\x7f\xbe\xfc\xdd\xdf\xbf\xfa\xed\x17\xff\xf9\xfd\x1f\xdf\x7d\x72\x32\x28\xa6\x32\xeb\x6b\x9c\xeb\xc0\x19\x27\xca\xd3\xe2\x5d\x3a\xc7\xc8\xdd\xeb\x3a\x07\x05\xf0\xbe\x9f\x43\x6e\xd2\x70\x81\x0b\x9c\x21\xd7\x6e\xe7\xfe\xe3\xe3\xa3\x4e\xd7\x39\x30\xd7\x15\x29\xca\x93\x54\xc3\xac\x3f\x15\x11\xb2\xc0\x09\xc5\x74\x4a\x78\xe4\x40\xc2\x48\x88\xb1\x60\x11\xca\xc0\x51\x31\x32\x06\xd5\xd8\xac\x2f\xf8\xe0\x02\x17\x69\xe2\x21\xd7\x86\x41\xa6\xfc\xb0\x18\x2d\x3c\xee\x47\x74\xb6\x74\xfe\x52\xdd\x74\x02\x12\xc7\x81\x13\x51\x95\x30\xb2\x70\x2a\x8f\x0b\x19\xa1\x1c\xc0\x5e\x32\x07\x25\x18\x8d\xc0\xb8\xfc\x8e\x03\x7e\x0b\x70\x79\xa3\x42\x49\x13\x0d\x4a\x86\x81\xe3\xfb\x61\xc4\xbd\xe7\x2a\x42\x46\x67\xd2\xe3\xa8\xfd\x89\x3f\x4b\xf1\x9d\x9b\xde\x9e\xb7\xb7\xdb\x7b\xfe\x71\x8a\x72\xf1\xce\xf7\xcc\x93\x75\xa1\x95\xad\xc1\xe4\x5a\x42\xc1\x95\x86\xbb\x0f\x87\x8f\xde\xff\xf9\x83\xe3\xf7\x20\x28\xe6\xdd\x89\x88\x26\x03\x3a\x25\x13\xf4\x27\x74\x7c\x67\x44\x14\xfe\xe0\xfb\xbd\xa7\xbb\xec\xbd\xc7\xf7\x58\x3c\xfc\xf1\xf0\xee\xd0\xfc\x0e\xef\xdf\xba\x3b\x3c\x7a\x7f\x38\x3c\x1a\x3e\xb4\x04\x43\x3f\x1a\x0e\x87\x0f\x0e\x4f\x87\x47\xc3\xc7\x59\x10\x74\x2c\xa4\xff\xc6\x1b\x05\xf4\x1b\xf0\x54\x68\xa2\xa9\xe0\x40\xc6\x63\x0c\xb5\x02\x1d\x23\xa8\x50\x22\x72\x20\x0a\xc6\x82\x31\x91\xa9\x41\xc9\x5f\xc9\xad\xfe\x76\x23\x9c\x6c\x1c\xbe\xec\xdb\xdf\xe5\x66\x06\x38\x3e\x7a\x74\x06\x97\xff\x13\x42\xbf\x7f\x60\x10\x0a\x20\xd8\xc0\x75\x59\xfd\x29\x2e\xb3\x2d\xb3\xc9\x2f\xaf\xc1\xb0\x61\xba\x96\x7e\xd9\xbf\xac\x73\x5e\xae\xe7\xb4\x7f\x8e\xeb\x9c\x05\x6d\x3d\xe7\xd1\x56\xce\xb7\x4d\x78\x0a\xea\xa3\x3a\xe7\x59\xe1\x96\x9b\x6f\xad\x0b\xe0\x65\x8d\xa7\xa0\x3c\xda\x36\x09\xa8\x73\x1e\x5d\x9b\xf3\x78\x03\xe7\xb5\x9d\x75\xad\x34\xf9\xb0\x9d\x00\xeb\xd3\xa4\xdf\xaf\x73\xed\xf7\xfb\xdf\x41\x9a\xd4\x10\xce\x1e\x1d\x1d\xbf\x0e\x02\xec\xdd\xae\x47\xb6\xe2\x3b\x8f\x69\x18\x03\x43\x12\x29\xd0\xc2\xd6\x7d\x5e\xed\x94\x4f\x60\x4a\x92\x84\xf2\xc9\x6a\xdd\x17\xea\xca\xdf\x65\xe3\xb2\xe9\x6e\x29\xdd\x36\xae\xcc\xca\x3c\x3d\xcd\x5d\x3e\x5b\x73\x97\x67\x64\xcd\xf4\xd7\xd5\x7d\x78\x72\x02\xd2\xac\x77\xe8\x76\x97\xba\xfb\x95\xee\xfe\x52\x77\x49\xab\x49\x8f\xcc\x8e\x6c\x7c\x93\xe5\x81\xcd\xca\xf8\xc6\x55\xa4\x33\x68\xd2\xd6\x49\xc7\x39\x4b\xbc\x49\x66\x49\xab\x49\x27\x42\xc1\xbc\x96\x54\xf3\x82\x33\xee\x2f\x8a\xbb\xac\x5f\xd2\x60\xb1\x4e\x7a\x51\x93\x2e\x65\x72\xc8\x26\x4e\x9c\xe3\x7c\x3b\x3e\xf7\xed\xcd\x38\xe5\xa1\xdd\x62\x42\x21\x64\xa4\xdc\xd2\x15\xe7\xbd\xca\x2b\xf7\x7b\x20\x91\x7d\x60\xff\x3e\xeb\xe5\x41\xa2\x82\x77\xe1\x45\x01\x35\x23\x12\xb2\x1e\xc4\x3d\x98\xf7\x60\x71\x67\xa7\x20\xab\x8c\xea\x30\x06\x77\x55\x00\x20\x24\x0a\x61\x77\xb0\xb3\x4c\xb7\x0c\x82\x4a\xe3\x79\x8d\x1e\xd7\xe8\xf7\x6b\xf4\x39\x04\x76\x5a\x35\xd2\x22\x27\x3d\xab\x91\x46\x12\xc9\x45\x53\xeb\xdb\x1b\xd5\xde\xdf\xa0\xf6\xbc\xa5\xb6\xe2\x87\x7e\x5b\xdf\x62\x75\x56\x6b\xa6\xb0\x77\xfb\xb5\x4d\xaf\xf8\xf3\x39\xb4\xdd\xb0\x65\x8a\x6b\xe6\x73\xf3\xad\xd7\xf6\xc9\x1a\x3f\x6c\x99\x62\x7d\x0e\x57\x65\xbe\x48\xd4\xa9\xe4\xb5\x24\x99\x3f\x19\xc0\x1c\x7c\xc8\x7a\x15\x69\xf1\x64\x00\x0b\xf0\x21\xee\x55\xd2\x35\x8c\x8c\xf2\x48\x64\xde\x6c\x0a\x01\x70\xcc\xe0\x27\x29\xba\x25\x1c\xb2\x01\x38\xff\x6f\x3a\xff\x52\xd4\x76\x6e\x35\x75\x53\x54\x8a\x4c\x70\x00\x8e\xd3\xab\xf9\x87\xcf\x88\x1a\x00\x4f\x19\x5b\x52\xe9\x74\xd2\x26\x95\x59\x3e\x80\xdd\x9a\x74\xde\xfa\x36\x21\xcb\xf3\xc0\x00\xb4\x4c\x71\x49\x1f\x27\xaa\x21\x9c\xa0\xa4\x22\x7a\x60\x5a\xcb\x43\x91\x72\x5d\x1b\xbc\x2a\x6f\xa6\x66\x00\x23\xb7\x51\x5b\xb6\x55\x35\xfd\x74\x00\x3a\xa6\xca\xab\xdd\xde\x90\x38\x56\x5e\xd1\x62\xd7\xcd\xf1\x6c\xb7\xed\x4d\xc9\xfc\xdc\x1c\xda\x20\x00\xe7\xf6\xae\x39\x05\xad\x67\xba\x8f\x74\x12\x6b\x08\x4a\x97\x53\xce\x51\x16\xc4\xbe\x65\x9d\xa0\xbe\x5b\x24\xc0\x21\xa3\xc8\xf5\x53\x0c\xb5\xdb\xf5\xb4\x48\xe0\x4d\x70\xea\xc8\xbe\x9f\x83\xcb\x10\x02\xf8\xc8\x9c\x08\x07\xbe\x7f\xe3\x05\x13\xa1\xf5\xa8\x17\x0b\xa5\xaf\xfc\xbc\xef\x55\xb1\xd0\x1f\xb5\x25\x05\x67\x82\x44\x10\x80\xdb\x85\xe0\xa0\xe6\x0a\xdf\x07\xc8\x0d\xe7\x42\xd3\xf1\xc2\x75\x6c\xa7\x0e\x86\x1d\x23\xa7\x5b\x67\xbc\xaa\x1e\xac\x80\x5a\xf0\xf0\x5e\xee\x27\xb7\xbb\x3a\x74\x2a\xd2\x30\x4e\x48\xe4\x36\x30\x42\xc2\xc2\x94\x11\x8d\x26\x98\xd5\x80\x42\xfd\xc0\x9c\x86\x66\x84\xb9\x6e\x37\x38\x78\x51\xab\x04\x0b\x38\x4e\x54\x19\xa0\x76\xd0\xc1\x87\x5d\xef\x56\x5b\x60\x85\x2b\x80\xdd\x8a\xe7\xaa\x07\xb7\x76\x77\xbb\x2b\xa9\x82\x3a\x16\x91\xaa\x67\x7c\x75\x12\xbc\xc0\x45\x3d\x83\xaa\x5a\xbc\xe1\x91\xe7\x64\xee\xd6\x47\x00\x52\x69\x4a\xc9\xb7\x47\xbf\x5a\x5a\x2f\x75\x0c\xc0\x49\x84\xd2\xad\xb1\x76\xb9\xb5\x2b\x24\x3f\x6f\x96\x53\x02\x07\xde\x34\x0f\x4d\x8c\xab\xda\xd3\x55\xb7\x66\xf1\xb2\xb8\xaa\xf3\xa6\xfb\x5d\x5b\x64\x23\x55\x3c\xb5\xcc\x58\xc5\x39\x5d\x24\x66\xe5\x79\xae\x04\x6f\x68\xb9\xea\x7a\x3a\x46\xee\xba\x12\x55\xd2\xca\xee\xb2\xd8\x05\x43\x8f\x89\x89\x5b\x57\x68\x6a\x6c\x60\x9d\x68\x04\x3d\x91\xea\x24\xd5\xdd\x86\x68\x83\x3d\x00\xc7\x69\x8c\x12\x86\x52\xbb\x1b\x84\x37\xf8\xbe\x7a\x75\xd1\xf2\xfd\x86\x0a\x34\x1d\x97\xe7\x79\x4e\xb7\xcd\xfa\x4a\xab\xc1\x0f\x75\x70\xe3\x85\x59\xed\xef\xd9\xde\xd1\x2c\x3c\xa7\x74\x8a\x6e\xf7\xea\xa3\x75\x93\x6c\x16\x6f\x63\x9a\xf9\xc2\x89\x0c\xa7\x68\x4b\xaa\x9c\x4b\xd5\xd1\x98\x1f\x43\x0d\xda\xc8\x9f\x2c\xb8\x99\xa1\x2b\x12\x94\x76\x66\x3d\xb0\x99\xbb\x26\x48\xa6\x41\x42\x08\xf2\xf1\x3b\x8d\x21\x3a\x06\x17\x3d\x21\xe9\x84\x72\xc2\x8e\x72\x80\x76\x4a\x59\xd9\x26\xd3\xc6\xb2\x00\x40\x2f\x91\x56\xd3\x3d\x1c\x93\x94\x69\xb7\xbb\xd3\x18\x37\x16\xcc\x73\xc4\xf1\x58\xa1\xfe\xa0\x67\xb7\xeb\xf2\xf1\xd9\x0a\x73\x66\x47\x35\x91\x13\xd4\x5e\x68\x17\xf3\xf3\xfc\xcd\x5e\xbc\x3a\x92\x6f\x02\x2b\x18\xf6\xdd\x9f\x49\xb3\xa2\xcf\x5c\xf6\x8b\xbd\xdc\xcf\x55\x97\xd8\xf4\x4f\xa6\x3c\x85\x3c\x72\x7f\x74\xf2\xf8\x91\xa7\xb4\xa4\x7c\x62\xb2\xa8\xed\xa1\x2a\x08\x03\xa8\xc5\xc3\xf7\x21\xed\x41\xd4\x83\xb0\x07\x59\x4b\x82\xf2\x08\xe7\x8d\xfd\xb6\xd8\x75\x25\x2a\x95\x4a\x1c\x98\x75\xb7\x3d\x68\x3a\x92\xdc\x12\x6f\xfe\xa4\x3d\xb8\x58\x0e\x2e\x5a\x83\x57\xdd\xee\xb5\xac\xaa\xdb\xd1\x09\x3b\x2d\xb9\xab\x7a\x1c\xab\xc6\x7d\x2a\x52\x85\xc7\x62\x86\x0f\xa9\xd2\xc8\x51\xba\xb8\x2e\x8b\xaa\x9c\x75\x3b\xd3\x4e\x99\xaa\xd7\x03\x3f\x4b\xae\x0d\x9d\xae\x83\x86\xb2\xa6\x3c\x89\x53\x31\x43\x9b\xc0\x15\x64\xc7\xea\x30\xf4\x4e\x6f\xd5\x98\x56\x3a\x44\x22\x4c\xbf\x01\x29\x4d\x4a\x9c\xe5\xbc\x1b\x28\x5b\x2c\xbd\x27\x32\x7e\x6d\x5b\xa3\xed\xb6\x92\x28\xfa\x36\x0c\x5d\x81\x71\x0a\x2b\x9d\x55\x2b\x37\x19\x69\x8b\x58\x15\x4d\xf1\x39\x8e\x4e\x44\x78\x81\xda\x75\x32\x35\xf0\x7d\xb3\x47\x34\x16\x57\xb3\x75\xf8\x53\xca\xa9\xb5\xd5\x69\xac\x1e\x99\xf2\x04\x17\x09\x72\xb3\xf2\x49\x5c\xb7\xd8\xe5\xd5\x5c\xf4\xb8\x10\xc0\x98\x30\x85\x1b\xf7\x2c\xa7\x52\x64\xc8\x1c\x43\xdd\xe8\xcb\xb6\x16\x8b\xef\xe7\xe1\x30\x5b\x1c\xea\x1e\x8c\xe9\x1c\xb2\x18\x39\xe4\x9f\x1f\x80\x2a\x10\xa9\x16\x63\xfb\x1d\x41\x0a\xb6\x79\xcd\x70\xa4\xb3\xbd\x60\x37\x06\x35\x8f\x46\x24\x32\x5e\xc6\xa3\x9e\x43\xcd\x88\xb4\xdc\x58\x9c\x37\x8c\x27\x8b\xdb\x6f\xd8\xdf\x6b\xbe\x92\x18\xce\x8c\xc2\x42\x6e\x53\xe0\xad\x9e\x90\x09\x85\x6b\x9a\xe3\x75\xc1\x32\x27\x92\x6b\xc5\xca\x60\xb6\x03\xb5\xad\xca\xaf\xef\xa7\x56\xbb\xb0\x55\xe7\xd5\xa6\x1d\xbe\xec\xdc\xf3\xe3\x5b\xb3\x88\x37\x54\x43\xc7\x56\x43\x67\x63\x35\x84\x24\x69\xe8\x7e\x15\x37\xdb\x37\x24\x44\x9a\xaa\xa9\x0a\x3b\x94\x48\x34\x1e\xe5\x3e\x73\x9d\x88\xce\xda\x79\x4f\x24\xcf\x0f\x56\xa7\x38\x37\xfd\x88\x73\x8e\x23\x65\x27\x5b\xf8\xa2\x07\x4f\xf3\x96\x0b\xb4\x30\x47\x6c\x41\x22\xcf\x59\xc5\xc8\x4f\x6d\xf6\xe3\x94\x41\x91\x18\x6d\x64\x1a\x0b\xae\x4f\xe8\x27\xc6\x0e\xe7\x66\xe3\xf8\xd7\x58\x94\x46\x22\x5a\x78\x24\x49\x90\x47\x87\x31\x65\x91\x6b\x20\x5a\x8b\x6c\xcb\x3d\x28\xa5\x55\x5f\x2d\xb8\x6e\x77\x4b\xaa\x77\x0a\x81\x4e\x0f\x88\x9c\x58\xa5\xea\xf5\x8b\xc9\x34\x5b\xc5\xa8\x67\x9a\x6e\xa0\x5c\x69\xc2\x43\x14\x63\xb8\xcb\xc4\x68\xb5\xf3\x5a\x7f\xda\x7a\x33\x80\x3d\xb3\xfa\xc4\xc8\x12\x18\x0b\xb9\xe1\xdc\xb7\x4c\xb7\x11\x13\xa3\x22\xe1\x8c\x1e\xf7\xa7\xf5\x59\xfc\xac\xb7\xa2\x16\x40\xdb\xf3\x40\x27\xff\xb0\xf4\x3c\xc1\x49\xa7\xc5\x72\xd5\x5d\xa3\xe7\xec\xe9\xc3\xe5\x99\xdc\x3c\x5c\x5e\x96\x4f\x19\x8e\x2e\xa8\x3e\x7b\xfa\x70\x8d\x58\x0a\x81\x11\x2d\x52\xf2\xf1\xe8\x39\x86\x86\xd3\x35\xd3\xee\xae\xf3\xc8\xb2\x41\x4f\x9b\xcb\x26\x20\x53\x68\xfd\xec\x5b\x07\x2b\xfa\x09\x0e\xc0\xf7\x34\x2a\xdd\x70\x7d\x77\xd5\xd7\xe6\xfc\x5c\x5f\x6d\x24\x86\x48\x67\x58\xbd\x97\x59\xae\x77\x39\xc2\x46\xcd\x1f\x96\xbd\xe4\xf5\xf4\x36\xda\x4f\x08\x20\x21\x52\xe1\x03\xde\x94\xf3\x54\x3a\x52\x5a\xba\x9d\x8a\xb1\xe3\x31\xe4\x13\x1d\x77\x7b\xb0\xb7\xdb\xda\xbd\x5b\xeb\x66\xf5\x6e\xc8\x69\x37\xbb\xeb\x6c\x78\xb1\x15\xea\x95\x9c\xb2\xad\x62\x8a\x5d\xfc\xda\xf5\x68\xf8\xbf\xb9\x1c\x6b\xeb\x70\xb1\x82\x57\xe5\xb8\x72\x22\x5c\xd6\x6c\x71\xb7\xd3\xc6\x2b\xde\xee\x75\xed\xd7\xdd\xf2\x73\xec\xbe\x9f\x7f\xd0\xdf\xd9\xf7\xed\xbf\x31\xfc\x37\x00\x00\xff\xff\x16\xf9\x70\x25\xd6\x20\x00\x00"),
		},
		"/terminal.html": &vfsgen۰CompressedFileInfo{
			name:             "terminal.html",
			modTime:          time.Date(2021, 3, 17, 10, 29, 46, 123756223, time.UTC),
			uncompressedSize: 2589,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x56\x51\x6f\x23\x35\x10\x7e\xef\xaf\x18\x0c\xa8\x1b\x2e\xf1\x26\x45\x48\x47\x9a\x2d\x77\x57\x8a\x04\x48\x1c\x52\x5b\x01\x12\x2f\x8e\x3d\x9b\xb8\x75\xec\x3d\x7b\x36\xe9\x82\xfa\xdf\x91\xbd\xc9\xee\xe6\x44\x7b\x3a\x74\x3c\x24\xb6\x67\xc6\xdf\xf7\x79\xec\xb1\x77\xf1\xd9\xf7\x6f\x2f\x6f\xfe\xf8\xf5\x0a\xd6\xb4\x31\x17\x27\x8b\xd8\x80\x11\x76\x55\x30\xb4\x2c\x1a\x50\xa8\x8b\x13\x80\xc5\x06\x49\x80\x5c\x0b\x1f\x90\x0a\x76\x7b\xf3\xc3\xe4\x25\xeb\x1d\x56\x6c\xb0\x60\x5b\x8d\xbb\xca\x79\x62\x20\x9d\x25\xb4\x54\xb0\x9d\x56\xb4\x2e\x14\x6e\xb5\xc4\x49\x1a\x8c\x41\x5b\x4d\x5a\x98\x49\x90\xc2\x60\x31\xe3\xd3\x01\xd0\x9a\xa8\x9a\xe0\xbb\x5a\x6f\x0b\xf6\xfb\xe4\xf6\xf5\xe4\xd2\x6d\x2a\x41\x7a\x69\x70\x80\xaa\xb1\x40\xb5\xc2\x76\x9e\xd1\xf6\x1e\x3c\x9a\x82\x69\xe9\x2c\x83\xb5\xc7\xb2\x60\xb9\x08\x01\x29\xe4\x84\x7e\xa3\xad\x30\x5c\x4b\xf7\xdd\xb6\x98\xb5\x73\x48\x93\xc1\x8b\x9b\xbd\x6f\x91\xb7\xe3\x63\xb4\x40\x8d\xc1\xb0\x46\xa4\x03\x66\x14\x17\xe6\x79\x2e\x95\xe5\x77\x41\xa1\xd1\x5b\xcf\x2d\x52\x6e\xab\x4d\xfe\x10\x99\x5e\x9d\xf1\x6f\xf9\x59\xae\x74\xa0\xd6\xc0\x37\xda\x72\x19\x02\xfb\xc4\xe0\x42\x29\x67\x43\x5e\xd6\xc6\x04\xe9\x11\xed\xa0\x7b\xcc\x99\x98\x62\x0f\xe0\x2b\xf8\x3b\xb5\x00\x1b\xe1\x57\xda\xce\x61\x7a\x9e\x0c\x8f\xe9\x3f\x6e\xfe\x38\xf5\x96\x4e\x35\x5d\xec\x1a\xf5\x6a\x4d\x73\x98\x4d\xa7\x5f\x0e\xc3\xd3\xdf\xe7\x49\xd9\x33\xb1\x00\x4b\x21\xef\x57\xde\xd5\x56\x4d\xa4\x33\xce\xcf\x61\x69\x84\xbc\x3f\xb8\x2b\xa1\x94\xb6\xab\x89\xc1\x92\xe6\xf0\x4d\xf5\xd0\x73\x2c\xf2\xbd\xf8\x45\xde\x9e\xc3\x93\x45\x54\x96\x96\xa5\xf4\x16\xb4\x2a\x58\xe2\x67\x17\x8b\x5c\xe9\x6d\xbb\x5e\xe9\x75\x45\x10\xbc\xfc\x40\x4e\xef\xde\xd5\xe8\x9b\x57\x5f\xf3\x33\x3e\x6b\x93\xda\x5a\x52\xfa\xee\x42\xc4\x6c\xb1\x3e\x0e\xf6\x99\x73\xf0\x09\x41\x0f\xfb\xaf\x29\xfe\xfe\x2f\xf4\x27\x4f\xd7\x7f\xe7\x92\x2e\x4c\xda\x3c\x4f\x3c\x06\xfd\x17\xbe\x9a\xf1\x19\x9f\x1e\x72\xbf\x14\x7b\xf3\x73\x34\xed\x69\xde\x0a\x0f\x51\xf8\x79\x37\xda\xe1\x32\x38\x79\x8f\x04\x05\x58\xdc\xc1\x6f\xb8\xbc\x4e\xe3\x8c\xed\xa2\x20\x06\x2f\xc0\x38\x29\x48\x3b\xcb\xd7\x2e\x10\xbc\x00\x96\x6e\x08\x36\x6a\x51\x3a\x04\xbe\xd4\x56\xf8\xe6\xa6\xa9\x10\x0a\x60\xc2\x7b\xd1\x2c\xeb\xb2\x44\xcf\xda\xc8\xb2\xb6\x32\xe2\x80\x58\x9e\x05\xf2\xd9\xb2\x2e\x47\x5d\x19\x78\xa4\xda\x5b\xb8\x26\xaf\xed\x8a\x97\xde\x6d\x2e\xd7\xc2\x5f\x3a\x85\x5c\x54\x95\x69\x32\x5b\x1b\x33\x4e\x1a\x6f\xb5\xa5\x97\xaf\x23\x7c\x82\x18\x0d\x2b\xac\x57\xe3\xac\xab\xd0\x42\xd1\xd1\x66\xb8\xa5\x9e\x2f\xd5\x60\xbb\xe6\xc3\x9d\x96\x1d\x7c\x00\xed\xbe\xfd\x8c\x4d\x98\x03\xf9\x1a\xc7\x9d\xa7\x0e\x78\x1d\x2b\xec\x7d\xbb\xac\x7d\x70\xfe\x4d\xbc\xae\x8e\x5d\x8f\xa3\xf3\x93\x01\x27\x77\x36\x3b\x55\x82\xc4\xe9\xb8\x57\x16\xc7\xbd\xb4\xe1\x2a\x02\x5a\x95\xb5\x22\x1f\xe8\xca\x4a\xa7\xd0\x67\x23\x8e\xa9\x97\xb1\x3f\x1f\xa6\xd3\xb8\x45\x09\x60\x74\x3e\xa0\x7c\x8f\xb1\x3d\x21\x43\xce\xa3\x6c\x7c\x1c\xe5\x2c\x52\xfe\x74\xfd\xf6\x17\x1e\xd2\x76\xe9\xb2\x19\xe4\x0e\x40\x3a\x13\xe6\x80\x5b\xe2\xb1\x37\x1e\x78\xbc\xdb\xed\x3d\xb1\xd7\x39\x1e\x47\xa3\xd1\xd3\xda\xd3\x4b\x33\x94\x9e\x0c\x43\xf1\xca\xc9\x7a\x83\x96\x78\xf2\x40\x01\xa9\x3d\x7f\x6a\x07\x2a\xb4\x59\x37\x65\x85\x74\x65\x30\x76\xdf\x34\x3f\xaa\xec\x34\x55\xf6\xe9\xe8\x58\x46\xa9\x29\xeb\x2c\x5f\x64\xac\xbd\xc3\xd9\x88\xb7\x79\xcd\x3a\x69\x43\x55\xfd\xcc\x4e\xc8\x41\xc7\xf0\x98\x6e\x30\x04\xb1\xc2\x27\x4f\x2a\x80\x2e\x21\x9a\x78\xdc\x65\xd0\x36\x90\xb0\x12\x5d\x09\xa9\x06\xde\xa4\x12\x1b\x86\xef\x99\x77\x5e\x13\x66\xfb\x5a\x3b\x4c\xef\xd7\x05\xf0\x08\x68\x02\x1e\x4d\x14\x06\x3d\xf5\xc1\x7d\xe8\xc9\x71\x3b\xd4\x2f\x8d\x0b\xcf\xa9\x1f\x88\x61\xd7\x18\x42\xbc\x01\xf6\x9f\x18\x84\x8a\x0d\x04\xa5\x48\x85\x81\xbc\x6b\xfa\x74\xff\x1b\x27\x7a\xef\xfc\x07\x32\x46\x4d\x15\x93\x24\x9d\x0d\xce\x20\x37\x6e\x05\x45\x01\xec\x30\x85\x1d\xa7\x6c\x10\x96\xc0\x9e\x58\xfa\xfe\x8d\xdd\xdf\xa9\x8b\xbc\x7d\x5c\xe3\x6b\x1b\xbf\x06\xff\x09\x00\x00\xff\xff\x7b\xb3\xad\x5b\x1d\x0a\x00\x00"),
		},
		"/terminal.ico": &vfsgen۰CompressedFileInfo{
			name:             "terminal.ico",
			modTime:          time.Date(2020, 5, 8, 3, 33, 2, 810682451, time.UTC),
			uncompressedSize: 67646,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x9d\x6f\x8c\x54\xd5\x19\xc6\x7f\xb3\x53\x98\x15\xa1\x5b\x6c\x41\xa5\xb5\x0b\x2a\x96\xda\x94\x54\x62\x0a\x1f\x48\xac\xda\x6e\xfa\x41\x49\x41\xa1\xc4\x34\x21\xa6\x9f\x1a\x1b\x03\xa9\x69\x8d\x29\x32\x09\x09\x35\x44\x63\xd9\x9a\xb4\x89\xa9\x7e\x10\xc3\x27\x3e\x61\xad\xab\xa0\x43\x69\x4b\x30\xd6\xc4\x18\x4b\x4d\x4c\xdc\x75\xb5\xda\x5a\xca\x82\x0a\x0b\xec\xee\x98\xd9\x7d\x87\x9d\x39\xcc\xee\xde\x73\xef\xfc\x39\x77\xf6\x79\x92\x37\xe4\x5d\xee\xb9\x77\xee\x79\xce\x79\xdf\x73\xde\xf3\xde\x73\x20\x43\x86\x7c\xbe\xf4\xef\x52\x56\x74\x66\x58\x0c\xac\x00\xf2\x50\xfa\xe3\xf8\xdf\x27\x90\xa1\xe7\x27\x8c\x4b\x9d\xd1\x01\x2c\x04\x6e\x06\xb6\x00\x0f\x01\xcf\x02\xef\x03\x45\x49\x22\x19\x01\x5e\x07\x7a\x81\x6d\xc0\x7a\x60\x39\x70\x19\xad\xc7\x55\xc0\xfd\xc0\x7e\xe0\x15\xe0\x6d\xe0\x9c\x38\x6b\xa8\x9c\xb4\xf6\xd0\x07\x3c\x05\xdc\x01\x74\x36\x99\xf7\x0c\xf0\x3d\xe0\x10\x70\x41\x9c\xb4\xbc\x3d\xec\x01\xbe\xde\x24\xee\xaf\x04\x1e\x06\x06\x55\xf7\xc1\xc8\x18\x70\x14\xd8\x00\xcc\x6d\x10\xef\x25\x1f\x7f\x2b\xf0\xbc\x3d\x4f\xf5\x1e\x9e\xfc\x1f\x78\x0c\xe8\x6e\x40\x9f\x7f\x10\xf8\x50\x75\x9c\x0a\xf9\x3b\xb0\x0e\xc8\x25\xe4\x3d\x0b\xac\xb5\xb1\x86\xfc\x7c\xfa\xc6\x05\x8f\x02\xd7\x24\xe0\x7f\xbd\xfc\x7c\xea\xc7\x05\x87\x62\xb6\x81\x65\xc0\xab\xaa\xc3\xb6\x68\x03\x8f\x7a\xfa\x82\x9c\xcd\x29\x54\x7f\xed\xe3\x0b\xd6\x79\xf0\xbf\xd1\xca\xd4\x63\x3c\xfa\x0e\xf0\x26\xb0\x0f\xd8\x29\x89\x24\xbb\x81\x83\xc0\x71\x60\x00\x38\x53\xa7\x31\x61\x94\x79\xc1\x52\xe0\xb5\x04\xcf\x29\xc7\x2a\xb7\x03\x77\x02\xd7\x02\x0b\x80\x39\x36\x87\x94\x44\x93\x4e\xe0\x0a\x60\x25\x70\x2f\xb0\x0b\x78\x1a\xf8\x6f\x02\x6e\x1e\x9b\x21\x3e\x70\x19\xf0\x44\xcc\xf9\x7d\xc9\x5e\xfc\x16\xb8\xde\xe2\x83\x42\xfd\x51\xf2\xcb\x3d\x36\x1f\x1b\x8d\x69\x8f\x37\x4c\x73\xff\xbb\x80\xd3\x31\xee\x7b\xcc\x7c\xc6\x5c\x51\xd4\xb4\x18\x6c\xc9\x4f\xfc\x3b\x06\x57\x47\xa7\x88\x15\x2f\x8f\x61\xf7\x4f\x98\xbd\x58\x26\x4a\x9a\x8e\x2c\x70\x9b\xad\xbb\x9d\xf3\x9c\x0f\xec\x71\xd6\x8c\x72\x66\xbb\x7d\xe7\x15\xdb\xcd\xaf\x0b\xad\xc3\x57\xad\x0d\xf8\xfa\xea\x3b\x9c\x31\xdf\xbb\x9e\xf7\x78\x0e\x58\xa4\xea\x0f\x02\xb7\xc5\xf0\x05\x4f\x55\x94\x5f\x05\x7c\xe2\x51\xf6\x3f\xb6\xfe\x2b\x84\xe3\x0b\x76\x7a\x8e\x09\xfb\x2a\x72\x48\x7e\x0c\x0c\x47\x2c\x77\xde\xd6\x82\x3a\x54\xed\xc1\x8d\x09\xfb\x3c\xe7\xe9\xcb\xad\x6c\xaf\x47\xb9\x97\x60\x3c\xc5\x4c\x08\x0f\x3d\x1e\xf1\x81\x93\xb6\xbe\x33\xc7\xd6\xf4\xa3\x94\xb9\x00\xdc\xad\x6a\x0e\x3a\x3e\xf0\xb4\x47\x5f\xde\x66\x6b\x43\x47\x23\x5e\x7f\x16\x58\xa3\x6a\x0e\x1a\xbb\x3c\xf8\xef\x35\x9b\x11\x35\xaf\xe3\x38\x70\x9d\xaa\x38\x68\xdc\xeb\xb1\x5e\x50\x1a\x03\xdc\xe7\xd1\x5e\xf6\x01\xf3\x55\xc5\x41\x63\xa5\xad\x19\x45\xcd\x2d\x7f\xc4\x73\xfd\x40\x08\x1b\x57\x98\x9d\x8e\xca\xe9\xcb\x11\xaf\x1b\xb5\xf1\x82\x10\x36\x3a\x6d\xed\x38\x2a\xff\x51\x7d\xc5\xc7\x36\x56\x10\xc2\x46\x87\xe5\x0f\xd4\x3b\x87\xe4\x23\xc5\xfc\x52\xc3\xff\xce\x06\xf0\xaf\x98\xaf\xf8\x17\xff\xe2\x5f\x10\xff\x82\xf8\x17\xc4\xbf\x20\xfe\x05\xf1\x2f\x88\x7f\x41\xfc\x0b\xe2\x5f\x10\xff\x82\xf8\x17\xc4\xbf\x20\xfe\x05\xf1\x2f\x88\x7f\x41\xfc\x0b\xe2\x5f\x10\xff\x82\xf8\x17\xc4\xbf\x20\xfe\x05\xf1\x2f\x88\x7f\x41\xfc\x0b\xe2\x5f\x10\xff\x82\xf8\x17\xda\x9f\xff\x1c\x30\x0f\xf8\x42\x83\xeb\x24\x63\x7b\x13\x2f\x68\x80\xcc\x0b\x6c\x2f\xbc\xd0\xf9\x2f\x71\xf1\x6d\xdb\x4b\xa4\x00\xfc\x05\xf8\x33\xf0\xb3\x06\xec\x2b\xf7\x45\xdb\xdb\x70\xaf\xed\x91\x7a\xa4\x01\x72\x18\x38\x00\x6c\x4d\x78\xd6\xce\x6c\xe1\x7f\x03\xf0\xaf\x29\xf6\xa2\xd9\x0f\xdc\x50\xa7\x7a\x58\x62\x7b\xa1\x9d\x69\xe2\x39\xad\x47\x6c\x2f\x56\xf1\x5f\x1b\xab\x22\xec\x4d\xd3\x07\x7c\xb3\x0e\x7e\xe5\x77\x2d\x3a\x6b\xe5\x48\x8b\xed\x40\xa8\xfc\x5f\x0e\x3c\xe3\xb1\x37\xed\x37\x12\x3c\x6b\x75\xc2\x73\x32\x92\xda\x81\xad\xe2\xff\x12\x94\xf8\xfc\xc0\xe3\x79\x2f\x26\xb0\x03\x9b\x5b\x7c\xde\xd2\x81\x16\x8e\x09\x43\xe5\x7f\xb5\xed\x1f\xea\xf3\xcc\x17\xec\x28\x7b\x5f\xfc\xaa\xc5\xfc\x1f\x6e\xe1\xf9\xec\xa1\xf2\xbf\x12\x18\x8a\xf1\xdc\x17\x63\xb4\x81\x9f\xb6\x98\xff\x83\x4d\x98\xd3\xa6\x8d\xff\xc5\xd6\x2f\xe2\x3c\xdb\x77\x4c\xd8\xe3\xb1\x9f\x7d\x23\xe4\x71\xf9\xff\x9a\xb8\x07\x38\x95\xa0\x0d\x44\x1d\x13\x2e\x32\xbb\xd1\x0a\xee\x07\x80\xef\x8a\xff\x29\xe7\x65\xbb\x63\xf6\xcd\x31\x4f\x3b\x50\xfa\xad\xff\x6c\xc1\xf9\xeb\xbf\xb4\xf3\x38\xc4\x7f\x6d\x7c\xc9\xce\x9a\x3a\xd7\x04\x3b\x70\x13\xf0\xc7\x8a\x33\x34\xcf\x36\x48\x4a\xbc\xff\xc9\xce\x58\x9d\xa3\xf8\x4f\xa4\x36\xf0\x78\x93\xec\xc0\x3c\x3b\xe7\xe4\x66\x3b\xcb\xbe\xde\xb2\x06\xf8\x96\xed\xbb\xac\xf8\x7f\x73\xed\xc0\x0a\x2d\xf7\xa5\x7a\xfd\x2f\x89\x1d\x88\x3b\x37\x14\xff\x61\xad\xff\x97\xdb\x40\x5c\x3b\xf0\x42\x1d\xd6\x0b\xc4\x7f\x6b\xf3\x3f\xca\xbe\x40\x76\x60\x76\xf2\x5f\x0f\x5f\x70\x10\xf8\x8e\xe8\x4f\x75\xfe\x57\x52\x3b\xf0\x9a\x8d\xf3\xc5\x7f\x7a\xf3\xff\x92\xda\x81\xd9\x7e\x2e\x71\x3b\xe4\x7f\x26\x99\x1b\x8e\x00\xbf\x10\xff\xa9\xcf\xff\x2d\xdb\x81\x33\x31\xe7\x04\x59\xf1\x9f\xfa\xfc\xef\x2f\x9b\x3d\xf7\xfd\xad\x7f\x03\xba\xc4\x7f\xaa\xf9\xcf\x58\xee\xee\x40\x8c\xdf\xfa\xb2\xe5\x7b\x8b\xff\x74\xf2\x5f\x7a\x87\x4d\x31\xb9\x2f\xc9\xef\xe5\xff\x53\xcd\x7f\xa9\xdf\x0f\xc6\xfc\x9d\x6f\xdb\xf7\x05\xe2\x3f\x7d\xfc\x77\x24\xb0\xf9\x25\xf9\x9f\xe5\x7e\x69\xfe\x9f\x4e\xfe\x93\xf4\xfb\xf7\x81\x2d\x81\x7d\x8b\x25\xfe\x9b\xd3\xef\x07\xac\x7c\x87\xc2\xbf\xa9\xe4\x7f\x53\x82\x7e\xff\x9e\xe5\xfb\x0b\xe9\xe3\x3f\x63\xdc\x0d\x24\xe0\x7e\x93\xdd\x47\x48\x17\xff\x19\xe3\xee\xbd\x04\x36\x7f\xb3\xb8\x4f\x2d\xff\x9b\x13\x70\x3f\x68\x6d\x47\x48\x1f\xff\x1a\xeb\xcd\x6e\xfe\x93\x8c\xf5\x06\x8d\x7b\x21\x9d\xfc\x47\xf9\xfe\x5f\xfd\xbe\x3d\xf9\xf7\xf9\xfe\x5f\xfd\xbe\xfd\xf8\xf7\xfd\xfe\xbf\xb2\xdf\x6f\x52\xbf\x4f\x3d\xff\x71\xbe\xff\x57\x6c\xa7\x7d\xf8\xf7\xfd\xfe\x5f\xb1\x9d\xf6\xe2\xdf\xe7\xfb\x7f\xc5\x76\xda\x73\xfc\x7f\x0f\x70\x5a\x63\xbd\x59\xcb\x7f\x27\xb0\x6b\x9a\x36\x70\x5c\x63\xbd\xa6\xf3\x3f\xd6\xe4\xf8\x4f\xa7\xd9\x81\xc3\x36\x1e\x38\x6b\xf3\x82\x67\x2c\x3e\x20\x34\x97\xff\xb7\x22\x5e\xf7\x99\xe5\x57\xd4\x0b\x8b\x6d\x4c\xb8\xda\xe6\x86\x97\x8b\xba\xba\x60\x0e\xb0\xcf\x83\xff\x3f\x78\x5c\xfb\x1b\x55\x6f\xf0\x58\x00\xbc\xe9\xc1\xe9\x03\x1e\xd7\x3e\x39\x8b\xf3\xaa\xd3\x82\x6b\x81\x77\x3c\x38\x2d\x8d\xb7\x3e\x89\x78\xed\x5f\x6d\x5f\x65\x21\x5c\xdc\x69\xfb\x11\x45\xcd\x99\xbc\x11\x78\x23\xe2\xf5\xa7\x6c\x4f\x25\x21\x5c\x6c\xf7\xe8\xfb\xcf\x02\xf3\x81\x97\x22\x5e\x7f\x1e\xf8\x81\xaa\x38\x58\x64\x80\x5e\x0f\xfe\x1f\xb2\x72\x7b\x3d\xca\xec\x09\x60\x8f\x33\xa1\x36\xae\x07\x5e\x8f\xc8\xe3\xb9\x8a\xf9\xdc\xcf\x81\x0b\x1e\xdf\x58\xac\x53\x55\x07\x87\xb9\xf6\xcd\xbc\xcf\x77\x52\xe5\xbd\x32\xd6\x02\x9f\x7a\x94\x3d\x06\x74\xab\xca\x83\xc2\x46\xe0\xa4\x07\x87\xaf\x00\x0b\xad\xec\x0d\xc0\xc7\x9e\xf1\xe0\xdd\xf2\x03\xc1\x60\x99\xf5\x49\x1f\xfe\xf6\x57\xc4\xd9\x17\xda\x99\x05\x3e\xe5\x4f\x00\x3f\x52\xd5\x07\x11\xef\x7b\xc2\x93\xbb\x92\xaf\xbf\xdf\xb9\xcf\x0f\x2d\xc6\xef\x73\x9f\x7f\x58\x1b\x90\x1d\x68\x0d\x16\xd9\x7c\xef\x84\x27\x6f\x87\x80\xab\x9c\x7b\x65\x81\x3c\x30\x1a\xc3\x0e\xec\xd6\x78\xa0\xa9\xe8\xb0\xb5\xb8\xe7\x3c\xd6\xef\x2a\xd7\xd7\xa7\x5a\xc7\x5b\x62\x67\xf2\xc5\x59\x1f\x3e\x16\xc8\xfe\xd7\xed\x8e\xc5\xc0\x83\x31\x6c\x75\x79\xad\xf7\xe1\x19\xf2\x6a\x6e\x8f\x99\xa7\x59\x9e\x1b\xee\xb1\x18\xd1\x4d\xd6\x9e\xb4\x5e\x90\x0c\xf3\x81\xeb\x6c\x9f\xf1\xbb\x2d\x56\x77\x3e\x26\x3f\xcf\x03\x57\x46\xb0\x2d\x3b\x63\xd8\x15\x37\x4e\x78\xca\xd6\x0b\x9e\xb4\x75\xc3\x2d\x66\x77\x24\x33\x4b\x0f\xb0\xcd\xce\x48\xdd\x67\x79\x31\x67\x3d\x62\x34\xb5\xe4\x43\xe0\xd6\x88\x6d\xee\x6a\xdb\x5f\xbd\x9e\xf9\x43\x9f\x99\xcd\x92\x4c\x2f\x1f\xd9\x5c\x7c\xb4\x8e\x75\x7f\xc1\x7c\x86\x4f\x5e\xd5\xf7\x63\xfa\x18\x49\x78\xd2\x17\xc1\xee\x53\x63\x3e\xf0\x6b\xdb\x4f\x53\x75\x98\x5e\x19\xb4\xf8\x6e\x1c\x7c\xc5\xda\x80\xec\x40\x3a\xe5\x55\x60\x7d\xc2\xbd\x50\xb3\xe6\x0b\xfa\x12\x8e\x09\x25\xcd\x93\x93\x36\x0f\x5b\x56\xc7\x79\xc8\xd5\x36\x2f\xf8\x40\xf5\x1b\xac\x8c\xd9\x1e\xf8\x1b\xed\xac\xbd\x46\xc4\x9d\x6e\xb7\x18\xd1\xa8\xea\x3b\x28\x39\x6d\xeb\x00\x4b\x9b\x10\x93\x58\x62\xb1\x62\x8d\x0b\xc2\x90\x52\x9f\xbf\xab\xc9\x67\x4e\x67\x6d\xcd\xe8\x80\xcd\x57\x3f\x4d\x18\x9f\x90\x44\x97\x61\xcb\xd9\x7d\xd7\x72\x3e\x96\xb7\x30\x3e\xb9\xd0\xf2\x07\xd6\x5a\x1e\xd1\x5e\x8b\x51\xbe\xe1\x91\x57\x2c\x99\x39\x76\x77\xd4\xe2\xb7\xbd\xf6\x7d\xe4\x2a\xb3\xf5\xb9\x40\x63\xd6\x37\x5a\x6e\xf9\x03\xf6\x8d\xc9\x5b\x9a\x3f\x44\x96\x33\xb6\x87\xfd\x23\xc0\x7d\x16\x0f\xbe\xa6\x11\x6b\x6b\xc5\xe1\x4a\xad\xab\x58\x1c\xaa\xd4\x73\xc5\x62\xa1\x52\xcf\x14\x8b\xd5\xc5\x77\x8c\x39\xfa\x70\xb5\xde\xed\xea\x43\xd5\x7a\xce\xd5\xfb\x1d\x1f\x53\x98\x5e\xcf\xe4\x6b\xbc\xd2\x45\xe4\xed\x1d\xca\x18\xb2\x77\xac\xd2\x6f\x99\xd4\x47\x1c\x7d\xac\xfa\x76\x13\x2f\xef\xe8\x99\x4a\xbd\x70\xa9\x9e\x9d\x4e\xef\xf7\xd7\x73\xd2\xdb\x4c\x1f\xee\x1a\x6f\x48\xdd\x65\xbd\x3f\x3b\xde\x29\xb3\x65\xbd\x90\x19\x6f\xa8\x99\xb2\x9e\xc7\xd5\x77\xf4\x57\xeb\xb7\xf4\x5b\x43\x35\xbd\xdb\xd1\xbb\x1c\x3d\xe7\xe8\x59\x47\xcf\x38\xf7\xa7\x4a\x2f\xc0\x56\x47\x5f\x53\xa9\xf7\xc3\xd7\xac\x1f\x8e\xeb\x43\x93\x7d\x70\xe2\xfd\x27\x2b\x65\x5c\x1f\x99\x6c\xf4\xe3\xfa\xd8\x64\xa7\x9a\xa8\x3f\x57\xcf\x3b\x7a\xe1\x62\x27\x9d\xd0\xfb\x1d\x7d\x68\x86\xff\x77\xf5\xc2\x0c\xcf\x77\x74\xf7\xf7\xbb\xba\xfb\xfe\xae\x5e\x98\xb4\x41\x35\xf5\xfc\xa4\x0d\x9a\x42\xcf\x3a\x7a\x6e\x06\xbd\x7b\x06\x7d\xc7\xb4\x7a\x21\x53\x6d\xaf\x5c\x7d\x28\x17\xd9\xbe\xd5\xb0\x97\xae\x3d\xad\xb2\xbf\xf9\x1a\xf6\x79\x87\xa3\xbb\xf6\xbd\xdb\xd1\x5d\xff\x90\xad\x7a\x5c\xbe\xa2\xc0\xb0\x05\x3e\x2a\xfd\x8f\xeb\x9f\xfa\xa7\xd7\x2f\xf1\x77\x8e\x7f\xec\x72\xfd\xe5\x88\xa3\xbb\xfe\xb7\x98\x77\xfc\x73\xd5\x03\xb3\xc5\x4b\xfc\xfb\x48\xf5\xed\x3e\x0f\x00\x00\xff\xff\x6f\x99\x99\xa3\x3e\x08\x01\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/index.html"].(os.FileInfo),
		fs["/remote.html"].(os.FileInfo),
		fs["/terminal.html"].(os.FileInfo),
		fs["/terminal.ico"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
