module github.com/blablatov/listen2rabbit4dirx

go 1.16

require (
	github.com/blablatov/listen2rabbit4dirx/call2handler v0.0.0-00010101000000-000000000000
	github.com/blablatov/tlsgorabbit v0.0.0-20221104181138-d714e80a9ae0
	github.com/streadway/amqp v1.0.0
)

replace github.com/blablatov/listen2rabbit4dirx/call2handler => ./call2handler
