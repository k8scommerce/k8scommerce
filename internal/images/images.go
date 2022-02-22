package images

import "k8scommerce/internal/images/config"

func NewImageResizer(cfg *config.ImageResizeConfig) ImageResizer {
	return &imageResizer{}
}

type ImageResizer interface {
	Resize()
}

type imageResizer struct{}

func (ir *imageResizer) Resize() {

}
