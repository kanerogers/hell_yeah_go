package private

import (
	"github.com/AdRoll/goamz/dynamodb"
)

type Table interface {
	GetDocument(key *dynamodb.Key, v interface{}) error
}
