// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	commentFieldNames          = builder.RawFieldNames(&Comment{})
	commentRows                = strings.Join(commentFieldNames, ",")
	commentRowsExpectAutoSet   = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	commentRowsWithPlaceHolder = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheCommentIdPrefix = "cache:comment:id:"
)

type (
	commentModel interface {
		Insert(ctx context.Context, data *Comment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Comment, error)
		Update(ctx context.Context, newData *Comment) error
		Delete(ctx context.Context, id int64) error
		FindByVideoId(ctx context.Context, videoId int64) ([]Comment, error)
	}

	defaultCommentModel struct {
		sqlc.CachedConn
		table string
	}

	Comment struct {
		Id        int64        `db:"id"`
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
		UserId    int64        `db:"user_id"`
		VideoId   int64        `db:"video_id"`
		Content   string       `db:"content"`
	}
)

func newCommentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCommentModel {
	return &defaultCommentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`comment`",
	}
}

func (m *defaultCommentModel) Delete(ctx context.Context, id int64) error {
	commentIdKey := fmt.Sprintf("%s%v", cacheCommentIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, commentIdKey)
	return err
}

func (m *defaultCommentModel) FindOne(ctx context.Context, id int64) (*Comment, error) {
	commentIdKey := fmt.Sprintf("%s%v", cacheCommentIdPrefix, id)
	var resp Comment
	err := m.QueryRowCtx(ctx, &resp, commentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCommentModel) FindByVideoId(ctx context.Context, videoId int64) ([]Comment, error) {
	var comments []Comment
	err := m.QueryRowsNoCacheCtx(ctx, &comments,
		fmt.Sprintf("select %s from %s  where video_id = ? and deleted_at is null ", commentRows, m.table), videoId)
	switch err {
	case nil:
		return comments, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCommentModel) Insert(ctx context.Context, data *Comment) (sql.Result, error) {
	commentIdKey := fmt.Sprintf("%s%v", cacheCommentIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, commentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.UserId, data.VideoId, data.Content)
	}, commentIdKey)
	return ret, err
}

func (m *defaultCommentModel) Update(ctx context.Context, data *Comment) error {
	commentIdKey := fmt.Sprintf("%s%v", cacheCommentIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, commentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.UserId, data.VideoId, data.Content, data.Id)
	}, commentIdKey)
	return err
}

func (m *defaultCommentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCommentIdPrefix, primary)
}

func (m *defaultCommentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCommentModel) tableName() string {
	return m.table
}
