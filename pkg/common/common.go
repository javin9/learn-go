package common

import "github.com/rupid/learn-gin/pkg/settings"

func main() {
	settings.SetConfig()
	settings.Test()
}
