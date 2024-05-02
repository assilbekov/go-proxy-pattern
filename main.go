package main

import (
	"fmt"
	"proxy/pkg"
)

var (
	admin = "admin"
	user  = "user"
	users = map[string]bool{
		admin: true,
		user:  false,
	}
)

func main() {
	db := &pkg.Database{}
	proxy := &pkg.ProxyDatabase{
		Users: users,
		DB:    db,
	}

	data, err := proxy.GetData(admin)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

	data, err = proxy.GetData(user)
	if err != nil {
		fmt.Println(err)
		return
	}
}
