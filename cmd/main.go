package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/astaxie/session"
)

var (
	// セッション情報を保存するためのmap
	sessions = make(map[string]bool)
	// セッション情報へのアクセスを同期するためのMutex
	sessionMutex   = &sync.Mutex{}
	globalSessions *session.Manager
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func loginHandler(responseWriter http.ResponseWriter, req *http.Request) {

	sessionID := "user123"

	sessionMutex.Lock()
	sessions[sessionID] = true

	for id := range sessions {
		fmt.Println("loginHandler: Currenct sessionID: ", id)
	}

	sessionMutex.Unlock()

	http.SetCookie(responseWriter, &http.Cookie{
		Name:   "session_id",
		Value:  sessionID,
		Path:   "/",
		MaxAge: 60,
	})

	_, err := responseWriter.Write([]byte("login successfully!"))
	if err != nil {
		panic(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	sessionMutex.Lock()
	_, ok := sessions[cookie.Value]
	sessionMutex.Unlock()

	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	_, err = w.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}
