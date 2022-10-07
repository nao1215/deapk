// Package android handles Android OS information.
package android

// API struct summarizes information for each android API level.
// API struct has codename, version, release date, API level.
type API struct {
	// Level is Android API level.
	Level uint64 `json:"level,omitempty"`
	// CodeName is Android OS codename that is different each version.
	CodeName string `json:"code_name,omitempty"`
	// Version is Android OS version.
	Version string `json:"version,omitempty"`
}

// NewAPI return API struct. The value set in the field depends
// on the API level (argument lv). For non-existent API levels,
// each field will be empty.
func NewAPI(lv uint64) *API {
	switch lv {
	case 1:
		return &API{
			Level:    1,
			CodeName: "(no code name)",
			Version:  "1.0",
		}
	case 2:
		return &API{
			Level:    2,
			CodeName: "(no code name)",
			Version:  "1.1",
		}
	case 3:
		return &API{
			Level:    3,
			CodeName: "Cupcake",
			Version:  "1.5",
		}
	case 4:
		return &API{
			Level:    4,
			CodeName: "Donut",
			Version:  "1.6",
		}
	case 5, 6, 7:
		return &API{
			Level:    lv,
			CodeName: "Eclair",
			Version:  "2.0 - 2.1",
		}
	case 8:
		return &API{
			Level:    8,
			CodeName: "Froyo",
			Version:  "2.2 - 2.2.3",
		}
	case 9, 10:
		return &API{
			Level:    lv,
			CodeName: "Gingerbread",
			Version:  "2.3 - 2.3.7",
		}
	case 11, 12, 13:
		return &API{
			Level:    lv,
			CodeName: "Honeycomb",
			Version:  "3.0 - 3.2.6",
		}
	case 14, 15:
		return &API{
			Level:    lv,
			CodeName: "Ice Cream Sandwich",
			Version:  "4.0 - 4.0.4",
		}
	case 16, 17, 18:
		return &API{
			Level:    lv,
			CodeName: "Jelly Bean",
			Version:  "4.1 - 4.3.1",
		}
	case 19, 20:
		return &API{
			Level:    lv,
			CodeName: "KitKat/4.4W",
			Version:  "4.4 - 4.4.4",
		}
	case 21, 22:
		return &API{
			Level:    lv,
			CodeName: "Lollipop",
			Version:  "5.0 - 5.1.1",
		}
	case 23:
		return &API{
			Level:    lv,
			CodeName: "Marshmallow",
			Version:  "6.0 - 6.0.1",
		}
	case 24, 25:
		return &API{
			Level:    lv,
			CodeName: "Nougat",
			Version:  "7.0 - 7.1.2",
		}
	case 26, 27:
		return &API{
			Level:    lv,
			CodeName: "Oreo",
			Version:  "8.0 - 8.1",
		}
	case 28:
		return &API{
			Level:    lv,
			CodeName: "Pie",
			Version:  "9",
		}
	case 29:
		return &API{
			Level:    lv,
			CodeName: "Q",
			Version:  "10",
		}
	case 30:
		return &API{
			Level:    lv,
			CodeName: "R",
			Version:  "11",
		}
	case 31:
		return &API{
			Level:    lv,
			CodeName: "S",
			Version:  "12",
		}
	case 32:
		return &API{
			Level:    lv,
			CodeName: "Sv2",
			Version:  "12L",
		}
	case 33:
		return &API{
			Level:    lv,
			CodeName: "Tiramisu",
			Version:  "13",
		}
	default:
		return &API{}
	}
}
