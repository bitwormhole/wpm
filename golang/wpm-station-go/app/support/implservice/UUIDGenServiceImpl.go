package implservice

import (
	"bytes"
	"crypto/sha256"
	"strconv"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/app/data/dxo"
	"github.com/bitwormhole/wpm/app/service"
)

// UUIDGenServiceImpl ...
type UUIDGenServiceImpl struct {
	markup.Component `id:"UUIDGenService" initMethod:"Init"`

	t0    util.Time
	index int64
}

func (inst *UUIDGenServiceImpl) _Impl() service.UUIDGenService {
	return inst
}

// Init ...
func (inst *UUIDGenServiceImpl) Init() error {
	inst.t0 = util.Now()
	inst.index = 0
	return nil
}

// GenerateUUID ...
func (inst *UUIDGenServiceImpl) GenerateUUID(seed string) dxo.UUID {

	inst.index++
	const sep = '\n'
	t0 := inst.t0
	now := util.Now()
	index := inst.index

	builder := bytes.Buffer{}
	builder.WriteString(seed)
	builder.WriteRune(sep)
	builder.WriteString(inst.time2str(t0))
	builder.WriteRune(sep)
	builder.WriteString(inst.time2str(now))
	builder.WriteRune(sep)
	builder.WriteString(inst.int2str(index))

	sum := inst.sum(builder.Bytes())
	return dxo.UUID(util.HexFromBytes(sum))
}

func (inst *UUIDGenServiceImpl) sum(b []byte) []byte {
	if b == nil {
		b = make([]byte, 0)
	}
	sum := sha256.Sum256(b)
	return sum[:]
}

func (inst *UUIDGenServiceImpl) int2str(n int64) string {
	return strconv.FormatInt(n, 10)
}

func (inst *UUIDGenServiceImpl) time2str(t util.Time) string {
	return strconv.FormatInt(t.Int64(), 10)
}
