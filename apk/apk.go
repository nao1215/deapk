// Package apk manage APK file information.
package apk

// APK is android package (*.apk) file information.
type APK struct {
	// Path is path to apk file.
	Path string
	// Package is android package (*.apk) itself.
	Package *Package
}

// Package is Android Package (*.apk) information.
type Package struct {
	// Name Android package(.apk) label.
	Name string
}

// NewAPK return APK struct.
func NewAPK(path string) *APK {
	return &APK{
		Path: path,
	}
}
