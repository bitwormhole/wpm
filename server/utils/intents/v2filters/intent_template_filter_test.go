package v2filters

import (
	"context"
	"testing"

	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/dto"
)

func TestPrepareCliParams(t *testing.T) {

	c := context.Background()
	o1 := &dto.Intent{}
	filter := &IntentTemplateFilter{}
	temp := &dto.IntentTemplate{}

	args := make([]string, 0)
	env := make(map[string]string)
	props := make(map[string]string)

	// props
	props["args.a1.1"] = "aa1"
	props["args.a3.1"] = "aa1"
	props["args.a2.2"] = "aa2"
	props["args.a2.1"] = "aa1"
	props["env.test.1"] = "et1"
	props["env.test.e3"] = "ete3"
	props["intent.test.wd"] = "/intent/test/wd"
	props["intent.test.command"] = "xxx"

	// env
	env["e1"] = "hello"
	env["${env.test.1}"] = "666"
	env["e3"] = "${env.test.e3}"

	// args
	args = append(args, "a1 ${args.a1.1}")
	args = append(args, "a2 ${args.a2.1} ${args.a2.2} ")
	args = append(args, "${args.a3.1}")

	temp.Command = "${intent.test.command}"
	temp.WD = "${intent.test.wd}"
	temp.Arguments = dxo.NewStringListCRLF(args)
	temp.Env = dxo.NewStringMap(env)

	o1.Properties = props
	o1.Template = temp

	o2, err := filter.prepareCliParams(c, o1)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(o2.Command)
	}
}
