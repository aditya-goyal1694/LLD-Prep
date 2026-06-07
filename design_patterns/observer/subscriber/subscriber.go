package subscriber

import "fmt"

type Subscriber interface {
    Update(data string) error
}


type User struct {
    Name string
}

func NewUserSubscriber(name string) Subscriber {
    return &User{
        Name: name,
    }
}

func (u *User) Update(data string) error {
    fmt.Printf("%s has been notified about %s.\n", u.Name, data)
    return nil
}