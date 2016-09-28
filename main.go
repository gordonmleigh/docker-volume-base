package main

import (
  "os"
  "github.com/docker/go-plugins-helpers/volume"
  log "github.com/Sirupsen/logrus"
)

const (
  driverName = "docker_volume_base"
)


func main() {
  log.WithFields(log.Fields{ "Name": driverName, "Pid": os.Getpid()}).Info("*** STARTED volume driver ***")

  d := NewBaseDriver("/tmp")
  h := volume.NewHandler(d)
  err := h.ServeUnix("root", driverName)

  if err != nil {
    log.Fatal(err)
  } else {
    log.Info("Started.")
  }
}
