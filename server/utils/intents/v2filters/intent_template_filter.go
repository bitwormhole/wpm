package v2filters

import (
	"context"
	"fmt"
	"strings"

	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/utils/intents"
	"github.com/bitwormhole/wpm/server/web/dto"
)

// IntentTemplateFilter ...
type IntentTemplateFilter struct {
	markup.Component `class:"intent-filter-registry"`

	ExecutableService      service.ExecutableService      `inject:"#ExecutableService"`
	LocalRepositoryService service.LocalRepositoryService `inject:"#LocalRepositoryService"`
	IntentTemplateService  service.IntentTemplateService  `inject:"#IntentTemplateService"`
}

func (inst *IntentTemplateFilter) _Impl() intents.FilterRegistry {
	return inst
}

// GetRegistrationList ...
func (inst *IntentTemplateFilter) GetRegistrationList() []*intents.FilterRegistration {
	reg := &intents.FilterRegistration{
		Enabled: true,
		Filter:  inst,
		Order:   0,
	}
	return []*intents.FilterRegistration{reg}
}

// Filter ...
func (inst *IntentTemplateFilter) Filter(c context.Context, o *dto.Intent) (*dto.Intent, error) {

	steps := make([]func(c context.Context, o *dto.Intent) (*dto.Intent, error), 0)

	steps = append(steps, inst.tryLoadRepository)
	steps = append(steps, inst.tryLoadWorktree)
	steps = append(steps, inst.tryLoadProject)
	steps = append(steps, inst.tryLoadFilePath)

	steps = append(steps, inst.prepareExec)
	steps = append(steps, inst.prepareProperties)
	steps = append(steps, inst.findTemplate)
	steps = append(steps, inst.prepareCliParams)

	o1 := o
	for _, step := range steps {
		if o1 == nil {
			break
		}
		o2, err := step(c, o1)
		if err != nil {
			return nil, err
		}
		o1 = o2
	}
	return o1, nil
}

func (inst *IntentTemplateFilter) load(c context.Context, o *dto.Intent) (*dto.Intent, error) {
	return o, nil
}

func (inst *IntentTemplateFilter) tryLoadWorktree(c context.Context, o *dto.Intent) (*dto.Intent, error) {
	// todo ...
	return o, nil
}

func (inst *IntentTemplateFilter) tryLoadRepository(c context.Context, o *dto.Intent) (*dto.Intent, error) {
	repo := o.Repository
	if repo == nil {
		return o, nil
	}
	id := repo.ID
	name := repo.Name
	ser := inst.LocalRepositoryService
	if id > 0 {
		repo2, err := ser.Find(c, id, nil)
		if err != nil {
			return nil, err
		}
		o.Repository = repo2
	} else if name != "" {
		repo2, err := ser.FindByName(c, name, nil)
		if err != nil {
			return nil, err
		}
		o.Repository = repo2
	} else {
		// return nil, fmt.Errorf("no executable (id|name) for intent")
		o.Repository = nil
	}
	return o, nil
}

func (inst *IntentTemplateFilter) tryLoadProject(c context.Context, o *dto.Intent) (*dto.Intent, error) {
	// todo ...
	return o, nil
}

func (inst *IntentTemplateFilter) tryLoadFilePath(c context.Context, o *dto.Intent) (*dto.Intent, error) {
	// todo ...
	return o, nil
}

func (inst *IntentTemplateFilter) makeTemplateSelector(o *dto.Intent) (*dto.IntentTemplate, error) {

	builder := &dto.IntentTemplate{}
	exe := o.Executable
	repo := o.Repository
	worktree := o.Worktree
	project := o.Project
	location := o.Location

	// repo|project|worktree|location
	if location != nil {
		builder.Target = "location"
	} else if project != nil {
		builder.Target = "project"
	} else if worktree != nil {
		builder.Target = "worktree"
	} else if repo != nil {
		builder.Target = "repository"
	} else {
		return nil, fmt.Errorf("no target for IntentTemplateSelectorBuilder")
	}

	// exe
	if exe == nil {
		return nil, fmt.Errorf("no executable")
	}

	builder.Executable = exe.Name
	builder.Action = "open"

	// done
	return builder, nil
}

func (inst *IntentTemplateFilter) makeError(msg string, o *dto.IntentTemplate) error {
	builder := strings.Builder{}
	if o != nil {
		builder.WriteString("action:")
		builder.WriteString(o.Action)
		builder.WriteString(" target:")
		builder.WriteString(o.Target)
		builder.WriteString(" with:")
		builder.WriteString(o.Executable)
	}
	return fmt.Errorf("%v [%v]", msg, builder.String())
}

func (inst *IntentTemplateFilter) findTemplate(c context.Context, o *dto.Intent) (*dto.Intent, error) {

	sel, err := inst.makeTemplateSelector(o)
	if err != nil {
		return nil, err
	}

	ser := inst.IntentTemplateService
	list, err := ser.ListByAET(c, sel)
	if err != nil {
		return nil, err
	}

	errNoTemp := inst.makeError("no intent-template to", sel)
	errMoreThanOne := inst.makeError("more than one intent-template to", sel)

	cnt := len(list)
	if cnt < 1 {
		return nil, errNoTemp
	} else if cnt > 1 {
		return nil, errMoreThanOne
	}

	temp := list[0]
	if temp == nil {
		return nil, errNoTemp
	}
	o.Template = temp

	return o, nil
}

