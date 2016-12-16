package model

import (
	"encoding/json"
	"io/ioutil"
)

// TODO add more doc:
type RawJson struct {
	SourceFilePath  string          `json:"SourceFile"`
	Name            string          `json:"FileName"`
	Directory       string          `json:"Directory"`
	Size            string          `json:"FileSize"`
	Type            string          `json:"FileType"`
	MIMEType        string          `json:"MIMEType"`
	Duration        string          `json:"Duration"`
	VideoCodecRaw   json.RawMessage `json:"VideoCodec,json.RawMessage"`
	VideoCodecIDRaw json.RawMessage `json:"VideoCodecID,json.RawMessage"`
	FrameRate       float32         `json:"VideoFrameRate,float32"`
	DisplayWidth    float32         `json:"DisplayWidth,float32"`
	DisplayHeight   float32         `json:"DisplayHeight,float32"`
	ImageWidth      float32         `json:"ImageWidth,float32"`
	ImageHeight     float32         `json:"ImageHeight,float32"`
	ImageSize       string          `json:"ImageSize"`
	AudioChannels   json.RawMessage `json:"AudioChannels,json.RawMessage"`

	// TODO try to read trivial entry as json.RawMessage (or similar) and try to
	// cast it by hand ?
}

// TODO add description
func readRawJSON(path string) (error, RawJson) {
	row, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var raw RawJson
	err = json.Unmarshal(row, &raw)
	if err != nil {
		return err, RawJson{}
	}

	return nil, raw
}
