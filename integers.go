package validate

type Integers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type IntComposer[T Integers] func(f Validator[T]) Validator[T]

func Int[T Integers]() IntComposer[T] {
	return Identity[Validator[T]]
}

func (f IntComposer[T]) and(second Validator[T]) IntComposer[T] {
	return func(first Validator[T]) Validator[T] {
		return mergeValidator(f(first), second)
	}
}

func (f IntComposer[T]) Compose() Validator[T]    { return f(nop[T]()) }
func (f IntComposer[T]) Min(min T) IntComposer[T] { return f.and(minimum("integer_min", min)) }
func (f IntComposer[T]) Max(max T) IntComposer[T] { return f.and(maximum("integer_max", max)) }
func (f IntComposer[T]) Choose(choices ...T) IntComposer[T] {
	return f.and(choose("integer_choose", choices))
}
