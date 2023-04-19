package vo

// Option ...
type Option struct {
	Base

	Options map[string]bool `json:"options"`
}
