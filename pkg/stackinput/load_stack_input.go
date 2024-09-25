package stackinput

import (
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	"github.com/pkg/errors"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/stackinput/fieldsextractor"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	PulumiConfigKey = "planton-cloud:stack-input"
	FilePathEnvVar  = "STACK_INPUT_FILE_PATH"
)

func LoadStackInput(ctx *pulumi.Context, stackInput proto.Message) error {
	stackInputString, ok := ctx.GetConfig(PulumiConfigKey)
	var jsonBytes []byte
	var err error

	if !ok {
		stackInputFilePath := os.Getenv("STACK_INPUT_FILE_PATH")
		if stackInputFilePath == "" {
			return errors.Errorf("stack-input not found in pulumi config %s or in %s environment variable",
				PulumiConfigKey, FilePathEnvVar)
		}
		inputFileBytes, err := os.ReadFile(stackInputFilePath)
		if err != nil {
			return errors.Wrap(err, "failed to read input file")
		}
		jsonBytes, err = yaml.YAMLToJSON(inputFileBytes)
		if err != nil {
			return errors.Wrap(err, "failed to load yaml to json")
		}
	} else {
		jsonBytes, err = yaml.YAMLToJSON([]byte(stackInputString))
		if err != nil {
			return errors.Wrap(err, "failed to load yaml to json")
		}
	}

	if err := protojson.Unmarshal(jsonBytes, stackInput); err != nil {
		return errors.Wrap(err, "failed to load json into proto message")
	}

	targetSpec, err := fieldsextractor.ExtractApiResourceSpecField(stackInput)
	if err != nil {
		return errors.Wrap(err, "failed to extract api resource spec field")
	}

	v, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages((*targetSpec).Interface()),
	)
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	if err = v.Validate((*targetSpec).Interface()); err != nil {
		return errors.Errorf("%s", err)
	}
	return nil
}
