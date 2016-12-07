package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

type FilmStat struct {
	// non-json field
	Id         string
	PrittyName string
	// json field
	SourceFilePath  string      `json:"SourceFile"`
	Name            string      `json:"FileName"`
	Directory       string      `json:"Directory"`
	Size            string      `json:"FileSize"`
	Type            string      `json:"FileType"`
	MIMEType        string      `json:"MIMEType"`
	DocType         string      `json:"DocType"`
	Duration        string      `json:"Duration"`
	VideoCodecID    string      `json:"VideoCodec"`
	FrameRate       float32     `json:"VideoFrameRate"`
	ImageWidth      int         `json:"ImageWidth"`
	ImageHeight     int         `json:"ImageHeight"`
	Compression     string      `json:"Compression"`
	Encoding        string      `json:"Encoding"`
	DisplayWidth    int         `json:"DisplayWidth"`
	DisplayHeight   int         `json:"DisplayHeight"`
	AudioCodecID    json.Number `json:"AudioCodecID,Number"`
	AudioSampleRate float32     `json:"AudioSampleRate"`
	AudioChannels   json.Number `json:"AudioChannels,Number"`
	TrackNumber     int         `json:"TrackNumber"`
	TrackType       string      `json:"TrackType"`
	CodecID         string      `json:"CodecID"`
	ImageSize       string      `json:"ImageSize"`
}

func FromJSON(path string) (error, FilmStat) {
	row, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var f FilmStat
	err = json.Unmarshal(row, &f)
	f.Id = extractId(&f)
	f.PrittyName = (strings.Split(f.Name, "."))[0]

	return err, f
}

func GetMovieByID(id string, films []FilmStat) (FilmStat, error) {
	// TODO TESTING PURPOSE, I DO NOT LIKE THIS AT ALL
	for _, f := range films {
		if f.Id == id {
			return f, nil
		}
	}
	return FilmStat{}, errors.New("Film with id: " + id + " not found.")
}

func extractId(f *FilmStat) string {
	replacer := strings.NewReplacer(
		" ", "",
		",", "",
		"-", "",
	)

	return replacer.Replace(f.Name)
}

/*
func (this FilmStat) ToJSON() string {
	b, err := json.Marshal(this)
	if err != nil {
		panic(err.Error())
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "   ")

	return out.String()
}
*/

func (this FilmStat) String() string {
	return "TODO"
}
