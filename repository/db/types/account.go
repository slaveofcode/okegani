package types

type AccountStatus string

const (
	AccountStatusRegistered AccountStatus = "REGISTERED"
	AccountStatusVerified   AccountStatus = "VERIFIED"
	AccountStatusSuspended  AccountStatus = "SUSPENDED"
)
