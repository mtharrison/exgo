# Go Formula Parser and Evaluator

Use this package to support simple formulas in your Go application.

Supports:

- Simple (non-nested) infix arithmetic operations: - + * / %

## API

### Structs

#### `exgo.Formula`

### Functions

#### `func exgo.Parse(string) exgo.Formula`
Parses a string formula and returns a formula struct.

#### `func (*exgo.Formula) exgo.Evaluate() float64`
Evaluates a formula and returns the resultant number
