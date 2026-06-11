# Adapter Pattern LLD Interview

## Problem Statement

You are building a **payment aggregation platform**.

Your system already supports a standard payment interface used internally across the company.

```go
type PaymentGateway interface {
    Pay(amount int) error
}
```

The checkout service, order service, and billing service all depend on this interface.

---

## Existing System

The checkout system expects every payment provider to implement:

```go
Pay(amount int) error
```

Example:

```go
checkout := NewCheckout(paymentGateway)

checkout.ProcessPayment(1000)
```

The checkout service should not know which payment provider is being used.

---

## New Requirement

The company wants to integrate a third-party payment provider called **FastPay**.

Unfortunately, FastPay was developed by another vendor and exposes the following API:

```go
type FastPay struct {}

func (f *FastPay) MakePayment(
    value float64,
) bool
```

Example:

```go
fastPay.MakePayment(1000.0)
```

Notice that:

* Method name is different.
* Return type is different.
* Signature is different.
* FastPay cannot be modified because it belongs to a third-party vendor.

---

## Objective

Integrate FastPay into the existing checkout system **without modifying**:

* Checkout
* Existing payment interfaces
* Third-party FastPay library

The checkout service should continue to work with:

```go
PaymentGateway
```

only.

---

## Future Requirements

Tomorrow additional third-party vendors may be introduced:

```text
QuickPay
RazorPaySDK
ExternalBankGateway
```

Each may expose completely different APIs.

The checkout service should remain unchanged.

---

## Example Usage

```go
fastPay := &FastPay{}

gateway := NewFastPayAdapter(fastPay)

checkout := NewCheckout(gateway)

checkout.ProcessPayment(1000)
```

Output:

```text
FastPay payment processed successfully
```