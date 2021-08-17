package vo

import (
	"github.com/bitwormhole/wpm/web/base"
	"github.com/bitwormhole/wpm/web/dto"
)

type TestVO struct {
	base.VO

	Locator *dto.TestGitRepoLocator
}
