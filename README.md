# micro-gen
ispecのマイクロサービスで使われるクリーンアーキテクチャのためのコード生成CLIです。コマンドを叩けば各レイヤーの構造体、インターフェース、コンストラクターが生成されます。

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

```controllers/user.go
package controllers

type UserController interface {
	CreateUser(c context.Context) (err error)
}
```

```usecases/user.go
package usecases

type UserUsecase interface {
	CreateUser(ipt inputs.CreateUser) (opt outputs.CreateUser)
}
type userUsecaseInteractor struct{}

func NewUserUsecaseInteractor() UserUsecaseInteractor {}

func (i userUsecaseInteractor) CreateUser(ipt inputs.CreateUser) (opt outputs.CreateUser) {}

```

