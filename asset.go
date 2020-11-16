package kreditbureau

type AssetClassification string

const (
	Normal      AssetClassification = "A = Normal (0-30)"
	Watch                           = "B = Watch (>30-90)"
	Substandard                     = "C = Substandard(>90-180)"
	Doubtful                        = "D = Doubtful(>180-360)"
	Loss                            = "E = Loss(>360)"
)
