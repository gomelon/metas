package bar

import (
	"github.com/gomelon/metas/mwire/testdata/foo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDefaultFoo,
	NewFooAOPWithBye,
	NewFooAOPWithGreet,

	wire.Bind(new(FooAOPForGreet), new(*DefaultFoo)),
	wire.Bind(new(FooAOPForBye), new(*FooAOPWithGreetImpl)),
	wire.Bind(new(foo.Foo), new(*FooAOPWithByeImpl)),
)
