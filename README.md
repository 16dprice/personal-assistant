Must have Go, Rust, and Hatch installed.

# Setup

## Get protoc

1. Run `brew install protobuf`

## Get Note Server running

1. in `packages/note-server`, run `cp .env.example .env`
1. run `docker compose -f ./docker/docker-compose.yml --env-file ./.env up -d` to bring up the DB container
2. create new database on that postgres server called `note_server_db`
3. run `go run src/seed/seed.go -reset` to run initial migrations. (Optionally, run `source .env` if you want to change env vars)
4. run `go run src/main.go`

## Build Rust Binary

1. in `packages/note-parser`, create directory `dist`
2. run `cd scripts && ./build.sh`

The file `np` in `dist` is a binary executable that will make requests to the gRPC server you started earlier. Running the following from the `packages/note-parser` directory will parse a note in the `notes` directory and make a request to the gRPC server, which will store the data in your DB: `./dist/np --path ./notes/1740771329\ How\ to\ Take\ Smart\ Notes.md`

## Run Python code

1. in `packages/note-agent-system` run `hatch env create`
2. run `hatch run python src/note_agent_system/main.py` and you should see a note as output if you run the command from the Rust command in the previous instruction set. This was obtained via a request to the gRPC server from the first instruction set.

