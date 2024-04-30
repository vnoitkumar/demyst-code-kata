# Code Kata

## How to run the project
Run the below command to start the project in docker 

```shell
$ make start
```

Incase you don't want to run it in docker, use the below command

```shell
$ go run main.go
```

And checkout the makefile for more useful commands

🔴 After running the app find the Trace (TRC) logs for the expected output as mentioned in the problem statement, refer the image below 🔴

![Log](/doc/images/log.png)


## How to run the test and coverage
Run the below command to run the test with coverage

```shell
$ make test_with_coverage
```

## Problem statement

The goal of the project is to build a command line tool.

Using Go, write a command line tool that consumes the first `20` `even` numbered TODO's in most performant way and output the `title` and whether it is `completed` or not.

- TODO at index 1 can be accessed at: <https://jsonplaceholder.typicode.com/todos/1>

- TODO at index 2 can be accessed at: <https://jsonplaceholder.typicode.com/todos/2>


## Flow diagram

![Flow diagram](/doc/images/flow-diagram.png)


## API Details

#### Sample Request

```shell
curl --request GET --url https://jsonplaceholder.typicode.com/todos/1 
```

#### Sample Response

```json
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
```