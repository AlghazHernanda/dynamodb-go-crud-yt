package env

import "os"

// get env
func GetEnv(env, defaultValue string) string {
	environment := os.Getenv(env)
	if environment == "" {
		//return default
		return defaultValue
	}
	//return environment
	return environment
}
