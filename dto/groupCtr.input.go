package dto

type GetGroupsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreateGroupIput struct {
	GroupName string   `json:"groupName"`
	Friends   []string `json:"friends"`
	Pics      []string `json:"pics"`
	Bio       string   `json:"bio"`
}
