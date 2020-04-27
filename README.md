# Crud template of the dynamodb

An easy template to access and get your data on aws DynamoDB.


## Getting Started
### Pre-requirements
In your machine install [AWS-CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html) and [CONFIGURE](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) to save configuration of the your database

### Install
If you already have golang installed you can install by running the command:
```sh
go get -u ./...
```

### Init server
To start the project run the command:
```sh
go run cmd/app/main.go
```
You can see in your terminal this log:
`service running on port  :8080`

## Usage

### Health
This route return life of the project
```text
    GET - http://localhost:8080/health
```

### Get_All
This route return all data in your database
```text
    GET - http://localhost:8080/product
```

#### Get_One
This route return specific data in your database
```text
    GET - http://localhost:8080/product/{ID}
```

#### Post
This route create on item in your database
```text
    POST - http://localhost:8080/product

    {
        "name": "product"
    }
```

#### Put
This route update on item in your database
```text
    PUT - http://localhost:8080/product/{ID}
    
    {
        "name": "product1"
    }
```

#### Delete
This route remove on item in your database
```text
    DELETE - http://localhost:8080/product/{ID}
```
