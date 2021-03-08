package orm

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ORMInfo struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (i ORMInfo) String() string {
	return fmt.Sprintf("ORMInfo : %v:%v", i.Host, i.Port)
}

const (
	ormHost string = "20.0.1.8"
	ormPort int    = 5431
)

var (
	ormInfoPayload []byte
	ormError       error
)

func init() {
	ormInfo := ORMInfo{
		Host: ormHost,
		Port: ormPort,
	}
	log.Print(ormInfo)
	ormInfoPayload, ormError = json.Marshal(ormInfo)
	if ormError != nil {
		log.Fatalf("Failed to Marshall ORMInfo, reason : %v, ormError")
	}
}

func ConnHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(ormInfoPayload)
}
