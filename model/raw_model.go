package model

import (
	"encoding/json"
	"io/ioutil"
)

type StreamCodec struct {
	CodecName     string `json:"codec_name"`
	CodecLongName string `json:"codec_long_name"`
	CodecType     string `json:"codec_type"`

	// in case Codec Type is video
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	DisplayRatio string `json:"display_aspect_ratio"`

	// in case Codec Type is audio
	Channels       int    `json:"channels"`
	ChannelsLayout string `json:"channels_layout"`

	Tags struct {
		Language string `json:"language"`
		Title    string `json:"title"`
	}
}

type FFormat struct {
	FileName       string `json:"filename"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	Duration       string `json:"duration"`
	ByteSize       string `json:"size"`

	Tags struct {
		Encoder string `json:"encoder"`
	}
}

type RawMovie struct {
	Streams []StreamCodec `json:"streams"`
	Format  FFormat       `json:"format"`
}

func Parse(moviePath string) (*RawMovie, error) {
	// TODO check path
	row, err := ioutil.ReadFile(moviePath)
	if err != nil {
		return nil, err
	}

	var f RawMovie
	err = json.Unmarshal(row, &f)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func ParseAll(directory string) ([]RawMovie, error) {
	// TODO check if argument is a directory
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	rawMovies := make([]RawMovie, len(files))
	for _, file := range files {
		filmP, err := Parse(directory + file.Name())
		if err != nil {
			return nil, err
		}
		rawMovies = append(rawMovies, (*filmP))
	}

	return rawMovies, nil
}
