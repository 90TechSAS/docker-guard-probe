package dockerapi

type DockerInfo struct {
	Containers         int         `json:"Containers"`
	Debug              int         `json:"Debug"`
	DockerRootDir      string      `json:"DockerRootDir"`
	Driver             string      `json:"Driver"`
	DriverStatus       [][]string  `json:"DriverStatus"`
	ExecutionDriver    string      `json:"ExecutionDriver"`
	ID                 string      `json:"ID"`
	IPv4Forwarding     int         `json:"IPv4Forwarding"`
	Images             int         `json:"Images"`
	IndexServerAddress string      `json:"IndexServerAddress"`
	InitPath           string      `json:"InitPath"`
	InitSha1           string      `json:"InitSha1"`
	KernelVersion      string      `json:"KernelVersion"`
	Labels             interface{} `json:"Labels"`
	MemTotal           int         `json:"MemTotal"`
	MemoryLimit        int         `json:"MemoryLimit"`
	NCPU               int         `json:"NCPU"`
	NEventsListener    int         `json:"NEventsListener"`
	NFd                int         `json:"NFd"`
	NGoroutines        int         `json:"NGoroutines"`
	Name               string      `json:"Name"`
	OperatingSystem    string      `json:"OperatingSystem"`
	RegistryConfig     struct {
		IndexConfigs struct {
			Docker_io struct {
				Mirrors  interface{} `json:"Mirrors"`
				Name     string      `json:"Name"`
				Official bool        `json:"Official"`
				Secure   bool        `json:"Secure"`
			} `json:"docker.io"`
		} `json:"IndexConfigs"`
		InsecureRegistryCIDRs []string `json:"InsecureRegistryCIDRs"`
	} `json:"RegistryConfig"`
	SwapLimit  int    `json:"SwapLimit"`
	SystemTime string `json:"SystemTime"`
}

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
