package global

import (
	"github.com/kalifun/go-programming-tour-book-blog/pkg/logger"
	"github.com/kalifun/go-programming-tour-book-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
