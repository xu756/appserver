package config

import (
	"testing"
)

func Test(t *testing.T) {
	Init("/Users/xu756/开发/小程序/appServer/configs/dev.yaml")
	t.Log(RunData)
}
