//go:build wireinject

package main

import "github.com/google/wire"

func InitializeApp() (*App, error){
	wire.Build(wire.NewSet(NewDB), wire.NewSet(NewApp))
	return nil, nil
}
