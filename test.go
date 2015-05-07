package main

import (
	// "github.com/gin-gonic/contrib/jwt"
	"github.com/kanerogers/hell_yeah_go/dynamo_crud"
	"github.com/kanerogers/hell_yeah_go/routing"
)

func main() {

	tableMaker := new(dynamo_crud.DynamoTableMaker)

	router := routing.Core(tableMaker)

	router.Run(":8000")
}
