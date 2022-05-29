package config

const (
	DSN               = "heartdancer:Heartdancer123.@tcp(47.98.120.35:3306)/heart?charset=utf8&parseTime=True"
	ServerPort        = ":8088"
	Redis_addr        = "47.98.120.35:6379"
	Redis_password    = "heartdancer"
	Redis_db          = 0
	Redis_ttl         = 60 * 60 * 24 * 30 // 30天
	Redis_reflash     = 60 * 60 * 24 * 15 // 15天
	UserTableName     = "t_user"
	VideoTableName    = "t_video"
	FavoriteTableName = "t_favorite"
	CommentTableName  = "t_comment"
	RelationTableName = "t_relation"
	QiniuAccessKey    = "EVmS6sY7STDCwq9Iw79geJEJalk2h0k4ql4r0V4s"
	QiniuSecretKey    = "Pp_AtsasRqDwfE6ClVIKlkKsNesWmHjK0Zqu3ouU"
)
