package config

const (
	DSN                = "root:123456@tcp(127.0.0.1:3306)/heart?charset=utf8mb4&parseTime=True"
	ServerPort         = ":8088"
	Redis_addr         = "127.0.0.1:6379"
	Redis_password     = ""
	Redis_db           = 0
	Redis_ttl          = 60 * 60 * 24 * 30 // 30天
	Redis_refresh      = 60 * 60 * 24 * 15 // 15天
	UserTableName      = "t_user"
	VideoTableName     = "t_video"
	FavoriteTableName  = "t_favorite"
	CommentTableName   = "t_comment"
	RelationTableName  = "t_relation"
	QiniuAccessKey     = "EVmS6sY7STDCwq9Iw79geJEJalk2h0k4ql4r0V4s"
	QiniuSecretKey     = "Pp_AtsasRqDwfE6ClVIKlkKsNesWmHjK0Zqu3ouU"
	QiniuBucket        = "douyin-heart"
	QiniuDomian        = "http://rcnco1agb.hd-bkt.clouddn.com"
	QiniuUrlExpire     = 60 * 60 * 24 * 2 // 2天
	RedisUrlExpireDiff = 60 * 60 * 2      //2小时 客户端拿到链接至少两小时内有效

	FEED_NUM = 20
)

var TEMP_FILE_DIR = ""
