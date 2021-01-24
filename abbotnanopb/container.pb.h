/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.4 */

#ifndef PB_ABBOT_CONTAINER_PB_H_INCLUDED
#define PB_ABBOT_CONTAINER_PB_H_INCLUDED
#include <pb.h>
#include "meta.pb.h"

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

/* Struct definitions */
typedef struct _abbot_CNICapArgs_DNSConfig {
    pb_callback_t servers;
    pb_callback_t searches;
    pb_callback_t options;
} abbot_CNICapArgs_DNSConfig;

typedef struct _abbot_CNICapArgs_DeviceID {
    pb_callback_t device_id;
} abbot_CNICapArgs_DeviceID;

typedef struct _abbot_CNICapArgs_IPAddresses {
    pb_callback_t ips;
} abbot_CNICapArgs_IPAddresses;

typedef struct _abbot_CNICapArgs_IPRange {
    pb_callback_t subnet;
    pb_callback_t range_start;
    pb_callback_t range_end;
    pb_callback_t gateway;
} abbot_CNICapArgs_IPRange;

typedef struct _abbot_CNICapArgs_InfinibandGUID {
    pb_callback_t infiniband_guid;
} abbot_CNICapArgs_InfinibandGUID;

typedef struct _abbot_CNICapArgs_MacAddress {
    pb_callback_t mac;
} abbot_CNICapArgs_MacAddress;

typedef struct _abbot_ContainerNetworkConfigEnsureRequest {
    pb_callback_t ipv4_subnet;
    pb_callback_t ipv6_subnet;
} abbot_ContainerNetworkConfigEnsureRequest;

typedef struct _abbot_ContainerNetworkConfigQueryRequest {
    char dummy_field;
} abbot_ContainerNetworkConfigQueryRequest;

typedef struct _abbot_ContainerNetworkConfigResponse {
    pb_callback_t ipv4_subnet;
    pb_callback_t ipv6_subnet;
} abbot_ContainerNetworkConfigResponse;

typedef struct _abbot_ContainerNetworkEnsureRequest_CniArgsEntry {
    pb_callback_t key;
    pb_callback_t value;
} abbot_ContainerNetworkEnsureRequest_CniArgsEntry;

typedef struct _abbot_ContainerNetworkStatusListResponse {
    pb_callback_t container_networks;
} abbot_ContainerNetworkStatusListResponse;

typedef struct _abbot_CNICapArgs_Bandwidth {
    int32_t ingress_rate;
    int32_t ingress_burst;
    int32_t egress_rate;
    int32_t egress_burst;
} abbot_CNICapArgs_Bandwidth;

typedef struct _abbot_CNICapArgs_PortMap {
    int32_t container_port;
    int32_t host_port;
    pb_callback_t protocol;
    pb_callback_t host_ip;
} abbot_CNICapArgs_PortMap;

typedef struct _abbot_ContainerNetworkDeleteRequest {
    pb_callback_t container_id;
    uint32_t pid;
} abbot_ContainerNetworkDeleteRequest;

typedef struct _abbot_ContainerNetworkEnsureRequest {
    pb_callback_t container_id;
    uint32_t pid;
    pb_callback_t cap_args;
    pb_callback_t cni_args;
} abbot_ContainerNetworkEnsureRequest;

typedef struct _abbot_ContainerNetworkQueryRequest {
    pb_callback_t container_id;
    uint32_t pid;
} abbot_ContainerNetworkQueryRequest;

typedef struct _abbot_ContainerNetworkRestoreRequest {
    pb_callback_t container_id;
    uint32_t pid;
} abbot_ContainerNetworkRestoreRequest;

typedef struct _abbot_ContainerNetworkStatusResponse {
    uint32_t pid;
    pb_callback_t interfaces;
} abbot_ContainerNetworkStatusResponse;

typedef struct _abbot_CNICapArgs {
    pb_size_t which_option;
    union {
        abbot_CNICapArgs_PortMap port_map_arg;
        abbot_CNICapArgs_Bandwidth bandwidth_arg;
        abbot_CNICapArgs_IPRange ip_range_arg;
        abbot_CNICapArgs_DNSConfig dns_config_arg;
        abbot_CNICapArgs_IPAddresses ip_addresses_arg;
        abbot_CNICapArgs_MacAddress mac_address_arg;
        abbot_CNICapArgs_InfinibandGUID infiniband_guid_arg;
        abbot_CNICapArgs_DeviceID device_id_arg;
    } option;
} abbot_CNICapArgs;

