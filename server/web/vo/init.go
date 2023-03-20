package vo

// Init ...
type Init struct {
	Base

	// check-update
	CheckUpdate About

	// check-setup
	Setup Setup

	// fetch project types
	ProjectType ProjectType

	// fetch executables
	Executable Executable
}
