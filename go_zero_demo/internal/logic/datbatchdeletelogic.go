package logic

import (
	"context"

	"pinggo/internal/svc"
	"pinggo/internal/types"

	"gitlab.deepwisdomai.com/infra/go-zero/core/logx"
)

type DatBatchDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDatBatchDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DatBatchDeleteLogic {
	return DatBatchDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DatBatchDeleteLogic) DatBatchDelete(req types.TrainDataBatchRequest) (*types.CommonResponse, error) {
	// todo: add your logic here and delete this line

	return &types.CommonResponse{}, nil
}
