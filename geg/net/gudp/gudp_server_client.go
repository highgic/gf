package main

import (
    "fmt"
    "time"
    "gitee.com/johng/gf/g/os/glog"
    "gitee.com/johng/gf/g/os/gtime"
    "gitee.com/johng/gf/g/net/gudp"
)

func main() {
    // Server
    go gudp.NewServer("127.0.0.1:8999", func(conn *gudp.Conn) {
        defer conn.Close()
        for {
            if data, _ := conn.Receive(-1); len(data) > 0 {
                if err := conn.Send(append([]byte("> "), data...)); err != nil {
                    glog.Error(err)
                }
            }
        }
    }).Run()

    time.Sleep(time.Second)

    // Client
    for {
        if conn, err := gudp.NewConn("127.0.0.1:8999"); err == nil {
            if b, err := conn.SendReceive([]byte(gtime.Datetime()), -1); err == nil {
                fmt.Println(string(b), conn.LocalAddr(), conn.RemoteAddr())
            } else {
                glog.Error(err)
            }
            conn.Close()
        } else {
            glog.Error(err)
        }
        time.Sleep(time.Second)
    }
}