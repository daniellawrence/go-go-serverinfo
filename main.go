package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL)
}

func json_hostinfo(w http.ResponseWriter, r *http.Request) {
	hostinfo, _ := host.HostInfo()
	b, _ := json.Marshal(hostinfo)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func json_swapinfo(w http.ResponseWriter, r *http.Request) {
	swapinfo, _ := mem.SwapMemory()
	b, _ := json.Marshal(swapinfo)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func json_meminfo(w http.ResponseWriter, r *http.Request) {
	memginfo, _ := mem.VirtualMemory()
	b, _ := json.Marshal(memginfo)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func json_cpuinfo(w http.ResponseWriter, r *http.Request) {
	cpuinfo, _ := cpu.CPUInfo()
	b, _ := json.Marshal(cpuinfo)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func json_diskpartitions(w http.ResponseWriter, r *http.Request) {
	diskinfo, _ := disk.DiskPartitions(true)
	b, _ := json.Marshal(diskinfo)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}
func json_diskIOCounters(w http.ResponseWriter, r *http.Request) {
	diskinfo, _ := disk.DiskIOCounters()
	b, _ := json.Marshal(diskinfo)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func embed(w http.ResponseWriter, r *http.Request) {
	data, _ := Asset("static/index.html")
	fmt.Println("embed!")
	fmt.Fprintf(w, string(data))
}

func main() {
	static_dir := http.FileServer(http.Dir("./static"))

	portPtr := flag.Int("port", 3001, "Port to bind go-go-server-info")
	bindPtr := flag.String("host", "0.0.0.0", "bind IP")

	flag.Parse()
	bind := fmt.Sprintf("%s:%d", *bindPtr, *portPtr)


	log.Println("Listening... " + bind)	
	http.Handle("/dev", static_dir)
	http.Handle("/dev/", static_dir)
	http.HandleFunc("/", embed)
	http.HandleFunc("/json/hostinfo/", json_hostinfo)
	http.HandleFunc("/json/cpuinfo/", json_cpuinfo)
	http.HandleFunc("/json/diskpartitions/", json_diskpartitions)
	http.HandleFunc("/json/diskiocounters/", json_diskIOCounters)
	http.HandleFunc("/json/meminfo/", json_meminfo)
	http.HandleFunc("/json/swapinfo/", json_swapinfo)
	err := http.ListenAndServe(bind, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
