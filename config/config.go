package config

const (
	DSN                = "root:123456@tcp(127.0.0.1:3306)/heart?charset=utf8mb4&parseTime=True"
	ServerPort         = ":5005"
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
	QiniuAccessKey     = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	QiniuSecretKey     = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	QiniuBucket        = "xxxxxxx"
	QiniuDomian        = "http://xxxx.xxxxt.clouddn.com"
	QiniuUrlExpire     = 60 * 60 * 24 * 2 // 2天
	RedisUrlExpireDiff = 60 * 60 * 2      //2小时 客户端拿到链接至少两小时内有效

	FEED_NUM = 20
)

var TEMP_FILE_DIR = ""
var RUNTIME_ENV = ""
