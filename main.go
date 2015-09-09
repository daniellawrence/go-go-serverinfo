package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
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


func main() {
	static_dir := http.FileServer(http.Dir("static"))
	log.Println("Listening... :3001")	
	http.Handle("/", static_dir)
	http.HandleFunc("/json/hostinfo/", json_hostinfo)
	http.HandleFunc("/json/cpuinfo/", json_cpuinfo)
	http.HandleFunc("/json/diskpartitions/", json_diskpartitions)
	http.HandleFunc("/json/diskiocounters/", json_diskIOCounters)
	http.ListenAndServe(":3001", nil)
}
