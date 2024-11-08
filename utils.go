package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ParseConfig(filename string) (map[string]string, error) {
	var service = map[string]string{}
	var port_address string
	var port_name string
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Echec d'ouverture ", err)
	}
	defer file.Close()
	for ; err != io.EOF; _, err = fmt.Fscanf(file, "%s %s\n", &port_name, &port_address) {
		if err != nil {
			log.Fatal("Echec de lecture ", err)
		}
		service[port_name] = port_address
	}
	return service, nil
}
