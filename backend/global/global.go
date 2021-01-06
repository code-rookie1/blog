package global

import (
	"blog/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)
var (
	GVB_DB      *gorm.DB
	GVB_CONFIG   config.Server
	GVB_VP       *viper.Viper
)
