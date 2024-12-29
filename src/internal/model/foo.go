package model

type Foo struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	CreatedAt string `json:"createdAt" bson:"createdAt"`
	UpdatedAt string `json:"updatedAt" bson:"updatedAt"`
}
