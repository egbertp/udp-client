# UDP client

This GoLang mini-project creates a binary file that sends packets with a counter on an UDP port and prints the output to the stdout.

I use this `udp-client` project to verify the correct working of my infrastructure.

## Usage

Download:

```
mkdir $GOPATH/src/github.com/egbertp/udp-client
cd $GOPATH/src/github.com/egbertp/udp-client
git clone https://github.com/egbertp/udp-client.git
```

### Build and Run

Install dependencies using [Glide Package Management for Go](https://glide.sh/)

```
$ glide install
```

Build the binary
```
$ make build
```

Release the app
```
$ make release
```

Run the app
```
$ ./udp-client --address 127.0.0.1 --port 37059

2018/02/16 14:03:40 Sending packet with payload: 0 to 127.0.0.1:37059
2018/02/16 14:03:41 Sending packet with payload: 1 to 127.0.0.1:37059
2018/02/16 14:03:42 Sending packet with payload: 2 to 127.0.0.1:37059
2018/02/16 14:03:43 Sending packet with payload: 3 to 127.0.0.1:37059
2018/02/16 14:03:44 Sending packet with payload: 4 to 127.0.0.1:37059
2018/02/16 14:03:45 Sending packet with payload: 5 to 127.0.0.1:37059
2018/02/16 14:03:46 Sending packet with payload: 6 to 127.0.0.1:37059
```

If you use this project together with my [udp-server](https://github.com/egbertp/udp-server) project, you can check if your firewalls are functioning properly

The output from the udp-server is:
```
2018/02/16 13:48:31 Received  0  from  127.0.0.1:52942
2018/02/16 13:48:32 Received  1  from  127.0.0.1:52942
2018/02/16 13:48:33 Received  2  from  127.0.0.1:52942
2018/02/16 14:03:40 Received  0  from  127.0.0.1:62656
2018/02/16 14:03:41 Received  1  from  127.0.0.1:62656
2018/02/16 14:03:42 Received  2  from  127.0.0.1:62656
2018/02/16 14:03:43 Received  3  from  127.0.0.1:62656
2018/02/16 14:03:44 Received  4  from  127.0.0.1:62656
2018/02/16 14:03:45 Received  5  from  127.0.0.1:62656
2018/02/16 14:03:46 Received  6  from  127.0.0.1:62656
```
