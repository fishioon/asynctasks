// Code generated by protoc-gen-go. DO NOT EDIT.
// source: target.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The HTTP method used to execute the task.
type HttpMethod int32

const (
	// HTTP method unspecified
	HttpMethod_HTTP_METHOD_UNSPECIFIED HttpMethod = 0
	// HTTP POST
	HttpMethod_POST HttpMethod = 1
	// HTTP GET
	HttpMethod_GET HttpMethod = 2
	// HTTP HEAD
	HttpMethod_HEAD HttpMethod = 3
	// HTTP PUT
	HttpMethod_PUT HttpMethod = 4
	// HTTP DELETE
	HttpMethod_DELETE HttpMethod = 5
	// HTTP PATCH
	HttpMethod_PATCH HttpMethod = 6
	// HTTP OPTIONS
	HttpMethod_OPTIONS HttpMethod = 7
)

var HttpMethod_name = map[int32]string{
	0: "HTTP_METHOD_UNSPECIFIED",
	1: "POST",
	2: "GET",
	3: "HEAD",
	4: "PUT",
	5: "DELETE",
	6: "PATCH",
	7: "OPTIONS",
}

var HttpMethod_value = map[string]int32{
	"HTTP_METHOD_UNSPECIFIED": 0,
	"POST":                    1,
	"GET":                     2,
	"HEAD":                    3,
	"PUT":                     4,
	"DELETE":                  5,
	"PATCH":                   6,
	"OPTIONS":                 7,
}

func (x HttpMethod) String() string {
	return proto.EnumName(HttpMethod_name, int32(x))
}

func (HttpMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_468528a86129e532, []int{0}
}

