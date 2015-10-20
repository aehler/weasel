package storage

import (
	"weasel/app/registry"
	"weasel/app/crypto"
	"path/filepath"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	MaxMemory = 1024 * 1024 * 1
	FileMode   = 0755
	FileMask   = "file[%d]"
	MetaMask   = "meta[%d]"
)

var storageInstance = registry.Registry.Storage("local")

func writeFiles(ent string, entId uint, request *http.Request) (Files, error) {

	var result Files

	if err := request.ParseMultipartForm(MaxMemory); err != nil {

		return nil, err
	}

	folder := generateFolderPath(ent, entId)

	if err := os.MkdirAll(fmt.Sprintf("%s%s", storageInstance.Dir, folder), FileMode); err != nil {

		return nil, err
	}

	for name, files := range request.MultipartForm.File {

		for _, file := range files {

			var num int

			fmt.Sscanf(name, FileMask, &num)

			filePath := fmt.Sprintf("%s%s", folder, fileNameToHash(file.Filename))

			localFile, err := os.Create(fmt.Sprintf("%s%s", storageInstance.Dir, filePath))

			md5hash := md5.New()

			mw := io.MultiWriter(md5hash, localFile)

			f, err := file.Open()

			if err != nil {

				return nil, err
			}

			size, err := io.Copy(mw, f)

			if err != nil {

				return nil, err
			}

			if err := localFile.Sync(); err != nil {

				return nil, err
			}

			if err := localFile.Close(); err != nil {

				return nil, err
			}

			result = append(result, File{
				Name:        file.Filename,
				Meta:        request.Form.Get(fmt.Sprintf(MetaMask, num)),
				ContentType: file.Header.Get("Content-Type"),
				MD5:         hex.EncodeToString(md5hash.Sum(nil)),
				Size:        uint(size),
				Path:        filePath,
				Entity:      ent,
				EntityId:    entId,
			})


		}
	}

	return result, nil
}

func generateFolderPath(objectName string, objectValue uint) string {

	path := time.Now().Format("2006/01/02/15")

	return fmt.Sprintf("/%s/%s/%s", path, objectName, crypto.Encrypt(fmt.Sprintf("%d", objectValue), "userdata"))
}

func fileNameToHash(fileName string) string {

	return fmt.Sprintf("%s%s", crypto.Unique(), filepath.Ext(fileName))
}
