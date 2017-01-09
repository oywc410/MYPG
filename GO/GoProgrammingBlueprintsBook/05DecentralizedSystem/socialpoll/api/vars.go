package main

import (
	"sync"
	"net/http"
)

var (
	varsLook sync.RWMutex
	vars map[*http.Request]map[string]interface{}
)

func OpenVars(r *http.Request) {
	varsLook.Lock()
	if vars == nil {
		vars = map[*http.Request]map[string]interface{}{}
	}
	vars[r] = map[string]interface{}{}
	varsLook.Unlock()
}

func CloseVars(r *http.Request) {
	varsLook.Lock()
	delete(vars, r)
	varsLook.Unlock()
}

func GetVar(r *http.Request, key string) interface{} {
	varsLook.RLock()
	value := vars[r][key]
	varsLook.RUnlock()
	return value
}

func SetVar(r *http.Request, key string, value interface{}) {
	varsLook.Lock()
	vars[r][key] = value
	varsLook.Unlock()
}