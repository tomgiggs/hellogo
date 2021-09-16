package logic

import (
	"context"

	"pinggo/internal/svc"
	"pinggo/internal/types"

	"gitlab.deepwisdomai.com/infra/go-zero/core/logx"
)

type DataRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) DataRecordLogic {
	return DataRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataRecordLogic) DataRecord() (*types.CommonResponse, error) {
	// todo: add your logic here and delete this line

	return &types.CommonResponse{}, nil
}
