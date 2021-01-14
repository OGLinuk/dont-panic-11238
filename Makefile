output: build
	./dont-panic-11238&

build:
	go build

services:
	docker-compose up -d --build --remove-orphans

clean:
	killall dont-panic-11238
	rm dont-panic-11238

clean-absolute:
	killall dont-panic-11238
	rm dont-panic-11238
	rm -rf DONTPANIC docker-compose.yml

clean-services:
	docker-compose down
