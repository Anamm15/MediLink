package constants

const (
	RedisKeyPrefix  = "medilink:%s"
	RedisKeyUser    = RedisKeyPrefix + "user:%s"
	RedisKeyPatient = RedisKeyPrefix + "patient:%s"
	RedisKeyDoctor  = RedisKeyPrefix + "doctor:%s"
	RedisKeyAdmin   = RedisKeyPrefix + "admin:%s"
)
