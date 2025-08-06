# AI Project Test Template

## 1. 测试目标
- 功能完整性：确保所有核心功能可用
- 响应准确性：确保输出与预期一致
- 易用性与集成性：确保易于集成和扩展
- 健壮性：确保异常输入下系统稳定
- 覆盖率：100%

## 2. 测试策略
- 单元测试：聚焦核心算法/逻辑

## 3. 单元测试函数模版
- golang
```golang
// Test_<FunctionName> 是一个针对 <FunctionName> 的单元测试模板。
// AI助手: 请根据待测试函数 <FunctionName> 的实际功能，填充以下测试用例。
func Test_<FunctionName>(t *testing.T) {
	// 1. 定义函数参数和返回值的类型
	// AI助手: 请根据待测试函数 <FunctionName> 的签名，定义 args 和 want 的类型。
	type args struct {
		// 参数1: 类型
		// 参数2: 类型
		// ...
	}

	// 2. 定义测试用例切片
	tests := []struct {
		name    string // 测试用例的简短描述，例如 "成功场景", "无效输入", "权限不足" 等。
		args    args   // 待测试函数的输入参数。
		want    // 预期返回的结果。
		wantErr bool   // 预期是否返回错误。
	}{
		// 3. 填充具体的测试用例
		// AI助手: 请根据 <FunctionName> 的功能需求，创建多个测试用例。
		// 每个用例都应覆盖一个特定的场景（成功、失败、边界条件等）。

		// 示例: 成功场景
		{
			name: "should_return_expected_result_for_valid_input", // 描述性名称，方便识别
			args: args{
				// 填充参数1: ...
				// 填充参数2: ...
			},
			want:    // 填充预期的成功结果
			wantErr: false, // 预期不返回错误
		},

		// 示例: 失败场景 - 无效输入
		{
			name: "should_return_error_for_invalid_input",
			args: args{
				// 填充参数以触发错误，例如: 参数为 nil, 空字符串或负数
			},
			want:    nil,   // 预期返回 nil 或零值
			wantErr: true,  // 预期返回错误
		},
		
		// AI助手: 请在此处添加更多测试用例，例如：
		// - 权限不足
		// - 依赖项返回错误
		// - 边界条件测试（例如：切片为空、数字为零）
	}

	// 4. 遍历并执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// AI助手: 如果 <FunctionName> 是一个方法，请在此处实例化接收者。
			// 例如: `h := &Helper{}` 或 `mockDB := &MockDB{}`
			// 
			// 在此调用待测试函数，并捕获返回值和错误。
			got, err := <FunctionName>(/* 传入 tt.args 中的参数 */)

			// 5. 错误断言
			if (err != nil) != tt.wantErr {
				t.Errorf("%s() error = %v, wantErr %v", <FunctionName>, err, tt.wantErr)
				return
			}

			// 6. 结果断言
			// AI助手: 根据 <FunctionName> 的返回值类型，选择合适的断言方法。
			// 如果是复杂类型或指针，使用 cmp.Diff 进行深度比较。
			// 如果是基本类型，使用 == 比较。
			// if diff := cmp.Diff(got, tt.want); diff != "" {
			//     t.Errorf("%s() got unexpected diff (-want +got):\n%s", <FunctionName>, diff)
			// }

			// if got != tt.want {
			//     t.Errorf("%s() got = %v, want %v", <FunctionName>, got, tt.want)
			// }
		})
	}
}
```


## 4. 自动化与报告
- AI 自动生成测试代码，自动运行并输出覆盖率、失败详情
