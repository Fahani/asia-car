AsiaCar wants to jump into the digital world and get remote control over its entire fleet; right now they have many manual processes and this generates many problems, queues of customers, vehicles that are in unknown state, etc. That is why AsiaCar asked us to implement an MVP, in order to understand if this digitalization is a viable solution.

Software they asked us to develop consists of two parts: for starters, we have an API that will be integrated into a mobile application, which will be used by AsiaCar operators. We currently have about 500 vehicle registrations per day and about 300 device installations.

Then, once the cars are installed, they will start sending telemetries, hence we have to be able to offer the external provider a json contract and a system to integrate with, where the providers can send us the levels of battery, fuel and mileage of each car. We currently have a fleet of ** 40,000 ** vehicles and we estimate that we will receive 30 telemetry values every second for each of them.

Please keep in mind that at the end of the year if all goes well we plan to expand our fleet to ** 400,000 ** vehicles so that all data ingestion will increase significantly and the system would have to be able to support it: keep scalability in mind while designing your software and architecture.

## Operations
The operations that the station operators may carry out will be the following:

### Infleet Vehicle
Once AsiaCar has bought a vehicle, this will arrive at the facilities for being registered in the system. An operator will take the chassis number and register the vehicle on the platform.

### Install Vehicle
Once the vehicle is registered on the platform, another operator will install the device in order to obtain the telemetries of the car in real time. What this process does is assigning a `device_serial_number` to our vehicle.

## Telemetry
In addition to the operations we have the actions on the telemetry that occur in real time, we can receive the following information:

### Mileage
```json
{
  "serial_number": "G-34567",
  "mileage": 1000,
  "unit": "km"
}
```

### Fuel
```json
{
  "serial_number": "G-34567",
  "fuel": 50,
  "type": "increment"
}
```

The `type` in the fuel message will inform us if it is an `increment` or `decrement`, when we react to this message. We must bear in mind that in our system fuel can never be negative.

### Battery
```json
{
  "serial_number": "G-34567",
  "battery": 75
}
```

In addition to receiving telemetries, we have to provide a way to query all the telemetries by chassis number and expose them in "real time".

## Model
All the actions that have been mentioned will be carried out on the `Vehicle` model that is exposed below:

| Vehicle              | Type     |
| -------------------- | -------- |
| chassis_nbr          | string   |
| license_plate        | string   |
| brand string         | string   |
| category string      | string   |
| infleet_date         | time     |
| device_serial_number | string   |
| battery_level        | int      |
| fuel_level           | int      |
| current_mileage      | int      |

On these data we have to run the following validations:
* `chassis_nbr`: Must be exactly 17 digits and alphanumeric.
* `license_plate`: Can only be alphanumeric.
* `category`: The car category will be defined by a standard called [ACRISS_CODE] (https://www.acriss.org/car-codes/), in the test we will evaluate that at least it is controlled that it is a maximum of 4 letters, but  you'll get a bonus if you do a validation based on the actual operation of the `ACRISS CODE`.
* `battery_level`, `fuel_level` and `current_mileage`: do not accept negative values.

## Requirements
* We prefer to have this software written in Golang, since it is with the main language we work with, but if you do not feel comfortable with it you can deliver it in the programming language you prefer.
* We do not expect any kind of specific API, but get ready to explain the choices you made.
* We do not want any real persistence system.
