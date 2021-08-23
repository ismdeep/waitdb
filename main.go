package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/ismdeep/args"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // load mysql driver
	"time"
)

const helpMsg = `Usage: waitdb -dialect <DIALECT> -dsn <DSN>`

type Writer struct {
	gorm.LogWriter
}

func (receiver *Writer) Println(v ...interface{}) {
}

func main() {
	dialect := ""
	dsn := ""
	if args.Exists("-dialect") {
		dialect = args.GetValue("-dialect")
	}
	if args.Exists("-dsn") {
		dsn = args.GetValue("-dsn")
	}

	if dsn == "" || dialect == "" || args.Exists("--help") {
		fmt.Println(helpMsg)
		return
	}

	fmt.Print("Connecting ")

	for {
		_ = mysql.SetLogger(gorm.Logger{
			LogWriter: &Writer{},
		})
		if _, err := gorm.Open(dialect, dsn); err != nil {
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
			continue
		}
		break
	}
	fmt.Println()
}
