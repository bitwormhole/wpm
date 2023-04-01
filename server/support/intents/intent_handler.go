package intents

import (
	"fmt"
	"os"
	"os/exec"
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
		return fmt.Errorf("intent is nil")
	}

	cli := i.CLI
	if cli == nil {
		return fmt.Errorf("intent.cli is nil")
	}

	c := inst.makeCmd(cli)
	if c == nil {
		return fmt.Errorf("intent-cmd is nil")
	}

	env := cli.Env
	if env != nil {
		c.Env = inst.makeEnv(env)
	}

	wd := strings.TrimSpace(cli.WD)
	if wd != "" {
		c.Dir = wd
	}

	return inst.run(c)
}

func (inst *IntentHandlerImpl) makeEnv(src map[string]string) []string {
	old := os.Environ()
	dst := make([]string, 0)
	dst = append(dst, old...)
	cnt := 0
	for k, v := range src {
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		if k == "" && v == "" {
			continue
		}
		str := k + "=" + v
		dst = append(dst, str)
		cnt++
	}
	if cnt == 0 {
		return nil
	}
	// sort.Strings(dst)
	return dst
}

func (inst *IntentHandlerImpl) makeCmd(req *dto.CommandRequest) *exec.Cmd {
	if req.Arguments == nil {
		return exec.Command(req.Command)
	}
	return exec.Command(req.Command, req.Arguments...)
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
