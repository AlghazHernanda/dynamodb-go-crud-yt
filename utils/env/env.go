package env

import "os"

// get env
func GetEnv(env, defaultValue string) string {
	environment := os.Getenv(env)
	if environment == "" {
		return defaultValue
	}
	return environment
}
