package asset

type Kind string

const (
	Unknown  Kind = "unknown"
	Image    Kind = "image"
	Document Kind = "document"
	Audio    Kind = "audio"
	Video    Kind = "video"
	Archive  Kind = "archive"
)

var (
	Kind_name = map[int32]Kind{
		0: Unknown,
		1: Image,
		2: Document,
		3: Audio,
		4: Video,
		5: Archive,
	}
	Kind_value = map[Kind]int32{
		Unknown:  0,
		Image:    1,
		Document: 2,
		Audio:    3,
		Video:    4,
		Archive:  5,
	}
)

func (k *Kind) String() string {
	return string(*k)
}

func (k *Kind) Int32() int32 {
	return Kind_value[*k]
}

// kind, err := l.getAssetKind(req.GetAsset().GetKind())
// if err != nil {
// 	return err
// }
// file.Kind = kind
