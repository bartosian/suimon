package ports

type Builder interface {
	Init() error
	Render() error
}
