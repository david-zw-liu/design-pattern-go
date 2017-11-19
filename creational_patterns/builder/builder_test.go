package builder

import (
	"testing"
)

func TestWaiterAndPizzaBuilder(t *testing.T) {
	var pizza *Pizza
	waiter := new(Waiter)

	hawaiianPizzaBuilder := new(HawaiianPizzaBuilder)
	spicyPizzaBuilder := new(SpicyPizzaBuilder)

	// Test HawaiianPizzaBuilder
	waiter.SetPizzaBuilder(hawaiianPizzaBuilder)
	waiter.ConstructPizza()
	pizza = waiter.GetPizza()
	if pizza.Dough != "cross" {
		t.Errorf("Expected pizza.Dough is \"cross\", but got \"%v\"", pizza.Dough)
	}
	if pizza.Sauce != "mild" {
		t.Errorf("Expected pizza.Sauce is \"mild\", but got \"%v\"", pizza.Sauce)
	}
	if pizza.Topping != "ham+pineapple" {
		t.Errorf("Expected pizza.Topping is \"ham+pineapple\", but got \"%v\"", pizza.Topping)
	}

	// Test SpicyPizzaBuilder
	waiter.SetPizzaBuilder(spicyPizzaBuilder)
	waiter.ConstructPizza()
	pizza = waiter.GetPizza()
	if pizza.Dough != "pan baked" {
		t.Errorf("Expected pizza.Dough is \"pan baked\", but got \"%v\"", pizza.Dough)
	}
	if pizza.Sauce != "hot" {
		t.Errorf("Expected pizza.Sauce is \"hot\", but got \"%v\"", pizza.Sauce)
	}
	if pizza.Topping != "pepperoni+salami" {
		t.Errorf("Expected pizza.Topping is \"pepperoni+salami\", but got \"%v\"", pizza.Topping)
	}
}
