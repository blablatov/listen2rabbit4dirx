module github.com/blablatov/listen2rabbit4dirx

go 1.16

replace github.com/blablatov/listen2rabbit4dirx/call2handler => ./call2handler

require (
	github.com/blablatov/listen2rabbit4dirx/call2handler v0.0.0-00010101000000-000000000000
	github.com/blablatov/tlsgorabbit v0.0.0-20221106115828-544d061a76fa
	github.com/streadway/amqp v1.0.0
)
