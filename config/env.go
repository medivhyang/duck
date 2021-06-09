package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

func LoadEnv(key string, i interface{}, decode DecodeFunc) error {
	if decode == nil {
		return errors.New("require decode func")
	}
	s := os.Getenv(key)
	if s == "" {
		return nil
	}
	return decode([]byte(s), i)
}

func StoreEnv(key string, i interface{}, encode EncodeFunc) error {
	if encode == nil {
		return errors.New("require encode func")
	}
	bs, err := encode(i)
	if err != nil {
		return err
	}
	return os.Setenv(key, string(bs))
}

func LoadOrStoreEnv(key string, i interface{}, decode DecodeFunc, encode EncodeFunc) error {
	s, ok := os.LookupEnv(key)
	if !ok {
		bs, err := encode(i)
		if err != nil {
			return err
		}
		return os.Setenv(key, string(bs))
	}
	return decode([]byte(s), i)
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvOrDefault(key string, defaultValue ...string) string {
	s, ok := os.LookupEnv(key)
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return s
}

func SetEnv(key string, value string) error {
	return os.Setenv(key, value)
}

func GetIntEnv(key string) (int, error) {
	return strconv.Atoi(GetEnv(key))
}

func GetIntEnvOrDefault(key string, defaultValue ...int) (int, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return 0, nil
	}
	return strconv.Atoi(v)
}

func SetIntEnv(key string, value int) error {
	return os.Setenv(key, strconv.Itoa(value))
}

func GetInt64Env(key string) (int64, error) {
	return strconv.ParseInt(GetEnv(key), 10, 64)
}

func GetInt64EnvOrDefault(key string, defaultValue ...int64) (int64, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return 0, nil
	}
	return strconv.ParseInt(v, 10, 64)
}

func SetInt64Env(key string, value int64) error {
	return os.Setenv(key, strconv.FormatInt(value, 10))
}

func GetFloat64Env(key string) (float64, error) {
	return strconv.ParseFloat(GetEnv(key), 64)
}

func GetFloat64EnvOrDefault(key string, defaultValue ...float64) (float64, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return 0, nil
	}
	return strconv.ParseFloat(v, 64)
}

func SetFloat64Env(key string, value float64, precision int) error {
	return os.Setenv(key, strconv.FormatFloat(value, 'f', precision, 64))
}

func GetBoolEnv(key string) (bool, error) {
	return strconv.ParseBool(GetEnv(key))
}

func GetBoolEnvOrDefault(key string, defaultValue ...bool) (bool, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return false, nil
	}
	return strconv.ParseBool(GetEnv(v))
}

func SetBoolEnv(key string, value bool) error {
	return os.Setenv(key, strconv.FormatBool(value))
}

func GetTimeEnv(key string, layout string, location *time.Location) (time.Time, error) {
	if location == nil {
		location = time.Local
	}
	return time.ParseInLocation(layout, GetEnv(key), location)
}

func GetTimeOrDefault(key string, layout string, location *time.Location, defaultValue ...time.Time) (time.Time, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return time.Time{}, nil
	}
	if location == nil {
		location = time.Local
	}
	return time.ParseInLocation(layout, v, location)
}

func SetTimeEnv(key string, value time.Time, layout string) error {
	return os.Setenv(key, value.Format(layout))
}

func GetTimeDurationEnv(key string) (time.Duration, error) {
	return time.ParseDuration(GetEnv(key))
}

func GetTimeDurationOrDefault(key string, defaultValue ...time.Duration) (time.Duration, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return time.Duration(0), nil
	}
	return time.ParseDuration(v)
}

func SetTimeDurationEnv(key string, value time.Duration) error {
	return os.Setenv(key, value.String())
}
