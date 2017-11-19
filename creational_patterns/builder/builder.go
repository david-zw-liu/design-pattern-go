package builder

type Pizza struct {
	Dough   string
	Sauce   string
	Topping string
}

func (p *Pizza) SetDough(dough string) {
	p.Dough = dough
}

func (p *Pizza) SetSauce(sauce string) {
	p.Sauce = sauce
}

func (p *Pizza) SetTopping(topping string) {
	p.Topping = topping
}

type PizzaBuilder struct {
	Pizza *Pizza
}

func (pb *PizzaBuilder) GetPizza() *Pizza {
	return pb.Pizza
}

func (pb *PizzaBuilder) CreateNewPizzaProduct() {
	pb.Pizza = &Pizza{}
}

type IPizzaBuilder interface {
	GetPizza() *Pizza
	CreateNewPizzaProduct()
	BuildDough()
	BuildSauce()
	BuildTopping()
}

type HawaiianPizzaBuilder struct {
	PizzaBuilder
}

func (hpb *HawaiianPizzaBuilder) BuildDough() {
	hpb.Pizza.SetDough("cross")
}

func (hpb *HawaiianPizzaBuilder) BuildSauce() {
	hpb.Pizza.SetSauce("mild")
}

func (hpb *HawaiianPizzaBuilder) BuildTopping() {
	hpb.Pizza.SetTopping("ham+pineapple")
}

func (hpb *HawaiianPizzaBuilder) GetPizza() *Pizza {
	return hpb.Pizza
}

type SpicyPizzaBuilder struct {
	PizzaBuilder
}

func (spb *SpicyPizzaBuilder) BuildDough() {
	spb.Pizza.SetDough("pan baked")
}

func (spb *SpicyPizzaBuilder) BuildSauce() {
	spb.Pizza.SetSauce("hot")
}

func (spb *SpicyPizzaBuilder) BuildTopping() {
	spb.Pizza.SetTopping("pepperoni+salami")
}

func (spb *SpicyPizzaBuilder) GetPizza() *Pizza {
	return spb.Pizza
}

type Waiter struct {
	PizzaBuilder IPizzaBuilder
}

func (w *Waiter) SetPizzaBuilder(pizzaBuilder IPizzaBuilder) {
	w.PizzaBuilder = pizzaBuilder
}

func (w *Waiter) GetPizza() *Pizza {
	return w.PizzaBuilder.GetPizza()
}

func (w *Waiter) ConstructPizza() {
	w.PizzaBuilder.CreateNewPizzaProduct()
	w.PizzaBuilder.BuildDough()
	w.PizzaBuilder.BuildSauce()
	w.PizzaBuilder.BuildTopping()
}
