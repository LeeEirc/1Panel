package log

import (
	"errors"
	"io"
	"os"
	"path"
)

const (
	RollingTimePattern = "0 0  * * *"
)

var (
	BufferSize         = 0x100000
	DefaultFileMode    = os.FileMode(0644)
	DefaultFileFlag    = os.O_RDWR | os.O_CREATE | os.O_APPEND
	ErrInvalidArgument = errors.New("error argument invalid")
	QueueSize          = 1024
	ErrClosed          = errors.New("error write on close")
)

type Config struct {
	TimeTagFormat string
	LogPath       string
	FileName      string
	MaxRemain     int
}

type Manager interface {
	Fire() chan string
	Close()
}

type RollingWriter interface {
	io.Writer
	Close() error
}

func FilePath(c *Config) (filepath string) {
	filepath = path.Join(c.LogPath, c.FileName) + ".log"
	return
}