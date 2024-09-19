package go_lsblk

import (
	"reflect"
	"strconv"
)

type BlockDeviceInfo struct {
	Name             string `json:"name" col:"NAME"`                 // device name
	KernelName       string `json:"KernelName" col:"KNAME"`          // internal kernel device name
	Path             string `json:"path" col:"PATH"`                 // path to the device node
	MajorMinor       string `json:"MajorMinor" col:"MAJ:MIN"`        // major:minor device number
	FSAvailable      string `json:"fsAvailable" col:"FSAVAIL"`       // filesystem size available
	FSSize           string `json:"fsSize" col:"FSSIZE"`             // filesystem size
	FSType           string `json:"fsType" col:"FSTYPE"`             // filesystem type
	FSUsed           string `json:"fsUsed" col:"FSUSED"`             // filesystem size used
	FSUsedPercent    string `json:"fsUsedPercent" col:"FSUSE%"`      // filesystem use percentage
	FSRoots          string `json:"FSRoots" col:"FSROOTS"`           // mounted filesystem roots
	FSVersion        string `json:"fsVersion" col:"FSVER"`           // filesystem version
	MountPoint       string `json:"mountPoint" col:"MOUNTPOINT"`     // where the device is mounted
	MountPoints      string `json:"mountPoints" col:"MOUNTPOINTS"`   // all locations where device is mounted
	Label            string `json:"label" col:"LABEL"`               // filesystem LABEL
	UUID             string `json:"uuid" col:"" col:"UUID"`          // filesystem UUID
	PTUUID           string `json:"ptuuid" col:"PTUUID"`             // partition table identifier (usually UUID)
	PTType           string `json:"pttype" col:"PTTYPE"`             // partition table type
	PartType         string `json:"partType" col:"PARTTYPE"`         // partition type code or UUID
	PartTypeName     string `json:"partTypeName" col:"PARTTYPENAME"` // partition type name
	PartLabel        string `json:"partLabel" col:"PARTLABEL"`       // partition LABEL
	PartUUID         string `json:"partUUID" col:"PARTUUID"`         // partition UUID
	PartFlags        string `json:"partFlags" col:"PARTFLAGS"`       // partition flags
	RA               string `json:"ra" col:"RA"`                     // read-ahead of the device
	RO               string `json:"ro" col:"RO"`                     // read-only device
	RM               string `json:"rm" col:"RM"`                     // removable device
	HotPlug          string `json:"hotplug" col:"HOTPLUG"`           // removable or hotplug device (usb, pcmcia, ...)
	Model            string `json:"model" col:"MODEL"`               // device identifier
	Serial           string `json:"serial" col:"SERIAL"`             // disk serial number
	Size             string `json:"size" col:"SIZE"`                 // size of the device
	State            string `json:"state" col:"STATE"`               // state of the device
	Owner            string `json:"owner" col:"OWNER"`               // user name
	Group            string `json:"group" col:"GROUP"`               // group name
	Mode             string `json:"mode" col:"MODE"`                 // device node permissions
	Alignment        string `json:"alignment" col:"ALIGNMENT"`       // alignment offset
	MinIO            string `json:"min-io" col:"MIN-IO"`             // minimum I/O size
	OptIO            string `json:"opt-io" col:"OPT-IO"`             // optimal I/O size
	PhyicalSector    string `json:"physicalSector" col:"PHY-SEC"`    // physical sector size
	LogicalSector    string `json:"logicalSector" col:"LOG-SEC"`     // logical sector size
	Rotation         string `json:"rotation" col:"ROTA"`             // rotational device
	Scheduler        string `json:"scheduler" col:"SCHED"`           // I/O scheduler name
	RQSize           string `json:"rqSize" col:"RQ-SIZE"`            // request queue size
	Type             string `json:"type" col:"TYPE"`                 // device type
	DiscAlign        string `json:"discAlign" col:"DISC-ALN"`        // discard alignment offset
	DiscGran         string `json:"discGran" col:"DISC-GRAN"`        // discard granularity
	DiscMax          string `json:"discMax" col:"DISC-MAX"`          // discard max bytes
	DiscZero         string `json:"discZero" col:"DISC-ZERO"`        // discard zeroes data
	WriteSame        string `json:"writeSame" col:"WSAME"`           // write same max bytes
	WWN              string `json:"wwn" col:"WWN"`                   // unique storage identifier
	Rand             string `json:"rand" col:"RAND"`                 // adds randomness
	ParentKernelName string `json:"ParentKernelName" col:"PKNAME"`   // internal parent kernel device name
	HCTL             string `json:"hctl" col:"HCTL"`                 // Host:Channel:Target:Lun for SCSI
	Tran             string `json:"tran" col:"TRAN"`                 // device transport type
	SubSystems       string `json:"subsystems" col:"SUBSYSTEMS"`     // de-duplicated chain of subsystems
	Revision         string `json:"revision" col:"REV"`              // device revision
	Vendor           string `json:"vendor" col:"VENDOR"`             // device vendor
	Zoned            string `json:"zoned" col:"ZONED"`               // zone model
	Dax              string `json:"dax" col:"DAX"`                   // dax-capable device
}

func (b BlockDeviceInfo) GetSize() (int64, error) {
	return strconv.ParseInt(b.Size, 10, 64)
}

func (b BlockDeviceInfo) MustGetSize() int64 {
	if b.Size == "" || b.Size == "0" {
		return 0
	}
	size, err := b.GetSize()
	if err != nil {
		panic(err)
	}
	return size
}

func (b BlockDeviceInfo) IsEmpty() bool {
	return reflect.DeepEqual(b, BlockDeviceInfo{})
}
