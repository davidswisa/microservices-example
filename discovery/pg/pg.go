package pg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PGInfo struct {
	Port int    `json:"port"`
	DB   string `json:"db"`
	Host string `json:"host"`
	Pass string `json:"pass"`
	User string `json:"user"`
	SSL  string `json:"ssl"`
}

func (i PGInfo) String() string {
	url := fmt.Sprintf("PGInfo : %v:%v/%v?sslmode=%v", i.Host, i.Port, i.DB, i.SSL)
	crd := fmt.Sprintf("%v\\%v", i.User, i.Pass)
	return fmt.Sprintf("%v | %v", url, crd)
}

const (
	pgHost     string = "20.10.1.4"
	pgPort     int    = 5432
	pgUser     string = "postgres"
	pgPassword string = "postgres"
	pgDatabase string = "postgres"
	pgSslMode  string = "disable"
)

var (
	pgInfoPayload []byte
	pgError       error
)

func init() {
	pgInfo := PGInfo{
		DB:   pgDatabase,
		Host: pgHost,
		Pass: pgPassword,
		Port: pgPort,
		SSL:  pgSslMode,
		User: pgUser,
	}
	log.Print(pgInfo)
	pgInfoPayload, pgError = json.Marshal(pgInfo)
	if pgError != nil {
		log.Fatalf("Failed to Marshall PGInfo, reason : %v", pgError)
	}
}

func ConnHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(pgInfoPayload)
}
