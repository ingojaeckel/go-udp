# go-udp

Very basic UDP server/client hello world example project inspired by http://www.minaandrawos.com/2016/05/14/udp-vs-tcp-in-golang/.

## Example Output

    $ go build && ./go-udp
    [srv] Handling new connection. Received 1 bytes from client:[23]; addr:127.0.0.1:65475
    [Alice] Received 2 bytes: [23 1]
    [Alice] Received 2 bytes: [23 2]
    [Alice] Received 2 bytes: [23 3]
    [Alice] Received 2 bytes: [23 4]
    [Alice] Received 2 bytes: [23 5]
    [srv] Handling new connection. Received 1 bytes from client:[23]; addr:127.0.0.1:61881
    [Bob] Received 2 bytes: [23 1]
    [Alice] Received 2 bytes: [23 6]
    [Bob] Received 2 bytes: [23 2]
    [Alice] Received 2 bytes: [23 7]
    [Bob] Received 2 bytes: [23 3]
    [Alice] Received 2 bytes: [23 8]
    [Bob] Received 2 bytes: [23 4]
    [Alice] Received 2 bytes: [23 9]
    [Bob] Received 2 bytes: [23 5]
    [Alice] Received 2 bytes: [23 10]
    [srv] Handling new connection. Received 1 bytes from client:[42]; addr:127.0.0.1:56092
    [Carol] Received 2 bytes: [42 1]
    [Bob] Received 2 bytes: [23 6]
    [Alice] Received 2 bytes: [23 11]
    [Carol] Received 2 bytes: [42 2]
    [Alice] Received 2 bytes: [23 12]
    [Bob] Received 2 bytes: [23 7]
    [Carol] Received 2 bytes: [42 3]
    [Alice] Received 2 bytes: [23 13]
    [Bob] Received 2 bytes: [23 8]
    [Carol] Received 2 bytes: [42 4]
    [Bob] Received 2 bytes: [23 9]
    [Alice] Received 2 bytes: [23 14]
    [Carol] Received 2 bytes: [42 5]
    [Alice] Received 2 bytes: [23 15]
    [Bob] Received 2 bytes: [23 10]
    [srv] Handling new connection. Received 1 bytes from client:[42]; addr:127.0.0.1:52379
    [Dan] Received 2 bytes: [42 1]
    [Carol] Received 2 bytes: [42 6]
    [Alice] Received 2 bytes: [23 16]
    [Bob] Received 2 bytes: [23 11]
    [Carol] Received 2 bytes: [42 7]
    [Dan] Received 2 bytes: [42 2]
    [Alice] Received 2 bytes: [23 17]
    [Bob] Received 2 bytes: [23 12]
    [Dan] Received 2 bytes: [42 3]
    [Carol] Received 2 bytes: [42 8]
    [Alice] Received 2 bytes: [23 18]
    [Bob] Received 2 bytes: [23 13]
    [Dan] Received 2 bytes: [42 4]
    [Carol] Received 2 bytes: [42 9]
    [Bob] Received 2 bytes: [23 14]
    [Alice] Received 2 bytes: [23 19]
    [Carol] Received 2 bytes: [42 10]
    [Dan] Received 2 bytes: [42 5]
    [Bob] Received 2 bytes: [23 15]
    [Alice] Received 2 bytes: [23 20]