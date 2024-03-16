package envconfig

import "github.com/kelseyhightower/envconfig"

type Env struct {
	// environment variable for server TODO 慢慢增加
	Port string `envconfig:"PORT" default:"3000"`
}

func Init(env *Env) error {
	return envconfig.Process("treehole", env)
}
