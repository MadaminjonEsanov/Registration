package config

import (
    "os"
    "github.com/joho/godotenv"
    "github.com/spf13/cast"
    "fmt"
)

type Config struct {
    HttpPort   string
    DBUser     string
    DBPassword string
    DBHost     string
    DBPort     int
    DBName     string
}

func NewConfig() *Config {
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }

    return &Config{
        HttpPort:   cast.ToString(getDefaultValue("HTTP_PORT", ":8080")),
        DBUser:     cast.ToString(getDefaultValue("DB_USER", "user")),
        DBPassword: cast.ToString(getDefaultValue("DB_PASSWORD", "password")),
        DBHost:     cast.ToString(getDefaultValue("DB_HOST", "localhost")),
        DBPort:     cast.ToInt(getDefaultValue("DB_PORT", 5432)),
        DBName:     cast.ToString(getDefaultValue("DB_NAME", "test_db")),
    }
}

func getDefaultValue(key string, defaultValue interface{}) interface{} {
    val, exists := os.LookupEnv(key)
    if exists {
        return val
    }
    return defaultValue
}
