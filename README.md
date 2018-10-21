# TO-DO

This is a simple server side implementation for to-do application through which `Task` managemant can be accomplished. This application provides REST API to perform CRUD operations on `Task`.

This application follows [go standard packaging](https://github.com/golang-standards/project-layout), 
uses `MongoDB` for data persistence, `uber's fx` for dependency injection and go standard vendoring for library dependencies.

## Ubiquitous language.

- A `Task` represents a note *(a brief record of points or ideas written down as an aid to memory.)*

- A `User` can have list of `task`

## Requirements

- Go version `go1.11.1` or upper.

- [MongoDB on your local machine](https://docs.mongodb.com/manual/installation/)


## Build and run


- Run the following command in root directory of repository.
```
go build -o cmd/todo/todo cmd/todo/*.go
```

- Now before running the application make sure that MongoDB is in running state and ready to accept connection, note down the `url` for the same.

- Run the following command in root directory of repository.
```
MONGO_DB_URL="your_mongo_db_url"  ./cmd/todo/todo -local
```
 - `-local` is the flag to ensure that server starts on `localhost` at port `8090`


## Usage
Once application goes in running state api's  can be hit on `http://localhost:8090`

### Following are the APIs

#### User Signup
Returns the `api_secret` for the user which needs to be note down for later usage. All further API request on `task` will need this api_secret.

|Method|Url|
|---|---|
|**POST**|`http://localhost:8090/user/signup`|

Body 
```
{
    "email": "akash@todo.com",
    "name": "akash"
}
```
Response

```
{
    "email":"akash@todo.com",
    "Name":"akash",
    "api_secret":"e6cc35cf-4e6c-4aab-ac24-ef9a7dab35da"
}
```

#### Create a task
Returns the `task` entity.

|Method|Url|
|---|---|
|**POST**|`http://localhost:8090/list/task`|


Header

|key|value|
|---|---|
|email|akash@todo.com|
|secret|e6cc35cf-4e6c-4aab-ac24-ef9a7dab35da|


Body
```
{
    "title" : "note 1",
    "body": "some description for note 1"
}
```

Response
```
{
    "id": "c2542fff-94ba-4dee-a96e-7ad9d6093a32",
    "title": "note 1",
    "body": "some description for note 1",
    "user_email": "akash@todo.com",
    "starred": false,
    "created_at": "2018-10-22T00:41:10.864637+05:30",
    "updated_at": "2018-10-22T00:41:10.864637+05:30"
}
```
#### List all task
Returns an array of `task` in response for the current user.
|Method|Url|
|---|---|
|**GET**|`http://localhost:8090/list`|
Header

|key|value|
|---|---|
|email|akash@todo.com|
|secret|e6cc35cf-4e6c-4aab-ac24-ef9a7dab35da|

Response
```
[
  {
      "id": "c2542fff-94ba-4dee-a96e-7ad9d6093a32",
      "title": "note 1",
      "body": "some description for note 1",
      "user_email": "akash@todo.com",
      "starred": false,
      "created_at": "2018-10-22T00:41:10.864637+05:30",
      "updated_at": "2018-10-22T00:41:10.864637+05:30"
  }
]
```

#### Similarly...

|Method|Url|description|needs secret|
|---|---|---|---|
|**GET**|`http://localhost:8090/list/task/{id}`|Get a task by id|yes|
|**PUT**|`http://localhost:8090/list/task`|Update a task - body must contain the `id` of existing task to be updated|yes|
|**DELETE**|`http://localhost:8090/list/task/{id}`|delete a task by `id`|yes|
