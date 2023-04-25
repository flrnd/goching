package translation

// Hexagram is an I Ching hexagram object
type Hexagram struct {
	ID           int      `json:"id"`
	Symbol       string   `json:"symbol"`
	Name         string   `json:"name"`
	Above        string   `json:"above"`
	Below        string   `json:"below"`
	Judgment     string   `json:"judgment"`
	Image        string   `json:"image"`
	Lines        []string `json:"lines"`
	BinaryString string   `json:"binaryString"`
}

// Translation is an Array of Hexagrams
type Translation []struct {
	Hexagrams Hexagram
}
