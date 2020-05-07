package core


type DumbRepository interface {
	Save(d *Dumb) error
}
