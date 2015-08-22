package raw

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

// RAF is the Fuji raw file format
type RAF struct {
	Header struct {
		Magic         [16]byte
		FormatVersion [4]byte
		CameraID      [8]byte
		Camera        [32]byte
		Dir           struct {
			Version [4]byte
			_       [20]byte
			Jpeg    struct {
				IDX int32
				Len int32
			}
		}
	}
	Jpeg []byte
	Exif *exif.Exif
}

// WriteJpeg writes the raw preview jpeg
func (r *RAF) WriteJpeg(w io.Writer) {
	w.Write(r.Jpeg)
}

// ReadRAF makes a new RAF from a file
func ReadRAF(fname string) *RAF {
	raf := new(RAF)

	f, err := os.Open(fname)
	defer f.Close()

	if err == nil {
		err = binary.Read(f, binary.BigEndian, &raf.Header)
	}

	jbuf := make([]byte, raf.Header.Dir.Jpeg.Len)
	f.ReadAt(jbuf, int64(raf.Header.Dir.Jpeg.IDX))
	raf.Jpeg = jbuf

	// Soon.
	// exif.RegisterParsers(mknote.Fuji)

	buf := bytes.NewBuffer(jbuf)
	raf.Exif, err = exif.Decode(buf)

	if err != nil {
		panic(err)
	}

	return raf
}
