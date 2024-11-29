# go

```sh
go clean -cache
go build -x -v --work 1-1.go
```

## 单元测试

`go test -v -coverprofile=coverage.out ./...` 是一个用于运行 Go 测试并生成代码覆盖率报告的命令。下面是各个选项的详细说明：

- `go test`: 这是运行 Go 测试的基本命令。
- `-v`: 这个选项表示以详细模式运行测试，它会输出每个测试的详细信息。
- `-coverprofile=coverage.out`: 这个选项生成一个覆盖率报告，并将结果写入 `coverage.out` 文件。
- `./...`: 这个选项表示递归地测试当前目录下的所有包及其子包。

### 使用步骤

1. **在你的项目根目录下运行命令**：

   打开终端或命令行工具，导航到你的 Go 项目根目录，然后运行以下命令：

   ```bash
   go test -v -coverprofile=coverage.out ./...
   ```
2. **查看生成的覆盖率报告**：

   执行命令后，会在你的项目根目录下生成一个名为 `coverage.out` 的文件。你可以使用以下命令查看覆盖率报告的摘要：

   ```bash
   go tool cover -func=coverage.out
   ```

   该命令会输出每个函数的覆盖率百分比以及总的覆盖率百分比。
3. **生成并查看 HTML 格式的覆盖率报告**：

   如果你希望以更直观的方式查看覆盖率报告，可以将其转换为 HTML 格式并在浏览器中打开：

   ```bash
   go tool cover -html=coverage.out -o coverage.html
   ```

   然后你可以在浏览器中打开 `coverage.html` 文件，查看代码覆盖率的详细情况。

---

在 Go 中，编写单元测试非常简单且直观。以下是编写 Go 单元测试的基本步骤和示例。

### 1. 编写测试文件

- 测试文件的文件名必须以 `_test.go` 结尾。
- 测试函数的名称必须以 `Test` 开头，后面跟着被测试函数的名称。

### 2. 导入测试包

导入 `testing` 包，该包提供了编写测试所需的函数和工具。

### 3. 编写测试函数

测试函数的签名为 `func TestXxx(t *testing.T)`，其中 `Xxx` 可以是任意名称，但通常是被测试函数的名称。

### 示例

假设我们有一个简单的函数 `Add`，它接受两个整数并返回它们的和：

```go
// filename: math.go
package math

func Add(a, b int) int {
    return a + b
}
```

我们要为这个函数编写一个单元测试。

```go
// filename: math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

### 运行测试

在你的项目目录下，运行以下命令来执行测试：

```bash
go test
```

你可以使用 `-v` 选项以详细模式运行测试：

```bash
go test -v
```

### 更多示例

#### 测试多个用例

你可以使用一个切片来存储多个测试用例，并在一个测试函数中迭代这些用例：

```go
// filename: math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    var tests = []struct {
        a, b     int
        expected int
    }{
        {1, 1, 2},
        {2, 3, 5},
        {0, 0, 0},
        {-1, -1, -2},
    }

    for _, tt := range tests {
        result := Add(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
        }
    }
}
```

#### 测试错误情况

如果你的函数可能返回错误，你可以使用 `t.Errorf` 或 `t.Fatalf` 来报告错误情况：

```go
// filename: math.go
package math

import "errors"

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

```go
// filename: math_test.go
package math

import "testing"

func TestDivide(t *testing.T) {
    _, err := Divide(1, 0)
    if err == nil {
        t.Errorf("expected an error but got none")
    }

    result, err := Divide(4, 2)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    expected := 2
    if result != expected {
        t.Errorf("Divide(4, 2) = %d; want %d", result, expected)
    }
}
```

### 使用 `t.Run` 进行子测试

你可以使用 `t.Run` 创建子测试，以便在一个测试函数中运行多个独立的测试：

```go
// filename: math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    var tests = []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 1, 1, 2},
        {"positive and negative", 2, -3, -1},
        {"zeros", 0, 0, 0},
        {"negative numbers", -1, -1, -2},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

### 结论

编写 Go 单元测试主要涉及以下几个步骤：

1. 编写以 `_test.go` 结尾的测试文件。
2. 导入 `testing` 包。
3. 编写以 `Test` 开头的测试函数。
4. 使用 `t.Errorf` 或 `t.Fatalf` 报告测试失败。

通过这些步骤，你可以为你的 Go 代码编写有效的单元测试，以确保代码的正确性。



```bash
echo 'export GOROOT=/usr/lib/go' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GO111MODULE=on' >> ~/.bashrc
source ~/.bashrc

```
