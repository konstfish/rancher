package management

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"

	v32 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v3 "github.com/rancher/rancher/pkg/generated/norman/management.cattle.io/v3"
	"github.com/rancher/rancher/pkg/types/config"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Amazonec2driver    = "amazonec2"
	Azuredriver        = "azure"
	DigitalOceandriver = "digitalocean"
	ExoscaleDriver     = "exoscale"
	HarvesterDriver    = "harvester"
	Linodedriver       = "linode"
	NutanixDriver      = "nutanix"
	OCIDriver          = "oci"
	OTCDriver          = "otc"
	OpenstackDriver    = "openstack"
	PacketDriver       = "packet"
	PhoenixNAPDriver   = "pnap"
	RackspaceDriver    = "rackspace"
	SoftLayerDriver    = "softlayer"
	Vmwaredriver       = "vmwarevsphere"
	GoogleDriver       = "google"
	OutscaleDriver     = "outscale"
)

/*var DriverData = map[string]map[string][]string{
	Amazonec2driver:    {"publicCredentialFields": []string{"accessKey"}, "privateCredentialFields": []string{"secretKey"}},
	Azuredriver:        {"publicCredentialFields": []string{"clientId", "subscriptionId", "tenantId", "environment"}, "privateCredentialFields": []string{"clientSecret"}, "optionalCredentialFields": []string{"tenantId"}},
	DigitalOceandriver: {"privateCredentialFields": []string{"accessToken"}},
	ExoscaleDriver:     {"privateCredentialFields": []string{"apiSecretKey"}},
	HarvesterDriver:    {"publicCredentialFields": []string{"clusterType", "clusterId"}, "privateCredentialFields": []string{"kubeconfigContent"}, "optionalCredentialFields": []string{"clusterId"}},
	Linodedriver:       {"privateCredentialFields": []string{"token"}, "passwordFields": []string{"rootPass"}},
	NutanixDriver:      {"publicCredentialFields": []string{"endpoint", "username", "port"}, "privateCredentialFields": []string{"password"}},
	OCIDriver:          {"publicCredentialFields": []string{"tenancyId", "userId", "fingerprint"}, "privateCredentialFields": []string{"privateKeyContents"}, "passwordFields": []string{"privateKeyPassphrase"}},
	OTCDriver:          {"privateCredentialFields": []string{"accessKeySecret"}},
	OpenstackDriver:    {"privateCredentialFields": []string{"password"}},
	PacketDriver:       {"privateCredentialFields": []string{"apiKey"}},
	PhoenixNAPDriver:   {"publicCredentialFields": []string{"clientIdentifier"}, "privateCredentialFields": []string{"clientSecret"}},
	RackspaceDriver:    {"privateCredentialFields": []string{"apiKey"}},
	SoftLayerDriver:    {"privateCredentialFields": []string{"apiKey"}},
	Vmwaredriver:       {"publicCredentialFields": []string{"username", "vcenter", "vcenterPort"}, "privateCredentialFields": []string{"password"}},
	GoogleDriver:       {"privateCredentialFields": []string{"authEncodedJson"}},
	OutscaleDriver:     {"publicCredentialFields": []string{"accessKey", "region"}, "privateCredentialFields": []string{"secretKey"}},
}*/

type DriverData struct {
	Name               string
	Url                map[string]DriverUrl
	UiUrl              string
	Whitelist          []string
	Active             bool
	Builtin            bool
	AddCloudCredential bool
	Annotations        map[string][]string
}

type DriverUrl struct {
	Url      string
	Checksum string
}

