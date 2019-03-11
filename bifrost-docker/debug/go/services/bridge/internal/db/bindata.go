// Code generated by go-bindata.
// sources:
// latest.sql
// migrations/01_init.sql
// migrations/02_payment_id.sql
// migrations/03_transaction_id.sql
// migrations/04_table_names.sql
// DO NOT EDIT!

package db

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

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xdb\x6e\xe2\x48\x13\xbe\xfe\x79\x8a\xba\x23\xd1\x6f\x18\xc8\x42\x26\x03\xca\x85\x07\x1c\x0d\x5a\x30\x19\x30\x3b\x13\x69\x25\xab\xb1\x0b\xa7\x15\xbb\xed\x74\xb7\x73\xd8\xa7\x5f\xb5\x0d\xc1\x27\x0e\x49\xd8\x4b\x77\x57\xd7\xe9\xab\xfa\xaa\xa0\xd1\xa8\x35\x1a\x70\x1b\x0a\xe9\x71\x9c\xff\x1c\x83\x4b\x24\x59\x12\x81\xe0\xc6\x41\x54\x6b\x34\x6a\xea\x7e\x18\x07\x11\xba\xb0\xe2\x61\xb0\x15\x78\x42\x2e\x68\xc8\xe0\x5b\xf3\xb2\xd9\xce\x48\x2d\x5f\x21\xf2\x6c\xf5\xbc\x20\x52\x9b\x1b\x16\x08\x49\x24\x06\xc8\xa4\x2d\x69\x80\x61\x2c\xe1\x1a\x5a\xfd\xe4\xca\x0f\x9d\x87\xf2\x29\x75\x7d\xb4\x29\xb3\x25\x27\x4c\x10\x47\xd2\x90\xd9\x02\x85\xd2\x5b\x16\x76\x7c\xaa\x54\x23\x73\x42\x97\x32\x0f\xae\xa1\xbe\xb0\x6e\xae\xea\xfd\x8d\x6d\xe6\x12\xee\xda\x4e\xc8\x56\x21\x0f\x28\xf3\x6c\x21\x39\x65\x9e\x80\x6b\x08\xd9\x5a\xc7\x3d\x3a\x0f\xf6\x2a\x66\xa9\xad\x65\xe8\x52\x54\xf7\x2b\xe2\x0b\xcc\x99\x09\x28\xb3\x03\x14\x82\x78\x89\xc0\x33\xe1\x8c\x32\x2f\x15\xe1\xe1\xb3\x2d\xd0\x89\x39\x95\xaf\x4a\xf9\x6a\xd5\x57\xa9\x54\x79\x32\x49\x80\x3d\x88\xfc\xc8\x13\x8f\x7e\x1f\xac\xd7\x08\x7b\x60\xfc\xb6\x0c\x73\x3e\x9a\x9a\x7d\x98\x3b\xf7\x18\x90\x1e\x34\xfa\x30\x7d\x66\xc8\x7b\x90\xe0\x30\x98\x19\xba\x65\x6c\x05\x61\x74\x03\xe6\xd4\x02\xe3\xf7\x68\x6e\xcd\x37\xfa\xe0\xd7\xc8\xfa\x01\xf3\xc1\x0f\x63\xa2\x2b\x1c\x1c\x22\x89\x1f\x7a\xfd\x5a\xde\xfa\x56\x4b\xc1\x8f\xc1\x74\x32\x31\x4c\x6b\xb7\x17\xe9\x3d\x4c\xcd\xb2\x0e\x18\xcd\xa1\x7e\x3b\xfe\x12\x79\xaa\x92\x22\x1e\x3a\xe8\xc6\x9c\xf8\xe0\x13\xe6\xc5\xc4\xc3\xba\x72\x23\x41\x02\x09\x77\xee\xed\x88\xc8\x7b\xb8\x86\x28\x5e\xfa\xd4\xd1\xf2\xee\x2a\x31\x17\x57\x24\xf6\xa5\x2d\xc9\xd2\x47\x11\x11\x07\x15\xa2\xf5\xc2\xed\x33\x95\xf7\x76\x48\xdd\x0c\x48\xb9\x58\xbd\x90\x47\x76\x40\x3d\x4e\x14\xa0\x62\x13\xa9\xa5\x7f\x1f\x1b\xdb\x38\x53\x27\xde\x82\x5d\x12\x2e\xf1\x21\x9b\xf8\x44\xbe\xa8\x0c\xce\x6a\x00\x00\xd4\x05\x89\x2f\x32\xc1\xc3\x5c\x8c\xc7\x5a\x72\x4a\xa2\xc8\xa7\xe8\xda\x44\x82\xaa\x54\x21\x49\x10\x81\xf2\x36\xf9\x84\x7f\x42\x86\xb5\x73\x95\x12\x7d\x6c\x19\xb3\x1d\x06\xa6\xbf\x4c\x75\x37\x5d\x7b\x54\x88\x8d\xa3\x83\xf4\x09\x5d\x3b\x22\xaf\xaa\xab\x3e\x17\x5c\x51\xdb\x36\xba\x25\xf5\x28\x2b\xc6\x17\x46\x98\x7a\x69\x53\x17\x9c\x7b\xc2\x89\x23\x91\xc3\x13\xe1\xaf\x94\x79\x67\x17\xdd\xee\x79\xe1\x45\x52\x13\x42\x54\xe5\x44\xf5\xf1\x5b\x5a\x8a\xcf\x88\xa7\x7a\x55\x86\x0f\xc8\x8e\x33\xa4\x68\x26\x16\xc7\xc9\x66\x99\xa5\x32\x90\xcb\xce\x39\x0c\x8d\x1b\x7d\x31\xb6\xa0\x6e\x7e\xd1\xeb\xbd\x5e\x49\xa8\x0c\x64\x29\x99\xc7\x21\xb9\x96\xb6\xa9\x6b\x0b\x7c\xdc\xe0\x39\x37\x7e\x2e\x0c\x73\xf0\x1e\x48\x37\x4f\x76\x68\x4e\x42\x9f\x5b\xfa\xcc\x4a\x29\xa3\x9d\x1c\x8c\xcc\xc1\xcc\x48\x1a\xfc\xfb\xdd\xfa\xc8\x9c\xc2\x64\x64\xfe\xa5\x8f\x17\xc6\xdb\xb7\xfe\x7b\xfb\x3d\xd0\x07\x3f\x0c\x68\xef\x0a\x3f\x6f\xf5\x24\x49\x48\x94\x0c\xe1\xfb\xdd\x31\xd9\x48\x7d\x3a\x90\x8c\x37\x8d\x25\xd4\x9a\xd4\x2d\x72\xa7\x48\xa6\xd7\xb6\x68\x3e\xd7\x73\x45\x6d\xdb\x9e\xa3\x4c\xa2\x87\xfc\x23\xd5\x7a\x64\x33\xb4\x5b\x25\xd1\x30\xe6\x0e\x56\x88\x76\x2f\x4b\xa2\xf1\x32\xa0\x52\xbe\xb7\x97\x45\xec\x38\x88\xee\xc1\x67\xa9\xb4\x8f\xae\x4a\x41\x4a\x3f\xe9\x11\xb2\x27\xf4\xc3\x08\xed\x17\x97\x57\xb1\x2e\x47\xa1\xa6\x82\xba\xdd\xd1\xfe\x9b\x5e\x56\x6f\x2a\x5a\x79\xc3\x39\x9b\x02\xf9\xa0\x9a\x32\x23\x94\xa0\xde\xdf\x0c\x4a\x3c\x8f\xf6\xa9\x18\xa1\x5a\xf3\x7f\xcd\x08\xd5\x56\x4f\x92\x84\xcf\x30\xc2\x1e\xb7\x12\x46\x28\xa2\x56\xc1\x08\x25\xaa\xa7\xee\xc6\xc3\x75\x91\x1c\xef\x57\x9a\xab\xa9\x39\x2e\x73\x11\xa4\x12\x83\xe9\x78\x31\x31\x15\x43\xa8\x2d\x68\x53\x85\x0c\x5f\xe4\x13\xf1\xcf\xea\xd5\x0c\x57\xef\xf5\x38\x7a\x8e\x4f\x84\x38\x3f\x44\x68\x27\x72\xbf\xa4\xf6\x28\xf7\xab\xe1\xa8\x76\x7f\x48\x24\x81\x55\xc8\x8f\x58\xf4\x60\xa8\x5b\xfa\x51\xbd\x32\xbd\xbd\x2b\x6f\x79\xd4\xd5\x32\xeb\xdc\x39\xdc\xcc\xa6\x13\x10\xd2\xa5\xac\x5f\x6b\xb5\x6d\xca\xa8\x6c\x8a\x47\xff\x7f\x17\xad\xf6\x55\xa3\xd5\x69\x5c\x74\xa1\x7d\xd5\xbb\xe8\xf4\x3a\x9d\x66\xb7\xfb\xf5\xdb\x55\xfb\xff\xad\x8b\x5a\xeb\xc2\xde\xa2\xb2\x5b\xfe\x6b\xfb\xb2\xd3\x4d\xe4\xff\xb0\xf3\xa9\xd8\xf3\xe6\xea\xeb\xb7\xf4\x4d\x27\x5d\x97\x6d\x46\x02\x14\xbb\x1f\x5c\xb5\x3b\x4a\xfc\xef\xe6\xae\x6c\xee\x5d\x2d\xdf\x97\xce\xf2\x5e\xa9\xf2\x99\x5d\x1f\xb5\xdc\x6a\xa8\xe5\x36\x3e\x6d\x3d\xc0\xb4\xc2\xec\xcb\xa3\x90\x09\xe4\x5d\xbb\xc4\xdc\x38\xaa\xba\xe7\xc6\xd8\x18\x58\x99\x9f\x28\x4d\x81\x7b\xfb\x4d\x83\xb6\x96\xfe\x1c\xd9\x5d\xb0\x7b\x37\x89\xf7\xa5\xb8\xbc\x46\xa8\xa4\xe6\x13\xb6\x4d\x64\x3a\xe6\xb5\xdc\x0c\xd7\x72\xa3\x59\x5b\x8f\x5e\x2d\x37\x6f\xb5\xcc\x74\xd5\x32\x23\xf2\x00\x14\x47\x92\xf8\xa7\xa1\xd8\xc1\x1d\x95\x50\x54\x52\x46\xf1\xdb\x8e\x1e\xf0\x75\xfb\xd3\xd8\x9c\x5b\x33\x7d\x64\x7e\x8c\x0e\x0b\xaa\x93\x49\xa9\x0f\x87\x19\xb5\x95\xd6\xe1\x76\x36\x9a\xe8\xb3\x3b\xf8\xd3\xb8\x53\xa0\x1e\x66\xef\x4c\x21\xc6\x8c\x3e\xc6\x78\xa2\x00\x8a\x86\xaa\x22\x28\xd9\x86\x85\x39\xfa\xb9\x30\xe0\x2c\x53\x2c\x07\xc7\x67\xb1\xa5\xb2\x4c\x61\x9f\x0e\x91\xa2\xdd\xaa\x80\x0e\xb9\xf2\x16\x5f\xf6\xe2\xfd\x11\x9e\xb0\xce\x3e\x12\xd5\xc7\x0a\xad\xd8\x6e\x27\x0c\xe2\x98\x5a\xab\x32\xbf\x27\x88\x5d\x7f\x72\x82\x13\x06\x91\x8f\x12\x13\x4f\xfe\x0d\x00\x00\xff\xff\x7e\x47\x0b\xb3\x11\x15\x00\x00")

