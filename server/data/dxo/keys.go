package dxo

// ProjectTypeKey ... like 'pom.xml'
type ProjectTypeKey string

////////////////////////////////////////////////////////////////////////////////

func (value ProjectTypeKey) String() string {
	return string(value)
}
