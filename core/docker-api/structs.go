package dockerapi

/*
	Docker API struct
*/
type DockerInfos struct {
	Containers         int         `json:"Containers"`
	Debug              bool        `json:"Debug"`
	DockerRootDir      string      `json:"DockerRootDir"`
	Driver             string      `json:"Driver"`
	DriverStatus       [][]string  `json:"DriverStatus"`
	ExecutionDriver    string      `json:"ExecutionDriver"`
	ID                 string      `json:"ID"`
	IPv4Forwarding     bool        `json:"IPv4Forwarding"`
	Images             int         `json:"Images"`
	IndexServerAddress string      `json:"IndexServerAddress"`
	InitPath           string      `json:"InitPath"`
	InitSha1           string      `json:"InitSha1"`
	KernelVersion      string      `json:"KernelVersion"`
	Labels             interface{} `json:"Labels"`
	MemTotal           int         `json:"MemTotal"`
	MemoryLimit        bool        `json:"MemoryLimit"`
	NCPU               int         `json:"NCPU"`
	NEventsListener    int         `json:"NEventsListener"`
	NFd                int         `json:"NFd"`
	NGoroutines        int         `json:"NGoroutines"`
	Name               string      `json:"Name"`
	OperatingSystem    string      `json:"OperatingSystem"`
	RegistryConfig     struct {
		IndexConfigs struct {
			Dockerio struct {
				Mirrors  interface{} `json:"Mirrors"`
				Name     string      `json:"Name"`
				Official bool        `json:"Official"`
				Secure   bool        `json:"Secure"`
			} `json:"docker.io"`
		} `json:"IndexConfigs"`
		InsecureRegistryCIDRs []string `json:"InsecureRegistryCIDRs"`
	} `json:"RegistryConfig"`
	SwapLimit  bool   `json:"SwapLimit"`
	SystemTime string `json:"SystemTime"`
}

/*
	Docker API struct
*/
type DockerVersion struct {
	APIVersion    string `json:"ApiVersion"`
	Arch          string `json:"Arch"`
	Experimental  bool   `json:"Experimental"`
	GitCommit     string `json:"GitCommit"`
	GoVersion     string `json:"GoVersion"`
	KernelVersion string `json:"KernelVersion"`
	Os            string `json:"Os"`
	Version       string `json:"Version"`
}

/*
	Docker API struct
*/
type ContainerShort struct {
	Command string   `json:"Command"`
	Created int      `json:"Created"`
	ID      string   `json:"Id"`
	Image   string   `json:"Image"`
	Labels  struct{} `json:"Labels"`
	Names   []string `json:"Names"`
	Ports   []struct {
		PrivatePort int    `json:"PrivatePort"`
		Type        string `json:"Type"`
	} `json:"Ports"`
	Status string `json:"Status"`
}

/*
	Docker API struct
*/
type ContainerShortSize struct {
	Command    string        `json:"Command"`
	Created    int           `json:"Created"`
	ID         string        `json:"Id"`
	Image      string        `json:"Image"`
	Labels     struct{}      `json:"Labels"`
	Names      []string      `json:"Names"`
	Ports      []interface{} `json:"Ports"`
	SizeRootFs float64       `json:"SizeRootFs"`
	SizeRw     float64       `json:"SizeRw"`
	Status     string        `json:"Status"`
}

