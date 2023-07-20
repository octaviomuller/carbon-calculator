package main

import (
	"example/carbon-calculator/equations"
	"fmt"
)

type (
	Form struct {
		Name           string
		Description    string
		Fields         []Field
		Expression     equations.EquationType
		ExpressionArgs map[string]string
	}

	Field struct {
		ID    string
		Label string
		Value interface{}
	}
)

var forms = []Form{
	{
		Name:        "Transporte público",
		Description: "Calculo por transport público",
		Expression:  equations.CalculatePublicTransport,
		ExpressionArgs: map[string]string{
			"vehicle":  "field1",
			"distance": "field2",
			"weight":   "field3",
		},
		Fields: []Field{
			{
				ID:    "field1",
				Label: "vehicle",
				Value: "bus",
			},
			{
				ID:    "field2",
				Label: "distance",
				Value: 800.00,
			},
			{
				ID:    "field3",
				Label: "weight",
				Value: 1000.00,
			},
		},
	},
}

func (p *Form) Calculate() float64 {
	var result float64
	var equation func(interface{}) float64 = equations.Equations[p.Expression]

	switch p.Expression {
	case equations.PublicTransport:
		{
			var vehicle equations.Vehicle = equations.Vehicle(p.CastFieldToArg(p.ExpressionArgs["vehicle"]).(string))
			var distance float64 = p.CastFieldToArg(p.ExpressionArgs["distance"]).(float64)
			var weight float64 = p.CastFieldToArg(p.ExpressionArgs["weight"]).(float64)

			publicTransportArgs := equations.CalculatePublicTransportArgs{
				Vehicle:  vehicle,
				Distance: distance,
				Weight:   weight,
			}

			result = equation(publicTransportArgs)
		}
	}

	return result
}

func (p *Form) CastFieldToArg(fieldID string) interface{} {
	for _, field := range p.Fields {
		if field.ID == fieldID {
			return field.Value
		}
	}

	return nil
}

func main() {
	for _, form := range forms {
		result := form.Calculate()
		fmt.Printf("Pegada de carbono: %f CO2eq", result)
	}
}
