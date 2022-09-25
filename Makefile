create_net:
	docker network create douyin-network

mysql:
	docker run --name douyin-mysql -v -v ./volume/mysql/data:/var/lib/mysql/ --network douyin-network -p 3306:3306 -e MYSQL_USER=douyin -e MYSQL_PASSWORD=douyin -e MYSQL_DATABASE=douyin -e MYSQL_ROOT_PASSWORD=douyin -d mysql:5.7

createdb:
	docker exec -it douyin-mysql  createdb --username=root --owner=root douyin

dropdb:
	docker exec -it douyin-mysql dropdb douyin

migrateup:
	migrate -path ./db/migration -database "mysql://douyin:douyin@(localhost:3306)/douyin" -verbose up

migratedown:
	migrate -path ./db/migration -database "mysql://douyin:douyin@(localhost:3306)/douyin" -verbose down

.PHONY: createdb dropdb migrateup migratedown create_net