/*
	Docker API struct
*/
type Container struct {
	AppArmorProfile string   `json:"AppArmorProfile"`
	Args            []string `json:"Args"`
	Config          struct {
		AttachStderr    bool        `json:"AttachStderr"`
		AttachStdin     bool        `json:"AttachStdin"`
		AttachStdout    bool        `json:"AttachStdout"`
		Cmd             []string    `json:"Cmd"`
		Domainname      string      `json:"Domainname"`
		Entrypoint      []string    `json:"Entrypoint"`
		Env             []string    `json:"Env"`
		ExposedPorts    struct{}    `json:"ExposedPorts"`
		Hostname        string      `json:"Hostname"`
		Image           string      `json:"Image"`
		Labels          struct{}    `json:"Labels"`
		MacAddress      string      `json:"MacAddress"`
		NetworkDisabled bool        `json:"NetworkDisabled"`
		OnBuild         interface{} `json:"OnBuild"`
		OpenStdin       bool        `json:"OpenStdin"`
		PortSpecs       interface{} `json:"PortSpecs"`
		StdinOnce       bool        `json:"StdinOnce"`
		Tty             bool        `json:"Tty"`
		User            string      `json:"User"`
		VolumeDriver    string      `json:"VolumeDriver"`
		Volumes         interface{} `json:"Volumes"`
		WorkingDir      string      `json:"WorkingDir"`
	} `json:"Config"`
	Created    string      `json:"Created"`
	Driver     string      `json:"Driver"`
	ExecDriver string      `json:"ExecDriver"`
	ExecIDs    interface{} `json:"ExecIDs"`
	HostConfig struct {
		Binds           []string      `json:"Binds"`
		BlkioWeight     int           `json:"BlkioWeight"`
		CapAdd          interface{}   `json:"CapAdd"`
		CapDrop         interface{}   `json:"CapDrop"`
		CgroupParent    string        `json:"CgroupParent"`
		ContainerIDFile string        `json:"ContainerIDFile"`
		CPUPeriod       int           `json:"CpuPeriod"`
		CPUQuota        int           `json:"CpuQuota"`
		CPUShares       int           `json:"CpuShares"`
		CpusetCpus      string        `json:"CpusetCpus"`
		CpusetMems      string        `json:"CpusetMems"`
		Devices         []interface{} `json:"Devices"`
		DNS             interface{}   `json:"Dns"`
		DNSSearch       interface{}   `json:"DnsSearch"`
		ExtraHosts      interface{}   `json:"ExtraHosts"`
		IpcMode         string        `json:"IpcMode"`
		Links           interface{}   `json:"Links"`
		LogConfig       struct {
			Config struct{} `json:"Config"`
			Type   string   `json:"Type"`
		} `json:"LogConfig"`
		LxcConf         []interface{} `json:"LxcConf"`
		Memory          int           `json:"Memory"`
		MemorySwap      int           `json:"MemorySwap"`
		NetworkMode     string        `json:"NetworkMode"`
		OomKillDisable  bool          `json:"OomKillDisable"`
		PidMode         string        `json:"PidMode"`
		PortBindings    struct{}      `json:"PortBindings"`
		Privileged      bool          `json:"Privileged"`
		PublishAllPorts bool          `json:"PublishAllPorts"`
		ReadonlyRootfs  bool          `json:"ReadonlyRootfs"`
		RestartPolicy   struct {
			MaximumRetryCount int    `json:"MaximumRetryCount"`
			Name              string `json:"Name"`
		} `json:"RestartPolicy"`
		SecurityOpt interface{} `json:"SecurityOpt"`
		UTSMode     string      `json:"UTSMode"`
		Ulimits     interface{} `json:"Ulimits"`
		VolumesFrom interface{} `json:"VolumesFrom"`
	} `json:"HostConfig"`
	HostnamePath    string `json:"HostnamePath"`
	HostsPath       string `json:"HostsPath"`
	ID              string `json:"Id"`
	Image           string `json:"Image"`
	LogPath         string `json:"LogPath"`
	MountLabel      string `json:"MountLabel"`
	Name            string `json:"Name"`
	NetworkSettings struct {
		Bridge                 string      `json:"Bridge"`
		EndpointID             string      `json:"EndpointID"`
		Gateway                string      `json:"Gateway"`
		GlobalIPv6Address      string      `json:"GlobalIPv6Address"`
		GlobalIPv6PrefixLen    int         `json:"GlobalIPv6PrefixLen"`
		HairpinMode            bool        `json:"HairpinMode"`
		IPAddress              string      `json:"IPAddress"`
		IPPrefixLen            int         `json:"IPPrefixLen"`
		IPv6Gateway            string      `json:"IPv6Gateway"`
		LinkLocalIPv6Address   string      `json:"LinkLocalIPv6Address"`
		LinkLocalIPv6PrefixLen int         `json:"LinkLocalIPv6PrefixLen"`
		MacAddress             string      `json:"MacAddress"`
		NetworkID              string      `json:"NetworkID"`
		PortMapping            interface{} `json:"PortMapping"`
		Ports                  struct{}    `json:"Ports"`
		SandboxKey             string      `json:"SandboxKey"`
		SecondaryIPAddresses   interface{} `json:"SecondaryIPAddresses"`
		SecondaryIPv6Addresses interface{} `json:"SecondaryIPv6Addresses"`
	} `json:"NetworkSettings"`
	Path           string `json:"Path"`
	ProcessLabel   string `json:"ProcessLabel"`
	ResolvConfPath string `json:"ResolvConfPath"`
	RestartCount   int    `json:"RestartCount"`
	State          struct {
		Dead       bool   `json:"Dead"`
		Error      string `json:"Error"`
		ExitCode   int    `json:"ExitCode"`
		FinishedAt string `json:"FinishedAt"`
		OOMKilled  bool   `json:"OOMKilled"`
		Paused     bool   `json:"Paused"`
		Pid        int    `json:"Pid"`
		Restarting bool   `json:"Restarting"`
		Running    bool   `json:"Running"`
		StartedAt  string `json:"StartedAt"`
	} `json:"State"`
	Volumes   struct{} `json:"Volumes"`
	VolumesRW struct{} `json:"VolumesRW"`
}

