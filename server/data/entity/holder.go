package entity

// Adapter 用来适配各种不同类型的 entity
type Adapter struct {
	ContentType    *ContentType
	Executable     *Executable
	IntentTemplate *IntentTemplate
	Media          *Media
	Package        *SoftwarePackage
	Project        *Project
	Repository     *LocalRepository
	Remote         *RemoteRepository
	Setting        *Setting
	Source         *Namespace
	User           *User
	Worktree       *Worktree
}

////////////////////////////////////////////////////////////////////////////////

// Holder 用来包装各种不同类型的 entity
type Holder struct {
	ID        int64
	TableName string
	Meta      Base
	Target    *Adapter
}

// // Get 取 target, 以 interface {} 的形式返回
// func (inst *Holder) Get() interface{} {
// 	t := inst.Target
// 	if t == nil {
// 		return nil
// 	}
// 	list := make([]any, 0)

// 	list = append(list, t.ContentType)
// 	list = append(list, t.Executable)
// 	list = append(list, t.IntentTemplate)
// 	list = append(list, t.Media)
// 	list = append(list, t.Package)
// 	list = append(list, t.Project)
// 	list = append(list, t.Repository)
// 	list = append(list, t.Remote)
// 	list = append(list, t.Setting)
// 	list = append(list, t.Source)
// 	list = append(list, t.User)
// 	list = append(list, t.Worktree)

// 	for _, item := range list {
// 		if !utils.InterfaceIsNil(item) {
// 			return item
// 		}
// 	}
// 	return nil
// }

// // SetContentType ...
// func (inst *Holder) SetContentType(o *ContentType) {
// 	inst.ID = int64(o.ID)
// 	inst.Meta = o.Base
// 	inst.TableName = o.TableName()
// 	inst.Target = &Adapter{
// 		ContentType: o,
// 	}
// }

// // SetMedia ...
// func (inst *Holder) SetMedia(o *Media) {
// 	inst.ID = int64(o.ID)
// 	inst.Meta = o.Base
// 	inst.TableName = o.TableName()
// 	inst.Target = &Adapter{
// 		Media: o,
// 	}
// }

// // SetSoftwarePackage ...
// func (inst *Holder) SetSoftwarePackage(o *SoftwarePackage) {
// 	inst.ID = int64(o.ID)
// 	inst.Meta = o.Base
// 	inst.TableName = o.TableName()
// 	inst.Target = &Adapter{
// 		Package: o,
// 	}
// }

// // SetProject ...
// func (inst *Holder) SetProject(o *Project) {
// 	inst.ID = int64(o.ID)
// 	inst.Meta = o.Base
// 	inst.TableName = o.TableName()
// 	inst.Target = &Adapter{
// 		Project: o,
// 	}
// }

// // SetLocalRepository ...
// func (inst *Holder) SetLocalRepository(o *LocalRepository) {
// 	inst.ID = int64(o.ID)
// 	inst.Meta = o.Base
// 	inst.TableName = o.TableName()
// 	inst.Target = &Adapter{
// 		Repository: o,
// 	}
// }

// // SetRemoteRepository ...
// func (inst *Holder) SetRemoteRepository(o *RemoteRepository) {
// 	inst.ID = int64(o.ID)
// 	inst.Meta = o.Base
// 	inst.TableName = o.TableName()
// 	inst.Target = &Adapter{
// 		Remote: o,
// 	}
// }

// // SetExecutable ...
// func (inst *Holder) SetExecutable(o *Executable) {
// 	inst.ID = int64(o.ID)
// 	inst.Meta = o.Base
// 	inst.TableName = o.TableName()
// 	inst.Target = &Adapter{
// 		Executable: o,
// 	}
// }

// // SetIntentTemplate ...
// func (inst *Holder) SetIntentTemplate(o *IntentTemplate) {
// 	inst.ID = int64(o.ID)
// 	inst.Meta = o.Base
// 	inst.TableName = o.TableName()
// 	inst.Target = &Adapter{
// 		IntentTemplate: o,
// 	}
// }
