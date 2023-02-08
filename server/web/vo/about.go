package vo

// About ...
type About struct {
	Base

	Name      string `json:"name"`
	Module    string `json:"module"`
	Title     string `json:"title"`
	Version   string `json:"version"`
	Revision  int    `json:"revision"`
	Copyright string `json:"copyright"`
	Profile   string `json:"profile"`
	User      string `json:"user"`
	Home      string `json:"home"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
}
