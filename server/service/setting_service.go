package service

import "github.com/bitwormhole/wpm/server/web/dto"

// SettingService ...
type SettingService interface {
	ListAll() map[string]string

	PutString(name, value string)
	GetString(name, defaultValue string) string

	PutSettings(o *dto.Settings) error
	GetSettings() (*dto.Settings, error)
}
