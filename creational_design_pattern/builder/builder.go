package builder

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(builder BuildProcess) {
	m.builder = builder
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetWheels().SetSeats().SetStructure()
}
