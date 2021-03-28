package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	stringConexao := "LMnRRDPCp3:EfDnEPTKcX@tcp(remotemysql.com:3306)/LMnRRDPCp3?charset=utf8&parseTime=True"

	_, err := gorm.Open(mysql.Open(stringConexao), &gorm.Config{})

	if err != nil {
		panic("Não conectado ao banco de dados")
	}
}
