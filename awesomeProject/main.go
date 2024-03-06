package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	endpoint = "http://localhost:8080/devices/%d/metrics"
	yamlFile = "registerdevices.yaml"
)

type Device struct {
	DeviceID     int       `yaml:"device_id"`
	Name         string    `yaml:"name"`
	Metric1      int       `yaml:"metric1"`
	Metric2      int       `yaml:"metric2"`
	Metric3      int       `yaml:"metric3"`
	IP           string    `yaml:"ip"`
	DateCreation time.Time `yaml:"date_creation"`
}

type Metrics struct {
	Metric1      uint    `json:"metric_1"`
	Metric1Value float64 `json:"metric_1_value"`
	Metric2      uint    `json:"metric_2"`
	Metric2Value float64 `json:"metric_2_value"`
	Metric3      uint    `json:"metric_3"`
	Metric3Value float64 `json:"metric_3_value"`
}

func main() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := sendMetrics()
			if err != nil {
				log.Printf("error sending metrics: %v", err)
			}
		}
	}
}

func sendMetrics() error {
	// Read device details from YAML file
	device, err := readYAML()
	if err != nil {
		return fmt.Errorf("error reading YAML: %v", err)
	}

	// Check if metrics slice is empty
	if device.Metric1 == 0 && device.Metric2 == 0 && device.Metric3 == 0 {
		return fmt.Errorf("no metrics found in the YAML file")
	}

	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return fmt.Errorf("error getting CPU usage: %v", err)
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Errorf("error getting memory usage: %v", err)
	}

	diskUsage, err := disk.Usage("/")
	if err != nil {
		return fmt.Errorf("error getting disk usage: %v", err)
	}

	// Iterate over metrics and send each one

	metrics := Metrics{
		Metric1:      uint(device.Metric1),
		Metric1Value: cpuUsage[0],
		Metric2:      uint(device.Metric2),
		Metric2Value: memInfo.UsedPercent,
		Metric3:      uint(device.Metric3),
		Metric3Value: diskUsage.UsedPercent,
	}

	// Send metrics to endpoint
	err = postMetrics(device.DeviceID, metrics)
	if err != nil {
		return fmt.Errorf("error posting metrics: %v", err)
	}

	return nil
}

func readYAML() (*Device, error) {
	yamlData, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %v", err)
	}

	var device struct {
		Device Device `yaml:"device"`
	}

	err = yaml.Unmarshal(yamlData, &device)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML data: %v", err)
	}

	fmt.Printf("Parsed device: %+v\n", device.Device)

	return &device.Device, nil
}

func postMetrics(deviceID int, metrics Metrics) error {
	endpointURL := fmt.Sprintf(endpoint, deviceID)
	jsonData, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	resp, err := http.Post(endpointURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
