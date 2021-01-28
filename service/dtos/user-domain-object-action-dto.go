package dtos

type UserDomainObjectActionDto struct {
	User   string `json:"user"`
	Domain string `json:"domain"`
	Object string `json:"object"`
	Action string `json:"action"`
}
