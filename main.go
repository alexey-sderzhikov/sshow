package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const CONNECTIONS = "connections.json"

type connection struct {
	Name    string
	Address string
}

func AddConnection(con connection) error {
	dataR, err := os.ReadFile(CONNECTIONS)
	if err != nil {
		log.Fatal(err)
	}

	conList := make([]connection, 0)
	if len(dataR) > 0 {
		err = json.Unmarshal(dataR, &conList)
		if err != nil {
			return err
		}
	}

	conList = append(conList, con)

	dataW, err := json.Marshal(conList)
	if err != nil {
		return err
	}

	err = os.WriteFile(CONNECTIONS, dataW, 0777)
	if err != nil {
		return err
	}

	return nil
}

func GetConnectionList() ([]connection, error) {
	data, err := os.ReadFile(CONNECTIONS)
	if err != nil {
		log.Fatal(err)
	}

	conList := make([]connection, 0)
	err = json.Unmarshal(data, &conList)
	if err != nil {
		return nil, err
	}

	return conList, nil
}

func main() {
	con := connection{
		Name:    "test",
		Address: "testAddress",
	}

	err := AddConnection(con)
	if err != nil {
		log.Fatal(err)
	}

	conList, err := GetConnectionList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(conList)
}
