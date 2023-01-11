package implservice

import (
	"context"
	"errors"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// RepositoryLocateServiceImpl ...
type RepositoryLocateServiceImpl struct {
	markup.Component `id:"RepositoryLocateService"`
}

func (inst *RepositoryLocateServiceImpl) _Impl() service.RepositoryLocateService {
	return inst
}

func (inst *RepositoryLocateServiceImpl) makeResult(ctx context.Context, dir fs.Path) (*dto.LocalRepository, error) {
	name := dir.Name()
	if name == ".git" {
		parent := dir.Parent()
		if parent != nil {
			name = parent.Name()
		}
	}
	o := &dto.LocalRepository{}
	o.DisplayName = name
	o.Name = name
	o.Path = dir.String()
	o.Ready = dir.IsDir()
	return o, nil
}

// Locate ...
func (inst *RepositoryLocateServiceImpl) Locate(ctx context.Context, path fs.Path) (*dto.LocalRepository, error) {
	for p := path; p != nil; p = p.Parent() {
		if p.IsDir() {
			p2 := p.GetChild(".git")
			if inst.IsRepositoryDirectory(ctx, p2) {
				return inst.makeResult(ctx, p2)
			} else if inst.IsRepositoryDirectory(ctx, p) {
				return inst.makeResult(ctx, p)
			}
		}
	}
	return nil, errors.New("no repository at path of " + path.String())
}

// IsRepositoryDirectory ...
func (inst *RepositoryLocateServiceImpl) IsRepositoryDirectory(ctx context.Context, path fs.Path) bool {

	// required
	config := path.GetChild("config")
	refs := path.GetChild("refs")
	objects := path.GetChild("objects")

	if !refs.IsDir() {
		return false
	}
	if !objects.IsDir() {
		return false
	}
	if !config.IsFile() {
		return false
	}

	// options
	head := path.GetChild("HEAD")
	logs := path.GetChild("logs")
	info := path.GetChild("info")
	hooks := path.GetChild("hooks")
	index := path.GetChild("index")
	rate := 0

	if head.IsFile() {
		rate++
	}
	if logs.IsDir() {
		rate++
	}
	if info.IsDir() {
		rate++
	}
	if hooks.IsDir() {
		rate++
	}
	if index.IsFile() {
		rate++
	}

	return rate >= 2
}
