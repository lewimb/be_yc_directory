migrate:
	go run db/migrates/up/up.go

rollback:
	go run db/migrates/down/down.go

run:
	go run server.go