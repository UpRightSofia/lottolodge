# Ledger

Loan Ledger

### Installation

copy `.env.example` to `.env` and edit the values

`find scripts -type f -exec chmod +x {} \;` enables the scripts

Make sure docker is installed and running

### Build and run docker image

`go run main.go` runs locally

`./scripts/build_and_tag_image.sh`

to build the image

then run 

`./scripts/start.sh`

to start all relevant containers

`./scripts/stop.sh` to stop them.