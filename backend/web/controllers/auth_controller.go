package controllers

//
//import (
//	"encoding/base64"
//	"net/http"
//	"strings"
//)
//
//const (
//	username = "admin"
//	password = "admin"
//)
//
//func BasicAuthMiddleware(w http.ResponseWriter, request *http.Request) {
//
//	authHeader := request.Header.Get("Authorization")
//	if authHeader == "" {
//		unauthorized(w)
//		return
//	}
//
//	auth := strings.SplitN(authHeader, " ", 2)
//	if len(auth) != 2 || auth[0] != "Basic" {
//		unauthorized(w)
//		return
//	}
//
//	payload, err := base64.StdEncoding.DecodeString(auth[1])
//	if err != nil {
//		unauthorized(w)
//		return
//	}
//
//	pair := strings.SplitN(string(payload), ":", 2)
//	if len(pair) != 2 || pair[0] != username || pair[1] != password {
//		unauthorized(w)
//		return
//	}
//
//	w.Write([]byte("You are authorized!"))
//
//}
//
//func unauthorized(w http.ResponseWriter) {
//	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
//	http.Error(w, "Unauthorized", http.StatusUnauthorized)
//}
