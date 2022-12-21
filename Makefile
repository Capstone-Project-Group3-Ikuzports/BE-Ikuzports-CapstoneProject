testuser:
	go test ./features/user... -coverprofile=cover.out && go tool cover -html=cover.out

testclass:
	go test ./features/class... -coverprofile=cover.out && go tool cover -html=cover.out

testlog:
	go test ./features/log... -coverprofile=cover.out && go tool cover -html=cover.out

testauth:
	go test ./features/auth... -coverprofile=cover.out && go tool cover -html=cover.out

testmentee:
	go test ./features/mentee... -coverprofile=cover.out && go tool cover -html=cover.out

run:
	go run main.go