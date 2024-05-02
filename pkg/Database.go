package pkg

type Database struct {
	//
}

func (d *Database) GetData(user string) ([]string, error) {
	return []string{"data1", "data2"}, nil
}
