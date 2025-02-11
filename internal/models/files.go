package models

import "time"

type FileInfo struct {
	Name        string
	Size        int64
	ContentType string
	URL         string
	UploadedAt  time.Time
}
