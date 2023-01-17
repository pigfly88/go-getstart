package basic

import (
	"fmt"
	"testing"
)

type Driver interface {
	connect()
}

type Mysql struct {
	connConfig string
}

func (v Mysql) connect() {
	fmt.Println("connecting to mysql.")
}

func conn(d Driver) {
	d.connect()
}

func TestMysql(t *testing.T) {
	var driver Driver = Mysql{"127.0.0.1:3306"}
	conn(driver)
}
