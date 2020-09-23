/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.2 */

#ifndef PB_ABBOT_PROTO_PB_H_INCLUDED
#define PB_ABBOT_PROTO_PB_H_INCLUDED
#include <pb.h>

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _abbot_RequestType {
    abbot_RequestType__INVALID_REQUEST = 0,
    abbot_RequestType_REQ_ENSURE_CTR_NETWORK = 1,
    abbot_RequestType_REQ_DELETE_CTR_NETWORK = 2,
    abbot_RequestType_REQ_QUERY_CTR_NETWORK = 3,
    abbot_RequestType_REQ_UPDATE_CTR_NETWORK_CONFIG = 4
} abbot_RequestType;

typedef enum _abbot_ResponseType {
    abbot_ResponseType__INVALID_RESPONSE = 0,
    abbot_ResponseType_RESP_DONE = 1,
    abbot_ResponseType_RESP_CTR_NETWORK_STATUS = 2
} abbot_ResponseType;

/* Struct definitions */
typedef struct _abbot_Request {
    abbot_RequestType kind;
    pb_callback_t body;
} abbot_Request;

typedef struct _abbot_Response {
    abbot_ResponseType kind;
    pb_callback_t body;
} abbot_Response;


/* Helper constants for enums */
#define _abbot_RequestType_MIN abbot_RequestType__INVALID_REQUEST
#define _abbot_RequestType_MAX abbot_RequestType_REQ_UPDATE_CTR_NETWORK_CONFIG
#define _abbot_RequestType_ARRAYSIZE ((abbot_RequestType)(abbot_RequestType_REQ_UPDATE_CTR_NETWORK_CONFIG+1))

#define _abbot_ResponseType_MIN abbot_ResponseType__INVALID_RESPONSE
#define _abbot_ResponseType_MAX abbot_ResponseType_RESP_CTR_NETWORK_STATUS
#define _abbot_ResponseType_ARRAYSIZE ((abbot_ResponseType)(abbot_ResponseType_RESP_CTR_NETWORK_STATUS+1))


/* Initializer values for message structs */
#define abbot_Request_init_default               {_abbot_RequestType_MIN, {{NULL}, NULL}}
#define abbot_Response_init_default              {_abbot_ResponseType_MIN, {{NULL}, NULL}}
#define abbot_Request_init_zero                  {_abbot_RequestType_MIN, {{NULL}, NULL}}
#define abbot_Response_init_zero                 {_abbot_ResponseType_MIN, {{NULL}, NULL}}

/* Field tags (for use in manual encoding/decoding) */
#define abbot_Request_kind_tag                   1
#define abbot_Request_body_tag                   2
#define abbot_Response_kind_tag                  1
#define abbot_Response_body_tag                  2

/* Struct field encoding specification for nanopb */
#define abbot_Request_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UENUM,    kind,              1) \
X(a, CALLBACK, SINGULAR, BYTES,    body,              2)
#define abbot_Request_CALLBACK pb_default_field_callback
#define abbot_Request_DEFAULT NULL

#define abbot_Response_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UENUM,    kind,              1) \
X(a, CALLBACK, SINGULAR, BYTES,    body,              2)
#define abbot_Response_CALLBACK pb_default_field_callback
#define abbot_Response_DEFAULT NULL

extern const pb_msgdesc_t abbot_Request_msg;
extern const pb_msgdesc_t abbot_Response_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define abbot_Request_fields &abbot_Request_msg
#define abbot_Response_fields &abbot_Response_msg

/* Maximum encoded size of messages (where known) */
/* abbot_Request_size depends on runtime parameters */
/* abbot_Response_size depends on runtime parameters */

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif
