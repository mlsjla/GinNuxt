package service

import (
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	MenuSet,
	RoleSet,
	UserSet,
	LoginSet,
	ThreadSet,
	CategorySet,
	PostSet,
	AttachmentSet,
	SettingSet,
	MobileCodeSet,
	CasbinRuleSet,
	RoleMenuSet,
) // end
