# foobar-tvs

This repository is under construction 🚧
### Database Design
You can the url of the database design [here](https://dbdiagram.io/d/63613e8e5170fb6441dcf57f).

![Image](doc/db_design.png)

### Initial Setup

- You need to setup [Go](https://go.dev/), [Docker](https://docker.com), [golang-migrate cli](https://github.com/golang-migrate/migrate) in your machine.

- Create mock docker network
```
docker network create mock-network
```

- deploy containers
```
docker-compose up -d
```

- table migration up
```
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/foobar-db?sslmode=disable" -verbose up  
```


### Acknowledgment
- [dbdiagram.io](https://dbdiagram.io)
