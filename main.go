package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/vmihailenco/msgpack"
)

var (
	mqttTopic   = "news"
	mqttHostURL = "tcp:/127.0.0.1:1883"
	radHeadrLen = 8
)

func bytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

/* 判断是否回应包。 */
func isAck(b []byte) bool {
	symbol := []byte{0x10, 0x16, 0x03, 0x18}
	radAckNameLen := int(b[radHeadrLen])
	radAckNameEnd := int(radHeadrLen + radAckNameLen)
	return bytes.Equal(symbol, b[radAckNameEnd+1:radAckNameEnd+5])
	// return b[0] == 4
	// return strings.Contains(string(b), string(symbol)) && b[0] != 4
}

/* 解码回应包。 */
func decodeAck(i int, v []byte) {
	radAckNameLen := int(v[radHeadrLen])
	radAckNameEnd := int(radHeadrLen + radAckNameLen)

	nameLen := int(v[radHeadrLen])
	name := string(v[radHeadrLen+2 : radAckNameEnd+1])
	temperature := fmt.Sprintf("%d.%d", v[radAckNameEnd+5], v[radAckNameEnd+6])
	batteryPrecent := v[radAckNameEnd+7]
	major, _ := bytesToIntU(v[radAckNameEnd+8 : radAckNameEnd+10])
	minor, _ := bytesToIntU(v[radAckNameEnd+10 : radAckNameEnd+12])
	mac := v[radAckNameEnd+12 : radAckNameEnd+18]

	fmt.Printf("ack -> name len: %2d, name: %9s, temperature: %5s ℃, battery precent: %3d%%, major: %d, minor: %d, mac: %X \n",
		nameLen, name, temperature, batteryPrecent, major, minor, mac)

	// fmt.Printf("ack -> %d, len: %d, [% X | % X |% X | % X | % X | % X | % X | % X | % X | % X | % X | % X | % X]\n",
	// 	i, len(v),
	// 	v[0], v[1:7], v[7],
	// 	v[radHeadrLen],
	// 	v[radHeadrLen+1],
	// 	v[radHeadrLen+2:radAckNameEnd+1],
	// 	v[radAckNameEnd+1:radAckNameEnd+5],
	// 	v[radAckNameEnd+5:radAckNameEnd+7],
	// 	v[radAckNameEnd+7:radAckNameEnd+8],
	// 	v[radAckNameEnd+8:radAckNameEnd+10],
	// 	v[radAckNameEnd+10:radAckNameEnd+12],
	// 	v[radAckNameEnd+12:radAckNameEnd+18],
	// 	v[radAckNameEnd+18:len(v)])
	// fmt.Printf("%d, len: %d, [% X]\n", i, len(v), v)
}

/* 解码广播数据格式。 */
func decodeReq(i int, v []byte) {
	if len(v) != 38 {
		return
	}

	/* uuid: 4-2-2-2-6 */
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		v[radHeadrLen+9:radHeadrLen+13],
		v[radHeadrLen+13:radHeadrLen+15],
		v[radHeadrLen+15:radHeadrLen+17],
		v[radHeadrLen+17:radHeadrLen+19],
		v[radHeadrLen+19:radHeadrLen+25])
	/* major: 2 */
	major, _ := bytesToIntU(v[radHeadrLen+25 : radHeadrLen+27])
	/* major: 2 */
	minor, _ := bytesToIntU(v[radHeadrLen+27 : radHeadrLen+29])
	/* rssi: 1 */
	rssi := int(v[radHeadrLen+29]) - 256

	fmt.Printf("req -> uuid: %s, major: %d, minor: %d, rssi: %d\n",
		uuid, major, minor, rssi)
	// fmt.Printf("req -> %d, len: %d, [% X | % X | % X | % X | % X | % X | % X |% X]\n",
	// 	i, len(v),
	// 	v[0], v[1:7], v[7],
	// 	v[radHeadrLen:radHeadrLen+9],
	// 	v[radHeadrLen+9:radHeadrLen+25],
	// 	v[radHeadrLen+25:radHeadrLen+27],
	// 	v[radHeadrLen+27:radHeadrLen+29],
	// 	v[radHeadrLen+29])
	// fmt.Printf("%d, len: %d, [% X]\n", i, len(v), v)
}

/* 订阅回调 */
func subCallBackFunc(c MQTT.Client, msg MQTT.Message) {
	if !c.IsConnected() {
		return
	}

	/* 解码 msgpack 包。 */
	var out map[string]interface{}
	err := msgpack.Unmarshal(msg.Payload(), &out)
	if err != nil {
		panic(err)
	}

	fmt.Printf("mqtt topic: [%s], version: [%v], msg id: [%v], time: [%v], ip: [%v], mac[%v]\n",
		msg.Topic(), out["v"], out["mid"], out["time"], out["ip"], out["mac"])

	devices := out["devices"]
	if reflect.TypeOf(devices).Kind() != reflect.Slice {
		panic(err)
	}

	/* 解码十六进制硬件广播数据。 */
	s := reflect.ValueOf(devices)
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i).Elem().Bytes()
		if isAck(v) {
			// decodeAck(i, v)
		} else {
			decodeReq(i, v)
		}
	}
	fmt.Println("------------")
}

/* 连接MQTT服务 */
func connMQTT(broker, user, passwd string) (bool, MQTT.Client) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(user)
	opts.SetPassword(passwd)

	mc := MQTT.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		return false, mc
	}

	return true, mc
}

func subscribe() {
	ok, mc := connMQTT(mqttHostURL, "", "")
	if !ok {
		fmt.Println("sub connMQTT failed")
		return
	}
	mc.Subscribe(mqttTopic, 0x00, subCallBackFunc)
}

func main() {
	subscribe()
	for {
		time.Sleep(time.Second)
	}
}
