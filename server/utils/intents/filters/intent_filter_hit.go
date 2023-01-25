package filters

import (
	"strings"

	"github.com/bitwormhole/wpm/server/web/dto"
)

func hit(i *dto.Intent, names []string) bool {
	if i == nil {
		return false
	}

	const ch1 = "\\"
	const ch2 = "/"
	name1 := i.Command
	name1 = strings.ReplaceAll(name1, ch1, ch2)
	name1s := strings.Split(name1, ch2)
	name1 = name1s[len(name1s)-1]

	name1 = strings.TrimSpace(name1)
	name1 = strings.ToLower(name1)
	if name1 == "" {
		return false
	}

	for _, name2 := range names {
		name2 = strings.TrimSpace(name2)
		name2 = strings.ToLower(name2)
		if name1 == name2 {
			return true
		}
	}

	return false
}
