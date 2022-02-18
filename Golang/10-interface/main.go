package main

//interface = contrato

type transporter interface{
	calculateFrete(value float32) float32
}

type correios struct{}

func (c correios) calculateFrete(value float32)float32{
	return value * 0.3
}

type fedex struct {}
func (f fedex)calculateFrete(value float32)float32{
	return value * 0.5
}


