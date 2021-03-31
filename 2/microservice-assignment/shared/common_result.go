package shared

type CommonResult struct {
	Result interface{}
	Error  error
}

func GenerateResult(data interface{}, err error) <-chan CommonResult {
	output := make(chan CommonResult)

	go func() {
		defer close(output)
		output <- CommonResult{Result: data, Error: err}
	}()

	return output
}
