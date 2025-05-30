# Protobuf Generator

## Language Support

Currently, only the following languages are supported:

* Go
* TypeScript
* Python

Additional languages will be added in the future.

## Usage

### /src/source

This is where you place your directories containing the **.proto** files.

All folders under `/src/source/*` should be lowercase and use `-` as a separator.
Internally, the name of the packages or modules will be converted to the appropriate naming convention for each target language.

### Docker

Clone the repository to any location and run `make compose` from the terminal.
This will generate the Protobuf files under `/src/generated`.
