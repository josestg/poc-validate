package validate

type Floats interface {
	~float32 | ~float64
}

type FloatValidator[T Floats] func(n T) error

func (f FloatValidator[T]) Evaluate(n T) error { return f(n) }

type FloatComposer[T Floats] func(f FloatValidator[T]) FloatValidator[T]

func Float[T Floats]() FloatComposer[T] { return Identity[FloatValidator[T]] }

func (f FloatComposer[T]) and(second FloatValidator[T]) FloatComposer[T] {
	return func(first FloatValidator[T]) FloatValidator[T] {
		return func(n T) error {
			if err := f(first).Evaluate(n); nil != err {
				return err
			}

			return second.Evaluate(n)
		}
	}
}

func (f FloatComposer[T]) Compose() FloatValidator[T] { return f(nop[T]()) }
func (f FloatComposer[T]) Min(min T) FloatComposer[T] { return f.and(minimum("float_min", min)) }
func (f FloatComposer[T]) Max(max T) FloatComposer[T] { return f.and(maximum("float_max", max)) }
func (f FloatComposer[T]) Choose(choices ...T) FloatComposer[T] {
	return f.and(choose("float_choose", choices))
}
