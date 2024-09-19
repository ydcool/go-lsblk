package go_lsblk

import (
	"bufio"
	"fmt"
	"os/exec"
	"reflect"
	"strings"
)

func ListBlockDevice() ([]BlockDeviceInfo, error) {
	cmd := exec.Command("lsblk", "-P", "-b", "-O")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute lsblk: %w", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	disks := make([]BlockDeviceInfo, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 || fields[0] == "NAME" {
			continue
		}
		disk := BlockDeviceInfo{}
		diskNotEmpty := false
		for _, kv := range fields {
			if strings.Contains(kv, "=") {
				parts := strings.Split(kv, "=")
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				dType := reflect.TypeOf(disk)
				for i := 0; i < dType.NumField(); i++ {
					dField := dType.Field(i)
					if dField.Tag.Get("col") == key {
						fieldValue := reflect.ValueOf(&disk).Elem().FieldByName(dField.Name)
						fieldValue.SetString(value)
						diskNotEmpty = true
						continue
					}
				}
			}
		}
		if diskNotEmpty {
			disks = append(disks, disk)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading output: %w", err)
	}

	return disks, nil
}
