package builder

type BikeBuilder struct {
	vehicle VehicleProduct
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.vehicle.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.vehicle.Seats = 2
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.vehicle.Structure = "Bike"
	return c
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.vehicle
}
