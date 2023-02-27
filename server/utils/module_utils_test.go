package utils

import (
	"fmt"
	"testing"

	"github.com/bitwormhole/starter"
	ginstarter "github.com/bitwormhole/starter-gin"
	gormstarter "github.com/bitwormhole/starter-gorm"
	"github.com/bitwormhole/starter/application"
)

func TestCheckModuleVersion(t *testing.T) {

	mods := []application.Module{}

	mods = append(mods, starter.Module())
	mods = append(mods, ginstarter.Module())
	mods = append(mods, gormstarter.Module())

	err := CheckModuleVersion(mods...)
	if err != nil {
		fmt.Println(err)
	}
}
