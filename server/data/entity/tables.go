package entity

////////////////////////////////////////////////////////////////////////////////

func getTableName(simpleName string) string {
	return "wpm_" + simpleName
}

////////////////////////////////////////////////////////////////////////////////

// ListAllTypes ...
func ListAllTypes() []any {

	list := make([]any, 0)

	list = append(list, new(ContentType))
	list = append(list, new(Executable))
	list = append(list, new(IntentTemplate))
	list = append(list, new(Location))
	list = append(list, new(Media))
	list = append(list, new(Project))
	list = append(list, new(LocalRepository))
	list = append(list, new(RemoteRepository))
	list = append(list, new(Setting))
	list = append(list, new(SoftwarePackage))
	list = append(list, new(InstalledFile))
	list = append(list, new(Namespace))
	list = append(list, new(SystemConfig))
	list = append(list, new(UserConfig))
	list = append(list, new(User))
	list = append(list, new(TreeRoot))

	return list
}

////////////////////////////////////////////////////////////////////////////////
