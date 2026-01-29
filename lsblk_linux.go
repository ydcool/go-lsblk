package go_lsblk

import (
	"bufio"
	"github.com/pkg/errors"
	"os/exec"
	"reflect"
	"strings"
)

func parseKeyValuePairs(line string) map[string]string {
	result := make(map[string]string)
	line = strings.TrimSpace(line)
	if line == "" {
		return result
	}

	var key, value string
	var inQuotes bool
	var current strings.Builder
	var state int // 0: looking for key, 1: in key, 2: looking for =, 3: looking for value, 4: in value

	for _, r := range line {
		switch state {
		case 0:
			if r != ' ' {
				current.WriteRune(r)
				state = 1
			}
		case 1:
			if r == '=' {
				key = current.String()
				current.Reset()
				state = 3
			} else {
				current.WriteRune(r)
			}
		case 3:
			if r == '"' {
				inQuotes = true
				state = 4
			} else if r != ' ' {
				current.WriteRune(r)
				state = 4
			}
		case 4:
			if inQuotes {
				if r == '"' {
					inQuotes = false
					value = current.String()
					current.Reset()
					result[key] = value
					state = 0
				} else {
					current.WriteRune(r)
				}
			} else {
				if r == ' ' {
					value = current.String()
					current.Reset()
					result[key] = value
					state = 0
				} else {
					current.WriteRune(r)
				}
			}
		}
	}

	// Handle last key-value pair if line ends without space
	if key != "" && (state == 4 || current.Len() > 0) {
		if current.Len() > 0 {
			value = current.String()
		}
		result[key] = value
	}

	return result
}

func ListBlockDevice() ([]BlockDeviceInfo, error) {
	cmd := exec.Command("lsblk", "--pairs", "--bytes", "--output-all")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to execute lsblk")
	}

	disks := make([]BlockDeviceInfo, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		fields := parseKeyValuePairs(line)
		if len(fields) < 2 || fields["NAME"] == "" {
			continue
		}
		disk := BlockDeviceInfo{}
		dType := reflect.TypeOf(disk)
		diskAssigned := false
		for key, value := range fields {
			if key == "" || value == "" || value == `""` {
				continue
			}

			value = strings.TrimSuffix(strings.TrimPrefix(value, `"`), `"`)
			for i := 0; i < dType.NumField(); i++ {
				dField := dType.Field(i)
				if dField.Tag.Get("col") == key {
					fieldValue := reflect.ValueOf(&disk).Elem().FieldByName(dField.Name)
					fieldValue.SetString(value)
					diskAssigned = true
					break
				}
			}
		}
		if diskAssigned {
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
