// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package migrator

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

// MigrationAssets contains SQL migration scripts and templates
var MigrationAssets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2021, 2, 14, 4, 34, 29, 116954550, time.UTC),
		},
		"/jobsdb": &vfsgen۰DirInfo{
			name:    "jobsdb",
			modTime: time.Date(2021, 2, 14, 4, 34, 29, 115498166, time.UTC),
		},
		"/jobsdb/000001_create_tables.down.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_tables.down.tmpl",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 114442065, time.UTC),
			uncompressedSize: 265,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\x8e\xb1\x8a\x83\x40\x10\x86\xfb\x7d\x8a\xbf\xb8\x56\x5f\xc0\xea\x0e\x3d\x10\x0e\x4e\xa2\x45\x52\x2d\x2b\x6e\x44\x91\x55\x76\x47\x50\x86\x79\xf7\x10\x13\x93\x90\x26\xe5\xcc\x7c\xff\xfc\x5f\x14\x21\xf5\xe3\x04\xe3\x56\xd8\xa5\x0b\xd4\xb9\x16\x8d\x21\x13\x2c\x05\xc5\xec\x8d\x6b\x2d\xe2\xf4\xbe\x11\x51\x00\x90\x1e\xfe\x0b\x54\xdf\x3f\x7f\x19\x98\xbf\xe2\xc2\xdb\x73\xb7\x88\xe8\x7e\xac\x83\x66\x8e\x45\x92\x4f\x9c\x0e\x64\x68\x7e\xd0\xcc\xd6\x35\x22\x4a\xed\x42\xb4\x4e\x16\x3b\x67\xf5\x75\x54\xb7\x77\xa7\x22\x7b\x3b\x24\xcf\x58\x3f\xce\xde\x99\x01\x64\xea\x61\x0f\x6c\xfd\xf9\x2f\xb2\x63\x5e\x56\x25\x98\x5f\x45\x36\x3c\x51\xea\x12\x00\x00\xff\xff\xa5\xa4\xe3\xe3\x09\x01\x00\x00"),
		},
		"/jobsdb/000001_create_tables.up.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_tables.up.tmpl",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 114706736, time.UTC),
			uncompressedSize: 1358,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x54\x51\x6f\xe2\x38\x10\x7e\xe7\x57\xcc\x03\x12\xad\x04\x9c\x74\xf7\x76\x7d\x0a\xe0\xdb\x4d\x0f\x02\x82\x74\xaf\x7d\x8a\x26\xf6\x00\xee\x1a\x3b\x67\x3b\x2d\x51\x94\xff\x7e\x72\x48\xba\x57\xb6\x2b\x35\x8f\x33\xdf\xf7\xcd\xcc\x97\x19\x4f\x26\x10\x6b\xe9\x25\x2a\x78\x21\xeb\xa4\xd1\x60\xf6\x70\x6f\x72\xb7\x98\x81\xc7\x5c\x91\x1b\x43\x8e\x8e\x04\x18\x0d\x9e\x4e\x85\x42\x4f\x20\xd0\x23\x14\xd6\xbc\x48\x41\x02\xf2\x0a\x08\xf9\xb1\xa7\x49\xed\x3c\x6a\x4e\x83\xc9\x04\xe6\x47\xe2\xdf\xe1\xd9\xe4\x4e\xe4\xbf\x39\xf2\x65\x31\x3d\x18\xd8\x1b\x0b\xa8\x14\xe0\x0b\x4a\x15\x8a\xbc\x57\x9e\x0e\x02\xf5\xde\x94\x56\xa3\x72\x97\x36\x06\xf3\x2d\x8b\x52\x06\x69\x34\x5b\x32\x88\xff\x82\x64\x9d\x02\x7b\x8c\x77\xe9\x0e\xea\x7a\xba\xb1\xb4\x97\xe7\xa6\xc9\x9e\x2f\x2c\xb8\x19\x00\x00\x48\x01\xb3\xf8\xcb\x8e\x6d\xe3\x68\x09\x9b\x6d\xbc\x8a\xb6\x4f\xf0\x37\x7b\x1a\xb7\x59\x53\x90\x45\x1f\x66\xfe\x16\x6d\xe7\x5f\xa3\xed\xcd\x1f\xbf\xdf\xb6\xc2\xc9\xc3\x72\x79\xc1\x08\xa3\x09\x66\xeb\xf5\x92\x45\xc9\x15\x2b\x2b\xb0\x52\x06\x05\xdc\xef\xd6\xc9\xec\x8a\xe7\x3c\x5a\x9f\x79\x79\x22\x48\xe3\x15\xdb\xa5\xd1\x6a\x73\x05\x21\x2d\xae\x00\xb7\x77\xdd\xe4\x79\xe0\x7b\x02\x5f\x15\x34\x86\xd2\x5d\x4c\x0e\x36\x76\x3f\x65\xb0\x58\xc3\x70\x08\x33\xf6\x25\x4e\x5a\xb1\xde\x9e\xa7\x0d\x0b\xb8\xac\xe5\x67\x81\xdf\xa6\xc3\x17\xed\x80\x25\x0f\xab\x9b\xb7\x40\xff\x8d\x5e\x51\x7a\xa9\x0f\xa3\xf1\xcf\x29\x3a\x13\x2f\x7f\x95\x74\x25\xe7\x44\x82\xc4\x47\xc9\x4e\x34\xb3\xe4\x6d\xf5\x11\x60\x8f\x52\x7d\x4c\xc5\xdc\x58\x4f\x62\x74\x7b\xd7\xe6\xd8\xe3\x9c\x6d\xd2\x78\x9d\xbc\x21\xff\xf9\xca\x12\x10\x65\xa1\x24\x0f\x63\x9a\xfc\x99\xb8\x87\x34\x44\x75\xa9\xd4\xdd\x80\x25\x0b\x18\x0e\x2f\x76\x2e\xd0\xa3\x23\xdf\x39\x07\x68\x09\xb4\xf1\xc0\x2d\xa1\x27\x01\x42\x5a\xe2\x5e\x55\xc1\xe1\x93\x3c\x74\x1b\xe1\xb8\x95\x85\x77\x63\xf0\x47\x6a\x29\x3d\xfc\x45\x62\xb7\xe8\x23\xd7\x05\x17\x3b\xd8\x97\x9a\xb7\xbc\xb0\xfc\x84\x62\x1a\x0a\xa7\x47\x82\x7f\x4b\xb2\xd5\xe5\x0f\x4a\xfd\x03\x7f\x2a\x9d\x07\x54\xaf\x58\xf5\x22\x6d\x21\xf1\xbe\xd5\xd2\x49\x7d\x68\x13\xe1\x36\x9c\x07\xc7\x8f\x74\xc2\xfe\x54\xdb\x22\xdf\xba\xb3\x2d\x0b\x11\x30\xed\x6d\xd1\x59\xba\xe0\xfd\xb5\x1e\x47\x0d\x39\x81\xa0\xbd\xd4\x24\xc6\xbd\xfe\xbb\xab\xfe\x4e\x15\x8c\x3a\xcb\xdc\x08\xbc\x01\xe9\xc3\xc2\xd3\xcf\xaa\x52\x0b\xc9\xc9\x75\xb3\x4a\x07\xd2\x01\x6a\xa0\x33\x9e\x0a\x45\xe1\x21\xc1\x37\xec\xb5\xb3\x7d\xf1\x1f\xf1\xbe\x0d\xf7\xe7\x60\x32\x09\x92\x75\x6d\x51\x1f\x08\xa6\x7d\x37\x4d\x13\xc2\xed\x2a\x2f\x53\xb6\xed\xde\x82\xba\x1e\xfe\xff\xfc\x73\x97\xd5\xf5\xb4\x69\x20\x5a\x2c\x60\xbe\x5e\x3e\xac\x12\xd0\xf4\x9a\x71\xa3\xca\x93\x86\x94\x3d\xa6\x77\x9f\x91\x69\x2f\xa8\xfc\xa4\x58\x5d\x93\x16\x4d\xf3\x5f\x00\x00\x00\xff\xff\x7d\x20\xde\x2b\x4e\x05\x00\x00"),
		},
		"/jobsdb/000002_alter_dataset_tables.down.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000002_alter_dataset_tables.down.tmpl",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 114920869, time.UTC),
			uncompressedSize: 802,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x91\x41\x6a\xf3\x30\x10\x85\xf7\x39\xc5\x2c\x0c\xf9\x7f\x68\x72\x01\xaf\x9c\x58\x6d\x03\x8e\x6c\x52\x99\xb6\x2b\x21\x5b\x93\xa0\x60\xec\x20\xc9\x34\x41\xe8\xee\xc5\x76\x1a\x48\x71\x42\x69\x67\x61\xf0\xcc\x9b\x6f\xc4\x7b\xce\x69\x51\xef\x10\xe6\xb1\xb0\xc2\xa0\x35\xde\x4f\x00\x00\xa2\x84\x91\x0d\xb0\x68\x91\x10\x70\x2e\x98\x67\x1a\xb7\xea\xe8\x3d\xdf\x37\x85\xe1\xce\xcd\xbd\x87\x78\x93\x66\xb0\x4c\x93\x7c\x4d\xa1\x35\xa8\xb9\x92\xe1\x8f\x97\x07\xcd\x79\xbb\xd4\x28\x2c\x4a\x2e\x2c\x18\xb4\x10\x93\xc7\x28\x4f\x18\xd0\x3c\x49\x7e\x49\xc4\xe3\x41\x69\x1c\x07\xf6\xc4\xd9\x0c\x34\x0e\x77\x61\xdf\x14\xdc\x58\x61\x91\xdb\xd3\x01\xa1\xfb\xf4\x9a\x38\x85\x20\x80\x05\x79\x5a\xd1\xfe\xbf\xab\xe5\x86\x44\x8c\x00\x7b\xcf\xc8\xb7\xbd\x8b\xa4\x7f\xf0\x0b\x10\x9a\xaf\xff\x5d\x35\xbf\x6a\xfa\x21\x94\x55\xf5\x6e\xfa\x30\x3e\xc6\x23\x96\xed\x3d\x81\x69\xcb\x12\x51\xa2\xbc\x25\x38\x1f\xe0\x1a\xad\x3e\xdd\x12\x6d\x85\xaa\x6e\x23\x44\xd1\x68\x8b\x72\xfa\x3f\xbc\x9a\x93\xb7\x25\xc9\xd8\x2a\xa5\x57\xdd\xd7\x67\x42\x41\xb6\x87\x4a\x95\x9d\x1d\x4d\xb1\xc7\xd2\x02\xeb\xba\x75\x5b\x55\x03\x82\xd0\x18\x82\xe0\xec\xff\xbd\x44\x7b\x53\xdb\xd1\x5c\x2f\x96\x8f\x25\x10\xfe\x95\xac\xe4\x80\x5d\x51\x16\x4e\x9c\xc3\x5a\x7a\xff\x19\x00\x00\xff\xff\xea\x0e\x28\x43\x22\x03\x00\x00"),
		},
		"/jobsdb/000002_alter_dataset_tables.up.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000002_alter_dataset_tables.up.tmpl",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 115142869, time.UTC),
			uncompressedSize: 718,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x91\x51\x6b\x9c\x40\x14\x85\xdf\xfd\x15\xe7\x41\x48\x0a\xd9\x85\x42\xe9\x8b\x4f\xee\x3a\xc9\x0a\x66\x14\x77\xb6\xd9\x3e\xc9\xb8\xde\x46\x83\xa8\xcc\x8c\x69\xc2\x30\xff\xbd\x8c\xdb\x6e\x28\x14\x52\xda\x37\xef\xf5\xdc\xef\xde\x73\xc6\x5a\x25\x87\x47\xc2\x3a\x91\x46\x6a\x32\xda\xb9\x00\x00\xe2\x4c\xb0\x12\x22\xde\x64\x0c\xd6\x86\xeb\x42\xd1\xb7\xee\xc5\xb9\xea\x69\xac\x75\x65\xed\xda\xb9\x9f\x9a\x6d\x9e\x1d\xee\x39\x4e\x8a\xa4\xa1\xa6\x92\x06\x9a\x0c\x12\x76\x1b\x1f\x32\x01\x9e\x3f\x5c\x7f\x88\xfe\x0d\x49\x2f\x53\xa7\xe8\x7f\x89\x49\xf2\x8b\x97\xde\x82\xe7\x02\xec\x98\xee\xc5\x1e\xb3\x26\x55\x75\x0d\x04\x3b\x8a\xa5\xcf\x0f\x59\x76\xd9\x72\xb5\xfa\x78\x15\x05\xcb\x96\xd5\x0a\xdb\x71\x78\x26\x65\xf0\x34\xd6\x95\x36\xd2\x10\xcc\x88\x67\xa9\x4e\xad\x54\x37\x68\xd4\x38\xc1\xbc\x4e\xf4\xf6\xbf\xf2\xe5\xbb\x37\x2e\xda\xf9\x8f\xde\xdf\x36\x89\xaf\x05\xc3\x97\xb8\xdc\xee\xe2\xf2\xfa\xf3\xa7\xbf\xb0\xfe\x0e\xd6\x7b\xf6\xcc\x4d\x7a\x97\x72\x11\x05\x17\x97\xf7\xdd\x63\x6b\xa0\x4d\xd7\xf7\xa8\xc9\x07\xd4\xa0\x7e\xc5\x68\x5a\x52\x7e\x52\x37\x35\xba\x41\x1b\x39\x9c\xe8\x06\xbd\xd4\x06\xe3\x40\x98\xa7\xc6\xbf\x3c\xbe\xfb\xb9\x73\x16\x2d\x2d\x79\xac\x17\x72\x92\x23\x0c\xb1\x61\x77\x29\x5f\xea\xa5\x57\xe6\xc5\xf9\x8a\xdf\x23\x8b\x2e\x0a\x76\xdc\xb2\x42\xa4\x39\xc7\xc3\x8e\x71\xe4\x62\xc7\xca\x3d\x84\xff\x1e\xe6\xbe\x3f\x0b\x19\x4f\x10\x86\x51\x60\x2d\x0d\x8d\x73\xc1\x8f\x00\x00\x00\xff\xff\x3e\x89\xb1\x86\xce\x02\x00\x00"),
		},
		"/jobsdb/000003_alter_journal_table.down.tmpl": &vfsgen۰FileInfo{
			name:    "000003_alter_journal_table.down.tmpl",
			modTime: time.Date(2021, 2, 14, 4, 34, 29, 115357015, time.UTC),
			content: []byte("\x2d\x2d\x20\x44\x72\x6f\x70\x20\x6a\x6f\x75\x72\x6e\x61\x6c\x20\x74\x61\x62\x6c\x65\x0a\x41\x4c\x54\x45\x52\x20\x54\x41\x42\x4c\x45\x20\x7b\x7b\x24\x2e\x50\x72\x65\x66\x69\x78\x7d\x7d\x5f\x6a\x6f\x75\x72\x6e\x61\x6c\x20\x44\x52\x4f\x50\x20\x43\x4f\x4c\x55\x4d\x4e\x20\x6f\x77\x6e\x65\x72\x3b"),
		},
		"/jobsdb/000003_alter_journal_table.up.tmpl": &vfsgen۰FileInfo{
			name:    "000003_alter_journal_table.up.tmpl",
			modTime: time.Date(2021, 2, 14, 4, 34, 29, 115568557, time.UTC),
			content: []byte("\x2d\x2d\x20\x41\x64\x64\x20\x6f\x77\x6e\x65\x72\x20\x63\x6f\x6c\x75\x6d\x6e\x20\x74\x6f\x20\x4a\x6f\x75\x72\x6e\x61\x6c\x20\x74\x61\x62\x6c\x65\x0a\x41\x4c\x54\x45\x52\x20\x54\x41\x42\x4c\x45\x20\x7b\x7b\x24\x2e\x50\x72\x65\x66\x69\x78\x7d\x7d\x5f\x6a\x6f\x75\x72\x6e\x61\x6c\x20\x41\x44\x44\x20\x43\x4f\x4c\x55\x4d\x4e\x20\x49\x46\x20\x4e\x4f\x54\x20\x45\x58\x49\x53\x54\x53\x20\x6f\x77\x6e\x65\x72\x20\x54\x45\x58\x54\x20\x4e\x4f\x54\x20\x4e\x55\x4c\x4c\x20\x44\x45\x46\x41\x55\x4c\x54\x20\x27\x27\x3b"),
		},
		"/node": &vfsgen۰DirInfo{
			name:    "node",
			modTime: time.Date(2021, 2, 14, 4, 34, 29, 116769480, time.UTC),
		},
		"/node/000001_create_event_schema.down.sql": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_event_schema.down.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 115852631, time.UTC),
			uncompressedSize: 279,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd2\xd5\xd5\xe5\xd2\xd5\xd5\x55\x70\x2d\x4b\xcd\x2b\x51\xf0\xcd\x4f\x49\xcd\x29\x06\x09\x70\x71\xb9\x04\xf9\x07\x28\x84\x38\x3a\xf9\xb8\x2a\x78\xba\x29\xb8\x46\x78\x06\x87\x04\x2b\xa4\x82\x94\xc5\xe7\x82\x95\x59\x43\xd5\x78\xfa\xb9\xb8\x46\x60\x57\x13\x5f\x5e\x94\x59\x92\x1a\x9f\x9d\x5a\x19\x9f\x99\x97\x92\x5a\x61\xcd\xc5\x05\xb3\x50\x21\x38\x39\x23\x35\x37\x51\x21\x2c\xb5\xa8\x38\x33\x3f\x0f\x9f\xa5\xc5\x60\x95\xf1\x65\x50\x95\xc4\xd8\x9b\x99\x02\xb3\x90\x18\xa5\x50\x0b\x32\x12\x8b\x33\xa0\xda\x00\x01\x00\x00\xff\xff\xcc\x7f\x3b\xaa\x17\x01\x00\x00"),
		},
		"/node/000001_create_event_schema.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_event_schema.up.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 116053446, time.UTC),
			uncompressedSize: 945,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x91\x4d\x6f\xe2\x30\x10\x86\xcf\xc9\xaf\x98\x1b\x44\x22\x97\x5d\x69\x0f\x70\x32\x60\x76\xbd\xcd\x07\x4d\x1c\x0a\x27\x2b\xc5\x83\x70\x0b\xa1\x4a\x0c\x2d\xaa\xfa\xdf\x2b\x87\xcf\x40\x11\x70\xcd\x3c\xf3\x4e\xfc\xbc\xae\xeb\xda\xae\xeb\x02\x5d\x61\xa6\xc1\x5f\x48\x9c\x15\xe6\x83\x6d\x77\x22\x4a\x38\x05\x4e\xda\x1e\x05\xd6\x83\x20\xe4\x40\x87\x2c\xe6\x31\xa0\x81\xc5\xbc\x84\xa1\x6e\x5b\x96\x92\x10\xd3\x88\x11\x0f\xfa\x11\xf3\x49\x34\x82\x07\x3a\x6a\xd8\x96\xb5\x5c\x2a\x09\x03\x12\x75\xfe\x91\xa8\xfe\xfb\x8f\x53\xa6\x04\x89\xe7\x99\xe1\x7b\xae\x34\x8a\x57\x5c\x1f\x88\x5f\x55\x62\x73\x48\xaf\xdf\x10\x38\x1d\xf2\x1f\x66\xe5\x4f\x08\x25\x31\xd3\x6a\xa2\x30\xaf\x72\xd0\xa5\x3d\x92\x78\x1c\x6a\x35\xb3\x32\xce\x31\xd5\x28\x45\xaa\x81\x33\x9f\xc6\x9c\xf8\xfd\x73\x36\x08\x9f\xea\x8e\xd3\xda\x1b\x60\x41\x97\x0e\x2f\x1b\x10\xfb\x67\x08\x95\x49\xfc\x80\x30\x38\x11\xb4\x07\x4c\xe8\x4e\x38\xc4\xe3\x29\xce\x53\x18\x60\x5e\xa8\x45\x76\x5d\x7a\x51\xf2\x62\xb5\xe5\x77\xde\xdb\xec\xef\xbd\xea\xe1\xcc\xde\xc5\x8a\xb6\x47\xa7\x69\x31\xbd\x58\xd2\x86\x81\xff\x71\x18\xb4\x2b\x83\x39\xea\x54\xa6\xfa\x74\x74\x28\xe5\xf3\xab\xd6\x6c\xbe\x14\x8b\xec\xd9\xe0\x13\x95\x17\x5a\x14\x88\xd9\xd5\x76\x0c\x3e\x4b\x6f\xa5\xef\xe8\x52\xc9\x43\x89\x67\xc2\xab\xa0\xd3\xda\x85\x26\x01\x7b\x4c\x6e\xca\x3e\xd2\x79\xf3\x9d\x06\x1c\x6d\x39\xad\xef\x00\x00\x00\xff\xff\x0d\x70\xf8\x26\xb1\x03\x00\x00"),
		},
		"/node/000002_create_col_counters_event_schema.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000002_create_col_counters_event_schema.up.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 116259960, time.UTC),
			uncompressedSize: 309,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x90\x4f\x4b\x03\x31\x10\x47\xcf\xcd\xa7\xf8\xdd\xaa\x87\x80\xe7\xf6\x94\xed\xa6\x12\x89\x59\xec\x26\xe0\x6d\x89\xdd\x80\x91\x36\xd1\xfc\xe9\x45\xfc\xee\x92\x05\x3d\x89\xf4\x38\xcc\xf0\xde\x63\x28\xa5\x84\x52\x0a\x36\xcf\xd8\xc5\x53\x3d\x07\x94\x88\x5c\x62\x72\x38\xc6\x1a\x8a\x4b\x19\xb3\x2d\xb6\x5d\x11\xc2\xa4\xe6\x07\x68\xd6\x49\x8e\x7c\x7c\x75\x67\x3b\x5d\x5c\xca\x3e\x86\x0c\xb2\x5a\xb1\xbe\xc7\x6e\x90\xe6\x51\x41\xec\xa1\x06\x0d\xfe\x2c\x46\x3d\xe2\x3d\xf9\x8b\x2d\x6e\x6a\x24\x3c\x8c\x83\xea\x96\xad\x32\x52\xa2\xe7\x7b\x66\xa4\xc6\xfa\xf3\x6b\xbd\xd9\xbc\xe5\x18\x5e\xb6\xff\x89\x7e\x3d\x6a\xd4\x07\x26\x94\x46\x0d\xfe\xa3\xba\xa9\x56\x3f\xc3\x28\xf1\x64\x38\x6e\xda\x70\x7b\x1d\xe7\x8f\xde\x12\x8b\x3d\x4d\xcb\x03\xd0\x89\xfb\x26\xf9\xc9\xbc\xdb\x92\xef\x00\x00\x00\xff\xff\x22\x3a\x04\x17\x35\x01\x00\x00"),
		},
		"/node/000002_del_col_counters_event_schema.down.sql": &vfsgen۰CompressedFileInfo{
			name:             "000002_del_col_counters_event_schema.down.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 116450338, time.UTC),
			uncompressedSize: 230,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x8f\x41\x0a\x83\x30\x10\x45\xf7\x9e\xe2\x5f\x20\x27\x70\x95\x5a\x0b\x82\xd5\xa2\x29\x74\x17\x82\x09\x34\xa0\x99\x36\x99\xf1\xfc\xc5\x9e\xa0\x74\xff\xff\x7b\x3c\xa5\x54\xa5\x94\x82\xf6\x1e\x0d\xad\xb2\x25\x30\xa1\x30\xe5\x80\x85\x24\x71\xc8\x05\xde\xb1\x3b\x56\x55\xa5\x7b\xd3\x4e\x30\xfa\xd4\xb7\x28\xcb\x33\x6c\xce\xee\x21\x97\x48\xa9\xe0\x3c\x8d\x37\x34\x63\x7f\xbf\x0e\xe8\x2e\x68\x1f\xdd\x6c\x66\xbc\x72\xdc\x1d\x07\x7b\x30\xea\x5f\xfe\xc3\x6c\x26\xdd\x0d\x06\x92\xe2\x5b\x82\x15\x89\xbe\xfe\x43\xcc\xc4\x6e\xb5\xdf\x86\xfa\x13\x00\x00\xff\xff\x73\xa7\x01\x8c\xe6\x00\x00\x00"),
		},
		"/node/000003_add_event_model_columns.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000003_add_event_model_columns.up.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 116643787, time.UTC),
			uncompressedSize: 553,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\xd0\xc1\x4a\xc3\x40\x10\x80\xe1\x73\xf3\x14\x73\xab\x1e\x16\x3c\xb7\xa7\x8d\x49\x25\xb2\xd9\x88\xd9\xa0\xb7\x30\x66\x07\x5d\xc9\xce\x4a\x76\xda\x8b\xf8\xee\x62\x41\x41\x7a\x2b\xf4\x3a\x03\xdf\x3f\x8c\x52\xaa\x50\x4a\x81\xf6\x3e\xf0\x2b\x4c\x69\xde\x47\xce\x20\x09\x22\x06\x16\x0c\x0c\x11\xb3\xd0\x02\x79\x7a\xa3\x88\x80\xec\x61\x4a\x7b\x16\x5a\x32\x04\x06\x3a\x10\xcb\x18\x93\xa7\x39\xff\x48\x45\xa1\x8d\xab\x1f\xc1\xe9\xd2\xd4\xff\xb7\xab\x95\xae\x2a\xb8\xed\xcc\xd0\x5a\x68\x76\x60\x3b\x07\xf5\x73\xd3\xbb\xfe\x17\xbf\xef\x3b\x5b\x1e\xe7\x76\x30\x06\xaa\x7a\xa7\x07\xe3\x60\xfd\xf9\xb5\xde\x6c\xde\x73\xe2\x97\xed\x59\x7c\x24\x41\x8f\x72\xb9\xc0\xc7\x12\x0e\x28\x34\x5e\x34\x22\x49\x70\x1e\x8f\xcf\x87\xb2\xb9\x6b\xac\xfb\xb3\x6f\xce\x13\x67\xcc\x32\x66\x22\x06\xd7\xb4\x75\xef\x74\xfb\x70\x7a\xb7\xed\x9e\xae\xae\xb7\xdf\x01\x00\x00\xff\xff\x31\x24\x19\x95\x29\x02\x00\x00"),
		},
		"/node/000003_remove_event_model_columns.down.sql": &vfsgen۰CompressedFileInfo{
			name:             "000003_remove_event_model_columns.down.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 116840334, time.UTC),
			uncompressedSize: 390,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\xcf\x41\x4a\xc6\x30\x10\xc5\xf1\x7d\x4e\xf1\x2e\x90\x13\x7c\xab\xaa\x15\x0a\xd5\x4a\x5b\xc1\x5d\x18\x9a\xc1\x06\x92\x89\x64\xa6\x3d\xbf\x28\x6e\x5c\xf6\x5b\x3f\x7e\x7f\x78\xde\x7b\xe7\xbd\xc7\xcc\xa5\x9e\x49\x3e\xb1\xd5\x7c\x14\x51\xd8\x4e\x06\x6a\x8c\x43\x39\xc2\x2a\x0a\x25\x31\x4a\x82\x42\x6a\xdc\xa0\xdb\xce\x85\x40\x12\xb1\xd5\x43\x8c\x9b\x22\x09\xf8\x64\xb1\x50\x6a\xe4\xac\x3f\x65\xe7\xba\x71\xed\x67\xac\xdd\xc3\xd8\xff\x5b\xf1\x34\x4f\x6f\x78\x9c\xc6\xf7\x97\x57\x0c\xcf\xe8\x3f\x86\x65\x5d\xfe\xba\xb7\xab\xac\xb0\x51\x24\xbb\x0e\xbf\x5a\x3a\xc9\x38\xdc\x85\xad\x1a\xe5\xf0\xfb\xff\xb2\xcd\xa4\x16\x94\x59\x6e\xee\x3b\x00\x00\xff\xff\x23\x26\xf4\xfe\x86\x01\x00\x00"),
		},
		"/warehouse": &vfsgen۰DirInfo{
			name:    "warehouse",
			modTime: time.Date(2021, 3, 24, 9, 2, 55, 656812442, time.UTC),
		},
		"/warehouse/000001_create_tables.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_tables.up.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 117132765, time.UTC),
			uncompressedSize: 4017,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd4\x57\x4f\x73\x9b\x3e\x10\x3d\xff\xf8\x14\x7b\xf0\x21\x9e\x89\x6f\xbf\xe9\x85\x13\xb6\xd5\x84\x16\x83\x07\x94\xd6\x39\x69\x54\x50\x12\xcd\x10\xf0\x80\xdc\xa6\xdf\xbe\x03\x22\x20\x30\xc8\x4a\xe2\x64\xda\x9b\xc7\xbb\xfb\xd8\x3f\x6f\xf5\x24\x6b\xb1\xb0\x16\x0b\xf8\xf5\x40\x4a\x41\xef\x79\x76\x4f\xee\x78\xca\xca\xea\x6f\x6b\x15\x22\x07\x23\xc0\xce\xd2\x43\xe0\x7e\x06\x3f\xc0\x80\x76\x6e\x84\xa3\x23\x7f\xb8\xb0\x00\x00\x78\x02\x4b\xf7\x2a\x42\xa1\xeb\x78\xb0\x0d\xdd\x8d\x13\xde\xc2\x57\x74\x7b\x59\x5b\xd3\x3c\xa6\x82\xe7\x19\x60\xb4\xc3\x35\x9a\x7f\xe3\x79\xd2\x56\xe6\x87\x22\x66\x84\x27\xf0\xcd\x09\x57\xd7\x4e\x78\xf1\xe9\xff\xf9\xc0\x27\x61\xa5\xe0\x59\x0d\xa1\x77\x2c\xe3\x07\xf6\x48\xe1\x4b\x14\xf8\xcb\x81\x89\x15\x45\x5e\xd4\x09\x34\xae\x82\x8a\x43\xa9\x62\xc9\xff\xef\x78\x51\x0a\xc2\x7e\xb2\x4c\x10\x2a\x00\xbb\x1b\x14\x61\x67\xb3\x6d\x2a\xa1\x1a\xa3\xc8\x05\x4d\xa5\xb5\xac\xda\xe1\xfa\xcd\xb7\xe2\x82\x51\xc1\x92\x5e\xc8\x20\xbd\xc3\x3e\x99\x76\x99\xdb\x96\xe5\x78\x18\x85\xcd\x48\x8e\x86\x56\x21\x38\xeb\x35\xac\x02\xef\x66\xe3\x0f\x46\xa6\xaf\x68\x32\x4c\x5b\xea\x64\xd4\x48\x0f\xec\x96\x50\xae\xbf\x46\xbb\x13\x84\x22\x3c\x21\x3c\x4b\xd8\x13\x04\xfe\x08\xdb\x5a\xba\x5c\x0e\x58\x71\xaa\x47\x20\x8d\x4d\xd2\xcd\xf4\xf1\xed\x16\xa9\x14\xb0\x2d\x6b\x1d\xc0\x6c\x06\x4b\x74\xe5\xfa\x75\xa5\xeb\x30\xd8\x4a\x3f\x05\xb1\x0a\x67\x44\xfc\xde\x33\xbb\x76\x42\xbb\x15\xda\x62\x37\xf0\xe1\xfb\x35\xf2\x21\xc0\xd7\x28\x8c\x00\x57\xbf\xb3\x43\x9a\xda\x16\xf2\xd7\x30\x9b\xd9\x56\xb7\x75\x69\x4e\x13\xe3\x95\xeb\x9c\x8d\xf6\x4d\x2d\x9c\x48\xd7\x96\x8b\x1f\xb6\x8b\xaa\x63\xd5\x29\x8d\xab\xa0\x3f\x52\x46\x32\xfa\xc8\xc6\xb2\x7a\xdd\x5a\x8d\xf0\x41\xe9\x62\x8f\x0c\xea\xe7\xab\x41\x57\x39\xe8\xa3\xcf\xc9\xfe\x0e\x97\x34\xfd\xef\xb7\x98\x74\xe9\xf5\x16\x43\xe5\xc4\xe4\x56\x5c\x2a\xc5\xcd\x55\xfa\x1d\xf6\x55\xbc\x09\xf7\x1a\x4f\x33\xe2\x19\x10\xa8\x4a\xa5\xdc\xd3\x98\x7d\x30\xc9\x4a\x41\x0b\x41\x74\xab\xc1\xb2\x44\x6b\x97\x08\x6d\xdf\xc7\xc2\x27\x8d\xc7\x72\xf3\x52\xe9\xaa\x2d\x6f\xd6\x28\x69\x7c\x62\xf1\x98\x7e\xf1\x47\x9e\xdd\x97\xea\x97\xde\x41\xb9\x9e\x99\xf7\xf7\x68\x96\xae\x27\xd3\xab\xae\x36\xeb\xe4\x92\x37\x45\x13\x49\x83\xde\x1e\xb7\xfb\x25\x6d\x73\x73\xac\xd1\xd3\x62\x1c\xfa\x05\xaa\xf9\x1c\x74\x0e\xbd\x94\x58\xe7\x90\x4b\x79\x8c\x19\x9e\x5a\xff\x0d\x03\x8c\x0e\xaf\x2e\xdf\x9e\x9f\x56\xaa\x0c\x97\x7b\x78\xf9\xec\x18\x27\x78\x85\xa4\xb9\x47\x76\x41\xef\xb0\x8b\xfd\x26\xbd\x50\xd6\xf4\x58\x2a\x7b\x9a\xf2\x15\x6d\xd5\x33\xbc\x07\x45\xd4\xb9\x4c\x89\xe1\x60\xda\x6a\xc8\x50\x02\xcd\xd3\x7e\x25\xe9\x55\xc4\x73\x50\x5f\x2a\x83\xd1\xcb\x4c\x7a\xbe\x81\xed\xff\x80\x8c\x9b\x3f\xf1\x4e\xde\x0f\xeb\xb1\xb5\x1c\x3c\xea\xa3\xee\x84\x3d\xfd\xa8\x69\x30\x06\xc1\x6d\xdb\xfa\x8f\x9c\xe7\xc1\x0d\x6f\x6f\xad\xfb\xdc\xfe\x13\x00\x00\xff\xff\x89\x94\x7c\x27\xb1\x0f\x00\x00"),
		},
		"/warehouse/000002_alter_wh_table_uploads.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000002_alter_wh_table_uploads.up.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 117314100, time.UTC),
			uncompressedSize: 201,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\xcc\x31\x0a\xc2\x30\x18\xc5\xf1\x3d\xa7\x78\xa3\x0e\x39\x41\xa7\xb4\x8d\x25\x10\x53\x31\x5f\xa1\x4e\x21\xd6\x80\x42\x4c\x84\xa6\x7a\x7d\x51\x5c\xc4\xc1\xf5\xbd\x1f\x7f\xc6\x39\xe3\x1c\x8f\xb3\x2b\xfe\x18\x83\x5b\x6e\x31\xfb\xd3\xfc\x9a\x99\xd0\x24\xf7\x20\x51\x6b\xf9\x03\x20\xda\x16\x4d\xaf\x87\xad\x81\xda\xc0\xf4\x04\x39\x2a\x4b\x16\x31\x4f\xbe\x5c\x72\x02\xc9\x91\xaa\x7f\x95\xf7\x39\xe5\xb8\x5c\x13\x4a\x2e\x3e\xba\x70\x0f\xa9\xcc\xa0\xc3\x4e\xa2\x56\x9d\x32\x84\xc1\x2a\xd3\xa1\x11\x96\x56\x5f\x46\xd8\x8f\x58\x57\xec\x19\x00\x00\xff\xff\x10\xd9\xe2\xb0\xc9\x00\x00\x00"),
		},
		"/warehouse/000003_wh_uploads_add_metadata_column.up.sql": &vfsgen۰FileInfo{
			name:    "000003_wh_uploads_add_metadata_column.up.sql",
			modTime: time.Date(2021, 2, 14, 4, 34, 29, 117486875, time.UTC),
			content: []byte("\x0a\x2d\x2d\x0a\x2d\x2d\x20\x77\x68\x5f\x75\x70\x6c\x6f\x61\x64\x73\x0a\x2d\x2d\x0a\x0a\x41\x4c\x54\x45\x52\x20\x54\x41\x42\x4c\x45\x20\x77\x68\x5f\x75\x70\x6c\x6f\x61\x64\x73\x20\x41\x44\x44\x20\x43\x4f\x4c\x55\x4d\x4e\x20\x49\x46\x20\x4e\x4f\x54\x20\x45\x58\x49\x53\x54\x53\x20\x6d\x65\x74\x61\x64\x61\x74\x61\x20\x4a\x53\x4f\x4e\x42\x20\x44\x45\x46\x41\x55\x4c\x54\x20\x27\x7b\x7d\x27\x3b\x0a"),
		},
		"/warehouse/000004_add_upd_col_to_wh_schemas_and_uniq_constraint.sql.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000004_add_upd_col_to_wh_schemas_and_uniq_constraint.sql.up.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 117702531, time.UTC),
			uncompressedSize: 460,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\x90\x41\x6b\xdc\x40\x0c\x85\xef\xfe\x15\x0f\x93\x43\x02\x75\xfb\x03\x72\x72\x62\xb7\x18\xb6\x76\x9b\xf5\x42\x6e\x46\xcc\x28\x9d\x01\x5b\xe3\xce\x68\x70\xf3\xef\xcb\x24\x81\xee\xa9\x3a\x49\x7a\xe2\xbd\x0f\x35\x4d\xd5\x34\x38\xdc\x92\x8c\xe3\x8d\x52\xd5\x34\x55\xd5\x9e\xe6\xfe\x09\x73\xfb\x70\xea\xaf\x24\xb4\x5d\x87\xc7\xe9\x74\xf9\x3e\x62\xf8\x8a\x71\x9a\xd1\x3f\x0f\xe7\xf9\x8c\xbc\x5b\x52\xb6\x0b\x29\xd4\x6f\x9c\x94\xb6\xfd\xbe\x2a\xc6\x64\x6d\x02\xc1\x04\x49\x1a\xc9\x8b\x42\xc3\xb5\xe5\xe1\xbc\x71\xa0\x75\x0d\x47\x42\x4e\x45\xb5\x01\x79\x4f\x1c\xf5\xf6\xdd\xf6\x8b\x97\x32\xdd\x21\x08\x24\x08\xf6\xe8\x37\x8a\xaf\x30\x61\xcd\x9b\x14\x60\x98\xc8\xa4\x5c\x82\x34\x92\x24\x32\xea\x83\x7c\x86\x7f\x81\x3a\xbe\x0e\xa7\x35\x32\xd9\x57\xf0\x1f\x9f\x34\x15\x55\x60\x03\x27\x48\x50\xe7\xe5\x57\xd5\x4d\xb8\xb9\xa9\x00\xe0\xa1\xff\x36\x8c\x6f\x5d\xa9\xff\x7e\x64\x3c\xcf\x4f\xed\x30\xce\xc8\xe2\x7f\x67\x5e\x0e\xb7\x78\xcb\xa2\xfe\xc5\x73\xc4\x65\x1c\x7e\x5e\xfa\xdb\x3a\x85\x1c\x0d\x2f\xde\xd6\x9f\x6a\xcb\x49\xbd\x50\xc1\x7c\x5f\x08\x6d\x9c\x76\x32\x5c\xdf\xdd\xbf\x85\xf6\xcf\x8f\xfd\x8f\x79\x98\xfe\x21\x1c\x05\x36\xa8\xe3\xf8\x01\x2e\x79\x5d\x3f\x8e\xc7\xae\x60\xff\x0d\x00\x00\xff\xff\x88\x0e\xff\xc3\xcc\x01\x00\x00"),
		},
		"/warehouse/000005_create_table_schema_versions.sql.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000005_create_table_schema_versions.sql.up.sql",
			modTime:          time.Date(2021, 2, 14, 4, 34, 29, 117911683, time.UTC),
			uncompressedSize: 1709,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xac\x94\xc1\x72\x9b\x3c\x14\x85\xf7\x7a\x8a\x3b\x8c\x17\xf6\x0c\xe4\xdf\xfc\xd3\x8d\x57\x32\x96\x09\x2d\x11\x54\x88\xd4\x59\x79\x54\xa3\x18\xcd\x38\x22\x23\xe1\x24\x7e\xfb\x8e\xc0\x06\xa7\x75\xda\x3a\xa9\x37\x06\x74\xf5\x9d\xa3\x7b\x0f\x04\x01\x0a\x02\x78\xae\x56\x76\x5d\xc9\x07\xb1\x7a\x92\xc6\xaa\x5a\x5b\x14\x04\x08\x85\x8c\x60\x4e\x80\xe3\x59\x42\x20\x5e\x00\x4d\x39\x90\x65\x9c\xf3\xfc\xcc\x0e\x18\x23\xf8\xeb\x9f\x2a\x61\x16\x47\x39\x61\x31\x4e\x20\x63\xf1\x0d\x66\x77\xf0\x85\xdc\xf9\x17\x30\x06\x0b\xaf\x68\xce\x24\x2d\x92\xe4\x12\x94\xad\x77\x66\x2d\x1d\xe7\x16\xb3\xf0\x1a\xb3\xf1\xa7\xff\x27\xef\x22\x69\xf1\x20\xed\xa3\x58\xcb\x0f\x93\x4a\x69\x1b\xa5\x45\xa3\x6a\xfd\x2f\x8c\x9d\xe2\x9a\xfd\xe3\xc7\xfd\x75\xbd\x87\xcf\x79\x4a\x67\xef\x02\x48\x63\x6a\x03\x9c\x2c\xf9\x25\xbb\xd6\x95\xd0\x1b\x59\xae\x44\x03\x3c\xbe\x21\x39\xc7\x37\x59\x2f\x3f\x99\xf6\xa1\x8d\xe9\x9c\x2c\xff\x18\xda\xd5\x69\x88\x56\x4a\x97\xf2\x05\x52\x7a\x36\xdd\xa7\x95\x4e\x27\x08\x60\x6d\xa4\x68\x24\x08\x68\x8c\xda\x6c\xa4\x81\xfb\x9d\x5e\xbb\x0e\x83\xd2\x56\x9a\xc6\x82\xd2\x4d\x7d\x8e\x76\x5f\x1b\x90\x4f\xd2\xec\x0f\x95\xff\xed\x1e\x4b\x87\xaa\xf5\x50\x6d\xa1\x11\xdf\xb7\xf2\x78\xa0\x94\x01\x23\x59\x82\x43\x02\x8b\x82\x86\x3c\x76\x3e\x85\x91\x55\xbd\xb3\xf2\x28\x50\x29\xdb\xd4\x66\x3f\x9e\xb4\x1d\x65\x84\x17\x8c\xe6\xc0\x59\x1c\x45\x84\xb5\xcf\x12\x4c\xa3\x02\x47\x04\xb2\x24\x8b\xf2\xaf\x09\xc2\x39\x1a\x8d\xd0\x8c\x44\x31\x6d\x0b\x62\x9a\x13\xc6\x21\xa6\x3c\x3d\xe3\x7c\xec\x9d\xf6\xc1\xf3\xbd\xfe\xd5\xf1\x7c\xf0\xfa\xf4\x7b\xbe\xf7\x3a\xbf\x6e\xf5\xe7\x08\xba\xdd\x2d\xc9\xf3\xbd\x36\x0c\x9e\xef\x0d\xe3\xf5\xba\x33\xdc\xe2\xa4\x20\xf9\x58\xcb\xe7\x2b\x55\xfa\xe0\xfe\x7b\xc9\xee\xb6\x17\xed\x6e\x5f\xeb\xfe\xfa\xcc\x29\x1f\x38\xad\x78\x77\xdd\xea\x77\x97\xdd\x28\x9c\x85\xc9\xf4\xa4\x8d\x6e\x6d\x8a\x08\x9d\xc3\x14\x8d\x46\x6f\x04\xe0\xb9\x92\x1a\x9a\x4a\x1a\x09\xca\x82\x38\x4c\x17\x6a\x03\xe7\x06\x3c\x20\x6c\xcb\x10\xda\x8a\x36\x3f\x57\xa0\xee\x1d\xa6\xe7\x8a\xad\x91\xa2\xdc\x83\x7c\x51\xb6\xb1\x6e\x49\x43\x59\x4b\x0b\xba\x6e\x2a\xa5\x37\x68\x9e\xc2\x68\xd4\xba\x1d\x26\xd9\xbe\x2e\x9d\xc3\x23\xe7\xcd\xc4\xac\x8e\x15\x78\xc1\x09\x3b\x86\x20\x65\x50\x64\xf3\x36\x7d\x14\x86\xc9\x5b\x0f\x16\x29\x03\x82\xc3\x6b\x60\xe9\x37\x20\x4b\x12\x16\x9c\x40\xc6\xd2\x90\xcc\x0b\x46\x7e\x13\xcc\xae\xa5\x64\x19\x92\xcc\x45\x18\x0d\x5f\x73\xa9\xa1\x76\xad\x3b\x1c\x4f\xef\xb6\xdb\x43\x31\x9d\xbb\xc3\x21\xf4\x23\x00\x00\xff\xff\xfd\x83\x8c\x84\xad\x06\x00\x00"),
		},
		"/warehouse/000006_add_wh_table_uploads_unique_constraint.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000006_add_wh_table_uploads_unique_constraint.up.sql",
			modTime:          time.Date(2021, 3, 12, 17, 10, 7, 174490900, time.UTC),
			uncompressedSize: 406,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x74\x90\xc1\x6a\xf3\x30\x10\x84\xcf\xd2\x53\xcc\xd1\x86\x10\xf8\xcf\x3f\x14\x64\x7b\x4b\x0d\x8e\xd5\xae\xe5\x86\x9e\x84\x83\x4c\x2b\x48\xec\xd4\xb5\xc8\xeb\x17\x5b\xb4\xa5\x69\x7b\x13\xa3\xd9\xfd\x76\xa6\xa0\x8a\x0c\xe1\x96\xf5\x0e\x97\x17\x3b\x77\x87\x63\x6f\xc3\xf9\x38\x76\xee\x4d\x8a\xfd\x1d\x31\xc1\x3b\xf8\x01\x89\x14\xa2\xa1\x8a\x72\x03\xef\xa4\x10\x62\x9d\x59\xd4\x0f\x7d\x7d\x0a\xd6\x7b\x5b\xb7\xbb\x8c\x38\x49\xa1\x1f\x89\x91\xdc\x2b\x36\xa5\x29\x75\x8d\xec\x69\xa1\xc4\xfd\xd6\xbb\x0d\x22\x70\xe8\x4e\x3d\x34\x17\xc4\x8b\xc3\x3b\x14\xd4\xe4\x29\x54\x83\x69\xbc\xd8\x21\x9c\x0e\xfd\xb4\x89\xeb\xe7\x6d\xa4\xaf\xf8\x28\x5d\xdf\x8d\x59\x2e\x1f\x29\x9e\xa7\x31\x9c\x7b\x77\x9d\x4a\xc4\x5c\xeb\xf0\xaf\x96\xed\x17\x15\x37\xf8\x27\x45\xfa\x5f\x4a\x55\x19\x62\x18\x95\x55\xf4\xa3\x29\xa8\xa2\x40\xae\xeb\xc6\xb0\x2a\x6b\x83\x30\xf8\xd7\xd0\x7f\xf3\xd8\xcf\xdc\x68\xeb\xf2\xa1\x25\x24\x7f\x35\x91\xca\xf7\x00\x00\x00\xff\xff\x89\x5e\x7d\x4a\x96\x01\x00\x00"),
		},
		"/warehouse/000007_add_wh_uploads_indexes.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000007_add_wh_uploads_indexes.up.sql",
			modTime:          time.Date(2021, 3, 16, 6, 28, 55, 290222259, time.UTC),
			uncompressedSize: 451,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x90\xc1\xaa\xc2\x30\x10\x45\xf7\xfd\x8a\xbb\xeb\x2b\x24\x5f\xf0\x56\xa2\x11\xba\x69\xc1\x76\xd1\x5d\x18\x9a\x80\x01\x4d\x0b\x99\xa2\x42\x3f\x5e\x8a\x58\xab\x18\x14\x57\xc3\xc0\xdc\x39\x67\x46\x4a\x9c\xf6\x7a\xe8\x0f\x1d\x99\x00\x29\x93\x64\xbd\x53\xab\x5a\x21\x2f\x36\xaa\x41\xbe\x45\x51\xd6\x50\x4d\x5e\xd5\xd5\x62\x52\x3b\xa3\x8d\x0d\xec\x3c\xb1\xeb\xfc\xd4\x7a\x3a\xda\xd0\x53\x6b\x75\x60\xe2\x21\x68\xe7\x8d\x3d\xa3\x2c\x96\x80\x3f\x67\x04\x9e\x83\x02\x73\x52\xe0\x16\xcd\xfe\xbf\xd5\x88\x3a\xbc\x85\x47\xc1\xbf\x01\xf9\xd2\xcf\xc7\x46\x45\xda\xce\xb7\xc4\xd6\x13\x5b\xf3\xd9\x6a\x5a\x79\xff\x82\x78\x15\xc6\x38\x22\xd5\xe9\x54\x1e\xe6\x59\x72\x0d\x00\x00\xff\xff\x8e\x18\xb9\x4e\xc3\x01\x00\x00"),
		},
		"/warehouse/000008_add_wh_tables_indices.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000008_add_wh_tables_indices.up.sql",
			modTime:          time.Date(2021, 3, 12, 17, 10, 7, 174911700, time.UTC),
			uncompressedSize: 563,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x94\x90\xc1\xaa\x83\x30\x10\x45\xf7\x7e\xc5\x2c\x15\x92\x2f\x70\xf5\x78\x4d\xc1\x8d\x42\x75\xe1\x6e\x48\x9b\x54\x07\x6c\x2c\x4d\xc4\x7e\x7e\x51\x4b\x35\x42\xc5\xee\x1c\x39\xb9\xf7\x70\x39\x87\xbe\xc6\xa6\x95\x0a\xaf\xd4\x68\x0b\x9c\x07\xc1\xff\x49\xfc\x15\x02\x92\xf4\x20\x4a\x48\x8e\x90\x66\x05\x88\x32\xc9\x8b\xdc\x87\xd1\x3a\x59\x91\xa9\xc6\x0b\x49\x21\x19\xa5\x9f\x90\xa5\xab\xcc\x70\xc5\x45\xf1\x0f\x1d\xa4\xd0\xb6\xdd\xe3\x32\x16\x28\x6d\x1d\x19\xe9\xa8\x35\xc3\xe9\xe4\xb9\xd1\x68\xe4\x4d\x7f\xad\x26\xc5\xe0\xf3\x9e\x81\x1f\xc0\x60\x4e\x18\xa4\xa6\x35\xa6\x7f\xdd\x7d\x88\xd9\x33\x88\xc7\x63\x5f\xbf\x3f\x47\x73\x27\x5d\x67\x3d\x39\x3f\x3d\x5c\xe2\x0c\x26\x7e\x56\x59\x0e\xb7\x47\xc5\xe3\x37\xa7\x5b\x2a\xf9\x2d\xdb\x93\x45\xf1\x2b\x00\x00\xff\xff\x42\xa7\x96\x5e\x33\x02\x00\x00"),
		},
		"/warehouse/000009_add_metadata_to_wh_staging_files.up.sql": &vfsgen۰FileInfo{
			name:    "000009_add_metadata_to_wh_staging_files.up.sql",
			modTime: time.Date(2021, 3, 24, 9, 3, 23, 170115027, time.UTC),
			content: []byte("\x0a\x2d\x2d\x0a\x2d\x2d\x20\x77\x68\x5f\x73\x74\x61\x67\x69\x6e\x67\x5f\x46\x69\x6c\x65\x73\x0a\x2d\x2d\x0a\x0a\x41\x4c\x54\x45\x52\x20\x54\x41\x42\x4c\x45\x20\x77\x68\x5f\x73\x74\x61\x67\x69\x6e\x67\x5f\x66\x69\x6c\x65\x73\x20\x41\x44\x44\x20\x43\x4f\x4c\x55\x4d\x4e\x20\x49\x46\x20\x4e\x4f\x54\x20\x45\x58\x49\x53\x54\x53\x20\x6d\x65\x74\x61\x64\x61\x74\x61\x20\x4a\x53\x4f\x4e\x42\x20\x44\x45\x46\x41\x55\x4c\x54\x20\x27\x7b\x7d\x27\x3b\x0a"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/jobsdb"].(os.FileInfo),
		fs["/node"].(os.FileInfo),
		fs["/warehouse"].(os.FileInfo),
	}
	fs["/jobsdb"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/jobsdb/000001_create_tables.down.tmpl"].(os.FileInfo),
		fs["/jobsdb/000001_create_tables.up.tmpl"].(os.FileInfo),
		fs["/jobsdb/000002_alter_dataset_tables.down.tmpl"].(os.FileInfo),
		fs["/jobsdb/000002_alter_dataset_tables.up.tmpl"].(os.FileInfo),
		fs["/jobsdb/000003_alter_journal_table.down.tmpl"].(os.FileInfo),
		fs["/jobsdb/000003_alter_journal_table.up.tmpl"].(os.FileInfo),
	}
	fs["/node"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/node/000001_create_event_schema.down.sql"].(os.FileInfo),
		fs["/node/000001_create_event_schema.up.sql"].(os.FileInfo),
		fs["/node/000002_create_col_counters_event_schema.up.sql"].(os.FileInfo),
		fs["/node/000002_del_col_counters_event_schema.down.sql"].(os.FileInfo),
		fs["/node/000003_add_event_model_columns.up.sql"].(os.FileInfo),
		fs["/node/000003_remove_event_model_columns.down.sql"].(os.FileInfo),
	}
	fs["/warehouse"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/warehouse/000001_create_tables.up.sql"].(os.FileInfo),
		fs["/warehouse/000002_alter_wh_table_uploads.up.sql"].(os.FileInfo),
		fs["/warehouse/000003_wh_uploads_add_metadata_column.up.sql"].(os.FileInfo),
		fs["/warehouse/000004_add_upd_col_to_wh_schemas_and_uniq_constraint.sql.up.sql"].(os.FileInfo),
		fs["/warehouse/000005_create_table_schema_versions.sql.up.sql"].(os.FileInfo),
		fs["/warehouse/000006_add_wh_table_uploads_unique_constraint.up.sql"].(os.FileInfo),
		fs["/warehouse/000007_add_wh_uploads_indexes.up.sql"].(os.FileInfo),
		fs["/warehouse/000008_add_wh_tables_indices.up.sql"].(os.FileInfo),
		fs["/warehouse/000009_add_metadata_to_wh_staging_files.up.sql"].(os.FileInfo),
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
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
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

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
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
