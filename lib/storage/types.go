package storage

import (
	"weasel/middleware/auth"
)

type File struct {
	Name string
	HashName string
	Bucket bucket
	MD5 string
	Path string
	AVCheck bool
	AVMessage string
	Size uint
	ContentType string
	Meta string
	Version uint
	Entity string
	EntityId uint
	Owner *auth.User
}

type Files []File
