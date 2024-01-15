package router

import (
	"github.com/WildEgor/g-cdn/internal/handlers"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(NewRouter, handlers.HandlersSet)
