package logic

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.deepwisdomai.com/infra/go-zero/core/stores/sqlx"
	model "pinggo/internal/models"
	"time"

	"gitlab.deepwisdomai.com/infra/go-zero/core/logx"
	"pinggo/internal/svc"
)

type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) HealthCheckLogic {
	return HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthCheckLogic) HealthCheck() error {
	// todo: add your logic here and delete this line
	conn := sqlx.NewMysql("root:cyl123@tcp(localhost:3306)/pingserver?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	m := model.NewDataUploadRecordModel(conn)
	r, err := m.Insert(model.DataUploadRecord{
		Id:     0,
		AppId:  "1000",
		Total:  30,
		Valid:  11,
		Size:   12852,
		Status: 1,
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		logx.Errorf("insert into db error:%v", err)
		return err
	}
	fmt.Println(r.LastInsertId())
	return nil
}
