package dtos

type UserRoleDomainAttributeDto struct {
	User      string `json:"user"`
	Role      string `json:"role"`
	Domain    string `json:"domain"`
	Attribute string `json:"attribute"`
}
