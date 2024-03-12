package storage

import (
	"github.com/MatRouillard/vehicle-server/storage/vehiclestore"
)

type Store interface {
	Vehicle() vehiclestore.Store
}
