package handler

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
	"util/tpl"
)

func tcpPort(w http.ResponseWriter, r *http.Request) {
	host := r.FormValue("host")
	port := r.FormValue("port")

	if host == "" || port == "" {
		if err := tpl.TcpPort.Execute(w, nil); err != nil {
			Fprint(w, "ERROR")
		}
		return
	}

	ports := strings.Split(port, ",")

	if len(ports) == 1 {
		ports = strings.Split(port, " ")
	}

	ports = func() (port []string) {
		for _, v := range ports {
			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}
			p, err := strconv.Atoi(v)
			if err != nil || p < 0 || p > 65535 {
				continue
			}

			port = append(port, v)
		}
		return
	}()

	data := ""
	if len(ports) == 0 {
		Fprint(w, "ERROR")
	}else if len(ports) == 1 {
		test := tcpGather(host, ports)
		data = test[ports[0]]
	} else {
		data = func() string {
			data, _ := json.Marshal(tcpGather(host, ports))
			return string(data)
		}()
	}
	Fprint(w, data)

}

func tcpGather(ip string, ports []string) map[string]string {
	results := make(map[string]string)
	for _, port := range ports {
		address := net.JoinHostPort(ip, port)
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results[port] = "failed"
		} else {
			if conn != nil {
				results[port] = "success"
				_ = conn.Close()
			} else {
				results[port] = "failed"
			}
		}
	}
	return results
}
