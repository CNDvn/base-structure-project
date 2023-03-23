package utils

type TPagination struct {
	Size int64
	Page int64
}

func (t *TPagination) GetLimit() int64 {
	if t.Size == 0 {
		return 20
	}
	return t.Size
}

func (t *TPagination) GetSkip() int64 {
	return (t.Page - 1) * t.Size
}
