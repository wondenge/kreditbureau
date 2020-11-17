package internal

type Group struct {
	GroupID   string
	GroupName string
	SubGroup  []SubGroup
}

type SubGroup struct {
	SubGroupID   string
	SubGroupName string
}


