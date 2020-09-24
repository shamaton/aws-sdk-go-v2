// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
	"time"
)

const ServiceID = "KMS"
const ServiceAPIVersion = "2014-11-01"

// AWS Key Management Service AWS Key Management Service (AWS KMS) is an encryption
// and key management web service. This guide describes the AWS KMS operations that
// you can call programmatically. For general information about AWS KMS, see the
// AWS Key Management Service Developer Guide
// (https://docs.aws.amazon.com/kms/latest/developerguide/). AWS provides SDKs that
// consist of libraries and sample code for various programming languages and
// platforms (Java, Ruby, .Net, macOS, Android, etc.). The SDKs provide a
// convenient way to create programmatic access to AWS KMS and other AWS services.
// For example, the SDKs take care of tasks such as signing requests (see below),
// managing errors, and retrying requests automatically. For more information about
// the AWS SDKs, including how to download and install them, see Tools for Amazon
// Web Services (http://aws.amazon.com/tools/). We recommend that you use the AWS
// SDKs to make programmatic API calls to AWS KMS. Clients must support TLS
// (Transport Layer Security) 1.0. We recommend TLS 1.2. Clients must also support
// cipher suites with Perfect Forward Secrecy (PFS) such as Ephemeral
// Diffie-Hellman (DHE) or Elliptic Curve Ephemeral Diffie-Hellman (ECDHE). Most
// modern systems such as Java 7 and later support these modes. Signing Requests
// Requests must be signed by using an access key ID and a secret access key. We
// strongly recommend that you do not use your AWS account (root) access key ID and
// secret key for everyday work with AWS KMS. Instead, use the access key ID and
// secret access key for an IAM user. You can also use the AWS Security Token
// Service to generate temporary security credentials that you can use to sign
// requests. All AWS KMS operations require Signature Version 4
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
// Logging API Requests AWS KMS supports AWS CloudTrail, a service that logs AWS
// API calls and related events for your AWS account and delivers them to an Amazon
// S3 bucket that you specify. By using the information collected by CloudTrail,
// you can determine what requests were made to AWS KMS, who made the request, when
// it was made, and so on. To learn more about CloudTrail, including how to turn it
// on and find your log files, see the AWS CloudTrail User Guide
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/). Additional
// Resources For more information about credentials and request signing, see the
// following:
//
//     * AWS Security Credentials
// (https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html) -
// This topic provides general information about the types of credentials used for
// accessing AWS.
//
//     * Temporary Security Credentials
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html) -
// This section of the IAM User Guide describes how to create and use temporary
// security credentials.
//
//     * Signature Version 4 Signing Process
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html) - This
// set of topics walks you through the process of signing a request using an access
// key ID and a secret access key.
//
// Commonly Used API Operations Of the API
// operations discussed in this guide, the following will prove the most useful for
// most applications. You will likely perform operations other than these, such as
// creating keys and assigning policies, by using the console.
//
//     * Encrypt ()
//
//
// * Decrypt ()
//
//     * GenerateDataKey ()
//
//     * GenerateDataKeyWithoutPlaintext ()
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []APIOptionFunc

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

func (o Options) GetCredentials() aws.CredentialsProvider {
	return o.Credentials
}

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetHTTPSignerV4() HTTPSignerV4 {
	return o.HTTPSignerV4
}

func (o Options) GetRegion() string {
	return o.Region
}

func (o Options) GetRetryer() retry.Retryer {
	return o.Retryer
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]APIOptionFunc, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

type APIOptionFunc func(*middleware.Stack) error

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
	}
	resolveAWSEndpointResolver(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("kms")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}
