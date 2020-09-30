# Boolean As a Service
A Golang API for booleans, supporting CRUD operations for booleans

Features include:
- Ability to create, get, update and delete a boolean
- Storing data in mysql database
- Service exposing RESTful end points

## Tech Stack Used:
- Golang
- MySQL
  - Gorm as orm library
- Docker

## Configuration

### MySQL
- __Mysql 8.0.21 or higher__ <br />
To install mysql on  macOS follow [install mysql](https://flaviocopes.com/mysql-how-to-install/) <br />

- __go 1.15 or higher__ <br />
To install go on macOS follow [install go](https://www.geeksforgeeks.org/how-to-install-golang-on-macos/) <br />


- Create a new user
```
mysql> CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';
```
- Create a database
```
mysql> CREATE DATABASE database_name;
```
- Give the created user all privileges on the created database
```
mysql> GRANT ALL PRIVILEGES ON database_name.* TO 'username'@'localhost' identified by 'password';
mysql> FLUSH PRIVILEGES;
```

## Installing and Running Service

### Build from source
 - Clone this repository and `cd` to the `booleans/src` directory where you cloned it.
 - Install the go module
 ```
 go mod download
 ```
 - Run
 ```
 DB_USER=<username> DB_PASS=<password> go run main.go 
```

### With Docker 
- Download the image from dockerhub
```
docker pull mishranant1/booleans
```
- Run the image
```
docker run -i -d -t -p 8080:8080 -e DB_USER=<username> -e DB_PASS=<password> mishranant1/booleans
```

## API
### Base URL
```
http://localhost:8080
```

### Create a new Boolean

```
POST /

request:

{
  "value": true,
  "key": "name" // this is optional
}

response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}
```
Usage:
 
```console
curl -X POST http://localhost:8080/[id] --header "Content-Type: application/json" --data '{"value": true, "label": "name"}'
```


### Retrieve a Boolean

```
GET /:id
response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": true,
  "key": "name"
}
```
Usage:
 
```console
curl -X GET http://localhost:8080/[id]
```

### Update a Boolean
```
PATCH /:id
request:

{
  "value":false,
  "key": "new name" // this is optional
}

response:

{
  "id":"b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "value": false,
  "key": "new name"
}
```

Usage:

```console
curl -X PATCH http://localhost:8080/[id] --header "Content-Type: application/json" --data '{"value": true, "label": "new name"}'
```

### Delete an existing Boolean
```
DELETE /:id
response:
HTTP 204 No Content
```

## Testing
To run test functions, run following commands from you `booleans` folder
### Test controllers
```
cd controllers
go test
```
### Test services

```
cd services
go test
```