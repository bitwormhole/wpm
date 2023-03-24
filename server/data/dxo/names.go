package dxo

// ExecutableGroupName ...
type ExecutableGroupName string

// ProjectGroupName ...
type ProjectGroupName string

// RepositoryGroupName ...
type RepositoryGroupName string

// UserName ...
type UserName string

// ProjectTypeName ... like 'java/pom'
type ProjectTypeName string

////////////////////////////////////////////////////////////////////////////////

func (value ProjectTypeName) String() string {
	return string(value)
}
