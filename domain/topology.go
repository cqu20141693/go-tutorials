package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Star 点赞
type Star struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	TopologyID string             `json:"id" bson:"id"`
	UserID     string             `json:"userId" bson:"userId"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
}

type Topology struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`

	Name string `json:"name"`
	Desc string `json:"desc"`

	Data     interface{} `json:"data"`
	UserID   string      `json:"userId" bson:"userId"`
	Username string      `json:"username" `

	Shared bool   `json:"shared"`
	Star   uint64 `json:"star" bson:"star,omitempty"`

	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"-" bson:"deletedAt,omitempty"`
}
