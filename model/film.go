package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
)

type FilmStat struct {
	Id              string
	SourceFilePath  string `json:"SourceFile"`
	Name            string `json:"FileName"`
	Directory       string `json:"Directory"`
	Size            string `json:"FileSize"`
	Type            string `json:"FileType"`
	MIMEType        string `json:"MIMEType"`
	DocType         string `json:"DocType"`
	Duration        string `json:"Duration"`
	VideoCodecID    string `json:"VideoCodecID"`
	FrameRate       int    `json:"VideoFrameRate"`
	ImageWidth      int    `json:"ImageWidth"`
	ImageHeight     int    `json:"ImageHeight"`
	DisplayWidth    int    `json:"DisplayWidth"`
	DisplayHeight   int    `json:"DisplayHeight"`
	AudioCodecID    string `json:"AudioCodecID"`
	AudioSampleRate int    `json:"AudioSampleRate"`
	AudioChannels   int    `json:"AudioChannels"`
	TrackNumber     int    `json:"TrackNumber"`
	TrackType       string `json:"TrackType"`
	CodecID         string `json:"CodecID"`
	ImageSize       string `json:"ImageSize"`
}

func FromJSON(path string) (error, FilmStat) {
	row, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var f FilmStat
	err = json.Unmarshal(row, &f)
	f.Id = extractId(&f)

	return err, f
}

func extractId(f *FilmStat) string {
	return strings.ToLower(strings.Replace(f.Name, " ", ".", -1))
}

func (this FilmStat) ToJSON() string {
	b, err := json.Marshal(this)
	if err != nil {
		panic(err.Error())
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "   ")

	return out.String()
}

func (this FilmStat) String() string {
	return "TODO"
}
