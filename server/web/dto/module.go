package dto

// Module 表示一个命令模板
type Module struct {
	// ID dxo.ExampleID `json:"id"`
	// Base

	Name     string `json:"name"`
	Version  string `json:"version"`
	Revision int    `json:"revision"`
	HexName  string `json:"hex_name"`
}
