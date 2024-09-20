package go_lsblk

import (
	"bufio"
	"github.com/pkg/errors"
	"os/exec"
	"reflect"
	"strings"
)

func ListBlockDevice() ([]BlockDeviceInfo, error) {
	cmd := exec.Command("lsblk", "-P", "-b", "-O")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to execute lsblk")
	}

	disks := make([]BlockDeviceInfo, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 || fields[0] == "NAME" {
			continue
		}
		disk := BlockDeviceInfo{}
		diskNotEmpty := false
		for _, kvPair := range fields {
			if !strings.Contains(kvPair, "=") {
				continue
			}
			parts := strings.Split(kvPair, "=")
			if len(parts) < 2 {
				continue
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if len(parts) > 2 {
				value = strings.Join(parts[1:], "=")
			}
			if key == "" || value == "" || value == `""` {
				continue
			}
			value = strings.TrimPrefix(value, `"`)
			value = strings.TrimSuffix(value, `"`)
			dType := reflect.TypeOf(disk)
			for i := 0; i < dType.NumField(); i++ {
				dField := dType.Field(i)
				if dField.Tag.Get("col") == key {
					fieldValue := reflect.ValueOf(&disk).Elem().FieldByName(dField.Name)
					fieldValue.SetString(value)
					diskNotEmpty = true
					break
				}
			}
		}
		if diskNotEmpty {
			disks = append(disks, disk)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.WithMessagef(err, "error reading output: %s", out)
	}

	return disks, nil
}

func ListBlockDeviceAsTree() ([]BlockDeviceInfo, error) {
	devices, err := ListBlockDevice()
	if err != nil {
		return nil, err
	}
	return buildTree("", devices), nil
}

func buildTree(pkName string, devices []BlockDeviceInfo) []BlockDeviceInfo {
	pDevs := make([]BlockDeviceInfo, 0)
	leftDevs := make([]BlockDeviceInfo, 0)
	for _, dev := range devices {
		if dev.ParentKernelName == pkName {
			pDevs = append(pDevs, dev)
		} else {
			leftDevs = append(leftDevs, dev)
		}
	}
	for i := range pDevs {
		pDevs[i].Children = buildTree(pDevs[i].KernelName, leftDevs)
	}
	return pDevs
}
