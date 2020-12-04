package file

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func StartFileServer()  {
	os.Mkdir("file", 0777)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("file"))))
	err := http.ListenAndServe(":88",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
func  ServeHTTP()  {
	fmt.Println("文件服务器启动成功！")
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ipv4 :=""
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				ipv4 = ipnet.IP.String()
			}

		}
	}
	fmt.Println("请在运行程序目录下找到flie文件夹，把要下载的文件放到file里面")
	fmt.Println("然后在手机打开：","http://"+ipv4+":88")
	fmt.Println("ctrl+c 退出程序")
}
