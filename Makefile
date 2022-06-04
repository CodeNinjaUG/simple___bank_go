postgres:    
	
	docker run --name  swaggie_bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=8bdc7axyzex -d postgres:12-alpine 

createdb:
	
	docker exec -it swaggie_bank createdb --username=root --owner=root  swaggie_bank

dropdb: 
	
	docker exec -it swaggie_bank dropdb swaggie_bank

migrateup:   
	
	migrate -path db/migrations -database "postgresql://root:8bdc7axyzex@localhost:5432/swaggie_bank?sslmode=disable" -verbose up

migratedown:  
	
	migrate -path db/migrations -database "postgresql://root:8bdc7axyzex@localhost:5432/swaggie_bank?sslmode=disable" -verbose down	
sqlc:
    
	sqlc generate

test:
    
	go test -v -co ver ./...


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test