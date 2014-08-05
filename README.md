# Serve - Simple HTTP File Server

`serve` is a simple HTTP static file server to serve files rooted on the
current working directory.

## Usage

    serve [port]

        port is optional, default to 9000.

## Examples

    cd /path/to/dir
    serve

`serve` will start to serving files under `/path/to/dir` on localhost port
9000.

    cd /path/to/dir
    serve 9001

`serve` will start to serving files under `/path/to/dir` on localhost port
9001.
