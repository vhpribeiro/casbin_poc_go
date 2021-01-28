package dtos

type UserObjectActionDto struct {
	User   string `json:"user"`
	Object string `json:"object"`
	Action string `json:"action"`
}
