package factory

type RocketInterface interface {
	GetName() string
	Start() bool
	Finish() bool
	Stop() bool
}

const (
	rocketNameRu  = "ru"
	rocketNameUSA = "usa"
	rocketNameCh  = "ch"
)

func NewRocketFactory(name string) RocketInterface {
	switch name {
	case rocketNameCh:
		return &RocketCh{}
		// return new(RocketCh)
	case rocketNameRu:
		return new(RocketRu)
	case rocketNameUSA:
		return new(RocketUSA)
	default:
		return nil
	}
}

type RocketRu struct{}

func (r *RocketRu) GetName() string { return rocketNameRu }
func (r *RocketRu) Start() bool     { return true }
func (r *RocketRu) Finish() bool    { return true }
func (r *RocketRu) Stop() bool      { return true }

type RocketCh struct{}

func (r *RocketCh) GetName() string { return rocketNameCh }
func (r *RocketCh) Start() bool     { return true }
func (r *RocketCh) Finish() bool    { return true }
func (r *RocketCh) Stop() bool      { return true }

type RocketUSA struct{}

func (r *RocketUSA) GetName() string { return rocketNameUSA }
func (r *RocketUSA) Start() bool     { return true }
func (r *RocketUSA) Finish() bool    { return true }
func (r *RocketUSA) Stop() bool      { return true }
