package models

import ()

type Image struct {
	UUID       string   `json:"UUID"`       //
	ImageId    string   `json:"imageid"`    //
	JSON       string   `json:"json"`       //
	Ancestry   string   `json:"ancestry"`   //
	Checksum   string   `json:"checksum"`   // tarsum+sha256
	Payload    string   `json:"payload"`    //
	URL        string   `json:"url"`        //
	Backend    string   `json:"backend"`    //
	Path       string   `json:"path"`       //
	Sign       string   `json:"sign"`       //
	Size       int64    `json:"size"`       //
	Uploaded   bool     `json:"uploaded"`   //
	Checksumed bool     `json:"checksumed"` //
	Encrypted  bool     `json:"encrypted"`  //
	Created    int64    `json:"created"`    //
	Updated    int64    `json:"updated"`    //
	Memo       []string `json:"memo"`       //
	Version    int64    `json:"version"`    //
}

func (i *Image) Has(image string) (bool, string, error) {
	UUID, err := GetUUID("image", image)
	if err != nil {
		return false, "", err
	}
	if len(UUID) <= 0 {
		return false, "", nil
	}

	err = Get(i, UUID)

	return true, UUID, err
}
