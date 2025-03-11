package utils

import (
	"encoding/base64"
	"fmt"
	"primeskills-test-api/pkg/config"
)

func GenMidtransAccessToken(cfg config.Config) string {
	key := fmt.Sprintf("%s:", cfg.Midtrans.ServerKey)
	return base64.StdEncoding.EncodeToString([]byte(key))
}
