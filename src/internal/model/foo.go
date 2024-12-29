package model

import "time"

type Foo struct {
	// The unique identifier for the foo item
	ID string `json:"id" bson:"_id,omitempty" example:"5f7b5e1b9b0b3a1b3c9b0b3a"`
	// The name of the foo item
	Name string `json:"name" bson:"name" example:"Example Foo"`
	// The description of the foo item
	Description string `json:"description" bson:"description" example:"This is an example foo item"`
	// The creation timestamp
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
