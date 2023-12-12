package service

import (
	"github.com/google/wire"
	"github.com/synctv-org/vendors/service/alist"
	"github.com/synctv-org/vendors/service/bilibili"
	"github.com/synctv-org/vendors/service/emby"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(bilibili.NewBilibiliService, alist.NewAlistService, emby.NewEmbyService)