func latestSqlBytes() ([]byte, error) {
	return bindataRead(
		_latestSql,
		"latest.sql",
	)
}

func latestSql() (*asset, error) {
	bytes, err := latestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "latest.sql", size: 5393, mode: os.FileMode(420), modTime: time.Unix(1530096833, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x41\x4f\xfa\x40\x10\xc5\xef\xfb\x29\xe6\x08\xf9\xff\x49\xd4\x08\x17\x4e\x55\x6a\x42\xac\x80\xb5\x3d\x70\x6a\x86\xdd\x49\x9d\xd8\xee\x36\xbb\x53\xc4\x6f\x6f\x30\x51\xda\x82\x9e\x7f\x2f\x33\xef\xbd\x99\xc9\x04\xfe\xd5\x5c\x7a\x14\x82\xbc\x51\xf7\x69\x1c\x65\x31\x64\xd1\x5d\x12\x43\x4a\x9a\x78\x4f\x66\x83\x1f\x35\x59\x81\x91\x02\x60\x03\x3b\x2e\x03\x79\xc6\xea\xbf\x02\x70\x0d\x79\x14\x76\xb6\x60\x03\x7b\xf4\xfa\x15\xfd\xe8\x66\x3a\x1d\x43\xbe\x5a\x3e\xe7\x31\xac\xd6\x19\xac\xf2\x24\x39\x8a\x1b\xef\x34\x85\x40\xa6\x40\x01\xe1\x9a\x82\x60\xdd\xf4\x25\x58\xb2\x2d\x0b\x71\x6f\x64\xfb\xf3\xba\xaa\x20\x28\x6d\xf8\x9d\x6f\xd2\xe5\x53\x94\x6e\xe1\x31\xde\xc2\x88\xcd\x58\x8d\xe7\xaa\x9f\xed\x85\xac\x64\x1e\x6d\x40\x7d\x74\xff\x9d\xed\x14\x4c\x4e\xb0\x1b\x6d\x76\xdb\xd9\x04\xe7\x56\xae\xaf\x06\x4e\x5d\xeb\x35\xfd\xe0\xe9\x6c\x80\xdb\x5d\xcd\x22\x7f\x35\x12\x5a\xad\x89\xcc\x50\xb2\x88\x1f\xa2\x3c\x39\xc9\x2a\x32\x25\xf9\xe3\x71\xd8\xca\x19\x25\xbb\xa7\xca\x35\x54\x1c\x8c\x07\xa1\x83\xf4\x56\x78\x0a\x6d\x25\x5f\xac\x57\xe9\x70\xca\xc5\x5a\xbb\x1f\xb4\x70\xef\x56\x2d\xd2\xf5\xe6\xf2\x07\xcd\xbb\x6c\x70\x81\xb9\xfa\x0c\x00\x00\xff\xff\x4d\x61\x55\x6b\x8b\x02\x00\x00")

func migrations01_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initSql,
		"migrations/01_init.sql",
	)
}

