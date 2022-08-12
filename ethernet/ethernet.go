package main

type EthernetHeader struct {
	DistMacAddr   []byte
	SourceMacAddr []byte
	Type          []byte
}

type EthernetFrame struct {
	Header EthernetHeader
	Data   []byte
	FCS    []byte
}

func NewEthernet(distMacAddr, sourceMacAddr, ethernetType, data, fcs []byte) EthernetFrame {
	header := EthernetHeader{
		DistMacAddr:   distMacAddr,
		SourceMacAddr: sourceMacAddr,
		Type:          ethernetType,
	}
	ethernet := EthernetFrame{
		Header: header,
		Data:   data,
		FCS:    fcs,
	}

	return ethernet
}
