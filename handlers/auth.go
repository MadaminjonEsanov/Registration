package handlers

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "registration/models"
    "registration/utils"
    //"github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func InitDB(connString string) {
    var err error
    db, err = pgxpool.Connect(context.Background(), connString)
    if err != nil {
        panic("failed to connect database")
    }
}

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)
    hashedPassword, _ := utils.HashPassword(user.Password)
    user.Password = hashedPassword

    query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
    _, err := db.Exec(context.Background(), query, user.Username, user.Password, user.Email)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error: %v", err)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var credentials models.User
    _ = json.NewDecoder(r.Body).Decode(&credentials)

    var user models.User
    query := `SELECT id, username, password, email FROM users WHERE username=$1`
    err := db.QueryRow(context.Background(), query, credentials.Username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        fmt.Fprintf(w, "Error: %v", err)
        return
    }

    if !utils.CheckPasswordHash(credentials.Password, user.Password) {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(user.Username)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
