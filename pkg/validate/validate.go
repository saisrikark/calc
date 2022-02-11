package validate

type OpType int

const (
	Add OpType = 1
	Div OpType = 2
	Mul OpType = 3
	Sub OpType = 4
)

func IsValid(op OpType, args []string) (bool, error) {
	argsLen := len(args)
	if argsLen < 2 {
		return false, insufficientArgsErr
	} else if !(op == Add || op == Mul) && argsLen > 2 {
		return false, excessiveArgsErr
	}
	return true, nil
}
