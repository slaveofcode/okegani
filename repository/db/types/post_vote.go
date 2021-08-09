package types

type PostVoteType string

const (
	// Positive
	PostVoteTypeSafe PostVoteType = "SAFE"

	// Middle
	PostVoteTypeBeCareful  PostVoteType = "BE_CAREFUL"
	PostVoteTypeSales      PostVoteType = "SALES"
	PostVoteTypeInsurance  PostVoteType = "INSURANCE"
	PostVoteTypeCreditCard PostVoteType = "CREDIT_CARD"
	PostVoteTypeSpam       PostVoteType = "SPAM"
	PostVoteTypeAnnoying   PostVoteType = "ANNOYING"

	// Negative
	PostVoteTypeDangerous  PostVoteType = "DANGEROUS"
	PostVoteTypeFraud      PostVoteType = "FRAUD"
	PostVoteTypeMisleading PostVoteType = "MISLEADING"
)
