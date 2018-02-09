# repo-error-status
Tool to inspect debian repositories and report result to influxdb.

# Dependency

- docker
- go & go packages

# Build
`
make
`

will produces an all-in-one docker image.

# Usage

`
docker run --rm IMAGE dump > config.json
`

to generate a configuration template file, edit it and use the configuraion with:

`
docker run --rm -v $PWD/config.json:/etc/repo-error-status/config.json IMAGE
`

