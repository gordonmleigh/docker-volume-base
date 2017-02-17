package driver

import (
	log "github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/volume"
)

// BaseDriver represents a simple volume plugin driver.
type BaseDriver struct {
	rootPath string
}

// NewBaseDriver creates a new instance of BaseDriver.
func NewBaseDriver(rootPath string) *BaseDriver {
	log.WithFields(log.Fields{"rootPath": rootPath}).Info("Creating new driver BaseDriver")

	d := BaseDriver{
		rootPath: rootPath,
	}
	return &d
}

// Create creates a new volume.
func (d *BaseDriver) Create(r volume.Request) volume.Response {
	log.WithFields(log.Fields{"Name": r.Name, "Options": r.Options}).Info("REQUEST: Create volume")
	return volume.Response{}
}

// List lists all volumes the driver knows of.
func (d *BaseDriver) List(r volume.Request) volume.Response {
	log.Info("REQUEST: List volumes")
	return volume.Response{}
}

// Get gets a specific volume.
func (d *BaseDriver) Get(r volume.Request) volume.Response {
	log.WithFields(log.Fields{"Name": r.Name}).Info("REQUEST: Get volume")
	if r.Name == "bad" {
		return volume.Response{Err: "not found"}
	}
	return volume.Response{Volume: &volume.Volume{Name: r.Name}}
}

// Remove deletes a specific volume.
func (d *BaseDriver) Remove(r volume.Request) volume.Response {
	log.WithFields(log.Fields{"Name": r.Name}).Info("REQUEST: Remove volume")
	return volume.Response{}
}

// Path gets the path of a given volume.
func (d *BaseDriver) Path(r volume.Request) volume.Response {
	log.WithFields(log.Fields{"Name": r.Name}).Info("REQUEST: Path")
	return volume.Response{Err: "not found"}
}

// Mount mounts a volume onto the local file system.
func (d *BaseDriver) Mount(r volume.MountRequest) volume.Response {
	log.WithFields(log.Fields{"Name": r.Name, "ID": r.ID}).Info("REQUEST: Mount")
	return volume.Response{Err: "not found"}
}

// Unmount removes a volume from the local file system.
func (d *BaseDriver) Unmount(r volume.UnmountRequest) volume.Response {
	log.WithFields(log.Fields{"Name": r.Name, "ID": r.ID}).Info("REQUEST: Unmount")
	return volume.Response{Err: "not found"}
}

// Capabilities returns the capabilities of the driver
func (d *BaseDriver) Capabilities(r volume.Request) volume.Response {
	log.Info("REQUEST: Capabilities")
	return volume.Response{}
}
