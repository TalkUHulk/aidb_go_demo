package aidb

type AIDBInput struct {
	FlowUUID string   `json:"flow_uuid"`
	Backend  []string `json:"backend"`
	Model    []string `json:"model"`
	Zoo      string   `json:"zoo"`
}

type AIDBFaceOutput struct {
	BBox     []float32   `json:"bbox"`
	Conf     float32     `json:"conf"`
	LandMark [][]float32 `json:"landmark"`
	Parsing  string      `json:"parsing"`
}

type AIDBObjectOutput struct {
	BBox  []float32 `json:"bbox"`
	Conf  float32   `json:"conf"`
	Label int       `json:"label"`
}

type AIDBOCROutput struct {
	Box        [][]float32 `json:"box"`
	Conf       float32     `json:"conf"`
	ConfRotate float32     `json:"conf_rotate"`
	Label      string      `json:"label"`
}

type AIDBClsOutput struct {
	Conf  float32 `json:"conf"`
	Label int     `json:"label"`
}

type AiDBOutput struct {
	Code      int                `json:"error_code"`
	Face      []AIDBFaceOutput   `json:"face"`
	Object    []AIDBObjectOutput `json:"object"`
	Ocr       []AIDBOCROutput    `json:"ocr"`
	Cls       []AIDBClsOutput    `json:"cls"`
	Anime     string             `json:"anime"`
	Tddfa     string             `json:"tddfa"`
	KeyPoints [][]float32        `json:"key_points"`
}
