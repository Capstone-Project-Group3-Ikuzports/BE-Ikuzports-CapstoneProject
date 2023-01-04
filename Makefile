testAllService:
	go test ./features/... -coverprofile=cover.out && go tool cover -html=cover.out

testUser:
	go test ./features/user... -coverprofile=cover.out && go tool cover -html=cover.out

testClubs:
	go test ./features/clubs... -coverprofile=cover.out && go tool cover -html=cover.out

testChat:
	go test ./features/chat... -coverprofile=cover.out && go tool cover -html=cover.out

testActivity:
	go test ./features/clubActivity... -coverprofile=cover.out && go tool cover -html=cover.out

testClubMember:
	go test ./features/clubMember... -coverprofile=cover.out && go tool cover -html=cover.out

testAuth:
	go test ./features/auth... -coverprofile=cover.out && go tool cover -html=cover.out

testCategory:
	go test ./features/category... -coverprofile=cover.out && go tool cover -html=cover.out

testEvent:
	go test ./features/event... -coverprofile=cover.out && go tool cover -html=cover.out

testGalery:
	go test ./features/galery... -coverprofile=cover.out && go tool cover -html=cover.out

testItemCategory:
	go test ./features/itemCategory... -coverprofile=cover.out && go tool cover -html=cover.out

testParticipant:
	go test ./features/participant... -coverprofile=cover.out && go tool cover -html=cover.out

testProduct:
	go test ./features/product... -coverprofile=cover.out && go tool cover -html=cover.out

testProductImage:
	go test ./features/productImage... -coverprofile=cover.out && go tool cover -html=cover.out

testTransaction:
	go test ./features/transaction... -coverprofile=cover.out && go tool cover -html=cover.out

run:
	go run main.go