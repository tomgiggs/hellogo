package logic

import (
	"context"

	"pinggo/internal/svc"
	"pinggo/internal/types"

	"gitlab.deepwisdomai.com/infra/go-zero/core/logx"
)

type DataBatchUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataBatchUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) DataBatchUploadLogic {
	return DataBatchUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataBatchUploadLogic) DataBatchUpload(req types.TrainDataBatchRequest) (*types.CommonResponse, error) {
	// todo: add your logic here and delete this line

	return &types.CommonResponse{}, nil
}
