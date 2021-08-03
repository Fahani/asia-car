# Asia Cars

### Install and compile
`mkdir -p build`

`go build -o build/asia_cars`


### Run the application
`cd build/`

`./asia_cars`

When these commands are executed, you can test the application on http://localhost:9000/status

### Run All Tests
```
go test ./test/...
```

### API Requests

#### In Fleet

```
curl --location --request POST 'http://localhost:9000/in-fleet' \
--header 'Content-Type: application/json' \
--data-raw '{
    "in_fleet_date": "2014-11-12T11:45:26.371Z",
    "chassis_number": "01234567890123456"
}'
```

#### Install

```
curl --location --request POST 'http://localhost:9000/install' \
--header 'Content-Type: application/json' \
--data-raw '{
    "device_serial_number": "abc",
    "chassis_number": "01234567890123456"
}'
```


#### Update Battery 

```
curl --location --request POST 'http://localhost:9000/update-battery' \
--header 'Content-Type: application/json' \
--data-raw '{
    "battery": 75,
    "device_serial_number": "abc"
}'
```


#### Update Fuel

```
curl --location --request POST 'http://localhost:9000/update-fuel' \
--header 'Content-Type: application/json' \
--data-raw '{
    "fuel": 75,
    "update_type": "increment",
    "device_serial_number": "abc"
}'
```


#### Update Mileage

```
curl --location --request POST 'http://localhost:9000/update-mileage' \
--header 'Content-Type: application/json' \
--data-raw '{
    "mileage": 75,
    "device_serial_number": "abc"
}'
```

#### Vehicle Details

```
curl --location --request POST 'http://localhost:9000/vehicle-details' \
--header 'Content-Type: application/json' \
--data-raw '{
    "chassis_number": "01234567890123456"
}'
```