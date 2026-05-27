// Copyright 2025 samber.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://github.com/samber/ro/blob/main/licenses/LICENSE.apache.md
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ro

// Pipe builds a composition of operators that will be chained to transform
// an observable stream. It provides a clean, declarative way to describe
// complex asynchronous operations.
//
// `PipeX()` should be favored over `Pipe()`, because it offers more type-safety.
//
// `PipeOp()` is the operator version of `Pipe()`.
func Pipe[First, Last any](source Observable[First], operators ...any) Observable[Last] {
	_ = "STUB: not implemented"
	return nil
}

// Since generic type can vary for each operator, we decided to use reflection and validate
// types at runtime. Type is not check for each message but only at Pipe() call.
// This is a peace of shit. If anybody find a better way to do it, please contribute!!

// check operator is a function with 1 input and 1 output

// check operator input implements Observable[T]

// check operator output implements Observable[T]

// check operator input implements source Observable[T]

// check operator output implements destination Observable[T]

// Pipe1 is a typesafe 🎉 implementation of Pipe, that takes a source and 1 operator.
//
// `PipeOp1()` is the operator version of `Pipe1()`.
func Pipe1[A, B any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
) Observable[B] {
	_ = "STUB: not implemented"
	return nil

	// Pipe2 is a typesafe 🎉 implementation of Pipe, that takes a source and 2 operators.
	//
	// `PipeOp2()` is the operator version of `Pipe2()`.
}

func Pipe2[A, B, C any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
) Observable[C] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe3 is a typesafe 🎉 implementation of Pipe, that takes a source and 3 operators.
//
// `PipeOp3()` is the operator version of `Pipe3()`.
func Pipe3[A, B, C, D any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
) Observable[D] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe4 is a typesafe 🎉 implementation of Pipe, that takes a source and 4 operators.
//
// `PipeOp4()` is the operator version of `Pipe4()`.
func Pipe4[A, B, C, D, E any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
) Observable[E] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe5 is a typesafe 🎉 implementation of Pipe, that takes a source and 5 operators.
//
// `PipeOp5()` is the operator version of `Pipe5()`.
func Pipe5[A, B, C, D, E, F any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
) Observable[F] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe6 is a typesafe 🎉 implementation of Pipe, that takes a source and 6 operators.
//
// `PipeOp6()` is the operator version of `Pipe6()`.
func Pipe6[A, B, C, D, E, F, G any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
) Observable[G] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe7 is a typesafe 🎉 implementation of Pipe, that takes a source and 7 operators.
//
// `PipeOp7()` is the operator version of `Pipe7()`.
func Pipe7[A, B, C, D, E, F, G, H any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
) Observable[H] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe8 is a typesafe 🎉 implementation of Pipe, that takes a source and 8 operators.
//
// `PipeOp8()` is the operator version of `Pipe8()`.
func Pipe8[A, B, C, D, E, F, G, H, I any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
) Observable[I] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe9 is a typesafe 🎉 implementation of Pipe, that takes a source and 9 operators.
//
// `PipeOp9()` is the operator version of `Pipe9()`.
func Pipe9[A, B, C, D, E, F, G, H, I, J any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
) Observable[J] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe10 is a typesafe 🎉 implementation of Pipe, that takes a source and 10 operators.
//
// `PipeOp10()` is the operator version of `Pipe10()`.
func Pipe10[A, B, C, D, E, F, G, H, I, J, K any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
) Observable[K] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe11 is a typesafe 🎉 implementation of Pipe, that takes a source and 11 operators.
//
// `PipeOp11()` is the operator version of `Pipe11()`.
func Pipe11[A, B, C, D, E, F, G, H, I, J, K, L any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
) Observable[L] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe12 is a typesafe 🎉 implementation of Pipe, that takes a source and 12 operators.
//
// `PipeOp12()` is the operator version of `Pipe12()`.
func Pipe12[A, B, C, D, E, F, G, H, I, J, K, L, M any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
) Observable[M] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe13 is a typesafe 🎉 implementation of Pipe, that takes a source and 13 operators.
//
// `PipeOp13()` is the operator version of `Pipe13()`.
func Pipe13[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
) Observable[N] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe14 is a typesafe 🎉 implementation of Pipe, that takes a source and 14 operators.
//
// `PipeOp14()` is the operator version of `Pipe14()`.
func Pipe14[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
) Observable[O] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe15 is a typesafe 🎉 implementation of Pipe, that takes a source and 15 operators.
//
// `PipeOp15()` is the operator version of `Pipe15()`.
func Pipe15[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
) Observable[P] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe16 is a typesafe 🎉 implementation of Pipe, that takes a source and 16 operators.
//
// `PipeOp16()` is the operator version of `Pipe16()`.
func Pipe16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
) Observable[Q] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe17 is a typesafe 🎉 implementation of Pipe, that takes a source and 17 operators.
//
// `PipeOp17()` is the operator version of `Pipe17()`.
func Pipe17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
) Observable[R] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe18 is a typesafe 🎉 implementation of Pipe, that takes a source and 18 operators.
//
// `PipeOp18()` is the operator version of `Pipe18()`.
func Pipe18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
) Observable[S] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe19 is a typesafe 🎉 implementation of Pipe, that takes a source and 19 operators.
//
// `PipeOp19()` is the operator version of `Pipe19()`.
func Pipe19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe20 is a typesafe 🎉 implementation of Pipe, that takes a source and 20 operators.
//
// `PipeOp20()` is the operator version of `Pipe20()`.
func Pipe20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
) Observable[U] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe21 is a typesafe 🎉 implementation of Pipe, that takes a source and 21 operators.
//
// `PipeOp21()` is the operator version of `Pipe21()`.
func Pipe21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
) Observable[V] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe22 is a typesafe 🎉 implementation of Pipe, that takes a source and 22 operators.
//
// `PipeOp22()` is the operator version of `Pipe22()`.
func Pipe22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
	operator22 func(Observable[V]) Observable[W],
) Observable[W] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe23 is a typesafe 🎉 implementation of Pipe, that takes a source and 23 operators.
//
// `PipeOp23()` is the operator version of `Pipe23()`.
func Pipe23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
	operator22 func(Observable[V]) Observable[W],
	operator23 func(Observable[W]) Observable[X],
) Observable[X] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe24 is a typesafe 🎉 implementation of Pipe, that takes a source and 24 operators.
//
// `PipeOp24()` is the operator version of `Pipe24()`.
func Pipe24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
	operator22 func(Observable[V]) Observable[W],
	operator23 func(Observable[W]) Observable[X],
	operator24 func(Observable[X]) Observable[Y],
) Observable[Y] {
	_ = "STUB: not implemented"
	return nil
}

