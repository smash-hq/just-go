package net_work

import (
	"fmt"
	log "github.com/corgi-kx/logcustom"
	"h7/tools"
	"html/template"
	"net/http"
)

// MyMux
// @Description: 创建一个实体
type MyMux struct {
}

// ServeHTTP
// @Description:
// @receiver receiver
// @param w
// @param r
func (receiver *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		func(w http.ResponseWriter, r *http.Request) {
			_, err := fmt.Fprintf(w, "Hello myroute!")
			if err != nil {
				return
			}
			return
		}(w, r)
		http.NotFound(w, r)
	}
	return
}

// Say
// @Description: 实现一个请求
// @receiver receiver
// @param w
// @param r
func (receiver MyMux) Say(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		log.Info("k", k)
		log.Info("v", v)
	}
	fmt.Fprintf(w, "hello! my hero!!!")
}

func (receiver MyMux) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		files, err := template.ParseFiles("login.gtpl")
		tools.CheckErr(err)
		log.Info(files.Execute(w, nil))
	} else {
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}
