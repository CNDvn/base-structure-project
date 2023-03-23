package resdto

import "gobase/pkg/schemas"

type TCreateVideoResDto struct {
	schemas.TVideo
	PreSignUrl string
}
