package logic_test

import (
	. "github.com/polaris1119/config"
	"github.com/polaris1119/logger"

	"logic"
	"testing"
)

func TestSendMail(t *testing.T) {
	logger.Init(ROOT+"/log", ConfigFile.MustValue("global", "log_level", "DEBUG"))

	err := logic.DefaultEmail.SendMail("中文test", "内容test content，收到？", []string{"xuxinhua@zhimadj.com"})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("successful")
	}
}