typedef struct _abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry {
    pb_callback_t key;
    bool has_value;
    abbot_ContainerNetworkStatusResponse value;
} abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry;


#ifdef __cplusplus
extern "C" {
#endif

/* Initializer values for message structs */
#define abbot_CNICapArgs_init_default            {0, {abbot_CNICapArgs_PortMap_init_default}}
#define abbot_CNICapArgs_PortMap_init_default    {0, 0, {{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_CNICapArgs_Bandwidth_init_default  {0, 0, 0, 0}
#define abbot_CNICapArgs_IPRange_init_default    {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_CNICapArgs_DNSConfig_init_default  {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_CNICapArgs_IPAddresses_init_default {{{NULL}, NULL}}
#define abbot_CNICapArgs_MacAddress_init_default {{{NULL}, NULL}}
#define abbot_CNICapArgs_InfinibandGUID_init_default {{{NULL}, NULL}}
#define abbot_CNICapArgs_DeviceID_init_default   {{{NULL}, NULL}}
#define abbot_ContainerNetworkEnsureRequest_init_default {{{NULL}, NULL}, 0, {{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_ContainerNetworkEnsureRequest_CniArgsEntry_init_default {{{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_ContainerNetworkRestoreRequest_init_default {{{NULL}, NULL}, 0}
#define abbot_ContainerNetworkDeleteRequest_init_default {{{NULL}, NULL}, 0}
#define abbot_ContainerNetworkConfigEnsureRequest_init_default {{{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_ContainerNetworkConfigQueryRequest_init_default {0}
#define abbot_ContainerNetworkQueryRequest_init_default {{{NULL}, NULL}, 0}
#define abbot_ContainerNetworkConfigResponse_init_default {{{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_ContainerNetworkStatusResponse_init_default {0, {{NULL}, NULL}}
#define abbot_ContainerNetworkStatusListResponse_init_default {{{NULL}, NULL}}
#define abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_init_default {{{NULL}, NULL}, false, abbot_ContainerNetworkStatusResponse_init_default}
#define abbot_CNICapArgs_init_zero               {0, {abbot_CNICapArgs_PortMap_init_zero}}
#define abbot_CNICapArgs_PortMap_init_zero       {0, 0, {{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_CNICapArgs_Bandwidth_init_zero     {0, 0, 0, 0}
#define abbot_CNICapArgs_IPRange_init_zero       {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_CNICapArgs_DNSConfig_init_zero     {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_CNICapArgs_IPAddresses_init_zero   {{{NULL}, NULL}}
#define abbot_CNICapArgs_MacAddress_init_zero    {{{NULL}, NULL}}
#define abbot_CNICapArgs_InfinibandGUID_init_zero {{{NULL}, NULL}}
#define abbot_CNICapArgs_DeviceID_init_zero      {{{NULL}, NULL}}
#define abbot_ContainerNetworkEnsureRequest_init_zero {{{NULL}, NULL}, 0, {{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_ContainerNetworkEnsureRequest_CniArgsEntry_init_zero {{{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_ContainerNetworkRestoreRequest_init_zero {{{NULL}, NULL}, 0}
#define abbot_ContainerNetworkDeleteRequest_init_zero {{{NULL}, NULL}, 0}
#define abbot_ContainerNetworkConfigEnsureRequest_init_zero {{{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_ContainerNetworkConfigQueryRequest_init_zero {0}
#define abbot_ContainerNetworkQueryRequest_init_zero {{{NULL}, NULL}, 0}
#define abbot_ContainerNetworkConfigResponse_init_zero {{{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_ContainerNetworkStatusResponse_init_zero {0, {{NULL}, NULL}}
#define abbot_ContainerNetworkStatusListResponse_init_zero {{{NULL}, NULL}}
#define abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_init_zero {{{NULL}, NULL}, false, abbot_ContainerNetworkStatusResponse_init_zero}

/* Field tags (for use in manual encoding/decoding) */
#define abbot_CNICapArgs_DNSConfig_servers_tag   1
#define abbot_CNICapArgs_DNSConfig_searches_tag  2
#define abbot_CNICapArgs_DNSConfig_options_tag   3
#define abbot_CNICapArgs_DeviceID_device_id_tag  1
#define abbot_CNICapArgs_IPAddresses_ips_tag     1
#define abbot_CNICapArgs_IPRange_subnet_tag      1
#define abbot_CNICapArgs_IPRange_range_start_tag 2
#define abbot_CNICapArgs_IPRange_range_end_tag   3
#define abbot_CNICapArgs_IPRange_gateway_tag     4
#define abbot_CNICapArgs_InfinibandGUID_infiniband_guid_tag 1
#define abbot_CNICapArgs_MacAddress_mac_tag      1
#define abbot_ContainerNetworkConfigEnsureRequest_ipv4_subnet_tag 1
#define abbot_ContainerNetworkConfigEnsureRequest_ipv6_subnet_tag 2
#define abbot_ContainerNetworkConfigResponse_ipv4_subnet_tag 1
#define abbot_ContainerNetworkConfigResponse_ipv6_subnet_tag 2
#define abbot_ContainerNetworkEnsureRequest_CniArgsEntry_key_tag 1
#define abbot_ContainerNetworkEnsureRequest_CniArgsEntry_value_tag 2
#define abbot_ContainerNetworkStatusListResponse_container_networks_tag 1
#define abbot_CNICapArgs_Bandwidth_ingress_rate_tag 1
#define abbot_CNICapArgs_Bandwidth_ingress_burst_tag 2
#define abbot_CNICapArgs_Bandwidth_egress_rate_tag 3
#define abbot_CNICapArgs_Bandwidth_egress_burst_tag 4
#define abbot_CNICapArgs_PortMap_container_port_tag 1
#define abbot_CNICapArgs_PortMap_host_port_tag   2
#define abbot_CNICapArgs_PortMap_protocol_tag    3
#define abbot_CNICapArgs_PortMap_host_ip_tag     4
#define abbot_ContainerNetworkDeleteRequest_container_id_tag 1
#define abbot_ContainerNetworkDeleteRequest_pid_tag 2
#define abbot_ContainerNetworkEnsureRequest_container_id_tag 1
#define abbot_ContainerNetworkEnsureRequest_pid_tag 2
#define abbot_ContainerNetworkEnsureRequest_cap_args_tag 3
#define abbot_ContainerNetworkEnsureRequest_cni_args_tag 4
#define abbot_ContainerNetworkQueryRequest_container_id_tag 1
#define abbot_ContainerNetworkQueryRequest_pid_tag 2
#define abbot_ContainerNetworkRestoreRequest_container_id_tag 1
#define abbot_ContainerNetworkRestoreRequest_pid_tag 2
#define abbot_ContainerNetworkStatusResponse_pid_tag 1
#define abbot_ContainerNetworkStatusResponse_interfaces_tag 2
#define abbot_CNICapArgs_port_map_arg_tag        1
#define abbot_CNICapArgs_bandwidth_arg_tag       2
#define abbot_CNICapArgs_ip_range_arg_tag        3
#define abbot_CNICapArgs_dns_config_arg_tag      4
#define abbot_CNICapArgs_ip_addresses_arg_tag    5
#define abbot_CNICapArgs_mac_address_arg_tag     6
#define abbot_CNICapArgs_infiniband_guid_arg_tag 7
#define abbot_CNICapArgs_device_id_arg_tag       8
#define abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_key_tag 1
#define abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_value_tag 2

/* Struct field encoding specification for nanopb */
#define abbot_CNICapArgs_FIELDLIST(X, a) \
X(a, STATIC,   ONEOF,    MESSAGE,  (option,port_map_arg,option.port_map_arg),   1) \
X(a, STATIC,   ONEOF,    MESSAGE,  (option,bandwidth_arg,option.bandwidth_arg),   2) \
X(a, STATIC,   ONEOF,    MESSAGE,  (option,ip_range_arg,option.ip_range_arg),   3) \
X(a, STATIC,   ONEOF,    MESSAGE,  (option,dns_config_arg,option.dns_config_arg),   4) \
X(a, STATIC,   ONEOF,    MESSAGE,  (option,ip_addresses_arg,option.ip_addresses_arg),   5) \
X(a, STATIC,   ONEOF,    MESSAGE,  (option,mac_address_arg,option.mac_address_arg),   6) \
X(a, STATIC,   ONEOF,    MESSAGE,  (option,infiniband_guid_arg,option.infiniband_guid_arg),   7) \
X(a, STATIC,   ONEOF,    MESSAGE,  (option,device_id_arg,option.device_id_arg),   8)
#define abbot_CNICapArgs_CALLBACK NULL
#define abbot_CNICapArgs_DEFAULT NULL
#define abbot_CNICapArgs_option_port_map_arg_MSGTYPE abbot_CNICapArgs_PortMap
#define abbot_CNICapArgs_option_bandwidth_arg_MSGTYPE abbot_CNICapArgs_Bandwidth
#define abbot_CNICapArgs_option_ip_range_arg_MSGTYPE abbot_CNICapArgs_IPRange
#define abbot_CNICapArgs_option_dns_config_arg_MSGTYPE abbot_CNICapArgs_DNSConfig
#define abbot_CNICapArgs_option_ip_addresses_arg_MSGTYPE abbot_CNICapArgs_IPAddresses
#define abbot_CNICapArgs_option_mac_address_arg_MSGTYPE abbot_CNICapArgs_MacAddress
#define abbot_CNICapArgs_option_infiniband_guid_arg_MSGTYPE abbot_CNICapArgs_InfinibandGUID
#define abbot_CNICapArgs_option_device_id_arg_MSGTYPE abbot_CNICapArgs_DeviceID

#define abbot_CNICapArgs_PortMap_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, INT32,    container_port,    1) \
X(a, STATIC,   SINGULAR, INT32,    host_port,         2) \
X(a, CALLBACK, SINGULAR, STRING,   protocol,          3) \
X(a, CALLBACK, SINGULAR, STRING,   host_ip,           4)
#define abbot_CNICapArgs_PortMap_CALLBACK pb_default_field_callback
#define abbot_CNICapArgs_PortMap_DEFAULT NULL

#define abbot_CNICapArgs_Bandwidth_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, INT32,    ingress_rate,      1) \
X(a, STATIC,   SINGULAR, INT32,    ingress_burst,     2) \
X(a, STATIC,   SINGULAR, INT32,    egress_rate,       3) \
X(a, STATIC,   SINGULAR, INT32,    egress_burst,      4)
#define abbot_CNICapArgs_Bandwidth_CALLBACK NULL
#define abbot_CNICapArgs_Bandwidth_DEFAULT NULL

#define abbot_CNICapArgs_IPRange_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   subnet,            1) \
X(a, CALLBACK, SINGULAR, STRING,   range_start,       2) \
X(a, CALLBACK, SINGULAR, STRING,   range_end,         3) \
X(a, CALLBACK, SINGULAR, STRING,   gateway,           4)
#define abbot_CNICapArgs_IPRange_CALLBACK pb_default_field_callback
#define abbot_CNICapArgs_IPRange_DEFAULT NULL

#define abbot_CNICapArgs_DNSConfig_FIELDLIST(X, a) \
X(a, CALLBACK, REPEATED, STRING,   servers,           1) \
X(a, CALLBACK, REPEATED, STRING,   searches,          2) \
X(a, CALLBACK, REPEATED, STRING,   options,           3)
#define abbot_CNICapArgs_DNSConfig_CALLBACK pb_default_field_callback
#define abbot_CNICapArgs_DNSConfig_DEFAULT NULL

#define abbot_CNICapArgs_IPAddresses_FIELDLIST(X, a) \
X(a, CALLBACK, REPEATED, STRING,   ips,               1)
#define abbot_CNICapArgs_IPAddresses_CALLBACK pb_default_field_callback
#define abbot_CNICapArgs_IPAddresses_DEFAULT NULL

#define abbot_CNICapArgs_MacAddress_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   mac,               1)
#define abbot_CNICapArgs_MacAddress_CALLBACK pb_default_field_callback
#define abbot_CNICapArgs_MacAddress_DEFAULT NULL

#define abbot_CNICapArgs_InfinibandGUID_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   infiniband_guid,   1)
#define abbot_CNICapArgs_InfinibandGUID_CALLBACK pb_default_field_callback
#define abbot_CNICapArgs_InfinibandGUID_DEFAULT NULL

#define abbot_CNICapArgs_DeviceID_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   device_id,         1)
#define abbot_CNICapArgs_DeviceID_CALLBACK pb_default_field_callback
#define abbot_CNICapArgs_DeviceID_DEFAULT NULL

#define abbot_ContainerNetworkEnsureRequest_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   container_id,      1) \
X(a, STATIC,   SINGULAR, UINT32,   pid,               2) \
X(a, CALLBACK, REPEATED, MESSAGE,  cap_args,          3) \
X(a, CALLBACK, REPEATED, MESSAGE,  cni_args,          4)
#define abbot_ContainerNetworkEnsureRequest_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkEnsureRequest_DEFAULT NULL
#define abbot_ContainerNetworkEnsureRequest_cap_args_MSGTYPE abbot_CNICapArgs
#define abbot_ContainerNetworkEnsureRequest_cni_args_MSGTYPE abbot_ContainerNetworkEnsureRequest_CniArgsEntry

#define abbot_ContainerNetworkEnsureRequest_CniArgsEntry_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   key,               1) \
X(a, CALLBACK, SINGULAR, STRING,   value,             2)
#define abbot_ContainerNetworkEnsureRequest_CniArgsEntry_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkEnsureRequest_CniArgsEntry_DEFAULT NULL

#define abbot_ContainerNetworkRestoreRequest_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   container_id,      1) \
X(a, STATIC,   SINGULAR, UINT32,   pid,               2)
#define abbot_ContainerNetworkRestoreRequest_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkRestoreRequest_DEFAULT NULL

#define abbot_ContainerNetworkDeleteRequest_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   container_id,      1) \
X(a, STATIC,   SINGULAR, UINT32,   pid,               2)
#define abbot_ContainerNetworkDeleteRequest_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkDeleteRequest_DEFAULT NULL

#define abbot_ContainerNetworkConfigEnsureRequest_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   ipv4_subnet,       1) \
X(a, CALLBACK, SINGULAR, STRING,   ipv6_subnet,       2)
#define abbot_ContainerNetworkConfigEnsureRequest_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkConfigEnsureRequest_DEFAULT NULL

#define abbot_ContainerNetworkConfigQueryRequest_FIELDLIST(X, a) \

#define abbot_ContainerNetworkConfigQueryRequest_CALLBACK NULL
#define abbot_ContainerNetworkConfigQueryRequest_DEFAULT NULL

#define abbot_ContainerNetworkQueryRequest_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   container_id,      1) \
X(a, STATIC,   SINGULAR, UINT32,   pid,               2)
#define abbot_ContainerNetworkQueryRequest_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkQueryRequest_DEFAULT NULL

#define abbot_ContainerNetworkConfigResponse_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   ipv4_subnet,       1) \
X(a, CALLBACK, SINGULAR, STRING,   ipv6_subnet,       2)
#define abbot_ContainerNetworkConfigResponse_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkConfigResponse_DEFAULT NULL

#define abbot_ContainerNetworkStatusResponse_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UINT32,   pid,               1) \
X(a, CALLBACK, REPEATED, MESSAGE,  interfaces,        2)
#define abbot_ContainerNetworkStatusResponse_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkStatusResponse_DEFAULT NULL
#define abbot_ContainerNetworkStatusResponse_interfaces_MSGTYPE abbot_NetworkInterface

#define abbot_ContainerNetworkStatusListResponse_FIELDLIST(X, a) \
X(a, CALLBACK, REPEATED, MESSAGE,  container_networks,   1)
#define abbot_ContainerNetworkStatusListResponse_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkStatusListResponse_DEFAULT NULL
#define abbot_ContainerNetworkStatusListResponse_container_networks_MSGTYPE abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry

#define abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   key,               1) \
X(a, STATIC,   OPTIONAL, MESSAGE,  value,             2)
#define abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_CALLBACK pb_default_field_callback
#define abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_DEFAULT NULL
#define abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_value_MSGTYPE abbot_ContainerNetworkStatusResponse

extern const pb_msgdesc_t abbot_CNICapArgs_msg;
extern const pb_msgdesc_t abbot_CNICapArgs_PortMap_msg;
extern const pb_msgdesc_t abbot_CNICapArgs_Bandwidth_msg;
extern const pb_msgdesc_t abbot_CNICapArgs_IPRange_msg;
extern const pb_msgdesc_t abbot_CNICapArgs_DNSConfig_msg;
extern const pb_msgdesc_t abbot_CNICapArgs_IPAddresses_msg;
extern const pb_msgdesc_t abbot_CNICapArgs_MacAddress_msg;
extern const pb_msgdesc_t abbot_CNICapArgs_InfinibandGUID_msg;
extern const pb_msgdesc_t abbot_CNICapArgs_DeviceID_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkEnsureRequest_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkEnsureRequest_CniArgsEntry_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkRestoreRequest_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkDeleteRequest_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkConfigEnsureRequest_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkConfigQueryRequest_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkQueryRequest_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkConfigResponse_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkStatusResponse_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkStatusListResponse_msg;
extern const pb_msgdesc_t abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define abbot_CNICapArgs_fields &abbot_CNICapArgs_msg
#define abbot_CNICapArgs_PortMap_fields &abbot_CNICapArgs_PortMap_msg
#define abbot_CNICapArgs_Bandwidth_fields &abbot_CNICapArgs_Bandwidth_msg
#define abbot_CNICapArgs_IPRange_fields &abbot_CNICapArgs_IPRange_msg
#define abbot_CNICapArgs_DNSConfig_fields &abbot_CNICapArgs_DNSConfig_msg
#define abbot_CNICapArgs_IPAddresses_fields &abbot_CNICapArgs_IPAddresses_msg
#define abbot_CNICapArgs_MacAddress_fields &abbot_CNICapArgs_MacAddress_msg
#define abbot_CNICapArgs_InfinibandGUID_fields &abbot_CNICapArgs_InfinibandGUID_msg
#define abbot_CNICapArgs_DeviceID_fields &abbot_CNICapArgs_DeviceID_msg
#define abbot_ContainerNetworkEnsureRequest_fields &abbot_ContainerNetworkEnsureRequest_msg
#define abbot_ContainerNetworkEnsureRequest_CniArgsEntry_fields &abbot_ContainerNetworkEnsureRequest_CniArgsEntry_msg
#define abbot_ContainerNetworkRestoreRequest_fields &abbot_ContainerNetworkRestoreRequest_msg
#define abbot_ContainerNetworkDeleteRequest_fields &abbot_ContainerNetworkDeleteRequest_msg
#define abbot_ContainerNetworkConfigEnsureRequest_fields &abbot_ContainerNetworkConfigEnsureRequest_msg
#define abbot_ContainerNetworkConfigQueryRequest_fields &abbot_ContainerNetworkConfigQueryRequest_msg
#define abbot_ContainerNetworkQueryRequest_fields &abbot_ContainerNetworkQueryRequest_msg
#define abbot_ContainerNetworkConfigResponse_fields &abbot_ContainerNetworkConfigResponse_msg
#define abbot_ContainerNetworkStatusResponse_fields &abbot_ContainerNetworkStatusResponse_msg
#define abbot_ContainerNetworkStatusListResponse_fields &abbot_ContainerNetworkStatusListResponse_msg
#define abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_fields &abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_msg

/* Maximum encoded size of messages (where known) */
/* abbot_CNICapArgs_size depends on runtime parameters */
/* abbot_CNICapArgs_PortMap_size depends on runtime parameters */
#define abbot_CNICapArgs_Bandwidth_size          44
/* abbot_CNICapArgs_IPRange_size depends on runtime parameters */
/* abbot_CNICapArgs_DNSConfig_size depends on runtime parameters */
/* abbot_CNICapArgs_IPAddresses_size depends on runtime parameters */
/* abbot_CNICapArgs_MacAddress_size depends on runtime parameters */
/* abbot_CNICapArgs_InfinibandGUID_size depends on runtime parameters */
/* abbot_CNICapArgs_DeviceID_size depends on runtime parameters */
/* abbot_ContainerNetworkEnsureRequest_size depends on runtime parameters */
/* abbot_ContainerNetworkEnsureRequest_CniArgsEntry_size depends on runtime parameters */
/* abbot_ContainerNetworkRestoreRequest_size depends on runtime parameters */
/* abbot_ContainerNetworkDeleteRequest_size depends on runtime parameters */
/* abbot_ContainerNetworkConfigEnsureRequest_size depends on runtime parameters */
#define abbot_ContainerNetworkConfigQueryRequest_size 0
/* abbot_ContainerNetworkQueryRequest_size depends on runtime parameters */
/* abbot_ContainerNetworkConfigResponse_size depends on runtime parameters */
/* abbot_ContainerNetworkStatusResponse_size depends on runtime parameters */
/* abbot_ContainerNetworkStatusListResponse_size depends on runtime parameters */
/* abbot_ContainerNetworkStatusListResponse_ContainerNetworksEntry_size depends on runtime parameters */

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif
