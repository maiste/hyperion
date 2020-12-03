# hyperion
 
## What is Hyperion ?
Hyperion is a Golang RESTApi server for the eos project and golactus.

## Build hyperion
To build Hyperion you can use the **make** command.

```sh
  # Build the project
  $ make

  # Equivalent to make
  $ make build

  # Build and run the project
  $ make run

  # Clean executable
  $ make clear
```
## Run inside docker
Hyperion is in the dockerhub. You can use it as follow:
```sh
docker run --rm -d --name hyperion -p <port>:8080 -v <path_to_template>:/templates maiste/hyperion:latest
```
