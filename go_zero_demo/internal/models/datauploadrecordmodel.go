package model

import (
	"database/sql"
	"fmt"
	"strings"

	"gitlab.deepwisdomai.com/infra/go-zero/core/stores/sqlc"
	"gitlab.deepwisdomai.com/infra/go-zero/core/stores/sqlx"
	"gitlab.deepwisdomai.com/infra/go-zero/core/stringx"
	"gitlab.deepwisdomai.com/infra/go-zero/tools/goctl/model/sql/builderx"
)

var (
	dataUploadRecordFieldNames          = builderx.RawFieldNames(&DataUploadRecord{})
	dataUploadRecordRows                = strings.Join(dataUploadRecordFieldNames, ",")
	dataUploadRecordRowsExpectAutoSet   = strings.Join(stringx.Remove(dataUploadRecordFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	dataUploadRecordRowsWithPlaceHolder = strings.Join(stringx.Remove(dataUploadRecordFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	DataUploadRecordModel interface {
		Insert(data DataUploadRecord) (sql.Result, error)
		FindOne(id int64) (*DataUploadRecord, error)
		Update(data DataUploadRecord) error
		Delete(id int64) error
	}

	defaultDataUploadRecordModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DataUploadRecord struct {
		Id        int64        `db:"id"`
		AppId     string       `db:"app_id"`     // appid
		Total     int64        `db:"total"`      // 总条数
		Valid     int64        `db:"valid"`      // 有效条数
		Size      int64        `db:"size"`       // 数据大小
		Status    int64        `db:"status"`     // 数据集状态
		CreatedAt sql.NullTime `db:"created_at"` // 创建时间
		UpdatedAt sql.NullTime `db:"updated_at"` // 更新时间
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
	}
)

func NewDataUploadRecordModel(conn sqlx.SqlConn) DataUploadRecordModel {
	return &defaultDataUploadRecordModel{
		conn:  conn,
		table: "`data_upload_record`",
	}
}

func (m *defaultDataUploadRecordModel) Insert(data DataUploadRecord) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, dataUploadRecordRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.AppId, data.Total, data.Valid, data.Size, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	return ret, err
}

func (m *defaultDataUploadRecordModel) FindOne(id int64) (*DataUploadRecord, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dataUploadRecordRows, m.table)
	var resp DataUploadRecord
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDataUploadRecordModel) Update(data DataUploadRecord) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dataUploadRecordRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.AppId, data.Total, data.Valid, data.Size, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	return err
}

func (m *defaultDataUploadRecordModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
