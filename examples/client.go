package main

import (
    . "github.com/zubairhamed/goap"
    "log"
)

/*	To test this example, also run examples/test_server.go */
func main() {
    log.Println("Starting Client..")
    client := NewClient()
    defer client.Close()

    client.Dial("udp", COAP_DEFAULT_HOST, COAP_DEFAULT_PORT)

    req := NewRequest(TYPE_CONFIRMABLE, GET, 12345)
    req.SetStringPayload("Hello, goap")
    req.SetRequestURI("/serviceD")

    // Sync Client Test
    log.Println("Sending Synchronous Message")
    resp, err := client.Send(req)
    if err != nil {
        log.Println(err)
    } else {
        log.Println("Got Synchronous Response:")
        log.Println(CoapCodeToString(resp.GetMessage().Code))
    }

    // Async Client Test
    req.IncrementMessageId()
    log.Println("Sending Asynchronous Message")
    client.SendAsync(req, func(resp *CoapResponse, err error){
        if err != nil {
            log.Println(err)
        } else {
            log.Println("Got Asynchronous Response:")
            log.Println(CoapCodeToString(resp.GetMessage().Code))
        }
    })

    // Discovery Test
    client.Discover(func(resp *CoapCode, err error){

    })
}
