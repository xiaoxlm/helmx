package kubetypes

import (
	"github.com/go-courier/helmx/constants"
)

type KubePodSpec struct {
	KubeVolumes                   `yaml:",inline"`
	KubeInitContainers            `yaml:",inline"`
	KubeContainers                `yaml:",inline"`
	KubeImagePullSecrets          `yaml:",inline"`
	KubeTolerations               `yaml:",inline"`
	PodOpts                       `yaml:",inline"`
	HostAliases                   []KubeHosts `yaml:"hostAliases,omitempty" json:"hostAliases,omitempty"`
	KubeTopologySpreadConstraints `yaml:",inline"`
	KubeAffinity                  `yaml:",inline"`
}

type PodOpts struct {
	RestartPolicy                 string            `json:"restartPolicy,omitempty" yaml:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds *int64            `json:"terminationGracePeriodSeconds,omitempty" yaml:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds         *int64            `json:"activeDeadlineSeconds,omitempty" yaml:"activeDeadlineSeconds,omitempty"`
	DNSPolicy                     string            `json:"dnsPolicy,omitempty" yaml:"dnsPolicy,omitempty"`
	DNSConfig                     *DNSConfig        `json:"dnsConfig,omitempty" yaml:"dnsConfig,omitempty"`
	NodeSelector                  map[string]string `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	ServiceAccountName            string            `json:"serviceAccountName,omitempty" yaml:"serviceAccountName,omitempty"`
	HostNetwork                   *bool             `json:"hostNetwork,omitempty" yaml:"hostNetwork,omitempty"`
}

type DNSConfig struct {
	Nameservers []string     `json:"nameservers,omitempty" yaml:"nameservers,omitempty"`
	Searches    []string     `json:"searches,omitempty" yaml:"searches,omitempty"`
	Options     []KubeOption `json:"options,omitempty" yaml:"options,omitempty"`
}

type KubeInitContainers struct {
	InitContainers []KubeContainer `yaml:"initContainers,omitempty"`
}

type KubeContainers struct {
	Containers []KubeContainer `yaml:"containers,omitempty"`
}

type KubeImagePullSecrets struct {
	ImagePullSecrets []KubeLocalObjectReference `yaml:"imagePullSecrets,omitempty"`
}

type KubeContainer struct {
	Name               string               `yaml:"name"`
	Command            []string             `yaml:"command,omitempty"`
	Args               []string             `yaml:"args,omitempty"`
	WorkingDir         string               `yaml:"workingDir,omitempty"`
	TTY                bool                 `yaml:"tty,omitempty"`
	Resources          ResourceRequirements `yaml:"resources,omitempty"`
	Lifecycle          *Lifecycle           `yaml:"lifecycle,omitempty"`
	ReadinessProbe     *Probe               `yaml:"readinessProbe,omitempty"`
	LivenessProbe      *Probe               `yaml:"livenessProbe,omitempty"`
	SecurityContext    *SecurityContext     `yaml:"securityContext,omitempty"`
	KubeImage          `yaml:",inline"`
	KubeContainerPorts `yaml:",inline"`
	KubeVolumeMounts   `yaml:",inline"`
	KubeEnv            `yaml:",inline"`
}

type KubeImage struct {
	Image           string               `yaml:"image,omitempty"`
	ImagePullPolicy constants.PullPolicy `yaml:"imagePullPolicy,omitempty"`
}

type KubeEnv struct {
	Env []KubeEnvVar `yaml:"env,omitempty"`
}

type KubeEnvVar struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type KubeContainerPorts struct {
	Ports []KubeContainerPort `yaml:"ports,omitempty"`
}

type KubeContainerPort struct {
	ContainerPort uint16             `yaml:"containerPort"`
	Protocol      constants.Protocol `yaml:"protocol,omitempty"`
}

type KubeVolumeMounts struct {
	VolumeMounts []KubeVolumeMount `yaml:"volumeMounts,omitempty"`
}

type KubeVolumeMount struct {
	Name      string `yaml:"name"`
	MountPath string `yaml:"mountPath"`
	SubPath   string `yaml:"subPath,omitempty"`
	ReadOnly  bool   `yaml:"readOnly,omitempty"`
}

type Lifecycle struct {
	PostStart *Handler `yaml:"postStart,omitempty"`
	PreStop   *Handler `yaml:"preStop,omitempty"`
}

type Probe struct {
	Handler   `yaml:",inline"`
	ProbeOpts `yaml:",inline"`
}

type ProbeOpts struct {
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty" yaml:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      int32 `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`
	PeriodSeconds       int32 `json:"periodSeconds,omitempty" yaml:"periodSeconds,omitempty"`
	SuccessThreshold    int32 `json:"successThreshold,omitempty" yaml:"successThreshold,omitempty"`
	FailureThreshold    int32 `json:"failureThreshold,omitempty" yaml:"failureThreshold,omitempty"`
}

type Handler struct {
	Exec      *ExecAction      `yaml:"exec,omitempty"`
	HTTPGet   *HTTPGetAction   `yaml:"httpGet,omitempty"`
	TCPSocket *TCPSocketAction `yaml:"tcpSocket,omitempty"`
}

type ExecAction struct {
	Command []string `yaml:"command,omitempty"`
}

type HTTPGetAction struct {
	Port        uint16       `yaml:"port"`
	Path        string       `yaml:"path,omitempty"`
	Host        string       `yaml:"host,omitempty"`
	Scheme      string       `yaml:"scheme,omitempty"`
	HTTPHeaders []HTTPHeader `yaml:"httpHeaders,omitempty"`
}

type HTTPHeader struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type TCPSocketAction struct {
	Port uint16 `yaml:"port"`
	Host string `yaml:"host,omitempty"`
}

type KubeHosts struct {
	Ip        string   `yaml:"ip" json:"ip"`
	HostNames []string `yaml:"hostnames" json:"hostNames"`
}
