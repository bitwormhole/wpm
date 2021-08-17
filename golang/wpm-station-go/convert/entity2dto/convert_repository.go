package entity2dto

import (
	"github.com/bitwormhole/wpm/data/entity"
	"github.com/bitwormhole/wpm/web/dto"
)

func ConvertRepository(src *entity.Repository) *dto.Repository {

	dst := &dto.Repository{}
	if src == nil {
		return dst
	}

	dst.Name = src.Name
	dst.DisplayName = src.DisplayName
	dst.Path = src.Path

	return dst
}

func ConvertRepositoryList(src []*entity.Repository) []*dto.Repository {
	dst := make([]*dto.Repository, 0)
	if src == nil {
		return dst
	}
	for _, item := range src {
		item2 := ConvertRepository(item)
		dst = append(dst, item2)
	}
	return dst
}
