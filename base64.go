package main

import (
	"encoding/base64"
	"fmt"
)

const (
	base64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
)

var coder = base64.NewEncoding(base64Table)

func Base64Encode(src []byte) []byte {         //编码
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {   //解码
	return coder.DecodeString(string(src))
}

func main() {
	a := `{"cer":"/usr/local/trpc/bin/u=198955208,3905140526\u0026fm=26\u0026gp=0.jpg.471542592","email":"913372477@qq.com","passWord":"123456789","appName":"xiaomi","packageName":"xiaomiPackName"}`
	ae := Base64Encode([]byte(a))
	fmt.Println(string(ae[:]))
	a1, _ := Base64Decode(ae)
	fmt.Println(string(a1[:]))
	if a == string(a1[:]) {
		fmt.Println("true")
	}
}

/*
output:
dx5iEaOhXhOpcaQx74rpD4Lm7sRxbMPpDCBo7s39PWAtXW30PiIt7kPuPk3rQkI0PiEbcWIqPiECGWzxQBr0PkIxQCcq1WIoFHJH7i2sPW3zPi3uPhOmOCYnDZBmOilhXWKyPybxQkbs2NLr7CQpGVOmOHJgbsQaGs5AOilhPWOyQk34QyfuOhqhDaJqWCLnEVOvOHgSDZ9nFVOmOHJgD4ngE4YXDZ0BOilhdMBgG40S3MLiFzugGZ3he2==
{"cer":"/usr/local/trpc/bin/u=198955208,3905140526\u0026fm=26\u0026gp=0.jpg.471542592","email":"913372477@qq.com","passWord":"123456789","appName":"xiaomi","packageName":"xiaomiPackName"}
true
*/
