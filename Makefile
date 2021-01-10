output: build
	./dont-panic-11238&
	docker-compose up --build --remove-orphans

build:
	go build

clean:
	killall dont-panic-11238
	rm dont-panic-11238
	docker-compose down
