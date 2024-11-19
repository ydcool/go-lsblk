package go_lsblk

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"strings"
)

type Int64String string

func (i Int64String) Int64() (int64, error) {
	return strconv.ParseInt(string(i), 10, 64)
}

func (i Int64String) MustInt64() int64 {
	size, err := i.Int64()
	if err != nil {
		panic(err)
	}
	return size
}

type BoolString string

func (b BoolString) Bool() (bool, error) {
	return strconv.ParseBool(string(b))
}

func (b BoolString) MustBool() bool {
	val, err := b.Bool()
	if err != nil {
		panic(err)
	}
	return val
}

type MajorMinorString string

func (m MajorMinorString) MajorMinor() (int64, int64, error) {
	ss := strings.Split(string(m), ":")
	if len(ss) != 2 {
		return 0, 0, errors.New("invalid major:minor format")
	}
	major, err := strconv.ParseInt(ss[0], 10, 64)
	if err != nil {
		return 0, 0, errors.WithMessage(err, "invalid major format")
	}
	minor, err := strconv.ParseInt(ss[1], 10, 64)
	if err != nil {
		return 0, 0, errors.WithMessage(err, "invalid minor format")
	}
	return major, minor, nil
}

func (m MajorMinorString) MustMajorMinor() (int64, int64) {
	major, minor, err := m.MajorMinor()
	if err != nil {
		panic(err)
	}
	return major, minor
}

type SCSIHCTL struct {
	Host    int
	Channel int
	Target  int
	Lun     int
}

func (s SCSIHCTL) String() string {
	return fmt.Sprintf("%d:%d:%d:%d", s.Host, s.Channel, s.Target, s.Lun)
}

type SCSIHCTLString string

func (s SCSIHCTLString) Parse() (SCSIHCTL, error) {
	ss := strings.Split(string(s), ":")
	if len(ss) != 4 {
		return SCSIHCTL{}, errors.New("invalid hctl format")
	}
	host, err := strconv.Atoi(ss[0])
	if err != nil {
		return SCSIHCTL{}, errors.WithMessage(err, "invalid host format")
	}
	channel, err := strconv.Atoi(ss[1])
	if err != nil {
		return SCSIHCTL{}, errors.WithMessage(err, "invalid channel format")
	}
	target, err := strconv.Atoi(ss[2])
	if err != nil {
		return SCSIHCTL{}, errors.WithMessage(err, "invalid target format")
	}
	lun, err := strconv.Atoi(ss[3])
	if err != nil {
		return SCSIHCTL{}, errors.WithMessage(err, "invalid lun format")
	}
	return SCSIHCTL{host, channel, target, lun}, nil
}

func (s SCSIHCTLString) MustParse() SCSIHCTL {
	hctl, err := s.Parse()
	if err != nil {
		panic(err)
	}
	return hctl
}

