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
			modTime:          time.Date(2020, 6, 11, 7, 15, 32, 592531103, time.UTC),
			uncompressedSize: 476,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x51\xc1\x4e\xc3\x30\x0c\xbd\xf7\x2b\x4c\xce\x84\xb0\x1b\x07\x27\x52\x35\x40\x70\x62\x42\x9b\x18\xc7\xac\xf5\x16\x4b\x69\x32\x3a\x6f\x13\x7f\x8f\xd2\x0e\x51\x09\x0e\x9c\xde\x8b\xfd\xde\x93\x1d\xe3\xd5\xfd\xcb\x7c\xf9\xbe\x78\x80\x20\x5d\x74\x15\x16\x80\xe8\xd3\xce\x2a\x4a\xca\x55\x15\x06\xf2\xad\xab\x00\xb0\x23\xf1\xd0\x04\xdf\x1f\x48\xac\x5a\x2d\x1f\xf5\x9d\xfa\x69\x24\xdf\x91\x55\x27\xa6\xf3\x3e\xf7\xa2\xa0\xc9\x49\x28\x89\x55\x67\x6e\x25\xd8\x96\x4e\xdc\x90\x1e\x1e\xd7\xc0\x89\x85\x7d\xd4\x87\xc6\x47\xb2\xb3\x9b\xdb\x49\x50\x10\xd9\x6b\xfa\x38\xf2\xc9\xaa\xb5\x5e\xd5\x7a\x9e\xbb\xbd\x17\xde\x44\x9a\xa4\x32\x59\x6a\x77\x34\xfa\x84\x25\x92\xab\x97\x6b\x5d\xef\x28\x09\x9a\xb1\x50\xa1\x19\x87\xaf\x70\x93\xdb\xcf\x41\x1a\x66\xee\x89\x62\xcc\x30\x51\x87\xd9\xd0\x3a\xc6\x02\x00\x18\x79\x24\x00\xe8\x21\xf4\xb4\xb5\xca\x08\xf5\x9d\x72\x6f\xb4\x81\xc2\x38\xf9\x88\xc6\x5f\xf4\xe6\xdb\xf0\xa7\x93\xd3\x36\x2b\x57\x2f\x9e\x61\xa0\xff\xb5\xf5\xd4\x65\x21\xe5\x5e\x07\x84\xf2\xb1\x70\x29\xfe\x8a\x40\x53\x66\x47\x33\x6e\x59\xd6\x2e\xc7\xfc\x0a\x00\x00\xff\xff\x39\x8e\x00\x46\xdc\x01\x00\x00"),
		},
		"/remote.html": &vfsgen۰CompressedFileInfo{
			name:             "remote.html",
			modTime:          time.Date(2020, 6, 11, 7, 15, 38, 79346835, time.UTC),
			uncompressedSize: 8494,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x3a\xcd\x73\xdb\xc6\xf5\x77\xfd\x15\xef\x87\x9f\x67\x08\x26\x24\x20\x39\x75\xe3\xd0\x84\x1a\x5a\x56\x62\x37\x96\xe5\xca\x52\x15\x4f\xa7\x69\x96\xc0\x23\xb1\xd6\x72\x17\xd9\x5d\xf0\x23\x16\x67\xda\x5b\x3a\x99\xe9\xa9\xd3\x53\xda\x43\x2f\xed\xa9\xd3\x53\x3b\xd3\x1e\xfa\xd7\xc4\x69\xff\x8b\xce\x2e\x3e\x08\x80\x1f\x91\xeb\xcc\x84\x07\x71\xf1\xf6\x7d\x7f\xec\xbe\x07\xaa\xff\x7f\x0f\x4e\x8f\xce\x9f\x3f\x3d\x86\x58\x4f\xd8\xe1\x5e\xdf\x7c\x01\x23\x7c\x1c\x38\xc8\x9d\xc3\xbd\xbd\x7e\x8c\x24\x3a\xdc\x03\xe8\x4f\x50\x13\x08\x63\x22\x15\xea\xc0\xb9\x38\xff\xa0\x7b\xd7\x59\x6d\x70\x32\xc1\xc0\x99\x52\x9c\x25\x42\x6a\x07\x42\xc1\x35\x72\x1d\x38\x33\x1a\xe9\x38\x88\x70\x4a\x43\xec\xda\x87\x0e\x50\x4e\x35\x25\xac\xab\x42\xc2\x30\x38\xf0\xf6\x2b\x8c\x62\xad\x93\x2e\x7e\x96\xd2\x69\xe0\x7c\xdc\xbd\x18\x74\x8f\xc4\x24\x21\x9a\x0e\x19\x56\xb8\x52\x0c\x30\x1a\x63\x46\xa7\xa9\x66\x78\x78\x86\x13\xa1\xd1\xa2\x48\xc1\xfa\x7e\x06\xdd\xeb\xfb\x99\x05\x7b\xfd\xa1\x88\x16\x16\x3f\xa2\x53\xa0\x51\xe0\x90\x24\xb1\x0c\x00\xfa\xf1\x3b\x87\xaf\xfe\xf0\xab\x6f\xfe\xf2\xcb\x57\xbf\xfe\xf2\xd5\x57\xff\xf8\xf7\xbf\xbe\xfa\xe6\xcf\x5f\xbe\xfa\xcd\x9f\xbe\xfe\xe2\x6f\xd0\x57\x13\xc2\x18\x28\xbd\x60\x18\x38\xa1\x60\x42\xf6\x60\x2c\xc9\xc2\x81\x69\x57\xe3\x5c\x07\xce\x04\x95\x22\x46\x9d\xbe\x6f\x91\x0f\xfb\x7e\xfc\x4e\xce\xda\x48\xcb\x69\x27\x44\x8e\x29\xef\xc1\x9d\x64\x0e\xfb\xc9\x3c\x17\x0e\xd0\x57\x09\xe1\x05\xd2\x90\x84\x57\x63\x29\x52\x1e\x75\x33\x59\x43\x96\xe2\x94\x0a\x86\xfa\x1e\x64\x90\x59\x4c\x99\x46\x23\x5e\xc5\x62\x16\x38\x12\x49\x24\x38\x5b\x38\x87\x67\x48\xa2\x53\xce\x16\x7d\xdf\xb0\x2c\xf9\x0f\x53\xad\x05\x87\xf7\x43\x46\xc3\x2b\x83\x3f\x92\xa8\xe2\x47\x93\xb1\x73\xf8\xf5\x17\x7f\x7f\xf5\xbb\xbf\x7e\xf3\xdb\x7f\xfe\xe7\xf7\x7f\x84\x51\xa2\x7a\xb9\x36\x85\x69\xa3\x44\x79\x5a\x7c\x40\xe7\x18\xb9\x07\x6d\x6b\xa1\xe1\xdc\xf7\x33\x9e\xdb\x44\x5c\xe1\x02\xa7\xc8\xb5\xdb\x7a\x78\x7a\x72\xdc\x6a\x3b\x87\xe6\x7b\x8d\x8a\xf2\x24\xd5\x30\xed\x4e\x44\x84\xcc\x38\x77\x32\x21\x3c\x72\x20\x61\x24\xc4\x58\xb0\x08\x65\xe0\xa8\x18\x19\x83\x72\x6f\xda\x15\xbc\x77\x85\x8b\x34\xf1\x90\x6b\x83\x20\x53\x7e\x94\xef\xe6\x4e\xf7\x23\x3a\x5d\xf9\x7f\x25\x6e\x32\x06\x89\xa3\xc0\x89\xa8\x4a\x98\x89\x60\xe1\x74\x21\x23\x94\x3d\x38\x48\xe6\xa0\x04\xa3\x11\x18\xaf\xdf\x73\xc0\x6f\x30\xcc\x17\x66\xa5\x42\x49\x13\x0d\x4a\x86\x81\xe3\xfb\x61\xc4\xbd\x17\x2a\x42\x46\xa7\xd2\xe3\xa8\xfd\xb1\x3f\x4d\xf1\xfd\xdb\xde\x81\x77\xb0\xdf\x79\xf1\x59\x8a\x72\xf1\xfe\x3b\xe6\xc9\xfa\xd0\xd2\x1e\xae\xd8\x64\x62\x42\xc1\x95\x86\xfb\x8f\x07\x4f\x3e\xfa\xc5\xa3\x93\x0f\x21\xc8\x15\x6f\x45\x44\x93\x1e\x9d\x90\x31\xfa\x63\x3a\xba\x37\x24\x0a\x7f\xf8\x83\xce\xd9\x3e\xfb\xf0\xf4\x01\x8b\x07\x3f\x19\xdc\x1f\x98\xcf\xd1\xc3\x3b\xf7\x07\xc7\x1f\x0d\x06\xc7\x83\xc7\x16\x60\xe0\xc7\x83\xc1\xe0\xd1\xd1\xf9\xe0\x78\x70\x3a\x0b\x82\xd6\x9e\xe5\xe9\xbf\xf5\x56\xce\xfb\x2d\x38\x13\x9a\x68\x2a\x38\x90\xd1\x08\x43\xad\x40\xc7\x08\x2a\x94\x88\x1c\x88\x82\x91\x60\x4c\xcc\x54\xaf\xc0\x2f\xe9\xd6\x3f\xfb\x11\x8e\xb7\x6e\x5f\x77\xed\xe7\x7a\x3b\x02\x9c\x1c\x3f\xb9\x80\xeb\xff\x89\x43\xb7\x7b\x68\x38\xe4\x8c\x60\x0b\xd6\x75\xf9\x27\xff\x9a\xee\xd0\x26\xfb\x7a\x03\x84\x2d\xea\x5a\xf8\x75\xf7\xba\x8a\x79\xbd\x19\xd3\xfe\x39\xa9\x62\xe6\xb0\xcd\x98\xc7\x3b\x31\xdf\x33\xe1\xc9\xa1\x4f\xaa\x98\x17\xb9\x5b\x6e\xbf\xbb\x29\x80\xd7\x15\x9c\x1c\xf2\x64\x97\x12\x50\xc5\x3c\xbe\x31\xe6\xc9\x16\xcc\x1b\x3b\xeb\x46\x69\xf2\x49\x33\x01\x36\xa7\x49\xb7\x5b\xc5\xea\x77\xbb\xdf\x43\x9a\x54\x38\x5c\x3c\x39\x3e\x79\x13\x0e\x70\x70\xb7\x1a\xd9\x12\xef\x32\xa6\x61\x0c\x0c\x49\xa4\x40\x0b\x5b\xf7\x59\xb5\x53\x3e\x86\x09\x49\x12\xca\xc7\xeb\x75\x9f\x8b\x2b\x3e\xd7\xb5\xaf\x6d\xab\x15\x75\xd3\xb8\x22\x2b\xb3\xf4\x34\xab\x4c\x5b\xb3\xca\x32\xb2\x62\xfa\x9b\xca\x3e\x7a\xf6\x0c\xa4\x39\xef\xd0\x6d\xaf\x64\x77\x4b\xd9\xdd\x95\xec\x02\x56\xa1\x1e\x9a\x6b\xd9\xf8\x66\x96\x05\x76\x56\xc4\x37\x2e\x23\x3d\x83\x3a\x6c\x13\x75\x9c\xa1\xc4\xdb\x68\x56\xb0\x0a\x75\x22\x14\xcc\x2b\x49\x35\xcf\x31\xe3\xee\x22\x5f\xcd\xba\x05\x0c\x16\x9b\xa8\x17\x15\xea\x82\x26\x63\x59\xe7\x13\x67\x7c\xbe\x1b\x9f\xfb\x76\x31\x4a\x79\x68\xaf\x98\x50\x08\x19\x29\xb7\x70\xc5\x65\xa7\xf4\xca\xc3\x0e\x48\x64\x1f\xdb\xbf\xcf\x3b\x59\x90\xa8\xe0\x6d\x78\x99\xb3\x9a\x12\x09\xb3\x0e\xc4\x1d\x98\x77\x60\x71\x6f\x2f\x07\xab\x19\xd5\x61\x0c\xee\x3a\x01\x40\x48\x14\xc2\x7e\x6f\x6f\x95\x6e\x33\x08\x4a\x89\x97\x15\x78\x5c\x81\x3f\xac\xc0\xe7\x10\x58\xb5\x2a\xa0\x45\x06\x7a\x5e\x01\x0d\x25\x92\xab\xba\xd4\xf7\xb6\x8a\x7d\xb8\x45\xec\x65\x43\x6c\x89\x0f\xdd\xa6\xbc\xc5\xba\x56\x1b\x54\x38\xb8\xfb\xc6\xa6\x97\xf8\x99\x0e\x4d\x37\xec\x50\x71\x83\x3e\xb7\xdf\x7d\x63\x9f\x6c\xf0\xc3\x0e\x15\xab\x3a\x2c\x8b\x7c\x91\xa8\x53\xc9\x2b\x49\x32\x7f\xda\x83\x39\xf8\x30\xeb\x94\xa0\xc5\xd3\x1e\x2c\xc0\x87\xb8\x53\x52\x57\x78\xcc\x28\x8f\xc4\xcc\x9b\x4e\x20\x00\x8e\x33\xf8\x69\x8a\x6e\xc1\x0e\x59\x0f\x9c\xff\x37\x93\x45\x41\x6a\x5b\xb7\x8a\xb8\x7c\x56\xe8\x81\xe3\x74\x2a\xfe\xe1\x53\xa2\x7a\xc0\x53\xc6\x56\x50\x3a\x19\x37\x41\x45\x96\xf7\x60\xbf\x42\x9d\x35\xbf\x75\x96\xc5\x50\xd0\x03\x2d\x53\x5c\xc1\x6d\x8f\x5f\x21\x4e\x50\x52\x11\x3d\x32\xbd\xe5\x91\x48\xb9\xae\x6c\x2e\x8b\xc5\xc4\x6c\x60\xe4\xd6\x6a\xcb\xf6\xaa\xa6\xa3\x0e\x40\xc7\x54\x79\x95\xe5\x2d\x89\x23\xe5\xe5\x4d\x76\xd5\x1c\xcf\xf6\xdb\xde\x84\xcc\x2f\xcd\x28\x08\x01\x38\x77\xf7\xcd\x28\xb4\x19\xe9\x21\xd2\x71\xac\x21\x28\x5c\x4e\x39\x47\x99\x03\xbb\x16\x75\x8c\xfa\x7e\x9e\x00\x47\x8c\x22\xd7\x67\x18\x6a\xb7\xed\x69\x91\xc0\xdb\xe0\x18\xce\x25\x6b\xdf\xcf\xb8\xcb\x10\x02\xf8\xd4\x0c\x9a\x3d\xdf\xbf\xf5\x92\x89\xd0\xba\xd4\x8b\x85\xd2\x4b\x3f\x6b\x7c\x55\x2c\xf4\xa7\x4d\x4a\xc1\x99\x20\x11\x04\xe0\xb6\x21\x38\xac\xf8\xc2\xf7\x01\x32\xcb\xb9\xd0\x74\xb4\x70\x1d\xdb\xab\x83\x41\xc7\xc8\x69\x57\x11\x97\xe5\x83\x25\x50\x0b\x1e\x3e\xc8\x1c\xe5\xb6\xd7\xb7\xce\x45\x1a\xc6\x09\x89\xdc\x76\xcd\x8e\x90\xb0\x30\x65\x44\xa3\x09\x67\xb9\xa1\x50\x3f\x32\x13\xd1\x94\x30\xd7\x6d\x07\x87\x2f\x2b\xb5\x60\x39\x8e\x12\x55\x84\xa8\x19\x76\xf0\x61\xdf\xbb\xd3\x24\x58\xc3\x0a\x60\xbf\xc4\x59\x76\xe0\xce\xfe\x7e\x7b\x2d\x59\x50\xc7\x22\x52\xd5\x9c\x2f\xa7\xc1\x2b\x5c\x54\x73\xa8\xac\xc6\x5b\x1e\x79\x41\xe6\x6e\x75\x07\x20\x95\xa6\x98\x7c\x3b\xfe\x55\x12\x7b\x25\xa3\x07\x4e\x22\x94\x6e\xec\x35\x0b\xae\x59\x23\xd9\xcc\x59\xa8\x04\x0e\xbc\x6d\x1e\xea\x3c\x96\x95\xa7\x65\xbb\x62\xf1\xaa\xbc\xca\x99\xd3\xfd\xbe\x2d\xb2\x91\xca\x9f\x1a\x66\xac\xf3\x39\x5f\x24\xe6\xec\x79\xa1\x04\xaf\x49\x59\xb6\x3d\x1d\x23\x77\x5d\x89\x2a\x69\xa4\x77\x51\xee\x82\xa1\xc7\xc4\xd8\xad\x0a\x34\x55\xd6\xb3\x4e\x34\x84\x9e\x48\x75\x92\xea\x76\x8d\xb4\x86\x1e\x80\xe3\xd4\x76\x09\x43\xa9\xdd\x2d\xc4\x5b\x7c\x5f\xbe\xbf\x68\xf8\x7e\x4b\x09\x9a\x9e\xcb\xf3\x3c\xa7\xdd\x44\x7d\xad\xe3\xe0\x47\x3a\xb8\xf5\xd2\x9c\xf7\x0f\x6c\xf7\x68\x8e\x9e\x73\x3a\x41\xb7\xbd\xfc\x74\x93\x92\xf5\xea\xad\xa9\x99\x1d\x9d\xc8\x70\x82\xb6\xa4\x0a\x5d\xee\xed\x55\x90\x18\x6a\xd0\x86\xfe\xd9\x82\x1b\x0d\x5d\x91\xa0\xb4\x9a\x75\xc0\x66\xee\x86\x20\x99\x16\x09\x21\xc8\xf6\xef\xd5\xb6\xe8\x08\x5c\xf4\x84\xa4\x63\xca\x09\x3b\xce\x18\x34\x53\xca\xd2\xd6\x91\xb6\x96\x05\x00\x7a\x89\xb4\x92\x1e\xe0\x88\xa4\x4c\x57\xcf\xa8\xc2\x82\x79\xc6\x71\x34\x52\xa8\x3f\xee\xd8\x0b\xbb\x78\x7c\xbe\x86\x3c\xb3\xbb\x9a\xc8\x31\x6a\x2f\xb4\xc7\xf9\x65\xf6\xc6\x30\x5e\xdf\xc9\xae\x81\x35\x1e\xf6\x9d\xa2\x49\xb3\xbc\xd3\x5c\x75\x8c\x9d\xcc\xcf\x65\x9f\x58\xf7\xcf\x4c\x79\x0a\x79\xe4\xfe\xf8\xd9\xe9\x13\x4f\x69\x49\xf9\xd8\x64\x51\xd3\x43\x65\x10\x7a\x50\x89\x87\xef\x43\xda\x81\xa8\x03\x61\x07\x66\x0d\x0a\xca\x23\x9c\xd7\x6e\xdc\xfc\xde\x95\xa8\x54\x2a\xb1\x67\xce\xdd\xe6\xa6\xe9\x49\x32\x4b\xbc\xf9\xd3\xe6\xe6\x62\xb5\xb9\x68\x6c\x2e\xdb\xed\x1b\x59\x55\xb5\xa3\x15\xb6\x1a\x74\xcb\x6a\x1c\xcb\xd6\x7d\x22\x52\x85\x27\x62\x8a\x8f\xa9\xd2\xc8\x51\xba\xb8\x29\x8b\xca\x9c\x75\x5b\x93\x56\x91\xaa\x37\x63\x7e\x91\xdc\x98\x75\xba\x89\x35\x14\x35\xe5\x49\x9c\x88\x29\xda\x04\x2e\x59\xb6\xac\x0c\x03\x6f\x75\xd6\x8d\x69\xa4\x43\x24\xc2\xf4\x5b\x38\xa5\x49\xc1\x67\xa5\x77\x8d\xcb\x0e\x4b\x1f\x88\x19\xbf\xb1\xad\xd1\x6e\x5b\x49\x14\x7d\x17\x86\xae\xb1\x71\x72\x2b\x9d\x75\x2b\xb7\x19\x69\x8b\x58\xe5\x6d\xf1\x25\x0e\x9f\x89\xf0\x0a\xb5\xeb\xcc\x54\xcf\xf7\xcd\x1d\x51\x3b\x5c\xcd\xd5\xe1\x4f\x28\xa7\xd6\x56\xa7\x76\x7a\xcc\x94\x27\xb8\x48\x90\x9b\x93\x4f\xe2\xa6\xc3\x2e\xab\xe6\xbc\xcb\x85\x00\x46\x84\x29\xdc\x7a\x67\x39\xa5\x20\x03\xe6\x18\xea\x5a\x63\xb6\xb3\x58\x7c\x3f\x0b\x87\xb9\xe2\x50\x77\x60\x44\xe7\x30\x8b\x91\x43\xf6\xb3\x06\x50\x05\x22\xd5\x62\x54\xfc\xf8\xb0\xfd\xcc\x70\xa4\xb3\xbb\x60\xb7\x06\x35\x8b\x46\x24\x66\xbc\x88\x47\x35\x87\xea\x11\x69\xb8\x31\x9f\x38\x8c\x27\xf3\xe5\xb7\xdc\xef\x15\x5f\x49\x0c\xa7\x46\x60\x4e\xb7\x2d\xf0\x56\x4e\xc8\x84\xc2\x0d\xdd\xf1\xa6\x60\x99\x99\xe4\x46\xb1\x32\x3c\x9b\x81\xda\x55\xe5\x37\xf7\x53\xa3\x5d\xd8\x29\x73\xb9\xed\x86\x2f\x5a\xf7\x6c\x80\xab\x17\xf1\x96\x6a\x68\xd9\x6a\x68\x6d\xad\x86\x90\x24\x35\xd9\xaf\xe3\x66\xfb\x8e\x84\x48\x53\x35\x65\x61\x87\x12\x89\xc6\xe3\xcc\x67\xae\x13\xd1\x69\x33\xef\x89\xe4\xd9\x68\x75\x8e\x73\xd3\x8f\x38\x97\x38\x54\x56\xd9\xdc\x17\x1d\x38\xcb\x5a\x2e\xd0\xc2\x0c\xd9\x82\x44\x9e\xb3\xce\x23\x9b\xdb\xec\x6f\x54\x86\x8b\xc4\x68\x2b\xd2\x48\x70\xfd\x8c\x7e\x6e\xec\x70\x6e\xd7\x06\xc0\xda\xa1\x34\x14\xd1\xc2\x23\x49\x82\x3c\x3a\x8a\x29\x8b\x5c\xc3\xa2\x71\xc8\x36\xdc\x83\x52\x5a\xf1\xe5\x81\xeb\xb6\x77\xa4\x7a\x2b\x27\x68\x75\x80\xc8\xb1\x15\xaa\xde\xbc\x98\x4c\xb3\x95\xef\x7a\xa6\xe9\x06\xca\x95\x26\x3c\x44\x31\x82\xfb\x4c\x0c\xd7\x3b\xaf\xcd\xd3\xd6\xdb\x01\x1c\x98\xd3\x27\x46\x96\xc0\x48\xc8\x2d\x73\xdf\x2a\xdd\x86\x4c\x0c\xf3\x84\x33\x72\xdc\x9f\x55\xb5\xf8\x79\x67\x4d\x2c\x80\xb6\xf3\x40\x2b\xfb\x6d\xe9\x45\x82\xe3\x56\x03\x65\xd9\xde\x20\xe7\xe2\xec\xf1\x6a\x2a\x37\x0f\xd7\xd7\xc5\xd3\x0c\x87\x57\x54\x5f\x9c\x3d\xde\x40\x96\x42\x60\x48\xf3\x94\x3c\x1d\xbe\xc0\xd0\x60\xba\x46\xed\xf6\x26\x8f\xac\x1a\xf4\xb4\x7e\x6c\x02\x32\x85\xd6\xcf\xbe\x75\xb0\xa2\x9f\x63\x0f\x7c\x4f\xa3\xd2\x35\xd7\xb7\xd7\x7d\x6d\xe6\xe7\xea\x69\x23\x31\x44\x3a\xc5\xf2\xcd\xcc\xea\xbc\xcb\x38\x6c\x95\xfc\x49\xd1\x4b\xde\x4c\x6e\xad\xfd\x84\x00\x12\x22\x15\x3e\xe2\x75\x3a\x4f\xa5\x43\xa5\xa5\xdb\x2a\x11\x5b\x1e\x43\x3e\xd6\x71\xbb\x03\x07\xfb\x8d\xdb\xbb\x71\x6e\x96\x6f\x87\x9c\x66\xb3\xbb\xc9\x86\x97\x3b\x59\xbd\x96\x53\x76\x55\x4c\x7e\x8b\xdf\xb8\x1e\x0d\xfe\xb7\x97\x63\xe5\x1c\xce\x4f\xf0\xb2\x1c\xd7\x26\xc2\x55\xcd\xe6\xab\xbd\x26\xbf\xfc\xfd\x5e\xdb\xfe\xc2\x5b\xfc\x22\xdb\xf7\xb3\x7f\x19\xd8\xeb\xfb\xf6\xdf\x23\xfe\x1b\x00\x00\xff\xff\x70\x1b\xd1\x9f\x2e\x21\x00\x00"),
		},
		"/terminal.html": &vfsgen۰CompressedFileInfo{
			name:             "terminal.html",
			modTime:          time.Date(2020, 5, 8, 3, 33, 2, 810333512, time.UTC),
			uncompressedSize: 2606,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x56\x61\x6f\x23\x35\x13\xfe\x9e\x5f\x31\xaf\x5f\x50\x37\x5c\xe2\x4d\x8a\x90\x8e\x34\x5b\x7a\x2d\x45\x02\x24\x0e\xa9\xad\x00\x89\x2f\x8e\x3d\x9b\xb8\x75\xec\x3d\x7b\x36\xe9\x82\xfa\xdf\x91\xbd\x49\x76\x73\xd0\x9e\x0e\x1d\x1f\x92\xf5\x7a\x3c\xcf\xf3\xcc\x78\xc6\xde\xf9\xff\xbe\x7d\x7b\x75\xfb\xdb\xcf\xd7\xb0\xa2\xb5\x39\x1f\xcc\xe3\x03\x8c\xb0\xcb\x82\xa1\x65\xe7\x83\xc1\x7c\x85\x42\x9d\x0f\x00\xe6\x6b\x24\x01\x72\x25\x7c\x40\x2a\xd8\xdd\xed\x77\xe3\xd7\xac\x33\x58\xb1\xc6\x82\x6d\x34\x6e\x2b\xe7\x89\x81\x74\x96\xd0\x52\xc1\xb6\x5a\xd1\xaa\x50\xb8\xd1\x12\xc7\xe9\x65\x04\xda\x6a\xd2\xc2\x8c\x83\x14\x06\x8b\x29\x9f\xf4\x80\x56\x44\xd5\x18\xdf\xd5\x7a\x53\xb0\x5f\xc7\x77\x6f\xc6\x57\x6e\x5d\x09\xd2\x0b\x83\x3d\x54\x8d\x05\xaa\x25\xb6\x7e\x46\xdb\x07\xf0\x68\x0a\xa6\xa5\xb3\x0c\x56\x1e\xcb\x82\xe5\x22\x04\xa4\x90\x13\xfa\xb5\xb6\xc2\x70\x2d\xdd\x37\x9b\x62\xda\xfa\x90\x26\x83\xe7\x82\x1e\xc7\x62\x89\x96\x60\xbf\x6a\x9e\xb7\x96\x63\xdc\x40\x8d\xc1\xb0\x42\xa4\x3d\x7a\x94\x19\x66\x79\x2e\x95\xe5\xf7\x41\xa1\xd1\x1b\xcf\x2d\x52\x6e\xab\x75\xfe\x18\xd1\x2e\x4e\xf9\xd7\xfc\x34\x57\x3a\x50\x3b\xc1\xd7\xda\x72\x19\x02\xfb\xc4\xe0\x42\x29\x67\x43\x5e\xd6\xc6\x04\xe9\x11\x6d\x6f\x78\xcc\x99\x98\xe2\x08\xe0\x0b\xf8\x33\x3d\x01\xd6\xc2\x2f\xb5\x9d\xc1\xe4\x2c\x4d\x3c\xa5\xff\xf4\x17\x8b\x61\x94\x46\x0b\xa7\x9a\x83\xc3\x0a\xf5\x72\x45\x33\x98\x4e\x26\x9f\xff\xcd\xe7\xff\x49\xde\x0b\x6b\x01\x16\x42\x3e\x2c\xbd\xab\xad\x1a\x4b\x67\x9c\x9f\xc1\xc2\x08\xf9\xb0\x37\x57\x42\x29\x6d\x97\x63\x83\x25\xcd\xe0\xab\xea\xb1\xe3\x98\xe7\xbb\x08\xe6\x79\x5b\x96\x83\x79\x54\x96\x62\x53\x7a\x03\x5a\x15\x2c\xf1\xb3\xf3\x79\xae\xf4\xa6\x0d\x5a\x7a\x5d\x11\x04\x2f\x3f\x90\xd8\xfb\x77\x35\xfa\xe6\xe2\x4b\x7e\xca\xa7\x6d\x66\xdb\x99\x94\xc3\xfb\x10\x31\x5b\xac\x8f\x83\x7d\xa1\x18\x3e\x21\xe8\xbe\x08\x34\xc5\xdf\x7f\x85\xfe\x6c\x89\xfd\x7b\x2e\xe9\xc2\xb8\xcd\xf3\xd8\x63\xd0\x7f\xe0\xc5\x94\x4f\xf9\x64\x9f\xfb\x85\xd8\x4d\xbf\x44\xd3\x96\xf4\x46\xf8\xd4\xc6\x67\x87\xb7\x2d\x2e\x82\x93\x0f\x48\x50\x80\xc5\x2d\xfc\x82\x8b\x9b\xf4\x9e\xb1\x6d\x14\xc4\xe0\x15\x18\x27\x05\x69\x67\xf9\xca\x05\x82\x57\xc0\xd2\x81\xc1\x86\x2d\xca\x01\x81\x2f\xb4\x15\xbe\xb9\x6d\x2a\x84\x02\x98\xf0\x5e\x34\x8b\xba\x2c\xd1\xb3\xb3\x41\x5a\x5a\xd6\x56\x46\x20\x10\x8b\xd3\x40\x3e\x5b\xd4\xe5\xf0\xd0\x07\x1e\xa9\xf6\x16\x6e\xc8\x6b\xbb\xe4\xa5\x77\xeb\xab\x95\xf0\x57\x4e\x21\x17\x55\x65\x9a\xcc\xd6\xc6\x8c\x92\xc8\x3b\x6d\xe9\xf5\x9b\x88\x9f\x20\x86\xfd\x16\xeb\xe4\x38\xeb\x2a\xb4\x50\x1c\x68\x33\xdc\x50\xc7\x97\x9a\xb0\x0d\xfa\x76\x77\xb2\x65\x7b\x1b\x40\xbb\x71\x3f\x62\x13\x66\x40\xbe\xc6\xd1\xc1\x52\x07\xbc\x89\x2d\xf6\xfe\xbc\xac\x7d\x70\xfe\x32\x1e\x5a\xc7\xa6\xa7\xe1\x2e\xfc\x96\x93\x3b\x9b\x9d\x28\x41\xe2\x64\xd4\x29\x8b\xef\x9d\xb4\x7e\x14\x01\xad\xca\x5a\x91\x8f\x74\x6d\xa5\x53\xe8\xb3\x21\xc7\x34\xca\xd8\xef\x8f\x93\x49\xdc\xa3\x04\x30\x3c\xeb\x51\xbe\xc7\xd8\x96\x48\x9f\xf3\x28\x1b\x1f\x47\x39\x8d\x94\x3f\xdc\xbc\xfd\x89\x87\xb4\x5d\xba\x6c\x7a\xb9\x03\x90\xce\x84\x19\xe0\x86\x78\x1c\x8d\x7a\x16\xef\xb6\x3b\x4b\x1c\x1d\x0c\x4f\xc3\xe1\xf0\x79\xed\xe9\xbe\xe9\x4b\x4f\x13\x7d\xf1\xca\xc9\x7a\x8d\x96\x78\xb2\x40\x01\xe9\x79\xf6\xdc\x0e\x54\x68\xb3\x83\xcb\x12\xe9\xda\x60\x1c\x5e\x36\xdf\xab\xec\x24\xb5\xf6\xc9\xf0\x58\x46\xa9\x29\x3b\xcc\x7c\x96\xb1\xf6\x10\x67\x43\xde\xe6\x35\x3b\x48\xeb\xab\xea\x3c\x0f\x42\xf6\x3a\xfa\x65\xba\xc6\x10\xc4\x12\x9f\xad\x54\x00\x5d\x42\x9c\xe2\x71\x97\x41\xdb\x40\xc2\x4a\x74\x25\xa4\x1e\xb8\x4c\x3d\xd6\x5f\xbe\x63\xde\x7a\x4d\x98\xed\x7a\x6d\xef\xde\xc5\x05\xf0\x04\x68\x02\x1e\x39\x0a\x83\x9e\xba\xc5\xdd\xd2\xc1\xf1\xb3\xaf\x5f\x1a\x17\x5e\x52\xdf\x13\xc3\x6e\x30\x84\x78\x02\xec\x3e\x26\x08\x15\xeb\x09\x4a\x2b\x15\x06\xf2\xae\xe9\xd2\xfd\x4f\x9c\xe8\xbd\xf3\x1f\xc8\x18\x35\x55\x4c\x92\x74\x36\x38\x83\xdc\xb8\x25\x14\x05\xb0\xbd\x0b\x3b\x4e\x59\x6f\x59\x02\x7b\x26\xf4\xdd\x25\xbb\x3b\x54\xe7\x79\x7b\xbb\xc6\xeb\x36\x7e\x1e\xfe\x15\x00\x00\xff\xff\x2c\xc1\x8f\x20\x2e\x0a\x00\x00"),
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
