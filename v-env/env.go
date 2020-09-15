package v_env

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
func GetStrEnv(key string) string {
	value := os.Getenv(key)
	value = strings.TrimSpace(value)
	if value == "" {
		log.Fatalf("env key \"%v\" not found", key)
	}
	return value
}


//noinspection GoUnusedExportedFunction
func GetIntEnv(key string) int {
	value := GetStrEnv(key)
	val, err := strconv.Atoi(value)
	if err != nil {
		err = errors.WithStack(err)
		log.Fatalf("env key \"%v\" is not int value\nerror:\n%+v", key, err)
	}
	return val
}


//noinspection GoUnusedExportedFunction
func GetBoolEnv(key string) bool {
	value := GetStrEnv(key)
	value = strings.ToLower(value)
	if value == "1" || value == "true" {
		return true
	}
	if value == "0" || value == "false" {
		return false
	}

	log.Fatalf("env value \"%v\", from key \"%v\" is not a valid boolean value", value, key)
	return false
}
