package vo

// Init ...
type Init struct {
	Base

	// check-update
	CheckUpdate About

	// check-setup
	Setup Setup

	// fetch Content types
	ContentType ContentType

	// fetch executables
	Executable Executable

	// fetch options
	Option Option
}
