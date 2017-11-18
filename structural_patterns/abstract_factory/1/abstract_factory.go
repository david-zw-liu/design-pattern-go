package abstract_factory_example_1

// define Architectures
type Architecture int

const (
	EMBER Architecture = iota
	ENGINOLA
)

// CPU
type CPU interface{}
type EmberCPU struct{}
type EnginolaCPU struct{}

// MMU
type MMU interface{}
type EmberMMU struct{}
type EnginolaMMU struct{}

// factory interface
type AbstractFactory interface {
	CreateCPU() CPU
	CreateMMU() MMU
}

// implement AbstractFactory interface
type EmberToolkit struct{}

func (EmberToolkit) CreateCPU() CPU {
	return EmberCPU{}
}
func (EmberToolkit) CreateMMU() MMU {
	return EmberMMU{}
}

type EnginolaToolkit struct{}

func (EnginolaToolkit) CreateCPU() CPU {
	return EnginolaCPU{}
}
func (EnginolaToolkit) CreateMMU() MMU {
	return EnginolaMMU{}
}

func GetFactory(architecture Architecture) AbstractFactory {
	emberToolkit := EmberToolkit{}
	enginolaToolkit := EnginolaToolkit{}
	var factory AbstractFactory

	switch architecture {
	case EMBER:
		factory = emberToolkit
	case ENGINOLA:
		factory = enginolaToolkit
	}

	return factory
}
