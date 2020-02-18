package vfactory

type veloFactory struct{}

func NewVeloFactory() Repository {
	return &veloFactory{}
}
