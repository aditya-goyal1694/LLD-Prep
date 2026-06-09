
You are building an **e-commerce checkout system**.

A customer can pay using different payment methods:

* Credit Card
* UPI
* PayPal

The checkout service should be able to process payments using any of these methods.

---

## Current Requirements

Each payment method has its own payment processing logic.

Examples:

* Credit Card → Card Network + Bank
* UPI → UPI Provider
* PayPal → PayPal Gateway

The checkout service should not know how a particular payment method works internally.

The checkout service should only know:

```go
Pay(amount int) error
```

and delegate the actual payment processing to the selected payment method.

---

## Future Requirements

The product team plans to add more payment methods:

* Wallet
* Net Banking
* Crypto

The checkout service should not require modification whenever a new payment method is introduced.

The system should be easily extensible for future payment providers.

---

## Example Usage

```go
paymentMethod := NewUPIPayment()

checkout := NewCheckout(paymentMethod)

checkout.Pay(1000)
```

Output:

```text
Processing payment of ₹1000 using UPI
```

Another example:

```go
paymentMethod := NewCreditCardPayment()

checkout := NewCheckout(paymentMethod)

checkout.Pay(1000)
```

Output:

```text
Processing payment of ₹1000 using Credit Card
```

---