package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	ssmClient := ssm.New(session.New())

	in := ssm.GetParametersByPathInput{Path: aws.String("/foo/bar")}
	out, err := ssmClient.GetParametersByPath(&in)
	if err != nil {
		log.Fatal(err)
	}

	for _, param := range out.Parameters {
		fmt.Println(aws.StringValue(param.Name))
	}
}
