package pipe

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bitwormhole/starter/util"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// Manager ...
type Manager struct {
	holders   []*Holder
	startUpAt util.Time
	idgen     dxo.PipeID
	locker    sync.Mutex
}

// Init ...
func (inst *Manager) Init() error {
	inst.startUpAt = util.Now()
	return nil
}

// Insert ...
func (inst *Manager) Insert(o *dto.PipeInfo) (*dto.PipeInfo, error) {

	inst.locker.Lock()
	defer func() {
		inst.locker.Unlock()
	}()

	inst.Trim()
	inst.idgen++

	o.CreatedAt = util.Now()
	o.ID = inst.idgen
	o.DesktopSessionID = inst.makeDSID(o)
	o.UUID = inst.makeUUID(o)
	o.Name = inst.makeName(o)

	h := &Holder{}
	h.Info = *o
	inst.holders = append(inst.holders, h)

	return o, nil
}

func (inst *Manager) makeDSID(o *dto.PipeInfo) dxo.DesktopSessionID {

	const nl = "\n"
	home := o.DesktopSessionHome
	user := o.DesktopSessionUser
	ts := inst.startUpAt.String()

	builder := &strings.Builder{}
	builder.WriteString(user)
	builder.WriteString(nl)
	builder.WriteString(home)
	builder.WriteString(nl)
	builder.WriteString(ts)
	builder.WriteString(nl)

	uuid := inst.buildUUID(builder)
	return dxo.DesktopSessionID(uuid)
}

func (inst *Manager) makeName(o *dto.PipeInfo) dxo.PipeName {
	name := fmt.Sprintf("pipe-%v", o.ID)
	return dxo.PipeName(name)
}

func (inst *Manager) makeUUID(o *dto.PipeInfo) dxo.UUID {

	const nl = "\n"
	now := time.Now()
	ts := inst.startUpAt
	id := int(inst.idgen)

	builder := &strings.Builder{}
	builder.WriteString(now.String())
	builder.WriteString(nl)
	builder.WriteString(ts.String())
	builder.WriteString(nl)
	builder.WriteString(strconv.Itoa(id))
	builder.WriteString(nl)
	builder.WriteString(o.DesktopSessionHome)
	builder.WriteString(nl)
	builder.WriteString(o.DesktopSessionUser)
	builder.WriteString(nl)

	return inst.buildUUID(builder)
}

func (inst *Manager) buildUUID(b *strings.Builder) dxo.UUID {
	src := b.String()
	dst := sha1.Sum([]byte(src))
	hex := util.HexFromBytes(dst[:])
	return dxo.UUID(hex)
}

// Remove ...
func (inst *Manager) Remove(id dxo.PipeID) error {
	h, err := inst.FindHolder(id)
	if err != nil {
		return err
	}
	h.deletedAt = util.Now()
	return nil
}

// WaitAndPull ...
func (inst *Manager) WaitAndPull(id dxo.PipeID, view *vo.Pipe, timeout time.Duration) error {

	h, err := inst.FindHolder(id)
	if err != nil {
		return err
	}

	ch := make(chan int)
	defer func() {
		h.Listener = nil
		close(ch)
	}()

	now := util.Now()
	h.Info.AttachedAt = now
	h.Listener = ch

	select {
	case <-ch:
		// fmt.Println(res)
	case <-time.After(timeout):
		// fmt.Println("timeout 1")
	}

	view.Packets = h.Packets
	h.Packets = nil

	return nil
}

// Push ...
func (inst *Manager) Push(id dxo.PipeID, view *vo.Pipe) error {

	holder, err := inst.FindHolder(id)
	if err != nil {
		return err
	}
	if holder == nil {
		return fmt.Errorf("no holder for pipe: %v", id)
	}

	holder.Packets = append(holder.Packets, view.Packets...)

	ch := holder.Listener
	if ch != nil {
		ch <- 666
	}

	return nil
}

// FindHolder ...
func (inst *Manager) FindHolder(id dxo.PipeID) (*Holder, error) {

	var holder *Holder
	all := inst.holders
	for _, h := range all {
		if h == nil {
			continue
		}
		if h.Info.ID == id {
			holder = h
			break
		}
	}
	errNoHolder := fmt.Errorf("no pipe with id:%v", id)
	if holder == nil {
		return nil, errNoHolder
	} else if holder.deletedAt > 0 {
		return nil, errNoHolder
	}
	return holder, nil
}

// Trim ...
func (inst *Manager) Trim() {
	src := inst.holders
	dst := make([]*Holder, 0)
	for _, item := range src {
		if item == nil {
			continue
		}
		if item.deletedAt > 0 {
			continue
		}
		dst = append(dst, item)
	}
	inst.holders = dst
}

// ListPipes ...
func (inst *Manager) ListPipes() []*dto.PipeInfo {
	src := inst.holders
	dst := make([]*dto.PipeInfo, 0)
	for _, h := range src {
		if h == nil {
			continue
		}
		info := &dto.PipeInfo{}
		*info = h.Info
		dst = append(dst, info)
	}
	return dst
}
