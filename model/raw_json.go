package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// TODO add more doc:
type RawJson struct {
	SourceFilePath      string          `json:"SourceFile"`
	Name                string          `json:"FileName"`
	Directory           string          `json:"Directory"`
	Size                string          `json:"FileSize"`
	Type                string          `json:"FileType"`
	MIMEType            string          `json:"MIMEType"`
	Duration            string          `json:"Duration"`
	VideoCodec          string          `json:"VideoCodec"`
	FrameRate           float32         `json:"VideoFrameRate,float32"`
	DisplayWidth        float32         `json:"DisplayWidth,float32"`
	DisplayHeight       float32         `json:"DisplayHeight,float32"`
	ImageWidth          float32         `json:"ImageWidth,float32"`
	ImageHeight         float32         `json:"ImageHeight,float32"`
	ImageSize           string          `json:"ImageSize"`
	AudioChannels       json.RawMessage `json:"AudioChannels,json.RawMessage"`
	AudioChannelsString string

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
	raw.AudioChannelsString = string(raw.AudioChannels)
	fmt.Println(string(raw.AudioChannels))

	return nil, raw
}
