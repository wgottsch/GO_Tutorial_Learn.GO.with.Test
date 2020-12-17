package errorhandling

// Refactor to constant errors
// before:

//Anywhere in different go Files
//var ErrNoOne = errors.New("This is Error number One")
//var ErrNoZwo = errors.New("This is Error number Zwo")
//var ErrNoThree = errors.New("This is Error number Three")
//var ErrNoFour = errors.New("This is Error number Four")
//var ErrNoFive = errors.New("This is Error number Five")
//var ErrNoSix = errors.New("This is Error number Six")

// after:
//

//const abc = errors.New("Not possible")

const (
	// comment: Statuscodes 100
	ErrNoOne = HandleErr("This is Error number One")
	ErrNoZwo = HandleErr("This is Error number Zwo")

	// comment: Statuscodes 200
	ErrNoThree = HandleErr("This is Error number Three")
	ErrNoFour  = HandleErr("This is Error number Four")

	//comment: Statuscodes 300
	ErrNoFive = HandleErr("This is Error number Five")
	ErrNoSix  = HandleErr("This is Error number Six")
)

type HandleErr string

func (e HandleErr) Error() string {
	return string(e)
}
