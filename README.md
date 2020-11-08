# decode-msgpack

golang sub mqtt topic and decode msgpack.

## 1. req

```shell
[...]$ go build main.go && ./main
mqtt topic: [news], version: [1.2.5.22], msg id: [1902], time: [1958], ip: [192.168.0.100], mac[246F2875D36C]
req -> uuid: fda50693-a4e2-4fb1-afcf-c6eb07647825, major: 530, minor: 837, rssi: -40
req -> uuid: fda50693-a4e2-4fb1-afcf-c6eb07647825, major: 533, minor: 830, rssi: -40
req -> uuid: fda50693-a4e2-4fb1-afcf-c6eb07647825, major: 531, minor: 869, rssi: -40
req -> uuid: fda50693-a4e2-4fb1-afcf-c6eb07647825, major: 533, minor: 848, rssi: -40
req -> uuid: fda50693-a4e2-4fb1-afcf-c6eb07647825, major: 530, minor: 871, rssi: -40
req -> uuid: fda50693-a4e2-4fb1-afcf-c6eb07647825, major: 530, minor: 807, rssi: -40
req -> uuid: fda50693-a4e2-4fb1-afcf-c6eb07647825, major: 534, minor: 816, rssi: -40
req -> uuid: fda50693-a4e2-4fb1-afcf-c6eb07647825, major: 529, minor: 855, rssi: -40
```

## 2. ack

```shell
[...]$ go build main.go && ./main
mqtt topic: [news], version: [1.2.5.22], msg id: [1838], time: [1893], ip: [192.168.0.100], mac[246F2875D36C]
ack -> name len:  9, name:  Co-watch, temperature: 27.61 ℃, battery precent: 100%, major: 1, minor: 2, mac: 52444C1A045F 
ack -> name len:  9, name:  Co-watch, temperature: 30.33 ℃, battery precent:  90%, major: 1, minor: 2, mac: 52444C1A0460 
ack -> name len:  9, name:  Co-watch, temperature: 28.67 ℃, battery precent:  89%, major: 1, minor: 2, mac: 52444C1A044E 
ack -> name len: 10, name: temptrack, temperature: 27.93 ℃, battery precent:  94%, major: 1, minor: 2, mac: 52444C1A0447 
ack -> name len:  9, name:  Co-watch, temperature: 28.60 ℃, battery precent:  76%, major: 1, minor: 2, mac: 52444C1A043E 
ack -> name len:  9, name:  Co-watch, temperature: 26.79 ℃, battery precent:  88%, major: 1, minor: 2, mac: 52444C1A0345 
ack -> name len:  9, name:  Co-watch, temperature: 28.88 ℃, battery precent:  88%, major: 1, minor: 2, mac: 52444C1A0461 
ack -> name len:  9, name:  Co-watch, temperature:  27.2 ℃, battery precent:  66%, major: 1, minor: 2, mac: 52444C1A044F 
ack -> name len:  9, name:  Co-watch, temperature: 28.85 ℃, battery precent:  89%, major: 1, minor: 2, mac: 52444C1A0454 
ack -> name len:  9, name:  Co-watch, temperature: 28.29 ℃, battery precent: 100%, major: 1, minor: 2, mac: 52444C1A0436 
ack -> name len:  9, name:  Co-watch, temperature: 27.32 ℃, battery precent:  87%, major: 1, minor: 2, mac: 52444C1A0450 
ack -> name len:  9, name:  Co-watch, temperature: 27.33 ℃, battery precent:  81%, major: 1, minor: 2, mac: 52444C1A046F 
```
