package equations

type Vehicle string

const (
	Bus    Vehicle = "bus"
	Subway Vehicle = "subway"
)

const CalculatePublicTransport = "public-transport"

type CalculatePublicTransportArgs struct {
	Vehicle  Vehicle
	Distance float64
	Weight   float64
}

func CalculatePublicTransportEquation(args interface{}) float64 {
	var vehicleFactor float64

	publicTransportArgs := args.(CalculatePublicTransportArgs)

	if publicTransportArgs.Vehicle == Bus {
		vehicleFactor = 3
	} else {
		vehicleFactor = 4
	}

	return ((publicTransportArgs.Distance * publicTransportArgs.Weight) / 10) * vehicleFactor
}
