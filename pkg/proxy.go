package pkg

import "fmt"

type ProxyDatabase struct {
	Users map[string]bool
	DB    *Database
}

func (d *ProxyDatabase) GetData(user string) ([]string, error) {
	if !d.Users[user] {
		return nil, fmt.Errorf("user %s is not allowed to access the data", user)
	}
	return d.DB.GetData(user)
}
