run-dev-auth: 
	go run main.go ./env/dev/.env.auth
run-dev-player: 
	go run main.go ./env/dev/.env.player

.PHONY: run-dev-auth, run-dev-player