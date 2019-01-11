# packer-post-processor-futong
Packer post processor to upload templates to Futong CMP


## Installation

Having your Go environment ready, run:

```
go get github.com/mengt/packer-post-processor-futong
cp /root/gobin/packer-post-processor-futong /usr/local/bin/
```
See [packer docs](https://www.packer.io/docs/extending/plugins.html) on how to install the plugin.

## Usage

The plugin provides a new postprocessor named `futong` with the following available options:

| Name                          | Description                                                    |
|-------------------------------|----------------------------------------------------------------|
| `api_url`                     | The Futong API endpoint URL. Eg. `https://some.host/api` |
| `api_insecure`                | A boolean specifying if wrong certificates should be trusted (eg. self-signed certs) |
| `api_username`                | The Username for the Futong API. Eg. `tmpluser` |
| `api_password`                | The password for the aforementioned user. Eg. `somepass` |
| `app_key`                     | Futong OAuth 1.0a app key |
| `app_secret`                  | Futong OAuth 1.0a app secret |
| `access_token`                | Futong OAuth 1.0a access token |
| `access_token_secret`         | Futong OAuth 1.0a access token secret |
| `datacenter`                  | The Futong datacenter name where to upload the template. Eg. `dclocation1` |
| `keep_input_artifact`         | Whether or not to keep the input artifact. Default is `false` |
| `template_name`               | The name of the template to be created in Futong. Defaults to the packer VM name. |
| `description`                 | Description of the template. Defaults to the same value as `name`. |
| `category`                    | Category for the template. Defaults to `OS`. |
| `disk_format`                 | The format of the disk being uploaded. If not specified it will be guessed from the input artifact. |
| `cpu`                         | CPU cores to use by the template. Defaults to `1`. |
| `hd_mb`                       | Size of the disk being uploaded. If not specified it will be guessed from the input artifact. |
| `ram_mb`                      | The amount of RAM in MB to be used by the template. Defaults to `1024`. |
| `login_user`                  | The login username for the template. If empty, we will set the same used by packer. |
| `login_password`              | The password for the `login_user` user. If empty, we will set the same used by packer. |
| `eth_driver`                  | The NIC driver to use in the template. Defaults to `E1000`. |
| `chef_enabled`                | Wether or not the template has the Futong Chef agent installed. Defaults to `false`. |
| `icon_url`                    | The URL of the icon for the template. No URL will be set by default. |
| `cpu_hotadd`                  | Whether or not the guest will support CPU hotplug. Defaults to `false`. |
| `ram_hotadd`                  | Whether or not the guest will support RAM hotplug. Defaults to `false`. |
| `disk_hotadd`                 | Whether or not the guest will support disks hotplug. Defaults to `false`. |
| `nic_hotadd`                  | Whether or not the guest will support NIC hotplug. Defaults to `false`. |
| `vnc_hotadd`                  | Whether or not the guest will support hot reconfigure of VNC remote access. Defaults to `false`. |
| `cpu_min`                     | Minimum vCPU that can be configured for this template. Defaults to nil. |
| `cpu_max`                     | Maximum vCPU that can be configured for this template. Defaults to nil. |
| `ram_min`                     | Minimum RAM that can be configured for this template. Defaults to nil. |
| `ram_max`                     | Maximum RAM that can be configured for this template. Defaults to nil. |
| `guest_setup`                 | Type of guest setup protocol available in the template, either `HYPERVISOR_TOOLS` or `CLOUD_INIT`. Defaults to nil. |
| `enable_hp_recommended`       | Whether or not to force only recommended HP for this template. Defaults to nil. |
| `guest_initial_password`      | Whether or not to generate a random initial password for the admin user on deploy. Defaults to nil (not doing it) |

## Example

Examples of usage:

```
"post-processors": [
    [
      {
        "type": "futong",
        "api_url": "http://IP/api",
        "api_username": "admin",
        "api_password": "futongcloud",
        "datacenter": "template-library-name",
        "template_name": "000-packertest",
        "cpu": "1",
        "description": "A packer test",
        "eth_driver": "E1000",
        "ram_mb": "1024",
        "icon_url": "{{user `icon`}}"
      }
    ]
```
