package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarBuilder(t *testing.T) {
	carBuilder := &CarBuilder{}
	m := ManufacturingDirector{}
	m.SetBuilder(carBuilder)
	m.Construct()
	carVehicle := carBuilder.GetVehicle()
	assert.Equal(t, 4, carVehicle.Wheels)
	assert.Equal(t, 5, carVehicle.Seats)
	assert.Equal(t, "Car", carVehicle.Structure)
}

func TestBuilder(t *testing.T) {
	carBuilder := &BikeBuilder{}
	m := ManufacturingDirector{}
	m.SetBuilder(carBuilder)
	m.Construct()
	carVehicle := carBuilder.GetVehicle()
	assert.Equal(t, 2, carVehicle.Wheels)
	assert.Equal(t, 2, carVehicle.Seats)
	assert.Equal(t, "Bike", carVehicle.Structure)
}
