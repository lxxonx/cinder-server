package dto

type GetGroupsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreateGroupInput struct {
	GroupName string   `json:"groupName"`
	Friends   []string `json:"friends"`
	Bio       string   `json:"bio"`
}

type LikeGroupInput struct {
	GroupName string `json:"groupName"`
}
