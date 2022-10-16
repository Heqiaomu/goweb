package global

import (
	"gitee.com/goweb/config"
	"gitee.com/goweb/tools/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"gorm.io/gorm"
	"sync"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_CONFIG config.Server
	GVA_Timer               timer.Timer = timer.NewTimerTask()

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)