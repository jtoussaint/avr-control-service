# AVR Control

A simple API for controlling an Audio Video Receiver (AVR), specifically a Denon S750H.  

The behavior of the Api can be found [here](avr.apib).

The service talks to an AVR that impliments this [protocol](http://assets.eu.denon.com/DocumentMaster/DE/AVR2113CI_1913_PROTOCOL_V02.pdf).

## Getting Started

After cloning this repo, run the `build.sh` script to build and push an image of the service locally.

### Unit tests

Tests for the service can be run by running `go test`.  Alternativly tests are run as part of the docker build in `build.sh`.

### Integration tests

Before running integration tests you need to have an authenticated session with github packages.  You can do this by creating a PAT then running the following command.  Replace / set `$GITHUB_PAT` and `$USER_NAME` before running the script.

`echo "${{ $GITHUB_PAT }}" | docker login docker.pkg.github.com -u $USER_NAME --password-stdin`

Integration tests can be run by running `integration-test.sh`.  These tests will pull a container that acts as a mock AVR from [https://github.com/jtoussaint/mock-denon-avr](https://github.com/jtoussaint/mock-denon-avr) and runs the tests found in the `postman-collections` folder.

### Running the service

The service can be run locally in disconnected mode by running `run-local.sh` after building the image.

The service requires the following environment vairables:

| Env Var  |      Description      |
|----------|:--------------|
| AVR_HOST | The hostname / ip address of the AVR. |
| AVR_PORT | The port that the AVR listens on (23) |

