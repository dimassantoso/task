# Simple API

## Requirements

- Golang version 1.5+
- other needed modules

## Environment Variables

 ```
    SERVICE_NAME=
    BASIC_AUTH=
    PORT=
    WRITE_DB_HOST=
    WRITE_DB_PORT=
    WRITE_DB_NAME=
    WRITE_DB_USER=
    WRITE_DB_PASSWORD=
    READ_DB_HOST=
    READ_DB_PORT=
    READ_DB_NAME=
    READ_DB_USER=
    READ_DB_PASSWORD=
    OMDB_BASE_URL=
    OMDB_API_KEY=
```

### How to run service ###
* open your terminal
* run `cd 2/microservice-assignment/`
* run `cp env.dists .env`
* run `make run` 

### How to run unit test ###
* run `make test`

### Collections ###
* run `cd example-go/schema/tools`
* import file `Assignment.postman_collection.json` to your **postman** for collection
* import file `Assignment-Env.postman_environment.json` to your **postman** for environment


### SQL ###
* run `cd example-go/schema/data`
* import file `dump_31-03-2021_22_51_41.sql` to your **MySQL database** for database