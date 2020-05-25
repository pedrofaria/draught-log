package formatter

type Config struct {
	MessageField          string
	LevelField            string
	TimestampField        string
	TimestampFormat       string
	PreFilterRegex        string
	PreFilterRegexReplace string
}