var driverData = []DriverData{
	{
		Name: "pinganyunecs",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "https://drivers.rancher.cn/node-driver-pinganyun/0.3.0/docker-machine-driver-pinganyunecs.tgz",
				Checksum: "f84ccec11c2c1970d76d30150916933efe8ca49fe4c422c8954fc37f71273bb5",
			},
			"arm64": DriverUrl{
				Url:      "https://drivers.rancher.cn/node-driver-pinganyun/0.3.0/docker-machine-driver-pinganyunecs.tgz",
				Checksum: "f84ccec11c2c1970d76d30150916933efe8ca49fe4c422c8954fc37f71273bb5",
			},
		},
		UiUrl:              "https://drivers.rancher.cn/node-driver-pinganyun/0.3.0/component.js",
		Whitelist:          []string{"drivers.rancher.cn"},
		Active:             false,
		Builtin:            false,
		AddCloudCredential: false,
		Annotations:        nil,
	},
	{
		Name: "aliyunecs",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "https://drivers.rancher.cn/node-driver-aliyun/1.0.4/docker-machine-driver-aliyunecs.tgz",
				Checksum: "5990d40d71c421a85563df9caf069466f300cd75723effe4581751b0de9a6a0e",
			},
			"arm64": DriverUrl{
				Url:      "https://drivers.rancher.cn/node-driver-aliyun/1.0.4/docker-machine-driver-aliyunecs.tgz",
				Checksum: "5990d40d71c421a85563df9caf069466f300cd75723effe4581751b0de9a6a0e",
			},
		},
		UiUrl:              "",
		Whitelist:          []string{"ecs.aliyuncs.com"},
		Active:             false,
		Builtin:            false,
		AddCloudCredential: false,
		Annotations:        nil,
	},
	{
		Name: "amazonec2",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
			"arm64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
		},
		UiUrl:              "",
		Whitelist:          []string{"iam.amazonaws.com", "iam.us-gov.amazonaws.com", "iam.%.amazonaws.com.cn", "ec2.%.amazonaws.com", "ec2.%.amazonaws.com.cn", "eks.%.amazonaws.com", "eks.%.amazonaws.com.cn", "kms.%.amazonaws.com", "kms.%.amazonaws.com.cn"},
		Active:             true,
		Builtin:            true,
		AddCloudCredential: true,
		Annotations:        nil,
	},
	{
		Name: "azure",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
			"arm64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
		},
		UiUrl:              "",
		Whitelist:          nil,
		Active:             true,
		Builtin:            true,
		AddCloudCredential: true,
		Annotations:        nil,
	},
	{
		Name: "cloudca",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "https://github.com/cloud-ca/docker-machine-driver-cloudca/files/2446837/docker-machine-driver-cloudca_v2.0.0_linux-amd64.zip",
				Checksum: "2a55efd6d62d5f7fd27ce877d49596f4",
			},
			"arm64": DriverUrl{
				Url:      "https://github.com/cloud-ca/docker-machine-driver-cloudca/files/2446837/docker-machine-driver-cloudca_v2.0.0_linux-arm64.zip",
				Checksum: "2a55efd6d62d5f7fd27ce877d49596f4",
			},
		},
		UiUrl:              "https://objects-east.cloud.ca/v1/5ef827605f884961b94881e928e7a250/ui-driver-cloudca/v2.1.2/component.js",
		Whitelist:          []string{"objects-east.cloud.ca"},
		Active:             false,
		Builtin:            false,
		AddCloudCredential: false,
		Annotations:        nil,
	},
	{
		Name: "cloudscale",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "https://github.com/cloudscale-ch/docker-machine-driver-cloudscale/releases/download/v1.2.0/docker-machine-driver-cloudscale_v1.2.0_linux_amd64.tar.gz",
				Checksum: "e33fbd6c2f87b1c470bcb653cc8aa50baf914a9d641a2f18f86a07c398cfb544",
			},
			"arm64": DriverUrl{
				Url:      "https://github.com/cloudscale-ch/docker-machine-driver-cloudscale/releases/download/v1.2.0/docker-machine-driver-cloudscale_v1.2.0_linux_arm64.tar.gz",
				Checksum: "",
			},
		},
		UiUrl:              "https://objects.rma.cloudscale.ch/cloudscale-rancher-v2-ui-driver/component.js",
		Whitelist:          []string{"objects.rma.cloudscale.ch"},
		Active:             false,
		Builtin:            false,
		AddCloudCredential: false,
		Annotations:        nil,
	},
	{
		Name: "digitalocean",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
			"arm64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
		},
		UiUrl:              "",
		Whitelist:          []string{"api.digitalocean.com"},
		Active:             true,
		Builtin:            true,
		AddCloudCredential: false,
		Annotations:        nil,
	},
	{
		Name: "exoscale",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
			"arm64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
		},
		UiUrl:              "",
		Whitelist:          []string{"api.exoscale.ch"},
		Active:             false,
		Builtin:            true,
		AddCloudCredential: false,
		Annotations:        nil,
	},
	{
		Name: "google",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
			"arm64": DriverUrl{
				Url:      "local://",
				Checksum: "",
			},
		},
		UiUrl:              "",
		Whitelist:          nil,
		Active:             false,
		Builtin:            true,
		AddCloudCredential: true,
		Annotations:        nil,
	},
	{
		Name: "harvester",
		Url: map[string]DriverUrl{
			"amd64": DriverUrl{
				Url:      "https://releases.rancher.com/harvester-node-driver/v0.6.5/docker-machine-driver-harvester-amd64.tar.gz",
				Checksum: "8de48b07dd2e8b7ee60ec99b8456925e9c16a7523affb61a5f1788868bb1f8f6",
			},
			"arm64": DriverUrl{
				Url:      "https://github.com/cloud-ca/docker-machine-driver-cloudca/files/2446837/docker-machine-driver-cloudca_v2.0.0_linux-arm64.zip",
				Checksum: "",
			},
		},
		UiUrl:              "",
		Whitelist:          []string{"releases.rancher.com"},
		Active:             true,
		Builtin:            true,
		AddCloudCredential: false,
		Annotations:        nil,
	},
}

