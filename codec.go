package contree

type Codec interface {
	Decode([]byte) *NodeConf
	Encode(*NodeConf) []byte
}
