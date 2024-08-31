package main

import "github.com/haodam/user-backend-golang/pkg/transports/https/routers"

func main() {

	r := routers.NewRouter()
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
