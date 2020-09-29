/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.2 */

#ifndef PB_ABBOT_META_PB_H_INCLUDED
#define PB_ABBOT_META_PB_H_INCLUDED
#include <pb.h>

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Struct definitions */
typedef struct _abbot_NetworkInterface {
    pb_callback_t name;
    pb_callback_t addresses;
    pb_callback_t hardware_address;
} abbot_NetworkInterface;


/* Initializer values for message structs */
#define abbot_NetworkInterface_init_default      {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define abbot_NetworkInterface_init_zero         {{{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}

/* Field tags (for use in manual encoding/decoding) */
#define abbot_NetworkInterface_name_tag          1
#define abbot_NetworkInterface_addresses_tag     2
#define abbot_NetworkInterface_hardware_address_tag 3

/* Struct field encoding specification for nanopb */
#define abbot_NetworkInterface_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   name,              1) \
X(a, CALLBACK, REPEATED, STRING,   addresses,         2) \
X(a, CALLBACK, SINGULAR, STRING,   hardware_address,   3)
#define abbot_NetworkInterface_CALLBACK pb_default_field_callback
#define abbot_NetworkInterface_DEFAULT NULL

extern const pb_msgdesc_t abbot_NetworkInterface_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define abbot_NetworkInterface_fields &abbot_NetworkInterface_msg

/* Maximum encoded size of messages (where known) */
/* abbot_NetworkInterface_size depends on runtime parameters */

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif