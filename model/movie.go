package model

type Movie struct {
	Id             string
	PrittyName     string
	ActualLocation string
	FileName       string
	Encoder        string
	Duration       string
	Size           string

	Video struct {
		Width  int
		Height int
		Ratio  string

		CodecName     string
		CodecLongName string
	}

	// Audio section

	// Subs section
}
