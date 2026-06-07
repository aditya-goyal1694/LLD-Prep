package main

import (
    "sync"
    "lld/design_patterns/singleton/orders"
    "lld/design_patterns/singleton/users"
    "lld/design_patterns/singleton/payments"
    "lld/design_patterns/singleton/inventory"
)

func main() {
    o := &orders.OrderService{}
    u := &users.UserService{}
    i := &inventory.InventoryService{}
    p := &payments.PaymentService{}
    
    var wg sync.WaitGroup
    wg.Add(4)
    
    go func() {
        defer wg.Done()
        o.PrintOrderDetails()
    }()
    go func() {
        defer wg.Done() 
        u.PrintUserDetails()
    }()
    go func() {
        defer wg.Done()
        i.PrintInventoryDetails()
    }()
    go func() {
        defer wg.Done()
        p.PrintPaymentDetails()
    }()
    
    wg.Wait()    
}
