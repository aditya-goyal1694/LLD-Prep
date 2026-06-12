# Decorator Pattern LLD Interview

## Problem Statement

You are building a **coffee ordering system**.

Customers can order a base coffee and optionally add multiple extras.

### Base Coffees

* Espresso
* Cappuccino
* Latte

### Add-ons

* Milk
* Whipped Cream
* Chocolate Syrup
* Caramel Syrup

Each add-on increases the final cost and modifies the description.

---

## Example Orders

### Example 1

Customer orders:

```text
Espresso
```

Output:

```text
Description: Espresso
Cost: ₹100
```

---

### Example 2

Customer orders:

```text
Espresso
+ Milk
+ Chocolate Syrup
```

Output:

```text
Description: Espresso, Milk, Chocolate Syrup
Cost: ₹140
```

---

### Example 3

Customer orders:

```text
Latte
+ Milk
+ Whipped Cream
+ Caramel Syrup
```

Output:

```text
Description: Latte, Milk, Whipped Cream, Caramel Syrup
Cost: ₹220
```

---

## Requirements

The system must support:

### Base Beverages

```text
Espresso
Latte
Cappuccino
```

### Add-ons

```text
Milk
Whipped Cream
Chocolate Syrup
Caramel Syrup
```

Multiple add-ons can be combined in any order.

Examples:

```text
Latte + Milk

Latte + Milk + Chocolate

Latte + Milk + Chocolate + Caramel

Cappuccino + Milk + Milk + Milk
```

All combinations should work.

---

## Future Requirements

The product team may introduce:

```text
Oat Milk
Vanilla Syrup
Hazelnut Syrup
Ice Cream
```

The design should allow new add-ons without modifying existing beverage classes.

