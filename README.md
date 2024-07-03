# fingo

Here goes a short and concise intro about the project.

## Setup

Install system runtimes with [asdf](https://asdf-vm.com):

```
asdf install
```

Then you should be able to set everything up with:

```
make setup
```

### Build

```
make deps.get
make build
```

It should build an executable at `./bin/fingo`.

### Run

It's possible to run the project without build it first (useful in development).

```
make run
```

## License

[Apache 2.0](LICENSE)
