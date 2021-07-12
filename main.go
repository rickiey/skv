package main

import (
	"context"
)

type SKV interface {
	GetKV(ctx context.Context, key string) ([]byte, error)
	PutKV(ctx context.Context, key string, v []byte) error
}

type CompressTypes int

const (
	CompressGzip CompressTypes = iota
	CompressZstd
	Compresslz4
)

func SetFsRepo(filePath, fileName string) {
	DefaultSKV = &FileKv{FilePath: filePath, FileName: fileName}
}

var DefaultSKV SKV = &FileKv{FilePath: defaultPath, FileName: defaultFile}

type MetaData struct {
	Version int    `json:"version"`
	Comment string `json:"comment"`
}

type RepoBody struct {
	Metadata MetaData          `json:"meta_data"`
	KVS      map[string][]byte `json:"kvs"`
}