func migrations01_initSql() (*asset, error) {
	bytes, err := migrations01_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_init.sql", size: 651, mode: os.FileMode(420), modTime: time.Unix(1530096833, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations02_payment_idSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xce\x31\x0b\xc2\x30\x10\x05\xe0\x3d\xbf\xe2\xc6\x16\xe9\x22\x74\xea\x74\x36\x11\x0b\x21\xd5\xf4\xe2\x5a\x82\x06\xc9\xd0\x6b\xad\x29\xe2\xbf\x17\x5c\xec\xa4\xe3\xe3\xc1\xf7\x5e\x51\xc0\x66\x88\xb7\xd9\xa7\x00\x6e\x12\xa8\x49\x59\x20\xdc\x69\x05\x5d\xe0\x44\xb3\xe7\x87\xbf\xa4\x38\x32\xa0\x94\x30\xf9\xd7\x10\x38\xf5\xf1\x0a\x67\xb4\xf5\x01\x6d\xb6\x2d\xcb\x1c\x8c\xd3\x1a\xa4\xda\xa3\xd3\xf4\x09\xd5\x5f\xaa\x6e\x4d\x47\x16\x1b\x43\x2b\xb5\x5f\x38\xde\x97\x00\xce\x34\x27\xa7\x20\xfb\x36\x79\x25\xc4\xfa\xac\x1c\x9f\xfc\x73\x43\xda\xf6\xb8\x92\x2b\xf1\x0e\x00\x00\xff\xff\xd0\xf6\xe1\x10\xeb\x00\x00\x00")

