// Code generated by goctl. DO NOT EDIT.
package types

type PublishRequest struct {
	Token    string `json:"token"`
	Title    string `json:"title"`
	ID       int64  `json:"id,optional"`
	PlayUrl  string `json:"play_url,optional"`
	CoverUrl string `json:"cover_url,optional"`
	Hash     string `json:"hash,optional"`
}

type PublishResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FeedRequest struct {
	LastTime string `form:"latest_time,optional"`
	Token    string `form:"token,optional"`
}

type FeedResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	NextTime   string  `json:"next_time"`
	VideoList  []Video `json:"video_list"`
}

type Video struct {
	Id            int64  `json:"id"`
	Author        Author `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type Author struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type PublishListRequest struct {
	Token  string `form:"token"`
	UserId string `form:"user_id"`
}

type PublishListResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}
