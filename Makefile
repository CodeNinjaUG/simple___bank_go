postgres:    
	
	docker run --name simplebank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=8bdc7axyzex -d postgres:12-alpine 

createdb:
	
	docker exec -it simplebank createdb --username=root --owner=root s_bank

dropdb: 
	
	docker exec -it simplebank dropdb s_bank

migrateup:   
	
	migrate -path db/migrations -database "postgresql://root:8bdc7axyzex@localhost:5432/s_bank?sslmode=disable" -verbose up

migratedown:  
	
	migrate -path db/migrations -database "postgresql://root:8bdc7axyzex@localhost:5432/s_bank?sslmode=disable" -verbose down	
sqlc:
    
	sqlc generate

test:
    
	go test -v -co ver ./...


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test