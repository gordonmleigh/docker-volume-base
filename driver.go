package main

import (
  "github.com/docker/go-plugins-helpers/volume"
  log "github.com/Sirupsen/logrus"
)

type BaseDriver struct {
  rootPath   string
}


func NewBaseDriver(rootPath string) *BaseDriver {
  log.WithFields(log.Fields{ "rootPath": rootPath }).Info("Creating new driver BaseDriver")

  d := BaseDriver{
    rootPath: rootPath,
  }
  return &d
}


func (d BaseDriver) Create(r volume.Request) volume.Response {
  log.WithFields(log.Fields{ "Name": r.Name, "Options": r.Options }).Info("REQUEST: Create volume")
  return volume.Response{}
}

func (d BaseDriver) List(r volume.Request) volume.Response {
  log.Info("REQUEST: List volumes")
  return volume.Response{}
}

func (d BaseDriver) Get(r volume.Request) volume.Response {
  log.WithFields(log.Fields{ "Name": r.Name }).Info("REQUEST: Get volume")
  return volume.Response{ Volume: &volume.Volume{ Name: r.Name }};
}

func (d BaseDriver) Remove(r volume.Request) volume.Response {
  log.WithFields(log.Fields{ "Name": r.Name }).Info("REQUEST: Remove volume")
  return volume.Response{}
}

func (d BaseDriver) Path(r volume.Request) volume.Response {
  log.WithFields(log.Fields{ "Name": r.Name }).Info("REQUEST: Path")
  return volume.Response{ Err: "not found" }
}

func (d BaseDriver) Mount(r volume.MountRequest) volume.Response {
  log.WithFields(log.Fields{ "Name": r.Name, "ID": r.ID }).Info("REQUEST: Mount")
  return volume.Response{ Err: "not found" }
}

func (d BaseDriver) Unmount(r volume.UnmountRequest) volume.Response {
  log.WithFields(log.Fields{ "Name": r.Name, "ID": r.ID }).Info("REQUEST: Unmount")
  return volume.Response{ Err: "not found" }
}

func (d BaseDriver) Capabilities(r volume.Request) volume.Response {
  log.Info("REQUEST: Capabilities")
  return volume.Response{}
}
