// Copyright 2025 samber.
//
// Licensed as an Enterprise License (the "License"); you may not use
// this file except in compliance with the License. You may obtain
// a copy of the License at:
//
// https://github.com/samber/ro/blob/main/licenses/LICENSE.ee.md
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package introspection

import (
	"go/ast"
)

const unsupportedExpr = "?"

type FunDesc struct {
	Name      string
	Pos       string
	Arguments []FunDescArgument
}

type FunDescArgument struct {
	Name string
	Pos  string
}

func (f *FunDescArgument) String() string { _ = "STUB: not implemented"; return "" }

func GetFunctionDescription(skipCaller int, skipArgs int) (*FunDesc, error) {
	_ = "STUB: not implemented"
	// Get caller information
	return nil, nil
}

// Parse the file

// Prepare output

// Find the specific call in the AST

// Extract function name

// Extract arguments

func getArgName(arg ast.Expr) string { _ = "STUB: not implemented"; return "" }

///////////////// Terminal args

// eg:
// PipeX(
//   source,
//   myOperator,
//   myPkg.myOperators.map,
// }

// eg:
// PipeX(
//   source,
//   myOperators["map"],
//   Map[int],
// }

// eg:
// PipeX(
//   source,
//   func() { ... },
// }

///////////////// Recursive args

// eg:
// PipeX(
//   source,
//   (myOperator),
// }

// eg:
// PipeX(
//   source,
//   myPkg.Mappers,
//   myPkg.operators.map,
// }

// eg: PipeX[A any, B any]()

// eg:
// PipeX(
//   source,
//   myPkg.Mappers["foo"],
//   myPkg.operators()["map"].func,
// }

// eg:
// PipeX(
//   source,
//   Map(...),
//   pkg.Map(...),
//   pkg.operators.Map(...),
//   pkg.operators["map"].fn(...),
// }

///////////////// Fallback: please open an issue if a case is not properly supported
