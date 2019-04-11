# micro-gen
Code generation CLI for clean architecture used by microservices. When you run a command, structures, interfaces and constructors for each layer are generated.

## Installation

```
go get -u github.com/ispec-inc/micro-gen
```

## Usage
```
$ micro-gen scafflod User CreateUser
create controllers/user.go
create usecases/user.go
create usecases/inputs/user.go
create usecases/outputs/user.go

$ micro-gen entity User
create entities/models/admin.go
create entities/repositories/admin.go
```

The generated `controllers/user.go` is below.

```controllers/user.go
package controllers

type UserController interface {
	CreateUser(c context.Context) (err error)
}
```

The generated `usecases/user.go` is below.
```usecases/user.go
package usecases

type UserUsecase interface {
	CreateUser(ipt inputs.CreateUser) (opt outputs.CreateUser)
}
type userUsecaseInteractor struct{}

func NewUserUsecaseInteractor() UserUsecaseInteractor {}

func (i userUsecaseInteractor) CreateUser(ipt inputs.CreateUser) (opt outputs.CreateUser) {}

```