// HTTP request.
//
// The task will be pushed to the worker as an HTTP request. If the worker
// or the redirected worker acknowledges the task by returning a successful HTTP
// response code ([`200` - `299`]), the task will be removed from the queue. If
// any other HTTP response code is returned or no response is received, the
// task will be retried according to the following:
//
// * User-specified throttling: [retry configuration][google.cloud.tasks.v2beta3.Queue.retry_config],
//   [rate limits][google.cloud.tasks.v2beta3.Queue.rate_limits], and the [queue's state][google.cloud.tasks.v2beta3.Queue.state].
//
// * System throttling: To prevent the worker from overloading, Cloud Tasks may
//   temporarily reduce the queue's effective rate. User-specified settings
//   will not be changed.
//
//  System throttling happens because:
//
//   * Cloud Tasks backs off on all errors. Normally the backoff specified in
//     [rate limits][google.cloud.tasks.v2beta3.Queue.rate_limits] will be used. But if the worker returns
//     `429` (Too Many Requests), `503` (Service Unavailable), or the rate of
//     errors is high, Cloud Tasks will use a higher backoff rate. The retry
//     specified in the `Retry-After` HTTP response header is considered.
//
//   * To prevent traffic spikes and to smooth sudden increases in traffic,
//     dispatches ramp up slowly when the queue is newly created or idle and
//     if large numbers of tasks suddenly become available to dispatch (due to
//     spikes in create task rates, the queue being unpaused, or many tasks
//     that are scheduled at the same time).
type HttpRequest struct {
	// Required. The full url path that the request will be sent to.
	//
	// This string must begin with either "http://" or "https://". Some examples
	// are: `http://acme.com` and `https://acme.com/sales:8080`. Cloud Tasks will
	// encode some characters for safety and compatibility. The maximum allowed
	// URL length is 2083 characters after encoding.
	//
	// The `Location` header response from a redirect response [`300` - `399`]
	// may be followed. The redirect is not counted as a separate attempt.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// The HTTP method to use for the request. The default is POST.
	HttpMethod HttpMethod `protobuf:"varint,2,opt,name=http_method,json=httpMethod,proto3,enum=api.HttpMethod" json:"http_method,omitempty"`
	// HTTP request headers.
	//
	// This map contains the header field names and values.
	// Headers can be set when the
	// [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	//
	// These headers represent a subset of the headers that will accompany the
	// task's HTTP request. Some HTTP request headers will be ignored or replaced.
	//
	// A partial list of headers that will be ignored or replaced is:
	//
	// * Host: This will be computed by Cloud Tasks and derived from
	//   [HttpRequest.url][google.cloud.tasks.v2beta3.HttpRequest.url].
	// * Content-Length: This will be computed by Cloud Tasks.
	// * User-Agent: This will be set to `"Google-Cloud-Tasks"`.
	// * X-Google-*: Google use only.
	// * X-AppEngine-*: Google use only.
	//
	// `Content-Type` won't be set by Cloud Tasks. You can explicitly set
	// `Content-Type` to a media type when the
	//  [task is created][google.cloud.tasks.v2beta3.CloudTasks.CreateTask].
	//  For example, `Content-Type` can be set to `"application/octet-stream"` or
	//  `"application/json"`.
	//
	// Headers which can have multiple values (according to RFC2616) can be
	// specified using comma-separated values.
	//
	// The size of the headers must be less than 80KB.
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// HTTP request body.
	//
	// A request body is allowed only if the
	// [HTTP method][google.cloud.tasks.v2beta3.HttpRequest.http_method] is POST, PUT, or PATCH. It is an
	// error to set body on a task with an incompatible [HttpMethod][google.cloud.tasks.v2beta3.HttpMethod].
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	// The mode for generating an `Authorization` header for HTTP requests.
	//
	// If specified, all `Authorization` headers in the [HttpRequest.headers][google.cloud.tasks.v2beta3.HttpRequest.headers]
	// field will be overridden.
	//
	// Types that are valid to be assigned to AuthorizationHeader:
	//	*HttpRequest_OauthToken
	//	*HttpRequest_OidcToken
	AuthorizationHeader  isHttpRequest_AuthorizationHeader `protobuf_oneof:"authorization_header"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HttpRequest) Reset()         { *m = HttpRequest{} }
func (m *HttpRequest) String() string { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()    {}
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_468528a86129e532, []int{0}
}

func (m *HttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRequest.Unmarshal(m, b)
}
func (m *HttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRequest.Marshal(b, m, deterministic)
}
func (m *HttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRequest.Merge(m, src)
}
func (m *HttpRequest) XXX_Size() int {
	return xxx_messageInfo_HttpRequest.Size(m)
}
func (m *HttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRequest proto.InternalMessageInfo

func (m *HttpRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HttpRequest) GetHttpMethod() HttpMethod {
	if m != nil {
		return m.HttpMethod
	}
	return HttpMethod_HTTP_METHOD_UNSPECIFIED
}

func (m *HttpRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type isHttpRequest_AuthorizationHeader interface {
	isHttpRequest_AuthorizationHeader()
}

type HttpRequest_OauthToken struct {
	OauthToken *OAuthToken `protobuf:"bytes,5,opt,name=oauth_token,json=oauthToken,proto3,oneof"`
}

type HttpRequest_OidcToken struct {
	OidcToken *OidcToken `protobuf:"bytes,6,opt,name=oidc_token,json=oidcToken,proto3,oneof"`
}

func (*HttpRequest_OauthToken) isHttpRequest_AuthorizationHeader() {}

func (*HttpRequest_OidcToken) isHttpRequest_AuthorizationHeader() {}

func (m *HttpRequest) GetAuthorizationHeader() isHttpRequest_AuthorizationHeader {
	if m != nil {
		return m.AuthorizationHeader
	}
	return nil
}

func (m *HttpRequest) GetOauthToken() *OAuthToken {
	if x, ok := m.GetAuthorizationHeader().(*HttpRequest_OauthToken); ok {
		return x.OauthToken
	}
	return nil
}

func (m *HttpRequest) GetOidcToken() *OidcToken {
	if x, ok := m.GetAuthorizationHeader().(*HttpRequest_OidcToken); ok {
		return x.OidcToken
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpRequest_OauthToken)(nil),
		(*HttpRequest_OidcToken)(nil),
	}
}

// Contains information needed for generating an
// [OAuth token](https://developers.google.com/identity/protocols/OAuth2).
// This type of authorization should generally only be used when calling Google
// APIs hosted on *.googleapis.com.
type OAuthToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	// to be used for generating OAuth token.
	// The service account must be within the same project as the queue. The
	// caller must have iam.serviceAccounts.actAs permission for the service
	// account.
	ServiceAccountEmail string `protobuf:"bytes,1,opt,name=service_account_email,json=serviceAccountEmail,proto3" json:"service_account_email,omitempty"`
	// OAuth scope to be used for generating OAuth access token.
	// If not specified, "https://www.googleapis.com/auth/cloud-platform"
	// will be used.
	Scope                string   `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthToken) Reset()         { *m = OAuthToken{} }
func (m *OAuthToken) String() string { return proto.CompactTextString(m) }
func (*OAuthToken) ProtoMessage()    {}
func (*OAuthToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_468528a86129e532, []int{1}
}

