start:
	@docker-compose up --build

stop:
	@docker-compose rm -v --stop --force
	@docker rmi book-tickets