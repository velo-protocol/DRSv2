package logic

type Logic interface {
	Init(configFilePath string) error
}
