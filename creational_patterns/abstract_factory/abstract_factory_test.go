package abstract_factory

import (
	"testing"
)

func TestGetFactory(t *testing.T) {
	emberFactory := GetFactory(EMBER)
	enginolaFactory := GetFactory(ENGINOLA)

	if _, ok := emberFactory.(EmberToolkit); !ok {
		t.Errorf("Expect get type `EmberToolkit, but got %T`", emberFactory)
	}

	emberCPU := emberFactory.CreateCPU()
	if _, ok := emberCPU.(EmberCPU); !ok {
		t.Errorf("Expect get type `EmberCPU, but got %T`", emberCPU)
	}

	emberMMU := emberFactory.CreateMMU()
	if _, ok := emberMMU.(EmberMMU); !ok {
		t.Errorf("Expect get type `EmberCPU, but got %T`", emberMMU)
	}

	if _, ok := enginolaFactory.(EnginolaToolkit); !ok {
		t.Errorf("Expect get type `EnginolaToolkit, but got %T`", enginolaFactory)
	}

	enginolaCPU := enginolaFactory.CreateCPU()
	if _, ok := enginolaCPU.(EnginolaCPU); !ok {
		t.Errorf("Expect get type `EnginolaCPU, but got %T`", enginolaCPU)
	}

	enginolaMMU := enginolaFactory.CreateMMU()
	if _, ok := enginolaMMU.(EnginolaMMU); !ok {
		t.Errorf("Expect get type `EnginolaCPU, but got %T`", enginolaMMU)
	}
}
