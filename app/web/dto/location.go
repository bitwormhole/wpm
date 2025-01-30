package dto

import "github.com/bitwormhole/wpm/app/data/dxo"

// Location ...
type Location struct {
	ID dxo.LocationID `json:"id"`

	Base

	Label       string `json:"label"`
	Description string `json:"description"`
	Tags        string `json:"tags"` // 以逗号分隔的一组标签

	AsFile   bool `json:"as_file"`
	AsFolder bool `json:"as_folder"`

	Category   dxo.Category  `json:"category"`
	MediaType  dxo.MediaType `json:"media_type"`
	RawURI     dxo.URI       `json:"raw_uri"`
	RegularURI dxo.URI       `json:"regular_uri"`
}
