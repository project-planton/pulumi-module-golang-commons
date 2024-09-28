package stackinput

import (
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	"github.com/pkg/errors"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/stackinput/fieldsextractor"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"os"
	"sigs.k8s.io/yaml"
)

const (
	FilePathEnvVar    = "STACK_INPUT_FILE_PATH"
	YamlContentEnvVar = "STACK_INPUT_YAML"
)

func LoadStackInput(stackInput proto.Message) error {
	var jsonBytes []byte
	var err error

	stackInputYaml := os.Getenv(YamlContentEnvVar)
	if stackInputYaml != "" {
		stackInputFilePath := os.Getenv(FilePathEnvVar)
		if stackInputFilePath == "" {
			return errors.Errorf("stack-input not found in %s or in %s environment variables",
				YamlContentEnvVar, FilePathEnvVar)
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
		jsonBytes, err = yaml.YAMLToJSON([]byte(stackInputYaml))
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
