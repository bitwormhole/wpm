package executables

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/utils"
	"github.com/bitwormhole/wpm/server/web/dto"
)

type myImportPresetExecutablesHanlder struct {
	context context.Context
	parent  *ExecutableImportServiceImpl
}

func (inst *myImportPresetExecutablesHanlder) Import() error {

	props, err := inst.loadProperties()
	if err != nil {
		return err
	}

	items, err := inst.loadItems(props)
	if err != nil {
		return err
	}

	return inst.insertItems(items)
}

func (inst *myImportPresetExecutablesHanlder) loadProperties() (collection.Properties, error) {
	osName := runtime.GOOS
	name := "preset-executables-" + osName + ".properties"
	res := inst.parent.AC.GetResources()
	text, err := res.GetText(name)
	if err != nil {
		return nil, err
	}
	return collection.ParseProperties(text, nil)
}

func (inst *myImportPresetExecutablesHanlder) loadItems(src collection.Properties) ([]*dto.Executable, error) {
	const (
		prefix = "executable."
		suffix = ".name"
	)
	dst := make([]*dto.Executable, 0)
	namelist := (&utils.PropertiesUtil{}).GetNameList(src, prefix, suffix)
	for _, name := range namelist {
		item, err := inst.loadItem(src, prefix, name)
		if err == nil {
			dst = append(dst, item)
		} else {
			// vlog.Warn(err)
			return nil, err
		}
	}
	return dst, nil
}

func (inst *myImportPresetExecutablesHanlder) loadItem(src collection.Properties, prefix, name string) (*dto.Executable, error) {

	key1 := prefix + name
	gett := src.Getter()
	o := &dto.Executable{}

	tags := gett.GetString(key1+".tags", "")
	o.Tags = dxo.StringList(tags)

	o.Name = gett.GetString(key1+".name", "")
	o.Title = gett.GetString(key1+".title", "")
	o.Description = gett.GetString(key1+".description", "")
	o.IconURL = gett.GetString(key1+".icon", "")
	o.OpenWithPriority = gett.GetInt(key1+".priority", 0)

	return o, nil
}

func (inst *myImportPresetExecutablesHanlder) insertItems(items []*dto.Executable) error {
	for _, item := range items {
		err := inst.insertItem(item)
		if err != nil {
			vlog.Warn(err)
		}
	}
	return nil
}

func (inst *myImportPresetExecutablesHanlder) insertItem(item *dto.Executable) error {
	err := inst.prepareItem(item)
	if err != nil {
		return err
	}
	ser := inst.parent.ExecutableService
	_, err = ser.Insert(inst.context, item)
	return err
}

func (inst *myImportPresetExecutablesHanlder) prepareItem(item *dto.Executable) error {

	name := item.Name
	cmd := exec.Command(name)
	fsSer := inst.parent.FileSystemService
	path := fsSer.Path(cmd.Path)

	if path == nil {
		return fmt.Errorf("no executable with name:" + name)
	}

	if !path.IsFile() {
		return fmt.Errorf("no executable file at: " + path.GetPath())
	}

	item.Path = path.GetPath()
	item.Name = path.GetName()
	return nil
}
