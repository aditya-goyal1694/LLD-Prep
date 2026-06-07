package logger

import (
	"fmt"
	"sync"
)

type Logger struct {}

var logger *Logger
var mu sync.Mutex

func GetLogger() *Logger {
    if (logger == nil) {
        mu.Lock()
        defer mu.Unlock()
        
        if (logger == nil) {
            logger = &Logger{}
        }
        
        return logger
    }
    return logger
}

// or in go
// var once sync.Once
// func GetLogger() *Logger {
//     once.Do(func() {
//         logger = &Logger{}
//     })

//     return logger
// }

func (l *Logger) LogInfo(msg string) {
    fmt.Printf("[INFO] %s\n", msg)
}


func (l *Logger) LogError(msg string) {
    fmt.Printf("[ERROR] %s\n", msg)
}
