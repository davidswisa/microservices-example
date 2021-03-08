run:
	docker-compose up -d --build
	# docker-compose up -d --build db kafka zookeeper  

run-prod:
	docker stop prod
	docker-compose up -d --build prod

run-cons:
	docker stop cons
	docker-compose up -d --build cons

run-orm:
	docker stop orm
	docker-compose up -d --build orm

run-ui:
	docker stop ui
	docker-compose up -d --build ui

run-querier:
	docker stop querier
	docker-compose up -d --build querier

prod:
	go run ./prod/main.go

cons:
	go run ./cons/main.go

querier:
	go run ./querier/main.go

orm:
	go run ./orm/main.go

test:
	scripts/send_request.sh
