package resdto

import "gobase/pkg/schemas"

type TCreateImageResDto struct {
	schemas.TImage
	PreSignUrl string
}
