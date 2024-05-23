package tts

import (
	"context"
	"errors"
)

type Format string

const (
	FormatLINEAR16 Format = "audio/l16"
	FormatMP3      Format = "audio/mpeg"
	FormatOGG      Format = "audio/ogg"
	FormatALAW     Format = "audio/alaw"
	FormatMULAW    Format = "audio/mulaw"
	FormatAAC      Format = "audio/aac"
	FormatFLAC     Format = "audio/flac"
	FormatWAV      Format = "audio/wav"
)

type AudioFile struct {
	Format Format `json:"mime"`
	Data   []byte `json:"data"`
}

var ErrUnsupportedFileFormat = errors.New("unsupported file format")

type TTSModel interface {
	GenerateSpeech(ctx context.Context, text string, fmt Format) (*AudioFile, error)
}
