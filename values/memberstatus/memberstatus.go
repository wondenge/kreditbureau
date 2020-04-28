package memberstatus

// Group Guarantee Member Status
type (
	MemberStatus string
)

const (
	A MemberStatus = "Active Member of the Group"
	B              = "Dormant Member of the Group"
	C              = "Exited the Group"
)
