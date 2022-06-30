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
	videoFieldNames          = builder.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheVideoIdPrefix = "cache:video:id:"
)

type (
	videoModel interface {
		Insert(ctx context.Context, data *Video) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Video, error)
		Update(ctx context.Context, newData *Video) error
		Delete(ctx context.Context, id int64) error
		GetVideoList(ctx context.Context, timeStamp int64) ([]*Video, error)
		GetVideoListByUserId(ctx context.Context, userId int64) ([]*Video, error)
	}

	defaultVideoModel struct {
		sqlc.CachedConn
		table string
	}

	Video struct {
		Id            int64        `db:"id"`
		CreatedAt     sql.NullTime `db:"created_at"`
		UpdatedAt     sql.NullTime `db:"updated_at"`
		DeletedAt     sql.NullTime `db:"deleted_at"`
		AuthorId      int64        `db:"author_id"`
		PlayUrl       string       `db:"play_url"`
		CoverUrl      string       `db:"cover_url"`
		FavoriteCount int64        `db:"favorite_count"`
		CommentCount  int64        `db:"comment_count"`
		PublishTime   time.Time    `db:"publish_time"`
		Title         string       `db:"title"`
	}
)

func newVideoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultVideoModel {
	return &defaultVideoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) Delete(ctx context.Context, id int64) error {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, videoIdKey)
	return err
}

func (m *defaultVideoModel) FindOne(ctx context.Context, id int64) (*Video, error) {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, id)
	var resp Video
	err := m.QueryRowCtx(ctx, &resp, videoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
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

func (m *defaultVideoModel) GetVideoList(ctx context.Context, timeStamp int64) ([]*Video, error) {
	var videos []*Video
	err := m.QueryRowsNoCacheCtx(ctx, &videos,
		fmt.Sprintf("select %s from %s where `publish_time` < ? or 1 =1 order by publish_time desc limit 30  ", videoRows, m.table), timeStamp)
	switch err {
	case nil:
		return videos, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) GetVideoListByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	var videos []*Video
	err := m.QueryRowsNoCacheCtx(ctx, &videos,
		fmt.Sprintf("select %s from %s where `author_id` = ? order by publish_time desc ", videoRows, m.table), userId)
	switch err {
	case nil:
		return videos, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) Insert(ctx context.Context, data *Video) (sql.Result, error) {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.AuthorId, data.PlayUrl, data.CoverUrl, data.FavoriteCount, data.CommentCount, data.PublishTime, data.Title)
	}, videoIdKey)
	return ret, err
}

func (m *defaultVideoModel) Update(ctx context.Context, data *Video) error {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.AuthorId, data.PlayUrl, data.CoverUrl, data.FavoriteCount, data.CommentCount, data.PublishTime, data.Title, data.Id)
	}, videoIdKey)
	return err
}

func (m *defaultVideoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheVideoIdPrefix, primary)
}

func (m *defaultVideoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideoModel) tableName() string {
	return m.table
}
