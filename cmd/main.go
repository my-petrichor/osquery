package main

import (
	"log"
	"os"

	"github.com/my-sakura/osquery/osquery/controller"
	//  flag "github.com/spf13/pflag"
)

var (
	osQueryRouterGroup = "api/v1/osquery"
	tables             = map[string]bool{
		"account_policy_data":              true,
		"acpi_tables":                      true,
		"ad_config":                        true,
		"alf":                              true,
		"alf_exceptions":                   true,
		"alf_explicit_auths":               true,
		"app_schemes":                      true,
		"apps":                             true,
		"apt_sources":                      true,
		"arp_cache":                        true,
		"asl":                              true,
		"atom_packages":                    true,
		"augeas":                           true,
		"authorization_mechanisms":         true,
		"authorizations":                   true,
		"authorized_keys":                  true,
		"azure_instance_metadata":          true,
		"azure_instance_tags":              true,
		"battery":                          true,
		"block_devices":                    true,
		"browser_plugins":                  true,
		"carbon_black_info":                true,
		"carves":                           true,
		"certificates":                     true,
		"chrome_extension_content_scripts": true,
		"chrome_extensions":                true,
		"cpu_time":                         true,
		"cpuid":                            true,
		"crashes":                          true,
		"crontab":                          true,
		"cups_destinations":                true,
		"cups_jobs":                        true,
		"curl":                             true,
		"curl_certificate":                 true,
		"device_file":                      true,
		"device_firmware":                  true,
		"device_hash":                      true,
		"device_partitions":                true,
		"disk_encryption":                  true,
		"disk_events":                      true,
		"dns_resolvers":                    true,
		"docker_container_fs_changes":      true,
		"docker_container_labels":          true,
		"docker_container_mounts":          true,
		"docker_container_networks":        true,
		"docker_container_ports":           true,
		"docker_container_processes":       true,
		"docker_container_stats":           true,
		"docker_containers":                true,
		"docker_image_history":             true,
		"docker_image_labels":              true,
		"docker_image_layers":              true,
		"docker_images":                    true,
		"docker_info":                      true,
		"docker_network_labels":            true,
		"docker_networks":                  true,
		"docker_version":                   true,
		"docker_volume_labels":             true,
		"docker_volumes":                   true,
		"ec2_instance_metadata":            true,
		"ec2_instance_tags":                true,
		"es_process_events":                true,
		"etc_hosts":                        true,
		"etc_protocols":                    true,
		"etc_services":                     true,
		"event_taps":                       true,
		"extended_attributes":              true,
		"fan_speed_sensors":                true,
		"file":                             true,
		"file_events":                      true,
		"firefox_addons":                   true,
		"gatekeeper":                       true,
		"gatekeeper_approved_apps":         true,
		"groups":                           true,
		"hardware_events":                  true,
		"hash":                             true,
		"homebrew_packages":                true,
		"ibridge_info":                     true,
		"interface_addresses":              true,
		"interface_details":                true,
		"interface_ipv6":                   true,
		"iokit_devicetree":                 true,
		"iokit_registry":                   true,
		"kernel_extensions":                true,
		"kernel_info":                      true,
		"kernel_panics":                    true,
		"keychain_acls":                    true,
		"keychain_items":                   true,
		"known_hosts":                      true,
		"last":                             true,
		"launchd":                          true,
		"launchd_overrides":                true,
		"listening_ports":                  true,
		"lldp_neighbors":                   true,
		"load_average":                     true,
		"location_services":                true,
		"logged_in_users":                  true,
		"magic":                            true,
		"managed_policies":                 true,
		"mdfind":                           true,
		"mdls":                             true,
		"memory_array_mapped_addresses":    true,
		"memory_arrays":                    true,
		"memory_device_mapped_addresses":   true,
		"memory_devices":                   true,
		"memory_error_info":                true,
		"mounts":                           true,
		"nfs_shares":                       true,
		"nvram":                            true,
		"oem_strings":                      true,
		"os_version":                       true,
		"osquery_events":                   true,
		"osquery_extensions":               true,
		"osquery_flags":                    true,
		"osquery_info":                     true,
		"osquery_packs":                    true,
		"osquery_registry":                 true,
		"osquery_schedule":                 true,
		"package_bom":                      true,
		"package_install_history":          true,
		"package_receipts":                 true,
		"pci_devices":                      true,
		"platform_info":                    true,
		"plist":                            true,
		"power_sensors":                    true,
		"preferences":                      true,
		"process_envs":                     true,
		"process_events":                   true,
		"process_memory_map":               true,
		"process_open_files":               true,
		"process_open_sockets":             true,
		"processes":                        true,
		"prometheus_metrics":               true,
		"python_packages":                  true,
		"quicklook_cache":                  true,
		"routes":                           true,
		"running_apps":                     true,
		"safari_extensions":                true,
		"sandboxes":                        true,
		"screenlock":                       true,
		"shared_folders":                   true,
		"sharing_preferences":              true,
		"shell_history":                    true,
		"signature":                        true,
		"sip_config":                       true,
		"smart_drive_info":                 true,
		"smbios_tables":                    true,
		"smc_keys":                         true,
		"socket_events":                    true,
		"ssh_configs":                      true,
		"startup_items":                    true,
		"sudoers":                          true,
		"suid_bin":                         true,
		"system_controls":                  true,
		"system_extensions":                true,
		"system_info":                      true,
		"temperature_sensors":              true,
		"time":                             true,
		"time_machine_backups":             true,
		"time_machine_destinations":        true,
		"ulimit_info":                      true,
		"uptime":                           true,
		"usb_devices":                      true,
		"user_events":                      true,
		"user_groups":                      true,
		"user_interaction_events":          true,
		"user_ssh_keys":                    true,
		"users":                            true,
		"virtual_memory_info":              true,
		"wifi_networks":                    true,
		"wifi_status":                      true,
		"wifi_survey":                      true,
		"xprotect_entries":                 true,
		"xprotect_meta":                    true,
		"xprotect_reports":                 true,
		"yara":                             true,
		"yara_events":                      true,
		"ycloud_instance_metadata":         true,
		"yum_sources":                      true,
	}
)

func main() {
	// r := gin.Default()
	osQuery := controller.New()
	// osQuery.Register(r.Group(osQueryRouterGroup))

	// log.Println(r.Run("0.0.0.0:8081"))
	if len(os.Args) != 2 {
		log.Printf("please input one arguments")
		return
	}
	var argument string

	argument = os.Args[1]
	switch argument {
	case "tables":
		osQuery.ListTable()
	default:
		if tables[argument] {
			osQuery.Table(argument)
		}
	}
}
