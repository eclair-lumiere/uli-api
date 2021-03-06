package serializer

import "go-crud/model"

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	Avatar    string `json:"avatar"`
	View	  uint64 `json:"view"`
	User	  User	 `json:"user"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	user,_:=model.GetUser(item.UserID)
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:	   item.VideoURL(),
		Avatar:	   item.AvatarURL(),
		User:	   BuildUser(user),
		View:	   item.View(),

		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) []Video {
	var videos []Video
	for _,item:=range items{
		video:=BuildVideo(item)
		videos=append(videos,video)
	}
	return videos
}
