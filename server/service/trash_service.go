package service

// TrashService ...
type TrashService interface {
	OnInsert()

	OnDelete()

	EnableAutoCleanBeforeInsert(en bool)

	Clean() error
}
