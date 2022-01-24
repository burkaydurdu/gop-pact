# Pact Notes

---

#### For Consumer
```shell
go test -v ./...  -tags=consumer
```

#### For Provider
```shell
go test -v ./...  -tags=provider
```

#### Run Broker
```shell
docker-compose build .
docker-compose up
```

### can-i-deploy
Before you deploy a new version of an application to a production environment,
you need to know whether or not the version you're about to deploy is compatible
with the versions of the other apps that already exist in that environment.
```shell
pact-broker can-i-deploy --pacticipant=GopClient --broker-base-url=http://localhost --version=2.5.2 --to=dev
```
---
🎉 Thank you! [Abdulsametileri](https://github.com/Abdulsametileri)