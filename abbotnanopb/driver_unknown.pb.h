/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.4 */

#ifndef PB_ABBOT_DRIVER_UNKNOWN_PB_H_INCLUDED
#define PB_ABBOT_DRIVER_UNKNOWN_PB_H_INCLUDED
#include <pb.h>

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

/* Struct definitions */
typedef struct _abbot_DriverUnknown {
    char dummy_field;
} abbot_DriverUnknown;


#ifdef __cplusplus
extern "C" {
#endif

/* Initializer values for message structs */
#define abbot_DriverUnknown_init_default         {0}
#define abbot_DriverUnknown_init_zero            {0}

/* Field tags (for use in manual encoding/decoding) */

/* Struct field encoding specification for nanopb */
#define abbot_DriverUnknown_FIELDLIST(X, a) \

#define abbot_DriverUnknown_CALLBACK NULL
#define abbot_DriverUnknown_DEFAULT NULL

extern const pb_msgdesc_t abbot_DriverUnknown_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define abbot_DriverUnknown_fields &abbot_DriverUnknown_msg

/* Maximum encoded size of messages (where known) */
#define abbot_DriverUnknown_size                 0

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif
