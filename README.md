# Receipt Processor Challenge

## Table of Contents
1. [Task](#task)
2. [How to Run](#how-to-run)
3. [Test API ](#test-api)
5. [Project Details](#project-details)
### Task
The task is to build a  web service that processes receipts according to the documented API specifications. The service should handle receipt submission, generate a unique ID for each receipt, and calculate points based on predefined rules.
### How to Run
To run the service, follow these steps:

#### 1. Clone the repository:

```bash
git clone git@github.com:samgo3/Receipt-Processor.git
cd Receipt-Processor
```

#### 2. Install Requirements
You will need the following installed on your local machine
- docker -- [install guide](https://docs.docker.com/get-docker/)
- [go](https://go.dev/doc/install) v1.21.1+ 
-  [swag](https://github.com/swaggo/swag) : This project uses  swag to generate swagger documentation for the apis.
    ```bash 
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

#### 3. How to Run

####  Using Makefile
Ensure that make is installed in the system
```bash
make all
make run
```


#### Using docker:
Start up docker service and run docker compose
```bash
docker compose up -d
```


### Test API
1. Using curl: 
    - Post Receipt to get id
    ```bash
    curl -X POST   -H "Content-Type: application/json"   -d '{
      "retailer": "M&M Corner Market",
      "purchaseDate": "2022-03-20",
      "purchaseTime": "14:33",
      "items": [
        {
          "shortDescription": "Gatorade",
          "price": "2.25"
        },{
          "shortDescription": "Gatorade",
          "price": "2.25"
        },{
          "shortDescription": "Gatorade",
          "price": "2.25"
        },{
          "shortDescription": "Gatorade",
          "price": "2.25"
        }
      ],
      "total": "9.00"
    }' http://localhost:5555/receipts/process
    ```
    - Get Points Info:
    Use the generated id to get the points information for the receipt
    ```bash
    curl http://localhost:5555/receipts/{id}/points
    ```

2. Use Swagger: Access the Swagger UI at [http://localhost:5555/swagger/index.html](http://localhost:5555/swagger/index.html) to test and execute POST and GET requests.

### Project Details
This project is built using the Go programming language and employs the Gorilla Mux package to route API calls. The service is designed to process receipts and calculate points based on predefined rules. It handles receipt processing entirely in-memory without relying on any external database.



