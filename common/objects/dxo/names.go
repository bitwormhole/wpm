package dxo

// ExecutableGroupName ...
type ExecutableGroupName string

// ProjectGroupName ...
type ProjectGroupName string

// RepositoryGroupName ...
type RepositoryGroupName string

// UserName ...
type UserName string

// ContentTypeName ... like 'text/html' , 'project/java-pom'
type ContentTypeName string

////////////////////////////////////////////////////////////////////////////////

func (value ContentTypeName) String() string {
	return string(value)
}