// Pipe25 is a typesafe 🎉 implementation of Pipe, that takes a source and 25 operators.
//
// `PipeOp25()` is the operator version of `Pipe25()`.
func Pipe25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z any](
	source Observable[A],
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
	operator22 func(Observable[V]) Observable[W],
	operator23 func(Observable[W]) Observable[X],
	operator24 func(Observable[X]) Observable[Y],
	operator25 func(Observable[Y]) Observable[Z],
) Observable[Z] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp is similar to Pipe, but can be used as an operator.
func PipeOp[First, Last any](operators ...any) func(Observable[First]) Observable[Last] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp1 is similar to Pipe1, but can be used as an operator.
func PipeOp1[A, B any](
	operator1 func(Observable[A]) Observable[B],
) func(Observable[A]) Observable[B] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp2 is similar to Pipe2, but can be used as an operator.
func PipeOp2[A, B, C any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
) func(Observable[A]) Observable[C] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp3 is similar to Pipe3, but can be used as an operator.
func PipeOp3[A, B, C, D any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
) func(Observable[A]) Observable[D] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp4 is similar to Pipe4, but can be used as an operator.
func PipeOp4[A, B, C, D, E any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
) func(Observable[A]) Observable[E] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp5 is similar to Pipe5, but can be used as an operator.
func PipeOp5[A, B, C, D, E, F any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
) func(Observable[A]) Observable[F] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp6 is similar to Pipe6, but can be used as an operator.
func PipeOp6[A, B, C, D, E, F, G any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
) func(Observable[A]) Observable[G] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp7 is similar to Pipe7, but can be used as an operator.
func PipeOp7[A, B, C, D, E, F, G, H any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
) func(Observable[A]) Observable[H] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp8 is similar to Pipe8, but can be used as an operator.
func PipeOp8[A, B, C, D, E, F, G, H, I any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
) func(Observable[A]) Observable[I] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp9 is similar to Pipe9, but can be used as an operator.
func PipeOp9[A, B, C, D, E, F, G, H, I, J any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
) func(Observable[A]) Observable[J] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp10 is similar to Pipe10, but can be used as an operator.
func PipeOp10[A, B, C, D, E, F, G, H, I, J, K any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
) func(Observable[A]) Observable[K] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp11 is similar to Pipe11, but can be used as an operator.
func PipeOp11[A, B, C, D, E, F, G, H, I, J, K, L any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
) func(Observable[A]) Observable[L] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp12 is similar to Pipe12, but can be used as an operator.
func PipeOp12[A, B, C, D, E, F, G, H, I, J, K, L, M any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
) func(Observable[A]) Observable[M] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp13 is similar to Pipe13, but can be used as an operator.
func PipeOp13[A, B, C, D, E, F, G, H, I, J, K, L, M, N any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
) func(Observable[A]) Observable[N] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp14 is similar to Pipe14, but can be used as an operator.
func PipeOp14[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
) func(Observable[A]) Observable[O] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp15 is similar to Pipe15, but can be used as an operator.
func PipeOp15[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
) func(Observable[A]) Observable[P] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp16 is similar to Pipe16, but can be used as an operator.
func PipeOp16[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
) func(Observable[A]) Observable[Q] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp17 is similar to Pipe17, but can be used as an operator.
func PipeOp17[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
) func(Observable[A]) Observable[R] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp18 is similar to Pipe18, but can be used as an operator.
func PipeOp18[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
) func(Observable[A]) Observable[S] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp19 is similar to Pipe19, but can be used as an operator.
func PipeOp19[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
) func(Observable[A]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp20 is similar to Pipe20, but can be used as an operator.
func PipeOp20[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
) func(Observable[A]) Observable[U] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp21 is similar to Pipe21, but can be used as an operator.
func PipeOp21[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
) func(Observable[A]) Observable[V] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp22 is similar to Pipe22, but can be used as an operator.
func PipeOp22[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
	operator22 func(Observable[V]) Observable[W],
) func(Observable[A]) Observable[W] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp23 is similar to Pipe23, but can be used as an operator.
func PipeOp23[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
	operator22 func(Observable[V]) Observable[W],
	operator23 func(Observable[W]) Observable[X],
) func(Observable[A]) Observable[X] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp24 is similar to Pipe24, but can be used as an operator.
func PipeOp24[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
	operator22 func(Observable[V]) Observable[W],
	operator23 func(Observable[W]) Observable[X],
	operator24 func(Observable[X]) Observable[Y],
) func(Observable[A]) Observable[Y] {
	_ = "STUB: not implemented"
	return nil
}

// PipeOp25 is similar to Pipe25, but can be used as an operator.
func PipeOp25[A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z any](
	operator1 func(Observable[A]) Observable[B],
	operator2 func(Observable[B]) Observable[C],
	operator3 func(Observable[C]) Observable[D],
	operator4 func(Observable[D]) Observable[E],
	operator5 func(Observable[E]) Observable[F],
	operator6 func(Observable[F]) Observable[G],
	operator7 func(Observable[G]) Observable[H],
	operator8 func(Observable[H]) Observable[I],
	operator9 func(Observable[I]) Observable[J],
	operator10 func(Observable[J]) Observable[K],
	operator11 func(Observable[K]) Observable[L],
	operator12 func(Observable[L]) Observable[M],
	operator13 func(Observable[M]) Observable[N],
	operator14 func(Observable[N]) Observable[O],
	operator15 func(Observable[O]) Observable[P],
	operator16 func(Observable[P]) Observable[Q],
	operator17 func(Observable[Q]) Observable[R],
	operator18 func(Observable[R]) Observable[S],
	operator19 func(Observable[S]) Observable[T],
	operator20 func(Observable[T]) Observable[U],
	operator21 func(Observable[U]) Observable[V],
	operator22 func(Observable[V]) Observable[W],
	operator23 func(Observable[W]) Observable[X],
	operator24 func(Observable[X]) Observable[Y],
	operator25 func(Observable[Y]) Observable[Z],
) func(Observable[A]) Observable[Z] {
	_ = "STUB: not implemented"
	return nil
}