var driverDefaults = map[string]map[string]string{
	HarvesterDriver: {"clusterType": "imported"},
	Vmwaredriver:    {"vcenterPort": "443"},
}

type machineDriverCompare struct {
	builtin            bool
	addCloudCredential bool
	url                string
	uiURL              string
	checksum           string
	name               string
	whitelist          []string
	annotations        map[string]string
}

func addMachineDrivers(management *config.ManagementContext) error {
	/*var arch = runtime.GOARCH

	for _, driver := range driverData {
		if err := addMachineDriver(driver.Name, driver.Url[arch], driver.UiUrl, driver.Checksum, driver.Whitelist, driver.Active, driver.Builtin, driver.AddCloudCredential, driver.Annotations, management); err != nil {
			return err
		}
	}*/

	if err := addMachineDriver("hetzner", "https://github.com/JonasProgrammer/docker-machine-driver-hetzner/releases/download/5.0.2/docker-machine-driver-hetzner_5.0.2_linux_arm64.tar.gz", "https://storage.googleapis.com/hcloud-rancher-v2-ui-driver/component.js", "", []string{"storage.googleapis.com"}, true, false, false, nil, management); err != nil {
		return err
	}
	return nil
}

func addMachineDriver(name, url, uiURL, checksum string, whitelist []string, active, builtin, addCloudCredential bool, driverAnnotations map[string][]string, management *config.ManagementContext) error {
	lister := management.Management.NodeDrivers("").Controller().Lister()
	cli := management.Management.NodeDrivers("")
	m, _ := lister.Get("", name)
	// annotations can have keys cred and password, values []string to be considered as a part of cloud credential
	annotations := map[string]string{}
	if m != nil {
		for k, v := range m.Annotations {
			annotations[k] = v
		}
	}
	for key, fields := range driverAnnotations {
		annotations[key] = strings.Join(fields, ",")
	}
	defaults := []string{}
	for key, val := range driverDefaults[name] {
		defaults = append(defaults, fmt.Sprintf("%s:%s", key, val))
	}
	if len(defaults) > 0 {
		annotations["defaults"] = strings.Join(defaults, ",")
	}
	if m != nil {
		old := machineDriverCompare{
			builtin:            m.Spec.Builtin,
			addCloudCredential: m.Spec.AddCloudCredential,
			url:                m.Spec.URL,
			uiURL:              m.Spec.UIURL,
			checksum:           m.Spec.Checksum,
			name:               m.Spec.DisplayName,
			whitelist:          m.Spec.WhitelistDomains,
			annotations:        m.Annotations,
		}
		new := machineDriverCompare{
			builtin:            builtin,
			addCloudCredential: addCloudCredential,
			url:                url,
			uiURL:              uiURL,
			checksum:           checksum,
			name:               name,
			whitelist:          whitelist,
			annotations:        annotations,
		}
		if !reflect.DeepEqual(new, old) {
			logrus.Infof("Updating node driver %v", name)
			m.Spec.Builtin = builtin
			m.Spec.AddCloudCredential = addCloudCredential
			m.Spec.URL = url
			m.Spec.UIURL = uiURL
			m.Spec.Checksum = checksum
			m.Spec.DisplayName = name
			m.Spec.WhitelistDomains = whitelist
			m.Annotations = annotations
			_, err := cli.Update(m)
			return err
		}
		return nil
	}

	logrus.Infof("Creating node driver %v", name)
	_, err := cli.Create(&v3.NodeDriver{
		ObjectMeta: v1.ObjectMeta{
			Name:        name,
			Annotations: annotations,
		},
		Spec: v32.NodeDriverSpec{
			Active:             active,
			Builtin:            builtin,
			AddCloudCredential: addCloudCredential,
			URL:                url,
			UIURL:              uiURL,
			DisplayName:        name,
			Checksum:           checksum,
			WhitelistDomains:   whitelist,
		},
	})

	return err
}

func isCommandAvailable(name string) bool {
	return exec.Command("command", "-v", name).Run() == nil
}
