/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.2 */

#ifndef PB_ABBOT_HOST_PB_H_INCLUDED
#define PB_ABBOT_HOST_PB_H_INCLUDED
#include <pb.h>
#include "meta.pb.h"
#include "driver_unknown.pb.h"
#include "driver_bridge.pb.h"
#include "driver_wireguard.pb.h"

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Struct definitions */
typedef struct _abbot_HostNetworkConfigEnsureRequest {
    pb_callback_t expected;
} abbot_HostNetworkConfigEnsureRequest;

typedef struct _abbot_HostNetworkStatusResponse {
    pb_callback_t actual;
} abbot_HostNetworkStatusResponse;

typedef struct _abbot_HostNetworkInterface {
    bool has_metadata;
    abbot_NetworkInterface metadata;
    pb_size_t which_config;
    union {
        abbot_DriverUnknown unknown;
        abbot_DriverBridge bridge;
        abbot_DriverWireguard wireguard;
    } config;
} abbot_HostNetworkInterface;


/* Initializer values for message structs */
#define abbot_HostNetworkInterface_init_default  {false, abbot_NetworkInterface_init_default, 0, {abbot_DriverUnknown_init_default}}
#define abbot_HostNetworkConfigEnsureRequest_init_default {{{NULL}, NULL}}
#define abbot_HostNetworkStatusResponse_init_default {{{NULL}, NULL}}
#define abbot_HostNetworkInterface_init_zero     {false, abbot_NetworkInterface_init_zero, 0, {abbot_DriverUnknown_init_zero}}
#define abbot_HostNetworkConfigEnsureRequest_init_zero {{{NULL}, NULL}}
#define abbot_HostNetworkStatusResponse_init_zero {{{NULL}, NULL}}

/* Field tags (for use in manual encoding/decoding) */
#define abbot_HostNetworkConfigEnsureRequest_expected_tag 1
#define abbot_HostNetworkStatusResponse_actual_tag 1
#define abbot_HostNetworkInterface_metadata_tag  1
#define abbot_HostNetworkInterface_unknown_tag   10
#define abbot_HostNetworkInterface_bridge_tag    11
#define abbot_HostNetworkInterface_wireguard_tag 12

/* Struct field encoding specification for nanopb */
#define abbot_HostNetworkInterface_FIELDLIST(X, a) \
X(a, STATIC,   OPTIONAL, MESSAGE,  metadata,          1) \
X(a, STATIC,   ONEOF,    MESSAGE,  (config,unknown,config.unknown),  10) \
X(a, STATIC,   ONEOF,    MESSAGE,  (config,bridge,config.bridge),  11) \
X(a, STATIC,   ONEOF,    MESSAGE,  (config,wireguard,config.wireguard),  12)
#define abbot_HostNetworkInterface_CALLBACK NULL
#define abbot_HostNetworkInterface_DEFAULT NULL
#define abbot_HostNetworkInterface_metadata_MSGTYPE abbot_NetworkInterface
#define abbot_HostNetworkInterface_config_unknown_MSGTYPE abbot_DriverUnknown
#define abbot_HostNetworkInterface_config_bridge_MSGTYPE abbot_DriverBridge
#define abbot_HostNetworkInterface_config_wireguard_MSGTYPE abbot_DriverWireguard

#define abbot_HostNetworkConfigEnsureRequest_FIELDLIST(X, a) \
X(a, CALLBACK, REPEATED, MESSAGE,  expected,          1)
#define abbot_HostNetworkConfigEnsureRequest_CALLBACK pb_default_field_callback
#define abbot_HostNetworkConfigEnsureRequest_DEFAULT NULL
#define abbot_HostNetworkConfigEnsureRequest_expected_MSGTYPE abbot_HostNetworkInterface

#define abbot_HostNetworkStatusResponse_FIELDLIST(X, a) \
X(a, CALLBACK, REPEATED, MESSAGE,  actual,            1)
#define abbot_HostNetworkStatusResponse_CALLBACK pb_default_field_callback
#define abbot_HostNetworkStatusResponse_DEFAULT NULL
#define abbot_HostNetworkStatusResponse_actual_MSGTYPE abbot_HostNetworkInterface

extern const pb_msgdesc_t abbot_HostNetworkInterface_msg;
extern const pb_msgdesc_t abbot_HostNetworkConfigEnsureRequest_msg;
extern const pb_msgdesc_t abbot_HostNetworkStatusResponse_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define abbot_HostNetworkInterface_fields &abbot_HostNetworkInterface_msg
#define abbot_HostNetworkConfigEnsureRequest_fields &abbot_HostNetworkConfigEnsureRequest_msg
#define abbot_HostNetworkStatusResponse_fields &abbot_HostNetworkStatusResponse_msg

/* Maximum encoded size of messages (where known) */
union abbot_HostNetworkInterface_config_size_union {char f11[(6 + abbot_DriverBridge_size)]; char f12[(6 + abbot_DriverWireguard_size)]; char f0[2];};
#define abbot_HostNetworkInterface_size          (6 + abbot_NetworkInterface_size + sizeof(abbot_HostNetworkInterface_config_size_union))
/* abbot_HostNetworkConfigEnsureRequest_size depends on runtime parameters */
/* abbot_HostNetworkStatusResponse_size depends on runtime parameters */

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif