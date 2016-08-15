#go-api

simple API written in go.

##Depenencies
to run you need go installed and `GOPATH` set.

##Run
```bash
go run *.go
```

if you want to specify a port different than 8080
```bash
PORT=3000 go run *.go
```

##to-do

- [ ] add UPDATE & DELETE routes.
- [ ] add a database to save the data.
- [ ] add JWT auth.
- [ ] dockerize.
- [ ] separate helpers & utils into packages.
- [ ] add tests.
