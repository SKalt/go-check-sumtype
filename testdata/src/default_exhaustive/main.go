// want package:`.*sumTypeFact\{T \[A B\]\}.*`
package default_exhaustive

//sumtype:decl
type T interface{ sealed() }

type A struct{}

func (a *A) sealed() {}

type B struct{}

func (b *B) sealed() {}

func main() {
	switch T(nil).(type) {
	case *A:
	default:
		println("legit catch all goes here")
	}
}