/*
	Docker API struct
*/
type ContainerStats struct {
	BlkioStats struct {
		IoMergedRecursive       []interface{} `json:"io_merged_recursive"`
		IoQueueRecursive        []interface{} `json:"io_queue_recursive"`
		IoServiceBytesRecursive []struct {
			Major int64  `json:"major"`
			Minor int64  `json:"minor"`
			Op    string `json:"op"`
			Value int64  `json:"value"`
		} `json:"io_service_bytes_recursive"`
		IoServiceTimeRecursive []interface{} `json:"io_service_time_recursive"`
		IoServicedRecursive    []struct {
			Major int64  `json:"major"`
			Minor int64  `json:"minor"`
			Op    string `json:"op"`
			Value int64  `json:"value"`
		} `json:"io_serviced_recursive"`
		IoTimeRecursive     []interface{} `json:"io_time_recursive"`
		IoWaitTimeRecursive []interface{} `json:"io_wait_time_recursive"`
		SectorsRecursive    []interface{} `json:"sectors_recursive"`
	} `json:"blkio_stats"`
	CPUStats struct {
		CPUUsage struct {
			PercpuUsage       []int64 `json:"percpu_usage"`
			TotalUsage        int64   `json:"total_usage"`
			UsageInKernelmode int64   `json:"usage_in_kernelmode"`
			UsageInUsermode   int64   `json:"usage_in_usermode"`
		} `json:"cpu_usage"`
		SystemCPUUsage int64 `json:"system_cpu_usage"`
		ThrottlingData struct {
			Periods          int64 `json:"periods"`
			ThrottledPeriods int64 `json:"throttled_periods"`
			ThrottledTime    int64 `json:"throttled_time"`
		} `json:"throttling_data"`
	} `json:"cpu_stats"`
	MemoryStats struct {
		Failcnt  int64 `json:"failcnt"`
		Limit    int64 `json:"limit"`
		MaxUsage int64 `json:"max_usage"`
		Stats    struct {
			ActiveAnon              int64 `json:"active_anon"`
			ActiveFile              int64 `json:"active_file"`
			Cache                   int64 `json:"cache"`
			HierarchicalMemoryLimit int64 `json:"hierarchical_memory_limit"`
			InactiveAnon            int64 `json:"inactive_anon"`
			InactiveFile            int64 `json:"inactive_file"`
			MappedFile              int64 `json:"mapped_file"`
			Pgfault                 int64 `json:"pgfault"`
			Pgmajfault              int64 `json:"pgmajfault"`
			Pgpgin                  int64 `json:"pgpgin"`
			Pgpgout                 int64 `json:"pgpgout"`
			Rss                     int64 `json:"rss"`
			RssHuge                 int64 `json:"rss_huge"`
			TotalActiveAnon         int64 `json:"total_active_anon"`
			TotalActiveFile         int64 `json:"total_active_file"`
			TotalCache              int64 `json:"total_cache"`
			TotalInactiveAnon       int64 `json:"total_inactive_anon"`
			TotalInactiveFile       int64 `json:"total_inactive_file"`
			TotalMappedFile         int64 `json:"total_mapped_file"`
			TotalPgfault            int64 `json:"total_pgfault"`
			TotalPgmajfault         int64 `json:"total_pgmajfault"`
			TotalPgpgin             int64 `json:"total_pgpgin"`
			TotalPgpgout            int64 `json:"total_pgpgout"`
			TotalRss                int64 `json:"total_rss"`
			TotalRssHuge            int64 `json:"total_rss_huge"`
			TotalUnevictable        int64 `json:"total_unevictable"`
			TotalWriteback          int64 `json:"total_writeback"`
			Unevictable             int64 `json:"unevictable"`
			Writeback               int64 `json:"writeback"`
		} `json:"stats"`
		Usage int64 `json:"usage"`
	} `json:"memory_stats"`
	Network struct {
		RxBytes   int64 `json:"rx_bytes"`
		RxDropped int64 `json:"rx_dropped"`
		RxErrors  int64 `json:"rx_errors"`
		RxPackets int64 `json:"rx_packets"`
		TxBytes   int64 `json:"tx_bytes"`
		TxDropped int64 `json:"tx_dropped"`
		TxErrors  int64 `json:"tx_errors"`
		TxPackets int64 `json:"tx_packets"`
	} `json:"network"`
	PrecpuStats struct {
		CPUUsage struct {
			PercpuUsage       []int64 `json:"percpu_usage"`
			TotalUsage        int64   `json:"total_usage"`
			UsageInKernelmode int64   `json:"usage_in_kernelmode"`
			UsageInUsermode   int64   `json:"usage_in_usermode"`
		} `json:"cpu_usage"`
		SystemCPUUsage int64 `json:"system_cpu_usage"`
		ThrottlingData struct {
			Periods          int64 `json:"periods"`
			ThrottledPeriods int64 `json:"throttled_periods"`
			ThrottledTime    int64 `json:"throttled_time"`
		} `json:"throttling_data"`
	} `json:"precpu_stats"`
	Read string `json:"read"`
}
