package model

import "time"

type Foo struct {
	// The unique identifier for the foo item
	ID string `json:"id,omitempty" bson:"_id,omitempty" example:"5f7b5e1b9b0b3a1b3c9b0b3a"`
	// The name of the foo item
	Name string `json:"name" bson:"name" example:"Example Foo" validate:"required,min=3,max=50"`
	// The description of the foo item
	Description string `json:"description" bson:"description" example:"This is an example foo item" validate:"required,min=10,max=500"`
	// The creation timestamp
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type CreateFooRequest struct {
	Name string `json:"name"`

	Description string `json:"description"`
}
