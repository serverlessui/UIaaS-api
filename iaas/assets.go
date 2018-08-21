package iaas

const (
	route53 = "route53.yml"
	s3site  = "s3site.yml"
)

// GetRouterTemplate() (string, error)
// GetSourceTemplate() (string, error)
// GetLambdaSource() (string, error)

//AWSTemplate is the implementation of IaasAsset interface
type AWSTemplate struct {
}

//Infrastructure is the interface to retrieve Iaas Assets
type Infrastructure interface {
	GetRoute53() (*string, error)
	GetS3Site() (*string, error)
}

//GetRoute53 method to retrieve Route53 Template
func (template AWSTemplate) GetRoute53() (*string, error) {
	return getAsset(route53)
}

//GetS3Site method to retrieve S3Site Template
func (template AWSTemplate) GetS3Site() (*string, error) {
	return getAsset(s3site)
}

//getAsset retrives []byte of file and converts it to a string
//todo- this can probably be factored out again
func getAsset(template string) (*string, error) {
	//todo- think through if it's worth factoring out into an interface for mocking?
	data, err := Asset(template)
	if err != nil {
		println("Error reading file ", template)
		return nil, err
	}
	s := string(data[:])
	return &s, nil
}
