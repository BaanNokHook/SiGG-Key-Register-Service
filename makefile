STUPIDSECRET=thisisadumbsecretyo

.PHONY: all

all:
	test

test:
	go fmt *.go
	keyregister_ENDPOINT=http://localhost:8080/api/0/ go test $$(glide novendor) -v

devenv:
	docker run -d --name keyregister-redis redis
	docker run -d --name keyregister-postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=keyregister postgres
	sleep 5 # Wait for postgres to bootup
	docker run -it --rm -e keyregister_SECRET_KEY='${STUPIDSECRET}' --link keyregister-postgres:postgres --link keyregister-redis:redis keyregister:latest upgrade
	docker run -d --name my-keyregister -e keyregister_SECRET_KEY='${STUPIDSECRET}' --link keyregister-redis:redis --link keyregister-postgres:postgres -p 8080:9000 keyregister:latest
	docker run -d --name keyregister-cron -e keyregister_SECRET_KEY='${STUPIDSECRET}' --link keyregister-postgres:postgres --link keyregister-redis:redis keyregister:latest run cron
	docker run -d --name keyregister-worker-1 -e keyregister_SECRET_KEY='${STUPIDSECRET}' --link keyregister-postgres:postgres --link keyregister-redis:redis keyregister:latest run worker

devclean:
	docker kill $$(docker ps -q -a --no-trunc --filter name=^keyregister)
	docker rm $$(docker ps -q -a --no-trunc --filter name=^keyregister)
