// Code generated by goctl. DO NOT EDIT.
package types

type FavoriteListRequest struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type VideoListResponse struct {
	StatusCode    int64        `json:"status_code"`
	StatusMsg     string       `json:"status_msg"`
	VideoInfoList []*VideoInfo `json:"video_list"`
}

type FavoriteRequest struct {
	Token      string `json:"token"`
	VideoId    int64  `json:"video_id"`
	ActionType int32  `json:"action_type"`
}

type FavoriteResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type VideoInfo struct {
	Id           int64    `json:"id"`
	Author       UserInfo `json:"author"`
	PlayUrl      string   `json:"play_url"`
	CoverUrl     string   `json:"cover_url"`
	FavoriteCunt int64    `json:"favorite_count"`
	CommentCount int64    `json:"comment_count"`
	IsFavorite   bool     `json:"is_favorite"`
	Title        string   `json:"title"`
}

type UserInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"gender"`
	FollowerCount int64  `json:"gender"`
	IsFollow      bool   `json:"gender"`
}
