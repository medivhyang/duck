package config

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (*Env) Get(key string) string {
	return GetEnv(key)
}

func (*Env) GetInt(key string) (int, error) {
	return GetIntEnv(key)
}

func (*Env) GetInt64(key string) (int64, error) {
	return GetInt64Env(key)
}

func (*Env) GetFloat64(key string) (float64, error) {
	return GetFloat64Env(key)
}

func (*Env) GetBool(key string) (bool, error) {
	return GetBoolEnv(key)
}

func (*Env) GetOrDefault(key string, defaultValue ...string) string {
	return GetEnvOrDefault(key, defaultValue...)
}

func (*Env) GetIntOrDefault(key string, defaultValue ...int) (int, error) {
	return GetIntEnvOrDefault(key, defaultValue...)
}

func (*Env) GetInt64OrDefault(key string, defaultValue ...int64) (int64, error) {
	return GetInt64EnvOrDefault(key, defaultValue...)
}

func (*Env) GetFloat64OrDefault(key string, defaultValue ...float64) (float64, error) {
	return GetFloat64EnvOrDefault(key, defaultValue...)
}

func (*Env) GetBoolOrDefault(key string, defaultValue ...bool) (bool, error) {
	return GetBoolEnvOrDefault(key, defaultValue...)
}

func (*Env) Load(contentType ContentType, key string, target interface{}) error {
	return LoadEnv(contentType, key, target)
}
