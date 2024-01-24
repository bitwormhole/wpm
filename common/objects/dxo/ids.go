package dxo

import "strconv"

// ExampleID ...
type ExampleID int

////////////////////////////////////////////////////////////////////////////////

// BackupID ...
type BackupID int

// MIMETypeID ...
type MIMETypeID int

// ExecutableID ...
type ExecutableID int

// ExecutableGroupID ...
type ExecutableGroupID int

// InstallationID ...
type InstallationID int64

// IntentID ...
type IntentID int

// IntentQueueID ...
type IntentQueueID int64

// IntentTemplateID ...
type IntentTemplateID int

// MediaID ...
type MediaID int

// NamespaceID ...
type NamespaceID int

// LocationID ...
type LocationID int

// ProjectID ...
type ProjectID int

// ContentTypeID ... like '8'
type ContentTypeID int

// InstalledFileID ...
type InstalledFileID int

// ProjectGroupID ...
type ProjectGroupID int

// ProfileID ...
type ProfileID int

// PlatformID ...
type PlatformID int

// LocalRepositoryID ...
type LocalRepositoryID int

// RemoteRepositoryID ...
type RemoteRepositoryID int

// MainRepositoryID ...
type MainRepositoryID int

// RepositoryGroupID ...
type RepositoryGroupID int

// SettingID ...
type SettingID int

// SetupID ...
type SetupID int

// SoftwarePackageID ...
type SoftwarePackageID int

// SoftwareSetID ...
type SoftwareSetID int

// UserID ...
type UserID int

// WorktreeID ...
type WorktreeID int

////////////////////////////////////////////////////////////////////////////////

func (id IntentQueueID) String() string {
	n := int64(id)
	return strconv.FormatInt(n, 10)
}