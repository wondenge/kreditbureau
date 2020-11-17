package internal

type ResidencyType string

const (
	Owned       ResidencyType = "A – Owned"
	Rented                    = "B – Rented"
	Mortgaged                 = "C – Mortgaged"
	FamilyOwned               = "D – Family Owned"
)
