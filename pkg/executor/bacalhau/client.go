package bacalhau

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bacalhau-project/bacalhau/pkg/models"

	"github.com/bacalhau-project/bacalhau/pkg/publicapi/apimodels"
	clientv2 "github.com/bacalhau-project/bacalhau/pkg/publicapi/client/v2"
	"github.com/lilypad-tech/lilypad/pkg/data"
	"github.com/lilypad-tech/lilypad/pkg/data/bacalhau"
)

type BacalhauClient struct {
	api clientv2.API
}

func NewBacalhauClient(apiHost string) (*BacalhauClient, error) {
	client := clientv2.NewHTTPClient(apiHost)
	apiClient := clientv2.NewAPI(client)
	return &BacalhauClient{api: apiClient}, nil
}

func (c *BacalhauClient) GetID() (string, error) {
	
	getNodeRequest := apimodels.GetAgentNodeRequest{}
	response, err := c.api.Agent().Node(context.Background(), &getNodeRequest)
	if err != nil {
		return "", err
	}
	return response.NodeState.Info.ID(), nil

}

func (c *BacalhauClient) Alive() (bool, error) {

	response, err := c.api.Agent().Alive(context.Background())
	if err != nil {
		return false, err
	}
	
	return response.IsReady(), nil
}

func (c *BacalhauClient) GetVersion() (string, error) {
	response, err := c.api.Agent().Version(context.Background())
	if err != nil {
		return "", fmt.Errorf("error getting bacalhau version %s", err.Error())
	}

	return response.BuildVersionInfo.GitVersion, nil
}

func (c *BacalhauClient) PostJob(job bacalhau.Job) (*apimodels.PutJobResponse, error) {
	translatedJob := translateJob(job)
	
	putJobRequest := apimodels.PutJobRequest{
		Job: translatedJob,
	}
	
	return c.api.Jobs().Put(context.Background(), &putJobRequest)
}

func (c *BacalhauClient) GetJob(jobID string) (*apimodels.GetJobResponse, error) {
	getJobRequest := apimodels.GetJobRequest{
		JobID: jobID,
	}
	response, err := c.api.Jobs().Get(context.Background(), &getJobRequest)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *BacalhauClient) GetNodes() ([]*models.NodeState, error) {
	
	getNodesRequest := apimodels.ListNodesRequest{}
	response, err := c.api.Nodes().List(context.Background(), &getNodesRequest)
	if err != nil {
		return nil, err
	}
	return response.Nodes, nil
}

func (c *BacalhauClient) GetMachineSpecs() ([]data.MachineSpec, error) {
	nodes, err := c.GetNodes()
	var specs []data.MachineSpec
	for _, node := range nodes {
		spec := data.MachineSpec{
			CPU:  int(node.Info.ComputeNodeInfo.MaxCapacity.CPU) * 1000, // convert float to "milli-CPU"
			RAM:  int(node.Info.ComputeNodeInfo.MaxCapacity.Memory),
			GPU:  int(node.Info.ComputeNodeInfo.MaxCapacity.GPU),
			Disk: int(node.Info.ComputeNodeInfo.MaxCapacity.Disk),
		}
		for _, gpu := range node.Info.ComputeNodeInfo.MaxCapacity.GPUs {
			spec.GPUs = append(spec.GPUs, data.GPUSpec{
				Name:   gpu.Name,
				Vendor: string(gpu.Vendor),
				VRAM:   int(gpu.Memory),
			})
		}
		specs = append(specs, spec)
	}
	if err != nil {
		return nil, err
	}
	return specs, nil
}

func (c *BacalhauClient) GetJobExecutions(jobID string) ([]*models.Execution, error) {
	getJobExecutionsRequest := apimodels.ListJobExecutionsRequest{
		JobID: jobID,
	}
	response, err := c.api.Jobs().Executions(context.Background(), &getJobExecutionsRequest)
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}

