package types

type PostStatus string
type PostVisibility string

const (
	PostVisibilityDraft     PostVisibility = "DRAFT"
	PostVisibilityPublished PostVisibility = "PUBLISHED"

	PostStatusOpen        PostStatus = "OPEN"
	PostStatusSolved      PostStatus = "SOLVED"
	PostStatusControversy PostStatus = "CONTROVERSY"
	PostStatusSuspended   PostStatus = "SUSPENDED"
)
