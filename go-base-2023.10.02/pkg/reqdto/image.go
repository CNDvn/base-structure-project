package reqdto

import "mime/multipart"

type TCreateImageReqDto struct {
	Title       string
	Description string
	Tags        []string
	FileName    string
	MimeType    string
	File        multipart.FileHeader
}
