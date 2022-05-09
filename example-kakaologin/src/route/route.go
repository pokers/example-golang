package route

import (
	handler "example-kakaologin/src/viewHandlers"
	"net/http"
)

func RegisterKakaoLogin() {
	http.HandleFunc("/login/", handler.BuildHandler(2, handler.KakaoLogin))
}
func RegisterAsserts() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
}
func RegisterKakaoAuth() {
	http.HandleFunc("/auth/", handler.BuildHandler(2, handler.KakaoAuth))
}
