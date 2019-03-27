package logic

import (
	. "db"

	"model"

	"golang.org/x/net/context"
)

type LearningMaterialLogic struct{}

var DefaultLearningMaterial = LearningMaterialLogic{}

func (LearningMaterialLogic) FindAll(ctx context.Context) []*model.LearningMaterial {
	objLog := GetLogger(ctx)

	learningMaterials := make([]*model.LearningMaterial, 0)
	err := MasterDB.Asc("type").Desc("seq").Find(&learningMaterials)
	if err != nil {
		objLog.Errorln("LearningMaterialLogic FindAll error:", err)
		return nil
	}

	return learningMaterials
}
