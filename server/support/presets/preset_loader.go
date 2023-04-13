package presets

import (
	"context"
	"runtime"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/vo"
)

type myPresetsLoader struct {
	ac           application.Context
	httpClientEx service.HTTPClientExService
	listFileName string
}

func (inst *myPresetsLoader) Load() (*vo.Online, error) {
	list, err := inst.loadPresetFileList()
	if err != nil {
		return nil, err
	}
	dst := &vo.Online{}
	for _, filename := range list {
		err := inst.loadPresetFile(filename, dst)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func (inst *myPresetsLoader) loadPresetFile(filename string, dst *vo.Online) error {

	// data, err := inst.ac.GetResources().GetBinary(filename)
	// if err != nil {
	// 	return err
	// }

	// src := &vo.Online{}
	// err = json.Unmarshal(data, src)
	// if err != nil {
	// 	return err
	// }

	ctx := context.Background()
	url := filename
	src, err := inst.httpClientEx.FetchOnlineDoc(ctx, url, nil)
	if err != nil {
		return err
	}

	dst.ContentTypes = append(dst.ContentTypes, src.ContentTypes...)

	dst.Executables = append(dst.Executables, src.Executables...)

	dst.IntentTemplates = append(dst.IntentTemplates, src.IntentTemplates...)

	dst.Mediae = append(dst.Mediae, src.Mediae...)

	dst.Sources = append(dst.Sources, src.Sources...)

	dst.Packages = append(dst.Packages, src.Packages...)

	return nil
}

func (inst *myPresetsLoader) loadPresetFileList() ([]string, error) {
	const (
		c1 = "\r"
		c2 = "\n"
	)
	listFileName := inst.getListFileName()
	text, err := inst.ac.GetResources().GetText(listFileName)
	if err != nil {
		return nil, err
	}
	text = strings.ReplaceAll(text, c1, c2)
	rows := strings.Split(text, c2)
	dst := make([]string, 0)
	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row == "" {
			continue // skip
		} else if strings.HasPrefix(row, "#") {
			continue // skip
		}
		dst = append(dst, row)
	}
	return dst, nil
}

func (inst *myPresetsLoader) getListFileName() string {
	const token = "{{os}}"
	os := runtime.GOOS
	name := inst.listFileName
	return strings.ReplaceAll(name, token, os)
}
