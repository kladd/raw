# raw - Decoding RAW images in Go

For now this is just for Fujifilm's RAW format. The header is read into a struct along with the embedded jpeg image. Exif data is parsed for the embedded jpeg.

```go
raf := raw.ReadRAF("/path/to/raw/file.RAF")

// Optionally write jpeg to disk
raf.WriteJpeg("/path/to/ouput.jpeg")

// Read some exif value
fl := raf.Exif.Get(exif.FocalLength)
```
