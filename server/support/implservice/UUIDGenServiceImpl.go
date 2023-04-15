package implservice

import (
	"bytes"
	"crypto/md5"
	"strconv"
	"strings"
	"sync"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
)

// UUIDGenServiceImpl ...
type UUIDGenServiceImpl struct {
	markup.Component `id:"UUIDGenService" initMethod:"Init"`

	t0     util.Time
	index  int64
	locker sync.Mutex
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

func (inst *UUIDGenServiceImpl) nextIndex() int64 {
	inst.locker.Lock()
	defer func() {
		inst.locker.Unlock()
	}()
	inst.index++
	return inst.index
}

// GenerateUUID ...
func (inst *UUIDGenServiceImpl) GenerateUUID(seed string) dxo.UUID {

	const sep = '\n'
	t0 := inst.t0
	now := util.Now()
	index := inst.nextIndex()

	builder := bytes.Buffer{}
	builder.WriteString(seed)
	builder.WriteRune(sep)
	builder.WriteString(inst.time2str(t0))
	builder.WriteRune(sep)
	builder.WriteString(inst.time2str(now))
	builder.WriteRune(sep)
	builder.WriteString(inst.int2str(index))

	sum := inst.sum(builder.Bytes())
	str := inst.stringifySum(sum)
	return dxo.UUID(str)
}

func (inst *UUIDGenServiceImpl) stringifySum(b []byte) string {
	const wantLength = 32
	str := util.HexFromBytes(b).String()
	if len(str) < wantLength {
		return str
	}
	cclist := []int{8, 4, 4, 4}
	builder := strings.Builder{}
	i1 := 0
	for _, cc := range cclist {
		i2 := i1 + cc
		part := str[i1:i2]
		builder.WriteString(part)
		builder.WriteString("-")
		i1 = i2
	}
	builder.WriteString(str[i1:])
	return builder.String()
}

func (inst *UUIDGenServiceImpl) sum(b []byte) []byte {
	if b == nil {
		b = make([]byte, 0)
	}
	sum := md5.Sum(b)
	return sum[:]
}

func (inst *UUIDGenServiceImpl) int2str(n int64) string {
	return strconv.FormatInt(n, 10)
}

func (inst *UUIDGenServiceImpl) time2str(t util.Time) string {
	return strconv.FormatInt(t.Int64(), 10)
}
