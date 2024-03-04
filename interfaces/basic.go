package interfaces

type Saver interface {
	Save() error
}

type Outputable interface {
	Saver
	Display()
}
