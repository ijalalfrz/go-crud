
### Prerequisites

What things you need to install the software and how to install them

```
Give examples
Golang v1.16
Go Mod
....
```

### Installing

A step by step series of examples that tell you have to get a development env running

Say what the step will be
- Create ENV file (.env) with this configuration:
```
APP_NAME=weight-service
PORT=9000
MONGODB_URL=mongodb://localhost:27017
MONGODB_DATABASE=weight-service
MONGODB_MIN_POOL_SIZE=50
MONGODB_MAX_POOL_SIZE=100
MONGODB_MAX_IDLE_CONNECTION_TIME_MS=10000
```

- Then run this command (Development Issues)
```
Give the example
...
$ make install
$ make run-dev
```

### Running the tests

Explain how to run the automated tests for this system
```sh
Give the example
...
$ make test-dev
```

### Running the tests (With coverage appear on)

Explain how to run the automated tests for this system
```sh
Give the example
...
$ make cover
```

