package fastpay

import "fmt"

type FastPay struct {}

func (f *FastPay) MakePayment(value float64) bool {
    fmt.Printf("The amount of %f has been deducted from your account.\n", value)
    return true
}