func translateJob(job bacalhau.Job) *models.Job {
	return &models.Job{
		ID:          job.Metadata.ID,
		Name:        job.Metadata.ID,
		Namespace:   "default",
		Type:        models.JobTypeBatch,
		Priority:    0,
		Count:       1,
		Constraints: translateConstraints(job.Spec.NodeSelectors),
		Meta:        make(map[string]string),
		Labels:      make(map[string]string),
		Tasks: []*models.Task{
			{
				Name: "main",
				Engine: &models.SpecConfig{
					Type:   job.Spec.EngineSpec.Type,
					Params: job.Spec.EngineSpec.Params,
				},
				Publisher: &models.SpecConfig{
					Type:   string(job.Spec.PublisherSpec.Type),
					Params: job.Spec.PublisherSpec.Params,
				},
				Env: translateEnvironmentVariables(job.Spec.Docker.EnvironmentVariables),
				InputSources: translateInputSources(job.Spec.Inputs),
				ResultPaths:  translateOutputSources(job.Spec.Outputs),
				
				ResourcesConfig: &models.ResourcesConfig{
					CPU:    job.Spec.Resources.CPU,
					Memory: job.Spec.Resources.Memory,
					Disk:   job.Spec.Resources.Disk,
					GPU:    job.Spec.Resources.GPU,
				},
				Network: &models.NetworkConfig{
					Type:    models.Network(job.Spec.Network.Type),
					Domains: job.Spec.Network.Domains,
				},
				Timeouts: &models.TimeoutConfig{
					ExecutionTimeout: job.Spec.Timeout,
					// TODO: add queue timeout and total timeout
				},
			},
		},
		State:      models.NewJobState(models.JobStateTypePending),
		CreateTime: job.Metadata.CreatedAt.UnixNano(),
		ModifyTime: time.Now().UnixNano(),
	}
}


func translateEnvironmentVariables(envVars []string) map[string]string {
	result := make(map[string]string)
	for _, env := range envVars {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result
}


func translateConstraints(constraints []bacalhau.LabelSelectorRequirement) []*models.LabelSelectorRequirement {
	if len(constraints) == 0 {
		return nil
	}

	translated := make([]*models.LabelSelectorRequirement, len(constraints))
	for i, constraint := range constraints {
		translated[i] = &models.LabelSelectorRequirement{
			Key:      constraint.Key,
			Operator: constraint.Operator,
			Values:   constraint.Values,
		}
	}
	return translated
	}


func translateInputSources(sources []bacalhau.StorageSpec) []*models.InputSource {
	if len(sources) == 0 {
		return nil
	}

	translated := make([]*models.InputSource, len(sources))
	for i, source := range sources {
		translated[i] = &models.InputSource{
			Source: &models.SpecConfig{
				Type: string(source.StorageSource),
			},
			Target: source.Path,
		}
		if source.StorageSource == bacalhau.StorageSourceIPFS {
			translated[i].Source.Params = map[string]interface{}{
				"cid": source.CID,
			}
		}
		if source.StorageSource == bacalhau.StorageSourceLocalDirectory {
				translated[i].Source.Params = map[string]interface{}{
					"SourcePath": source.SourcePath,
					"ReadWrite":  source.ReadWrite,
				}
		}
		if source.StorageSource == bacalhau.StorageSourceS3 {
			translated[i].Source.Params = map[string]interface{}{
				"Bucket": source.S3.Bucket,
				"Key":        source.S3.Key,
				"Endpoint":  source.S3.Endpoint,
				"ChecksumSHA256": source.S3.ChecksumSHA256,
			}
		}
		if source.StorageSource == bacalhau.StorageSourceURLDownload {
			translated[i].Source.Params = map[string]interface{}{
				"URL": source.URL,
			}
		}
	}
	return translated
}

func translateOutputSources(sources []bacalhau.StorageSpec) []*models.ResultPath {
	if len(sources) == 0 {
		return nil
	}

	translated := make([]*models.ResultPath, len(sources))
	for i, source := range sources {
		translated[i] = &models.ResultPath{
			Name: source.Name,
			Path: source.Path,
		}
	}
	return translated
}

