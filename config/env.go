package config

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (*Env) Get(key string) string {
	return GetEnv(key)
}

func (*Env) GetOrDefault(key string, defaultValue ...string) string {
	return GetEnvOrDefault(key, defaultValue...)
}

func (*Env) Load(contentType ContentType, key string, target interface{}) error {
	return LoadEnv(contentType, key, target)
}
