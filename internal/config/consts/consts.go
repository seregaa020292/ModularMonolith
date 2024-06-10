package consts

import "time"

const (
	ServerPort = "8080"

	HttpRateRequestLimit = 100
	HttpRateWindowLength = 1 * time.Minute
)

var (
	SensitiveDataMask  = "*****"
	SensitiveFilerKeys = []string{"password"}
)
