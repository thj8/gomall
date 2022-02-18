package main

import (
	"fmt"
	"net/http"
)

func main() {
	//h := thj(http.HandlerFunc(indexhanlder))
	//h = thj2(h)
	//mux := http.NewServeMux()
	//mux.Handle("/", h)
	//http.ListenAndServe(":1234", mux)

	http.HandleFunc("/", indexhanlder)
	http.Handle("/thj", http.HandlerFunc(indexhanlder))
	http.ListenAndServe(":1234", nil)
}

func indexhanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld")
}

//
//func thj(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// do something
//		fmt.Println("do something...before.........")
//
//		ctx := r.Context()
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
//
//func thj2(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// do something
//		fmt.Println("do something2...before.........")
//
//		ctx := r.Context()
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
