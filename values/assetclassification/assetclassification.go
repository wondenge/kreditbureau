package assetclassification

type (
	AssetClassification string
)

const (
	A AssetClassification = "Normal (0-30)"
	B                     = "Watch (>30-90)"
	C                     = "Substandard(>90-180)"
	D                     = "Doubtful(>180-360)"
	E                     = "Loss(>360)"
)
