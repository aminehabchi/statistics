# Linear Regression and Pearson Correlation Calculator

This project reads a dataset from a file and calculates the Linear Regression Line and Pearson Correlation Coefficient.

## How It Works

The program reads data from a text file where each line contains a numerical value. The line numbers represent the x-axis values, and the numbers on each line represent the y-axis values of a graph.

It then calculates:

- **Linear Regression Line**: `y = mx + b`
  - The values for `m` (slope) and `b` (intercept) are calculated using linear regression formulas.
  - Output format: `y = <m>x + <b>`, both values printed with 6 decimal places.

- **Pearson Correlation Coefficient**:
  - This measures the linear relationship between the `x` and `y` values.
  - Output format: Pearson Correlation Coefficient: `<r>`, printed with 10 decimal places.

