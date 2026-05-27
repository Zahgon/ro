// Copyright 2025 samber.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://github.com/samber/ro/blob/main/licenses/LICENSE.apache.md
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package roproc

import (
	"time"

	"github.com/samber/ro"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/sensors"
)

// NewVirtualMemoryWatcher creates an observable that emits virtual memory statistics at regular intervals.
func NewVirtualMemoryWatcher(interval time.Duration) ro.Observable[*mem.VirtualMemoryStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewSwapMemoryWatcher creates an observable that emits swap memory statistics at regular intervals.
func NewSwapMemoryWatcher(interval time.Duration) ro.Observable[*mem.SwapMemoryStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewSwapDeviceWatcher creates an observable that emits swap device information at regular intervals.
func NewSwapDeviceWatcher(interval time.Duration) ro.Observable[*mem.SwapDevice] {
	_ = "STUB: not implemented"
	return nil
}

// NewCPUInfoWatcher creates an observable that emits CPU information statistics at regular intervals.
func NewCPUInfoWatcher(interval time.Duration) ro.Observable[cpu.InfoStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewDiskUsageWatcher creates an observable that emits disk usage statistics at regular intervals.
func NewDiskUsageWatcher(interval time.Duration, mountpointOrDevicePath string) ro.Observable[*disk.UsageStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewDiskIOCountersWatcher creates an observable that emits disk I/O counters at regular intervals.
func NewDiskIOCountersWatcher(interval time.Duration, names ...string) ro.Observable[map[string]disk.IOCountersStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewDiskPartitionWatcher creates an observable that emits disk partition information at regular intervals.
func NewDiskPartitionWatcher(interval time.Duration) ro.Observable[disk.PartitionStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewHostInfoWatcher creates an observable that emits host information at regular intervals.
func NewHostInfoWatcher(interval time.Duration) ro.Observable[*host.InfoStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewHostUserWatcher creates an observable that emits host user information at regular intervals.
func NewHostUserWatcher(interval time.Duration) ro.Observable[host.UserStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewLoadAverageWatcher creates an observable that emits load average statistics at regular intervals.
func NewLoadAverageWatcher(interval time.Duration) ro.Observable[*load.AvgStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewLoadMiscWatcher creates an observable that emits miscellaneous load statistics at regular intervals.
func NewLoadMiscWatcher(interval time.Duration) ro.Observable[*load.MiscStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewNetConnectionsWatcher creates an observable that emits network connection statistics at regular intervals.
func NewNetConnectionsWatcher(interval time.Duration) ro.Observable[net.ConnectionStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewNetConntrackWatcher creates an observable that emits network conntrack statistics at regular intervals.
func NewNetConntrackWatcher(interval time.Duration, perCPU bool) ro.Observable[net.ConntrackStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewNetFilterCountersWatcher creates an observable that emits network filter counters at regular intervals.
func NewNetFilterCountersWatcher(interval time.Duration) ro.Observable[net.FilterStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewNetIOCountersWatcher creates an observable that emits network I/O counters at regular intervals.
func NewNetIOCountersWatcher(interval time.Duration, perNIC bool) ro.Observable[net.IOCountersStat] {
	_ = "STUB: not implemented"
	return nil
}

// NewSensorsTemperatureWatcher creates an observable that emits sensor temperature statistics at regular intervals.
func NewSensorsTemperatureWatcher(interval time.Duration, perNIC bool) ro.Observable[sensors.TemperatureStat] {
	_ = "STUB: not implemented"
	return nil
}
