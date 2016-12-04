package model

type FilmStat struct {
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
	DisplayHeight   int    `json:"DisplayHeigth"`
	AudioCodecID    string `json:"AudioCodecID"`
	AudioSampleRate int    `json:"AudioSampleRate"`
	AudioChannel    int    `json:"AudioChannel"`
	TrackNumber     int    `json:"TrackNumber"`
	TrackType       string `json:"TrackType"`
	CodecID         string `json:"CodecID"`
	ImageSize       string `json:"ImageSize"`
}
