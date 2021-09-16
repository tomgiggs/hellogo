package logic

import (
	"context"

	"gitlab.deepwisdomai.com/infra/go-zero/core/logx"
	"pinggo/internal/svc"
)

type MetricsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMetricsLogic(ctx context.Context, svcCtx *svc.ServiceContext) MetricsLogic {
	return MetricsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MetricsLogic) Metrics() error {
	// todo: add your logic here and delete this line

	return nil
}
