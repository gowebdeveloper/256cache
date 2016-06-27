package main

import (
	"net/http"
	"io/ioutil"
)

func mapdoing(w http.ResponseWriter, r *http.Request) {

	result:=""
	switch r.Method {
	case "GET":
		if ob,ok:=db[r.URL.Path[1:]];ok{
			result=ob
		}else {
			w.WriteHeader(400)
			w.Write([]byte("null"))
			return
		}

	case "PUT":
		bs, err:=ioutil.ReadAll(r.Body)
		if(err!=nil){
			w.Header().Del("Content-Length")
			w.Header().Del("Content-Type")
			w.Header().Del("Date")
			w.WriteHeader(400)
			w.Write([]byte("need body"))
			return
		}else {
			if _,ok:=db[r.URL.Path[1:]];ok{
				result="replace"
			}else {
				result="insert"
			}
			db[r.URL.Path[1:]] = string(bs)
		}
	}

	w.WriteHeader(200)
	w.Write([]byte(result))


}