type BlockDeviceInfo struct {
	Name               string           `json:"name" col:"NAME"`                      // device name
	KernelName         string           `json:"kernelName" col:"KNAME"`               // internal kernel device name
	Path               string           `json:"path" col:"PATH"`                      // path to the device node
	MajorMinor         MajorMinorString `json:"majorMinor" col:"MAJ:MIN"`             // major:minor device number
	FSAvailable        Int64String      `json:"fsAvailable" col:"FSAVAIL"`            // filesystem size available
	FSSize             Int64String      `json:"fsSize" col:"FSSIZE"`                  // filesystem size
	FSType             string           `json:"fsType" col:"FSTYPE"`                  // filesystem type
	FSUsed             Int64String      `json:"fsUsed" col:"FSUSED"`                  // filesystem size used
	FSUsedPercentage   string           `json:"fsUsedPercent" col:"FSUSE%"`           // filesystem use percentage
	FSRoots            string           `json:"fsRoots" col:"FSROOTS"`                // mounted filesystem roots
	FSVersion          string           `json:"fsVersion" col:"FSVER"`                // filesystem version
	MountPoint         string           `json:"mountPoint" col:"MOUNTPOINT"`          // where the device is mounted
	MountPoints        string           `json:"mountPoints" col:"MOUNTPOINTS"`        // all locations where device is mounted
	Label              string           `json:"label" col:"LABEL"`                    // filesystem LABEL
	UUID               string           `json:"uuid"  col:"UUID"`                     // filesystem UUID
	PartitionTableUUID string           `json:"partitionTableUUID" col:"PTUUID"`      // partition table identifier (usually UUID)
	PartitionTableType string           `json:"partitionTableType" col:"PTTYPE"`      // partition table type
	PartitionType      string           `json:"partitionType" col:"PARTTYPE"`         // partition type code or UUID
	PartitionTypeName  string           `json:"partitionTypeName" col:"PARTTYPENAME"` // partition type name
	PartitionLabel     string           `json:"partitionLabel" col:"PARTLABEL"`       // partition LABEL
	PartitionUUID      string           `json:"partitionUUID" col:"PARTUUID"`         // partition UUID
	PartitionFlags     string           `json:"partitionFlags" col:"PARTFLAGS"`       // partition flags
	ReadAhead          Int64String      `json:"readAhead" col:"RA"`                   // read-ahead of the device
	ReadOnly           BoolString       `json:"readOnly" col:"RO"`                    // read-only device
	Removable          BoolString       `json:"removable" col:"RM"`                   // removable device
	HotPlug            BoolString       `json:"hotplug" col:"HOTPLUG"`                // removable or hotplug device (usb, pcmcia, ...)
	Model              string           `json:"model" col:"MODEL"`                    // device identifier
	Serial             string           `json:"serial" col:"SERIAL"`                  // disk serial number
	Size               Int64String      `json:"size" col:"SIZE"`                      // size of the device
	State              string           `json:"state" col:"STATE"`                    // state of the device
	Owner              string           `json:"owner" col:"OWNER"`                    // user name
	Group              string           `json:"group" col:"GROUP"`                    // group name
	Mode               string           `json:"mode" col:"MODE"`                      // device node permissions
	Alignment          Int64String      `json:"alignment" col:"ALIGNMENT"`            // alignment offset
	MinimumIO          Int64String      `json:"minimumIO" col:"MIN-IO"`               // minimum I/O size
	OptimalIO          Int64String      `json:"optimalIO" col:"OPT-IO"`               // optimal I/O size
	PhysicalSector     Int64String      `json:"physicalSector" col:"PHY-SEC"`         // physical sector size
	LogicalSector      Int64String      `json:"logicalSector" col:"LOG-SEC"`          // logical sector size
	Rotational         BoolString       `json:"rotational" col:"ROTA"`                // rotational device
	Scheduler          string           `json:"scheduler" col:"SCHED"`                // I/O scheduler name
	RequestQueueSize   Int64String      `json:"requestQueueSize" col:"RQ-SIZE"`       // request queue size
	Type               string           `json:"type" col:"TYPE"`                      // device type
	DiscardAlignment   Int64String      `json:"discardAlignment" col:"DISC-ALN"`      // discard alignment offset
	DiscardGranularity Int64String      `json:"discardGranularity" col:"DISC-GRAN"`   // discard granularity
	DiscardMaxBytes    Int64String      `json:"discardMaxBytes" col:"DISC-MAX"`       // discard max bytes
	DiscardZero        Int64String      `json:"discardZero" col:"DISC-ZERO"`          // discard zeroes data
	WriteSame          Int64String      `json:"writeSame" col:"WSAME"`                // write same max bytes
	WWN                string           `json:"wwn" col:"WWN"`                        // unique storage identifier
	Randomness         BoolString       `json:"randomness" col:"RAND"`                // adds randomness
	ParentKernelName   string           `json:"parentKernelName" col:"PKNAME"`        // internal parent kernel device name
	HCTL               SCSIHCTLString   `json:"hctl" col:"HCTL"`                      // Host:Channel:Target:Lun for SCSI
	Transport          string           `json:"transport" col:"TRAN"`                 // device transport type
	SubSystems         string           `json:"subsystems" col:"SUBSYSTEMS"`          // de-duplicated chain of subsystems
	Revision           string           `json:"revision" col:"REV"`                   // device revision
	Vendor             string           `json:"vendor" col:"VENDOR"`                  // device vendor
	Zoned              string           `json:"zoned" col:"ZONED"`                    // zone model
	Dax                string           `json:"dax" col:"DAX"`                        // dax-capable device

	Children []BlockDeviceInfo `json:"children"` // children devices
}

func (b BlockDeviceInfo) IsEmpty() bool {
	return reflect.DeepEqual(b, BlockDeviceInfo{})
}
