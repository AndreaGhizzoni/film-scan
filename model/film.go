package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
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
	Audio struct {
		Channels string
	}
}

func extractFromJSON(jsonRaw RawJson) Film {
	var f Film
	replacer := strings.NewReplacer(
		" ", "",
		",", "",
		"-", "",
	)
	f.Id = replacer.Replace(jsonRaw.Name)
	f.PrittyName = (strings.Split(jsonRaw.Name, "."))[0]
	f.Duration = jsonRaw.Duration
	if jsonRaw.ImageWidth != 0 && jsonRaw.ImageHeight != 0 {
		f.Size.Width = int(jsonRaw.ImageWidth)
		f.Size.Height = int(jsonRaw.ImageHeight)
	} else if jsonRaw.DisplayWidth != 0 && jsonRaw.DisplayHeight != 0 {
		f.Size.Width = int(jsonRaw.DisplayWidth)
		f.Size.Height = int(jsonRaw.DisplayHeight)
	} else if jsonRaw.ImageSize != "" {
		split := strings.Split(jsonRaw.ImageSize, "x")
		width, _ := strconv.Atoi(split[0])
		height, _ := strconv.Atoi(split[1])
		f.Size.Width = width
		f.Size.Height = height
	}
	f.File.Source = jsonRaw.SourceFilePath
	f.File.Directory = jsonRaw.Directory
	f.File.Size = jsonRaw.Size
	f.File.Type = jsonRaw.Type
	f.File.MimeType = jsonRaw.MIMEType
	f.File.VideoCodec = jsonRaw.VideoCodec
	f.File.FrameRate = jsonRaw.FrameRate

	f.Audio.Channels = jsonRaw.AudioChannelsString

	return f
}

//TODO add doc
func FromJSON(path string) (error, Film) {
	err, rawJson := readRawJSON(path)
	if err != nil {
		return err, Film{}
	}

	return nil, extractFromJSON(rawJson)
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

// TODO add doc
func (this Film) ToJSON() string {
	b, err := json.Marshal(this)
	if err != nil {
		panic(err.Error())
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "   ")

	return out.String()
}

func (this Film) String() string {
	return "TODO"
}
