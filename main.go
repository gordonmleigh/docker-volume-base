package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/volume"
	"github.com/gordonmleigh/docker-volume-base/driver"
)

const (
	driverName = "docker_volume_base"
)

func main() {
	log.WithFields(log.Fields{"Name": driverName, "Pid": os.Getpid()}).Info("*** STARTED volume driver ***")

	d := driver.NewBaseDriver("/tmp")
	h := volume.NewHandler(d)
	err := h.ServeTCP(driverName, ":8080", nil)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Started.")
	}
}
