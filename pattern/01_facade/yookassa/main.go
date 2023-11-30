package yookassa

type PaymentResponse struct {
}

func Payment() *PaymentResponse {
	return &PaymentResponse{}
}

// тут много-много всяких методов, которые нам не нужны в нашем приложении
// либо нужны, но работать с ними напрямую сложно, а нам нужен интерфейс попроще
