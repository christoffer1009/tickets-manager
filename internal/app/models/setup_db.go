// models/setupdb.go
package models

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB() (*gorm.DB, error) {
	// Configuração do banco de dados usando Viper para ler do arquivo de configuração
	viper.SetConfigName("database_config")
	viper.AddConfigPath("./../../configs/")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo de configuração: %w", err)
	}

	// Configurar informações de conexão com o banco de dados
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// Abrir conexão com o banco de dados
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	// Migrar esquema do banco de dados (criar tabelas, etc.)
	if err := db.AutoMigrate(&Tecnico{}, &Cliente{}, &Ticket{}); err != nil {
		return nil, fmt.Errorf("erro ao migrar esquema do banco de dados: %w", err)
	}

	return db, nil
}
