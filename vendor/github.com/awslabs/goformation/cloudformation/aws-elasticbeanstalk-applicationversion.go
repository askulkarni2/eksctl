package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSElasticBeanstalkApplicationVersion AWS CloudFormation Resource (AWS::ElasticBeanstalk::ApplicationVersion)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html
type AWSElasticBeanstalkApplicationVersion struct {

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-applicationname
	ApplicationName *Value `json:"ApplicationName,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-description
	Description *Value `json:"Description,omitempty"`

	// SourceBundle AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
	SourceBundle *AWSElasticBeanstalkApplicationVersion_SourceBundle `json:"SourceBundle,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplicationVersion) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::ApplicationVersion"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSElasticBeanstalkApplicationVersion) MarshalJSON() ([]byte, error) {
	type Properties AWSElasticBeanstalkApplicationVersion
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSElasticBeanstalkApplicationVersion) UnmarshalJSON(b []byte) error {
	type Properties AWSElasticBeanstalkApplicationVersion
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSElasticBeanstalkApplicationVersion(*res.Properties)
	}

	return nil
}

// GetAllAWSElasticBeanstalkApplicationVersionResources retrieves all AWSElasticBeanstalkApplicationVersion items from an AWS CloudFormation template
func (t *Template) GetAllAWSElasticBeanstalkApplicationVersionResources() map[string]AWSElasticBeanstalkApplicationVersion {
	results := map[string]AWSElasticBeanstalkApplicationVersion{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSElasticBeanstalkApplicationVersion:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ElasticBeanstalk::ApplicationVersion" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSElasticBeanstalkApplicationVersion{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSElasticBeanstalkApplicationVersionWithName retrieves all AWSElasticBeanstalkApplicationVersion items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticBeanstalkApplicationVersionWithName(name string) (AWSElasticBeanstalkApplicationVersion, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSElasticBeanstalkApplicationVersion:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ElasticBeanstalk::ApplicationVersion" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSElasticBeanstalkApplicationVersion{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSElasticBeanstalkApplicationVersion{}, errors.New("resource not found")
}
