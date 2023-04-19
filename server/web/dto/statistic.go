package dto

// Statistic ...
type Statistic struct {
	Base

	Backups         int64 `json:"backups"`
	ContentTypes    int64 `json:"content_types"`
	Executables     int64 `json:"executables"`
	IntentTemplates int64 `json:"intent_templates"`
	Mediae          int64 `json:"mediae"`
	PlugIns         int64 `json:"plug_ins"`
	Projects        int64 `json:"projects"`
	Repositories    int64 `json:"repositories"`
}
