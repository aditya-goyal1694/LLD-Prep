You are building a backend service for an e-commerce platform.

The application writes logs from various modules:

- Order Service
- Payment Service
- Inventory Service
- User Service

All logs must be written through a single centralized logger instance.

Example:
```
[INFO] Order created
[ERROR] Payment failed
[INFO] Inventory updated
```

### Requirements

The logger should support:
```Go
LogInfo(message string)
LogError(message string)
```

### Constraint

The entire application must use exactly one logger instance.

This means:
```Go
logger1 := GetLogger()
logger2 := GetLogger()
```
must always refer to the same object.

### Future Requirement

The application becomes highly concurrent.

100 goroutines may call:
```Go
GetLogger()
```
at the same time.

The design must still guarantee:

Exactly one logger instance.