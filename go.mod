module github.com/blablatov/listen2rabbit4dirx

go 1.16

require (
	github.com/blablatov/listen2rabbit4dirx/call2handler v0.0.0-00010101000000-000000000000
	github.com/pandeptwidyaop/gorabbit v1.0.0-alpha
	github.com/streadway/amqp v1.0.0
)

replace github.com/blablatov/listen2rabbit4dirx/call2handler => ./call2handler
