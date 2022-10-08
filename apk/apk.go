// Package apk manage APK file information.
package apk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/shogo82148/androidbinary/apk"
)

var (
	// ErrNotOpenAPK : failed to open apk file
	ErrNotOpenAPK = errors.New("failed to open apk file")
)

// APK is android package (*.apk) file information.
type APK struct {
	// Path is path to apk file.
	Path string
	// Package is android package (*.apk) itself.
	Package *Package
}

// Package is Android Package (*.apk) information.
type Package struct {
	// Basic is basic information for apk file.
	Basic *Basic `json:"basic,omitempty"`
	// Metadata is metadata from AndroidManifest.
	Metadata []*Metadata `json:"metadata,omitempty"`
}

// Basic is basic information for apk file.
type Basic struct {
	// PackageName is package name (bundle id).
	PackageName string `json:"package_name,omitempty"`
	// ApplicationName is android application name.
	ApplicationName string `json:"application_name,omitempty"`
	// Version is application version.
	Version string `json:"version,omitempty"`
	// MainActivity is the activity that loads first and the rest of your application.
	MainActivity string `json:"main_activity,omitempty"`
	// SDK is android sdk information.
	SDK *SDK `json:"sdk,omitempty"`
}

// Metadata is metadata from AndroidManifest.
type Metadata struct {
	// Name is unique name to identify the value
	Name string `json:"name,omitempty"`
	// Value is value assigned to the item.
	Value string `json:"value,omitempty"`
}

// SDK is android sdk information
type SDK struct {
	// Minimun is supported minimum SDK versions
	Minimum int32 `json:"minimum,omitempty"`
	// Target is target SDK version.
	Target int32 `json:"target,omitempty"`
	// Maximum is supported maximum SDK versions. It's deprecated attribute from android 2.0.1.
	Maximum int32 `json:"maximum,omitempty"`
}

// NewAPK return APK struct.
func NewAPK(path string) *APK {
	return &APK{
		Path: path,
		Package: &Package{
			Basic: &Basic{
				SDK: &SDK{},
			},
			Metadata: []*Metadata{},
		},
	}
}

// Parse parses the files contained in the *.apk file
// and sets the metadata into an APK struct.
func (a *APK) Parse() error {
	apk, err := apk.OpenFile(a.Path)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrNotOpenAPK, a.Path)
	}
	defer apk.Close()

	a.setBasicInfo(apk)
	a.setMetadata(apk)
	return nil
}

// setBasicInfo extract basic information from the apk file and
// set its contents into an APK struct.
func (a *APK) setBasicInfo(apk *apk.Apk) {
	a.setSDK(apk)
	a.Package.Basic.PackageName = apk.PackageName()

	var err error
	a.Package.Basic.ApplicationName, err = apk.Label(nil)
	if err != nil {
		a.Package.Basic.ApplicationName = "(unknown)"
	}

	a.Package.Basic.Version, err = apk.Manifest().VersionName.String()
	if err != nil {
		a.Package.Basic.Version = "(unknown)"
	}

	a.Package.Basic.MainActivity, err = apk.MainActivity()
	if err != nil {
		a.Package.Basic.MainActivity = "(unknown)"
	}
}

func (a *APK) setSDK(apk *apk.Apk) {
	var err error
	a.Package.Basic.SDK.Minimum, err = apk.Manifest().SDK.Min.Int32()
	if err != nil {
		a.Package.Basic.SDK.Minimum = -1
	}

	a.Package.Basic.SDK.Target, err = apk.Manifest().SDK.Target.Int32()
	if err != nil {
		a.Package.Basic.SDK.Target = -1
	}

	a.Package.Basic.SDK.Maximum, err = apk.Manifest().SDK.Max.Int32()
	if err != nil || a.Package.Basic.SDK.Maximum == 0 {
		a.Package.Basic.SDK.Maximum = -1
	}
}

func (a *APK) setMetadata(apk *apk.Apk) {
	for _, v := range apk.Manifest().App.MetaData {
		name, err := v.Name.String()
		if err != nil {
			name = "(can not parse metadata name)"
		}

		value, err := v.Value.String()
		if err != nil {
			name = "(can not parse metadata value)"
		}
		a.Package.Metadata = append(a.Package.Metadata, &Metadata{Name: name, Value: value})
	}
}

// Print write apk information at io.Writer (e.g. STDOUT)
func (a *APK) Print(w io.Writer) {
	fmt.Fprintf(w, "[Application]\n")
	fmt.Fprintf(w, " name           : %s\n", a.Package.Basic.ApplicationName)
	fmt.Fprintf(w, " version        : %s\n", a.Package.Basic.Version)
	fmt.Fprintf(w, " main activity  : %s\n", a.Package.Basic.MainActivity)
	fmt.Fprintf(w, " package        : %s\n", a.Package.Basic.PackageName)
	fmt.Fprintf(w, "[SDK]\n")
	fmt.Fprintf(w, " target version : %d\n", a.Package.Basic.SDK.Target)
	fmt.Fprintf(w, " max version    : %d (deprecated attribute)\n", a.Package.Basic.SDK.Maximum)
	fmt.Fprintf(w, " min version    : %d\n", a.Package.Basic.SDK.Minimum)

	if len(a.Package.Metadata) != 0 {
		fmt.Fprintf(w, "[Metadata]\n")
	}
	for _, v := range a.Package.Metadata {
		fmt.Fprintf(w, "android:name       : %s\n", v.Name)
		fmt.Fprintf(w, "android:value      : %s\n", v.Value)
	}
}

// PrintJSON write apk information in json format
func (a *APK) PrintJSON(w io.Writer) error {
	j, err := json.MarshalIndent(a.Package, "", "\t")
	if err != nil {
		return err
	}
	fmt.Fprintln(w, string(j))
	return nil
}
