package equations

type EquationType string

const (
	PublicTransport EquationType = "public-transport"
)

var Equations = map[EquationType]func(interface{}) float64{
	PublicTransport: CalculatePublicTransportEquation,
}
