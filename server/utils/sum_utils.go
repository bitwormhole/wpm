package utils

import (
	"crypto/sha256"
	"os"

	"bitwormhole.com/starter/afs"
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

// ComputeFileSHA256sumAFS ...
func ComputeFileSHA256sumAFS(file afs.Path) (util.Hex, error) {
	md := sha256.New()
	src, err := file.GetIO().OpenReader(&afs.Options{Read: true, File: true, Flag: os.O_RDONLY})
	if err != nil {
		return "", err
	}
	err = util.PumpStream(src, md, nil)
	sum := md.Sum(nil)
	hex := util.HexFromBytes(sum)
	return hex, nil
}

// ComputeSHA256sum ...
func ComputeSHA256sum(data []byte) util.Hex {
	sum := sha256.Sum256(data)
	hex := util.HexFromBytes(sum[:])
	return hex
}
