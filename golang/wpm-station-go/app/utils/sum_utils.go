package utils

import (
	"crypto/sha256"

	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/util"
)

// ComputeFileSHA256sum ...
func ComputeFileSHA256sum(file fs.Path) (util.Hex, error) {
	md := sha256.New()
	src, err := file.GetIO().OpenReader(nil)
	if err != nil {
		return "", err
	}
	err = util.PumpStream(src, md, nil)
	sum := md.Sum(nil)
	hex := util.HexFromBytes(sum)
	return hex, nil
}
