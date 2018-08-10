package main

import (
	"os"
	"practice/mysqlBackup/common"
	"log"
)

// 入口函数
func main() {
	if len(os.Args)>1 {
		flag := common.OpFlag{
			Tables: false,
			Datum: false,
			Views: false,
			Funcs: false,
		}
		switch os.Args[1] {
		case "table":
			flag.Tables = true
		case "data":
			flag.Tables = true
			flag.Datum = true
		case "views":
			flag.Views = true
		case "funcs":
			flag.Funcs = true
		default:
			log.Fatal("Your arg must be in:table, data, views or funcs.");
		}
	}else{
		flag := common.OpFlag{
			Tables: true,
			Datum: true,
			Views: true,
			Funcs: true,
		}
	}

	err := backup.Export(flag);
}