func migrations02_payment_idSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations02_payment_idSql,
		"migrations/02_payment_id.sql",
	)
}

func migrations02_payment_idSql() (*asset, error) {
	bytes, err := migrations02_payment_idSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/02_payment_id.sql", size: 235, mode: os.FileMode(420), modTime: time.Unix(1530096833, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations03_transaction_idSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x08\x4a\x4d\x4e\xcd\x2c\x4b\x4d\x09\x48\xac\xcc\x4d\xcd\x2b\x51\x70\x74\x71\x51\x28\x29\x4a\xcc\x2b\x4e\x4c\x2e\xc9\xcc\xcf\x8b\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x33\xd1\x54\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\xf7\xd3\x77\x54\xb7\xe6\xe2\x42\x36\xd9\x25\xbf\x3c\x0f\xaf\xd9\x2e\x41\xfe\x01\x68\x86\x5b\x73\x01\x02\x00\x00\xff\xff\x95\xd0\x87\x98\x9c\x00\x00\x00")

func migrations03_transaction_idSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations03_transaction_idSql,
		"migrations/03_transaction_id.sql",
	)
}

func migrations03_transaction_idSql() (*asset, error) {
	bytes, err := migrations03_transaction_idSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/03_transaction_id.sql", size: 156, mode: os.FileMode(420), modTime: time.Unix(1530096833, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations04_table_namesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x08\x4a\x4d\x4e\xcd\x2c\x4b\x4d\x09\x48\xac\xcc\x4d\xcd\x2b\x51\x08\x72\xf5\x73\xf4\x75\x55\x08\xf1\x57\x28\x82\xca\xc4\x17\x40\xa4\xac\x51\xf4\x05\xa7\xe6\x95\x84\x14\x25\xe6\x15\x27\x26\x97\x64\xe6\xe7\x21\xe9\x2b\x4e\xcd\x2b\x89\x2f\x41\x48\x59\x73\x71\x21\xdb\xef\x92\x5f\x9e\x87\x62\x12\xba\x3d\x48\x46\xa1\x39\x0e\xd5\x05\xe8\xf6\x20\xe9\x43\x73\x9c\x35\x17\x20\x00\x00\xff\xff\xb7\xf3\x31\x25\x01\x01\x00\x00")

func migrations04_table_namesSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations04_table_namesSql,
		"migrations/04_table_names.sql",
	)
}

func migrations04_table_namesSql() (*asset, error) {
	bytes, err := migrations04_table_namesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/04_table_names.sql", size: 257, mode: os.FileMode(420), modTime: time.Unix(1530096833, 0)}
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
	"latest.sql":                       latestSql,
	"migrations/01_init.sql":           migrations01_initSql,
	"migrations/02_payment_id.sql":     migrations02_payment_idSql,
	"migrations/03_transaction_id.sql": migrations03_transaction_idSql,
	"migrations/04_table_names.sql":    migrations04_table_namesSql,
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
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"01_init.sql":           &bintree{migrations01_initSql, map[string]*bintree{}},
		"02_payment_id.sql":     &bintree{migrations02_payment_idSql, map[string]*bintree{}},
		"03_transaction_id.sql": &bintree{migrations03_transaction_idSql, map[string]*bintree{}},
		"04_table_names.sql":    &bintree{migrations04_table_namesSql, map[string]*bintree{}},
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
