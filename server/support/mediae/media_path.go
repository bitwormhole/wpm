package mediae

import (
	"fmt"
	"strings"

	"bitwormhole.com/starter/afs"
	"github.com/bitwormhole/wpm/server/web/dto"
)

func makeMediaWebPath(o *dto.Media) (string, error) {

	// like: "/media/t1/t2/sum/filename"

	if o == nil {
		return "", fmt.Errorf("media object is nil")
	}

	name := o.Name
	sum := o.SHA256SUM.String()
	ctype := o.ContentType

	builder := strings.Builder{}
	builder.WriteString("/media/")
	builder.WriteString(ctype)
	builder.WriteString("/")
	builder.WriteString(sum)
	builder.WriteString("/")
	builder.WriteString(name)
	path := builder.String()

	if name == "" || sum == "" || ctype == "" {
		return "", fmt.Errorf("bad media path: " + path)
	}

	return path, nil
}

func makeMediaLocalPath(o *dto.Media, parent afs.Path) (afs.Path, error) {

	// like: "/media/sum[0:2]/sum[2:]"

	if o == nil || parent == nil {
		return nil, fmt.Errorf("param is nil")
	}

	sum := o.SHA256SUM.String()
	if len(sum) < 20 {
		return nil, fmt.Errorf("bad sum value: " + sum)
	}

	if parent.GetName() == "media" {
		parent = parent.GetParent()
	}

	builder := strings.Builder{}
	builder.WriteString("media/")
	builder.WriteString(sum[0:2])
	builder.WriteString("/")
	builder.WriteString(sum[2:])
	path := builder.String()

	return parent.GetChild(path), nil
}
