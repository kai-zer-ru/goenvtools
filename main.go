package GoEnvTools

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type GoEnv struct {
	EnvFileName string
}

func (env *GoEnv) InitEnv() error {
	if env.EnvFileName == "" {
		env.EnvFileName = ".env"
	}
	return godotenv.Load(env.EnvFileName)
}

func (env *GoEnv) GetEnvString(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func (env *GoEnv) GetEnvUint32(key string, defaultValue uint32) uint32 {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return uint32(valInt)
}

func (env *GoEnv) GetEnvInt64(key string, defaultValue int64) int64 {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return int64(valInt)
}

func (env *GoEnv) GetEnvInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return valInt
}

func (env *GoEnv) GetEnvBool(key string, defaultValue bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return env.convertInterfaceToBool(val)
}

func (env *GoEnv) convertInterfaceToBool(value interface{}) bool {
	switch value.(type) {
	case string:
		v := value.(string)
		if v == "true" || v == "1" {
			return true
		} else {
			return false
		}
	case int64:
	case int32:
	case int16:
	case uint32:
	case uint64:
	case uint16:
		if value == 1 {
			return true
		} else {
			return false
		}
	case bool:
		return value.(bool)
	default:
		return false
	}
	return false
}
