package validate

type Integers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type IntValidator[T Integers] func(v T) error

func (f IntValidator[T]) Evaluate(n T) error { return f(n) }

type IntComposer[T Integers] func(f IntValidator[T]) IntValidator[T]

func Int[T Integers]() IntComposer[T] {
	return Identity[IntValidator[T]]
}

func (f IntComposer[T]) and(second IntValidator[T]) IntComposer[T] {
	return func(first IntValidator[T]) IntValidator[T] {
		return func(n T) error {
			if err := f(first).Evaluate(n); nil != err {
				return err
			}

			return second.Evaluate(n)
		}
	}
}

func (f IntComposer[T]) Compose() IntValidator[T] { return f(nop[T]()) }
func (f IntComposer[T]) Min(min T) IntComposer[T] { return f.and(minimum("integer_min", min)) }
func (f IntComposer[T]) Max(max T) IntComposer[T] { return f.and(maximum("integer_max", max)) }
func (f IntComposer[T]) Choose(choices ...T) IntComposer[T] {
	return f.and(choose("integer_choose", choices))
}
