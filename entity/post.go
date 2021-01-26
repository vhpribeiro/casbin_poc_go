package entity

type Post struct {
	ID    int64  `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Text  string `json:"text" bson:"text"`
}
