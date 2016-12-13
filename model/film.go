package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

// TODO add more doc:
// TODO instead of using nested structs, separate if methods are needed
// non-json field
type Film struct {
	Id         string
	PrittyName string
	Duration   string
	Size       struct {
		Width  int
		Height int
	}
	File struct {
		Source     string
		Directory  string
		Size       string
		Type       string
		MimeType   string
		VideoCodec string
		FrameRate  float32
	}
}

// TODO add more doc:
// TODO separate this structure and related method to separated file
// json field
type FilmStat struct {
	SourceFilePath string  `json:"SourceFile"`
	Name           string  `json:"FileName"`
	Directory      string  `json:"Directory"`
	Size           string  `json:"FileSize"`
	Type           string  `json:"FileType"`
	MIMEType       string  `json:"MIMEType"`
	Duration       string  `json:"Duration"`
	VideoCodec     string  `json:"VideoCodec"`
	FrameRate      float32 `json:"VideoFrameRate,float32"`
	DisplayWidth   float32 `json:"DisplayWidth,float32"`
	DisplayHeight  float32 `json:"DisplayHeight,float32"`
	// TODO get also: ImageWidth, ImageHeight and ImageSize
	// TODO try to read trivial entry as json.RawMessage (or similar) and try to
	// cast it by hand ?
}

func extractFromJSON(jsonRaw FilmStat) Film {
	var f Film
	replacer := strings.NewReplacer(
		" ", "",
		",", "",
		"-", "",
	)
	f.Id = replacer.Replace(jsonRaw.Name)
	f.PrittyName = (strings.Split(jsonRaw.Name, "."))[0]
	f.Duration = jsonRaw.Duration
	f.Size.Width = int(jsonRaw.DisplayWidth)
	f.Size.Height = int(jsonRaw.DisplayHeight)
	f.File.Source = jsonRaw.SourceFilePath
	f.File.Directory = jsonRaw.Directory
	f.File.Size = jsonRaw.Size
	f.File.Type = jsonRaw.Type
	f.File.MimeType = jsonRaw.MIMEType
	f.File.VideoCodec = jsonRaw.VideoCodec
	f.File.FrameRate = jsonRaw.FrameRate

	return f
}

func FromJSON(path string) (error, Film) {
	row, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	var raw FilmStat
	err = json.Unmarshal(row, &raw)
	if err != nil {
		return err, Film{}
	}

	return nil, extractFromJSON(raw)
}

func GetMovieByID(id string, films []Film) (Film, error) {
	// TODO TESTING PURPOSE, I DO NOT LIKE THIS AT ALL
	for _, f := range films {
		if f.Id == id {
			return f, nil
		}
	}
	return Film{}, errors.New("Film with id: " + id + " not found.")
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
