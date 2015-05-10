package common

import (
	"github.com/AdRoll/goamz/dynamodb"
)

type Table interface {
	GetDocument(key *dynamodb.Key, v interface{}) error
	PutDocument(key *dynamodb.Key, v interface{}) error
	DeleteDocument(key *dynamodb.Key) error
}
