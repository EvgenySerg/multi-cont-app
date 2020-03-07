package storage

type RedisRepo struct {
	url string
	user string
	password string
}

func NewRedisRepo(url string, user string, password string) *RedisRepo{
	return &RedisRepo{
		url:      url,
		user:     user,
		password: password,
	}
}

func (r *RedisRepo) Save(){

}