package config

type OutputType string

const (
	OutputJPEG OutputType = "jpeg"
	OutputPNG  OutputType = "png"
	OutputWEBP OutputType = "webp"
	OutputAUTO OutputType = "auto"
)

type ColorSpace string

const (
	BlackAndWhite ColorSpace = "bw"
	StandardRGB   ColorSpace = "srgb"
)

type ImageResizeConfig struct {
	EnableResizer   bool
	ImageSettings   []ImageSetting
	WatermarkConfig WatermarkConfig
}

type ImageSetting struct {
	SizeTag        string
	Width          int
	Height         int
	Quality        int
	PngCompression int `json:"PngCompression,omitempty,optional"`
	HasWatermark   bool
	OutputType     string
	ColorSpace     string `json:"ColorSpace,omitempty,optional"`
}

type WatermarkConfig struct {
	Margin         int     `json:"Margin,omitempty,optional"`
	DPI            int     `json:"DPI,omitempty,optional"`
	TextWidth      int     `json:"TextWidth,omitempty,optional"`
	Opacity        float64 `json:"Opacity,omitempty,optional"`
	NoReplicate    bool    `json:"NoReplicate,omitempty,optional"`
	Text           string  `json:"Text,omitempty,optional"`
	Font           string  `json:"Font,omitempty,optional"`
	RGBColor       string  `json:"RGBColor,omitempty,optional"`
	RemoteImageURL string  `json:"RemoteImageURL,omitempty,optional"`
}
