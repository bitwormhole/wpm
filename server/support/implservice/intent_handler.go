package implservice

import (
	"fmt"
	"os/exec"
	"sort"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentHandlerImpl ...
type IntentHandlerImpl struct {
	markup.Component `id:"IntentHandlerService"`
}

func (inst *IntentHandlerImpl) _Impl() service.IntentHandlerService {
	return inst
}

// HandleIntent ...
func (inst *IntentHandlerImpl) HandleIntent(i *dto.Intent) error {

	if i == nil {
		return nil
	}

	params := i
	if params == nil {
		return fmt.Errorf("intent-cli is nil")
	}

	c := inst.makeCmd(params)
	env := params.Env
	if env != nil {
		c.Env = inst.makeEnv(env)
	}

	wd := strings.TrimSpace(params.WD)
	if wd != "" {
		c.Dir = wd
	}

	return inst.run(c)
}

func (inst *IntentHandlerImpl) makeEnv(src map[string]string) []string {
	dst := make([]string, 0)
	for k, v := range src {
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		str := k + "=" + v
		dst = append(dst, str)
	}
	sort.Strings(dst)
	return dst
}

func (inst *IntentHandlerImpl) makeCmd(p *dto.Intent) *exec.Cmd {
	if p.Arguments == nil {
		return exec.Command(p.Command)
	}
	return exec.Command(p.Command, p.Arguments...)
}

func (inst *IntentHandlerImpl) handleError(x any) {
	if x == nil {
		return
	}
	err, ok := x.(error)
	if ok && err != nil {
		vlog.Error(err)
	}
}

func (inst *IntentHandlerImpl) run(c *exec.Cmd) error {

	vlog.Info("execute ", c.Path)
	go func() {
		err := inst.run2(c)
		inst.handleError(err)
	}()
	return nil
}

func (inst *IntentHandlerImpl) run2(c *exec.Cmd) error {

	defer func() {
		x := recover()
		inst.handleError(x)
	}()

	err := c.Start()
	if err != nil {
		return err
	}

	return c.Wait()
}