func (inst *IntentTemplateFilter) prepareCliParams(c context.Context, o *dto.Intent) (*dto.Intent, error) {

	p := o.Properties
	tmp := o.Template
	pr := &myPropertiesResolver{}
	pr.init(p)

	args := tmp.Arguments.Array()
	env := tmp.Env.Map()

	o.Command = pr.resolveString(tmp.Command)
	o.Arguments = pr.resolveStringArray(args)
	o.Env = pr.resolveStringMap(env)
	o.WD = pr.resolveString(tmp.WD)

	err := pr.err
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (inst *IntentTemplateFilter) prepareProperties(c context.Context, o *dto.Intent) (*dto.Intent, error) {
	maker := &myIntentPropertiesMaker{}
	p, err := maker.make(o)
	if err != nil {
		return nil, err
	}
	o.Properties = p
	return o, nil
}

func (inst *IntentTemplateFilter) prepareExec(c context.Context, o *dto.Intent) (*dto.Intent, error) {
	ex := o.Executable
	if ex == nil {
		return nil, fmt.Errorf("no executable for intent")
	}
	id := ex.ID
	name := ex.Name
	ser := inst.ExecutableService
	if id > 0 {
		ex2, err := ser.Find(c, id)
		if err != nil {
			return nil, err
		}
		o.Executable = ex2
	} else if name != "" {
		ex2, err := ser.FindByName(c, name)
		if err != nil {
			return nil, err
		}
		o.Executable = ex2
	} else {
		return nil, fmt.Errorf("no executable (id|name) for intent")
	}
	return o, nil
}

////////////////////////////////////////////////////////////////////////////////

type myPropertiesResolver struct {
	properties map[string]string
	err        error
}

func (inst *myPropertiesResolver) init(p map[string]string) {
	if p == nil {
		p = make(map[string]string)
	}
	inst.properties = p
}

func (inst *myPropertiesResolver) handleError(err string) {
	if err == "" {
		return
	}
	inst.err = fmt.Errorf(err)
}

func (inst *myPropertiesResolver) resolveWithName(name string) string {
	value := ""
	table := inst.properties
	name = strings.TrimSpace(name)
	if table != nil {
		value = table[name]
	}
	if value == "" {
		inst.handleError("no property with name [" + name + "]")
	}
	return value
}

func (inst *myPropertiesResolver) resolveString(src string) string {
	const (
		token1 = "${"
		token2 = "}"
	)
	builder := strings.Builder{}
	text := src
	for {
		i1 := strings.Index(text, token1)
		i2 := strings.Index(text, token2)
		if i1 < 0 {
			break
		} else if i2 < 0 {
			inst.handleError("bad token: " + src)
			return ""
		} else if i2 < i1 {
			inst.handleError("bad token: " + src)
			return ""
		}
		t1 := text[0:i1]
		t2 := text[i1+len(token1) : i2]
		t3 := text[i2+len(token2):]
		t22 := inst.resolveWithName(t2)
		builder.WriteString(t1)
		builder.WriteString(t22)
		text = t3
	}
	builder.WriteString(text)
	return builder.String()
}

func (inst *myPropertiesResolver) resolveStringMap(src map[string]string) map[string]string {
	dst := make(map[string]string)
	for k1, v1 := range src {
		k2 := inst.resolveString(k1)
		v2 := inst.resolveString(v1)
		dst[k2] = v2
	}
	return dst
}

func (inst *myPropertiesResolver) resolveStringArray(src []string) []string {
	dst := make([]string, 0)
	for _, value1 := range src {
		value2 := inst.resolveString(value1)
		dst = append(dst, value2)
	}
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type myIntentPropertiesMaker struct{}

func (inst *myIntentPropertiesMaker) make(o *dto.Intent) (map[string]string, error) {

	dst := make(map[string]string)
	err := inst.makeExe(o.Executable, dst)
	if err != nil {
		return nil, err
	}

	err = inst.makeRepo(o.Repository, dst)
	if err != nil {
		return nil, err
	}

	err = inst.makeWorktree(o.Worktree, dst)
	if err != nil {
		return nil, err
	}

	err = inst.makeProject(o.Project, dst)
	if err != nil {
		return nil, err
	}

	err = inst.makeLocation(o.Location, dst)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

func (inst *myIntentPropertiesMaker) makeExe(o *dto.Executable, dst map[string]string) error {

	const prefix = "executable."
	if o == nil {
		return nil
	}

	dst[intents.ExecutableName] = o.Name
	dst[intents.ExecutablePath] = o.Path
	dst[intents.ExecutableTitle] = o.Title

	return nil
}

func (inst *myIntentPropertiesMaker) makeRepo(o *dto.LocalRepository, dst map[string]string) error {
	if o == nil {
		return nil
	}
	dst[intents.RepositoryName] = o.Name
	dst[intents.RepositoryPath] = o.Path
	dst[intents.RepositoryDotGitPath] = o.DotGitPath
	dst[intents.RepositoryWorkDir] = o.WorkingPath
	dst[intents.RepositoryConfigFile] = o.ConfigFile
	dst[intents.RepositoryCoreDir] = o.RepositoryPath
	return nil
}

func (inst *myIntentPropertiesMaker) makeWorktree(o *dto.Worktree, dst map[string]string) error {
	if o == nil {
		return nil
	}
	dst[intents.WorktreeName] = o.Name
	dst[intents.WorktreeDotGitPath] = o.DotGitPath
	dst[intents.WorktreeWorkDir] = o.WorkingDir
	return nil
}

func (inst *myIntentPropertiesMaker) makeProject(o *dto.Project, dst map[string]string) error {
	if o == nil {
		return nil
	}
	dst[intents.ProjectName] = o.Name
	dst[intents.ProjectFullPath] = o.Path
	return nil
}

func (inst *myIntentPropertiesMaker) makeLocation(o *dto.File, dst map[string]string) error {
	if o == nil {
		return nil
	}
	dst[intents.LocationName] = o.Name
	dst[intents.LocationPath] = o.Path
	dst[intents.LocationType] = o.Type
	return nil
}

////////////////////////////////////////////////////////////////////////////////