func (m *OAuthToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthToken.Unmarshal(m, b)
}
func (m *OAuthToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthToken.Marshal(b, m, deterministic)
}
func (m *OAuthToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthToken.Merge(m, src)
}
func (m *OAuthToken) XXX_Size() int {
	return xxx_messageInfo_OAuthToken.Size(m)
}
func (m *OAuthToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthToken.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthToken proto.InternalMessageInfo

func (m *OAuthToken) GetServiceAccountEmail() string {
	if m != nil {
		return m.ServiceAccountEmail
	}
	return ""
}

func (m *OAuthToken) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

// Contains information needed for generating an
// [OpenID Connect
// token](https://developers.google.com/identity/protocols/OpenIDConnect).
// This type of authorization can be used for many scenarios, including
// calling Cloud Run, or endpoints where you intend to validate the token
// yourself.
type OidcToken struct {
	// [Service account email](https://cloud.google.com/iam/docs/service-accounts)
	// to be used for generating OIDC token.
	// The service account must be within the same project as the queue. The
	// caller must have iam.serviceAccounts.actAs permission for the service
	// account.
	ServiceAccountEmail string `protobuf:"bytes,1,opt,name=service_account_email,json=serviceAccountEmail,proto3" json:"service_account_email,omitempty"`
	// Audience to be used when generating OIDC token. If not specified, the URI
	// specified in target will be used.
	Audience             string   `protobuf:"bytes,2,opt,name=audience,proto3" json:"audience,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OidcToken) Reset()         { *m = OidcToken{} }
func (m *OidcToken) String() string { return proto.CompactTextString(m) }
func (*OidcToken) ProtoMessage()    {}
func (*OidcToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_468528a86129e532, []int{2}
}

func (m *OidcToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OidcToken.Unmarshal(m, b)
}
func (m *OidcToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OidcToken.Marshal(b, m, deterministic)
}
func (m *OidcToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OidcToken.Merge(m, src)
}
func (m *OidcToken) XXX_Size() int {
	return xxx_messageInfo_OidcToken.Size(m)
}
func (m *OidcToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OidcToken.DiscardUnknown(m)
}

var xxx_messageInfo_OidcToken proto.InternalMessageInfo

func (m *OidcToken) GetServiceAccountEmail() string {
	if m != nil {
		return m.ServiceAccountEmail
	}
	return ""
}

func (m *OidcToken) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.HttpMethod", HttpMethod_name, HttpMethod_value)
	proto.RegisterType((*HttpRequest)(nil), "api.HttpRequest")
	proto.RegisterMapType((map[string]string)(nil), "api.HttpRequest.HeadersEntry")
	proto.RegisterType((*OAuthToken)(nil), "api.OAuthToken")
	proto.RegisterType((*OidcToken)(nil), "api.OidcToken")
}

func init() { proto.RegisterFile("target.proto", fileDescriptor_468528a86129e532) }

var fileDescriptor_468528a86129e532 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x6f, 0x9b, 0x40,
	0x10, 0x0d, 0xc6, 0x1f, 0xf1, 0x60, 0xa5, 0x68, 0x9b, 0xb6, 0x28, 0x55, 0x25, 0xe4, 0x13, 0xea,
	0x81, 0x56, 0xee, 0xa1, 0x55, 0x6e, 0x6e, 0xbc, 0x2d, 0x96, 0x9a, 0x80, 0xf0, 0xa6, 0x97, 0x1e,
	0xd0, 0x66, 0x59, 0x85, 0x95, 0x63, 0x96, 0xc2, 0x12, 0x89, 0xfe, 0xd1, 0xfe, 0x9d, 0x6a, 0xc1,
	0x90, 0x9e, 0x73, 0x9b, 0x37, 0x6f, 0xde, 0xbe, 0x99, 0x07, 0xb0, 0x50, 0xb4, 0xbc, 0xe7, 0xca,
	0x2f, 0x4a, 0xa9, 0x24, 0x32, 0x69, 0x21, 0x96, 0x7f, 0x47, 0x60, 0x05, 0x4a, 0x15, 0x31, 0xff,
	0x5d, 0xf3, 0x4a, 0x21, 0x1b, 0xcc, 0xba, 0x7c, 0x70, 0x0c, 0xd7, 0xf0, 0xe6, 0xb1, 0x2e, 0xd1,
	0x47, 0xb0, 0x32, 0xa5, 0x8a, 0xe4, 0xc0, 0x55, 0x26, 0x53, 0x67, 0xe4, 0x1a, 0xde, 0xd9, 0xea,
	0x85, 0x4f, 0x0b, 0xe1, 0x6b, 0xe1, 0x75, 0xdb, 0x8e, 0x21, 0x1b, 0x6a, 0xf4, 0x19, 0x66, 0x19,
	0xa7, 0x29, 0x2f, 0x2b, 0xc7, 0x74, 0x4d, 0xcf, 0x5a, 0xbd, 0x1b, 0xa6, 0x8f, 0x36, 0x7e, 0xd0,
	0xf1, 0x38, 0x57, 0x65, 0x13, 0xf7, 0xd3, 0x08, 0xc1, 0xf8, 0x4e, 0xa6, 0x8d, 0x33, 0x76, 0x0d,
	0x6f, 0x11, 0xb7, 0x35, 0x5a, 0x81, 0x25, 0x69, 0xad, 0xb2, 0x44, 0xc9, 0x3d, 0xcf, 0x9d, 0x89,
	0x6b, 0x78, 0xd6, 0xd1, 0x3e, 0x5c, 0xd7, 0x2a, 0x23, 0xba, 0x1d, 0x9c, 0xc4, 0xd0, 0x4e, 0xb5,
	0x08, 0x7d, 0x00, 0x90, 0x22, 0x65, 0x47, 0xc9, 0xb4, 0x95, 0x9c, 0x75, 0x12, 0x91, 0xb2, 0x5e,
	0x31, 0x97, 0x3d, 0xb8, 0xb8, 0x84, 0xc5, 0xff, 0x1b, 0xe9, 0x14, 0xf6, 0xbc, 0xe9, 0x53, 0xd8,
	0xf3, 0x06, 0x9d, 0xc3, 0xe4, 0x91, 0x3e, 0xd4, 0xbc, 0xbd, 0x7f, 0x1e, 0x77, 0xe0, 0x72, 0xf4,
	0xc5, 0xf8, 0xfa, 0x1a, 0xce, 0xb5, 0xb3, 0x2c, 0xc5, 0x1f, 0xaa, 0x84, 0xcc, 0x93, 0xee, 0x9a,
	0xe5, 0x4f, 0x80, 0xa7, 0x05, 0xd1, 0x0a, 0x5e, 0x55, 0xbc, 0x7c, 0x14, 0x8c, 0x27, 0x94, 0x31,
	0x59, 0xe7, 0x2a, 0xe1, 0x07, 0x2a, 0xfa, 0xa4, 0x5f, 0x1e, 0xc9, 0x75, 0xc7, 0x61, 0x4d, 0x69,
	0xcf, 0x8a, 0xc9, 0x62, 0xf0, 0x6c, 0xc1, 0xf2, 0x17, 0xcc, 0x87, 0x2b, 0x9e, 0xf5, 0xec, 0x05,
	0x9c, 0xd2, 0x3a, 0x15, 0x3c, 0x67, 0xfd, 0xcb, 0x03, 0x7e, 0x5f, 0x01, 0x3c, 0x7d, 0x54, 0xf4,
	0x16, 0xde, 0x04, 0x84, 0x44, 0xc9, 0x35, 0x26, 0x41, 0xb8, 0x49, 0x6e, 0x6f, 0x76, 0x11, 0xbe,
	0xda, 0x7e, 0xdb, 0xe2, 0x8d, 0x7d, 0x82, 0x4e, 0x61, 0x1c, 0x85, 0x3b, 0x62, 0x1b, 0x68, 0x06,
	0xe6, 0x77, 0x4c, 0xec, 0x91, 0x6e, 0x05, 0x78, 0xbd, 0xb1, 0x4d, 0xdd, 0x8a, 0x6e, 0x89, 0x3d,
	0x46, 0x00, 0xd3, 0x0d, 0xfe, 0x81, 0x09, 0xb6, 0x27, 0x68, 0x0e, 0x93, 0x68, 0x4d, 0xae, 0x02,
	0x7b, 0x8a, 0x2c, 0x98, 0x85, 0x11, 0xd9, 0x86, 0x37, 0x3b, 0x7b, 0xa6, 0x13, 0x64, 0xf2, 0xe0,
	0xcb, 0x9c, 0xdf, 0x4b, 0x9f, 0x56, 0x4d, 0xce, 0x14, 0xad, 0xf6, 0xd5, 0xdd, 0xb4, 0xfd, 0x4f,
	0x3f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xd2, 0x0d, 0x69, 0xb7, 0x02, 0x00, 0x00,
}
