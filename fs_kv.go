package main

import (
	"context"
	"encoding/json"
	"io/fs"
	"os"
)

type FileKv struct {
	FilePath     string
	FileName     string
	Compress     bool
	CompressType string
}

var defaultFile string = "skv-tmp-default-file.json"
var defaultPath string = "/tmp/simpleKV-default-tmp-dir/"

func (f *FileKv) GetKV(ctx context.Context, key string) ([]byte, error) {

	err := os.MkdirAll(f.FilePath, fs.ModeType|fs.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return nil, err
		}
	}

	fsSource := os.DirFS(f.FilePath)

	fileRaw, err := fs.ReadFile(fsSource, f.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			newfile, err := os.Create(f.FilePath + f.FileName)
			if err != nil {
				return nil, err
			}
			initBody := RepoBody{
				Metadata: MetaData{
					Comment: "init",
				},
			}

			ins, _ := json.Marshal(initBody)
			os.WriteFile(f.FilePath+f.FileName, ins, fs.ModePerm|fs.ModeType)

			newfile.Close()
			return nil, nil
		}
		return nil, err
	}

	body := RepoBody{}

	err = json.Unmarshal(fileRaw, &body)
	if err != nil {
		return nil, err
	}

	return body.KVS[key], nil
}

func (f *FileKv) PutKV(ctx context.Context, key string, v []byte) error {
	err := os.MkdirAll(f.FilePath, fs.ModeType|fs.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	fsSource := os.DirFS(f.FilePath)

	fileRaw, err := fs.ReadFile(fsSource, f.FileName)
	if err != nil {
		if os.IsNotExist(err) {
			newfile, err := os.Create(f.FilePath + f.FileName)
			if err != nil {
				return err
			}

			initBody := RepoBody{
				Metadata: MetaData{
					Comment: "init",
				},
			}

			ins, _ := json.Marshal(initBody)
			os.WriteFile(f.FilePath+f.FileName, ins, fs.ModePerm|fs.ModeType)
			newfile.Close()
			return nil
		}
		return err
	}

	body := RepoBody{}

	err = json.Unmarshal(fileRaw, &body)
	if err != nil {
		return err
	}

	if body.KVS == nil {
		body.KVS = make(map[string][]byte, 8)
	}

	body.KVS[key] = v

	finalJson, _ := json.Marshal(body)

	return os.WriteFile(f.FilePath+f.FileName, finalJson, fs.ModePerm|fs.ModeType)
}
