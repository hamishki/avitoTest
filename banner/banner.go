package banner

import "encoding/json"

type Banner struct {
	Id        int    `json:"banner_id"`
	Data      []byte `json:"data"`
	FeatureID int    `json:"feature_id"`
	TagsIDs   []int  `json:"tag_ids"`
	IsActive  bool   `json:"is_active"`
}

func (b Banner) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

func (b Banner) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &b)
}
