package config

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (*Env) Get(key string) string {
	return GetEnv(key)
}

func (*Env) GetInt(key string) (int, error) {
	return GetEnvInt(key)
}

func (*Env) GetInt64(key string) (int64, error) {
	return GetEnvInt64(key)
}

func (*Env) GetFloat64(key string) (float64, error) {
	return GetEnvFloat64(key)
}

func (*Env) GetBool(key string) (bool, error) {
	return GetEnvBool(key)
}

func (*Env) GetOrDefault(key string, defaultValue ...string) string {
	return GetEnvOrDefault(key, defaultValue...)
}

func (*Env) GetIntOrDefault(key string, defaultValue ...int) (int, error) {
	return GetEnvIntOrDefault(key, defaultValue...)
}

func (*Env) GetInt64OrDefault(key string, defaultValue ...int64) (int64, error) {
	return GetEnvInt64OrDefault(key, defaultValue...)
}

func (*Env) GetFloat64OrDefault(key string, defaultValue ...float64) (float64, error) {
	return GetEnvFloat64OrDefault(key, defaultValue...)
}

func (*Env) GetBoolOrDefault(key string, defaultValue ...bool) (bool, error) {
	return GetEnvBoolOrDefault(key, defaultValue...)
}

func (*Env) Load(contentType ContentType, key string, target interface{}) error {
	return LoadEnv(contentType, key, target)
}
