package logic

import (
	"model"

	"golang.org/x/net/context"

	. "db"

	"github.com/polaris1119/logger"
)

type DownloadLogic struct{}

var DefaultDownload = DownloadLogic{}

func (DownloadLogic) FindAll(ctx context.Context) []*model.Download {
	downloads := make([]*model.Download, 0)
	err := MasterDB.Desc("seq").Find(&downloads)
	if err != nil {
		logger.Errorln("DownloadLogic FindAll Error:", err)
	}

	return downloads
}
