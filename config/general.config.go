package config

import "time"

var NOW = time.Now().UTC()
var NOWF = time.Now().UTC().Format("2006-01-02 15:04:05")