package config

import (
	"log"
    "testing"
)

func TestGetDB (t *testing.T ){
	db, err := GetDB()
	if err != nil{
		t.Error("Error")
	}
	log.Println(db)
}