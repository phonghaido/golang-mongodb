# GO API Server

## Configuration
There is only one value in the configuration file (config.txt), which is **MONGODB_URI**.
The value of **MONGODB_URI** is *mongodb+srv://<username>:<passowrd>@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true*

## Run Application
The build file of the application can be found under **bin** directory.
To run the application, simply run this command
```bash
   ./bin/golang-mongodb
```
To build the application, run this command
```bash
   go build -o /bin/golang-mongodb
```

## Endpoints
- POST `/v1/mongodb`
  - Fetch data from Mongo DB. JSON body with `startDate`, `endDate`, `minCount`, `maxCount` is obligatory
  - Example:
    ```sh
    curl --request POST \
      --url http://localhost:8080/v1/mongodb \
      --header 'Content-Type: application/json' \
      --data '{"startDate":"2016-01-27", "endDate":"2020-01-28", "minCount": 100, "maxCount": 30000}'
    ```
- POST `/v1/in-memory`
  - Set key-value data in application's In-Memory. JSON body with `key`, `value` is obligatory
  - Example:
    ```sh
    curl --request POST \
      --url http://localhost:8080/v1/in-memory \
      --header 'Content-Type: application/json' \
      --data '{"key":"key-test-1", "value":"value-test-1"}'
    ```
- GET `/v1/in-memory?key=<the-key-of-the-data>`
  - Get data with key
  - Example:
    ```sh
    curl --request GET \
      --url http://localhost:8080/v1/in-memory?key=key-test-1
    ```