package sniff

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"time"
	"fmt"
	"os"
)

var (
	device       string = "any"
	snapshotLen  uint32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = -1 * time.Second
	handle       *pcap.Handle
)

func main() {

	sniff()

}

func sniff() {

	f, _ := os.Create("/var/log/srm.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)
	defer f.Close()

	// Open the device for capturing
	handle, err = pcap.OpenLive(device, int32(snapshotLen), promiscuous, timeout)

	var filter string = "tcp and port 8080"

	err = handle.SetBPFFilter(filter)

	if err != nil {
		fmt.Printf("Error opening device %s: %v", device, err)
		os.Exit(1)
	}
	defer handle.Close()

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)

		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())

	}

}
