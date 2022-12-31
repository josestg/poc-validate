package validate

type Floats interface {
	~float32 | ~float64
}

type FloatComposer[T Floats] func(f Validator[T]) Validator[T]

func Float[T Floats]() FloatComposer[T] { return Identity[Validator[T]] }

func (f FloatComposer[T]) and(second Validator[T]) FloatComposer[T] {
	return func(first Validator[T]) Validator[T] {
		return func(n T) error {
			if err := f(first).Evaluate(n); nil != err {
				return err
			}

			return second.Evaluate(n)
		}
	}
}

func (f FloatComposer[T]) Compose() Validator[T]      { return f(nop[T]()) }
func (f FloatComposer[T]) Min(min T) FloatComposer[T] { return f.and(minimum("float_min", min)) }
func (f FloatComposer[T]) Max(max T) FloatComposer[T] { return f.and(maximum("float_max", max)) }
func (f FloatComposer[T]) Choose(choices ...T) FloatComposer[T] {
	return f.and(choose("float_choose", choices))
}
