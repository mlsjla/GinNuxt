package api

import "github.com/google/wire"

var APISet = wire.NewSet(
	UploadSet,
	LoginSet,
	MenuSet,
	RoleSet,
	UserSet,
	ThreadSet,
	CategorySet,
	PostSet,
	AttachmentSet,
	SettingSet,
	MobileCodeSet,
	CasbinRuleSet,
	RoleMenuSet,
) // end
