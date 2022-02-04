# Iot Middleware

## Prime Directive:
‘Regardless of what we discover, we understand and truly believe that everyone did the best job he or she could, given what was known at the time, his or her skills and abilities, the resources available, and the situation at hand.’

## Dependencies:
This project depends on the following projects to run correctly (install with the package manager of your distro of choice):

```sh
$ sudo pacman -S go docker npm
$ pip install --user docker-compose
```

Dev tools:

```sh
$ pip install --user pre-commit
$ pre-commit install
```


## Makefile targets:
* __native_run:__ Run middleware application on native mode.
* __native_build:__ Build middleware application for Linux-x86.
* __docker_run:__ Run all application containers.
* __docker_build:__ Build middleware application on a container.

## Running the project:

Start InfluxDB:

```sh
$ docker-compose up
```

Install web dependencies:

```sh
$ cd web
$ npm install
```

Build web application:

```sh
$ cd web
$ npm run build
```

Start Golang Middleware:

```sh
$ go run . --dev
```

## Running the device simulator:

```sh
$ pip install -r tools/requirements.txt --user
$ python3 tools/simulador.py --mf --tf --cm
```
