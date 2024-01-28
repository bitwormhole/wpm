package repositories

import (
	"crypto/sha1"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/server/data/entity"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/security-gorm/rbacdb"
)

// LocalRepositoryConvertor ...
type LocalRepositoryConvertor struct{}

// ListE2D ...
func (inst *LocalRepositoryConvertor) ListE2D(src []*entity.LocalRepository) []*dto.LocalRepository {
	dst := make([]*dto.LocalRepository, 0)
	for _, item1 := range src {
		item2 := inst.E2D(item1)
		dst = append(dst, item2)
	}
	return dst
}

// E2D ...
func (LocalRepositoryConvertor) E2D(src *entity.LocalRepository) *dto.LocalRepository {

	dst := new(dto.LocalRepository)
	rbacdb.CopyBaseFieldsFromEntityToDTO(&src.BaseEntity, &dst.BaseDTO)
	dst.ID = src.ID

	// todo ...
	dst.Name = src.Name
	dst.DisplayName = src.DisplayName
	dst.Description = src.Description

	dst.Path = src.Path
	dst.ConfigFile = src.ConfigFile

	// dst.ConfigFile = src.ConfigFile
	// dst.WorkingPath = src.WorkingPath
	// dst.RepositoryPath = src.RepositoryPath
	// dst.DotGitPath = src.DotGitPath
	// dst.Path = src.Path

	dst.Bare = src.Bare
	dst.IconURL = "todo..."

	return dst
}

// D2E ...
func (inst *LocalRepositoryConvertor) D2E(src *dto.LocalRepository) *entity.LocalRepository {

	dst := new(entity.LocalRepository)
	rbacdb.CopyBaseFieldsFromDtoToEntity(&src.BaseDTO, &dst.BaseEntity)
	dst.ID = src.ID

	// todo ...
	dst.Name = src.Name
	dst.DisplayName = src.DisplayName
	dst.Description = src.Description

	// dst.ConfigFile = src.ConfigFile
	// dst.RepositoryPath = src.RepositoryPath
	// dst.DotGitPath = src.DotGitPath

	dst.ConfigFile = src.ConfigFile
	dst.Path = src.Path
	dst.RegularPath = dst.ComputeRegularPath()

	dst.Bare = src.Bare
	dst.URN = inst.computeURN(src)

	return dst
}

func (LocalRepositoryConvertor) computeURN(src *dto.LocalRepository) dxo.URN {
	path := src.RegularPath
	sum := sha1.Sum([]byte(path))
	hex := lang.HexFromBytes(sum[:])
	str := "urn:localrepo:" + hex.String()
	return dxo.URN(str)
}
