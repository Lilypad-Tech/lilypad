package bacalhau

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bacalhau-project/bacalhau/pkg/models"

	"github.com/Lilypad-Tech/lilypad/v2/pkg/data"
	"github.com/Lilypad-Tech/lilypad/v2/pkg/data/bacalhau"
	"github.com/bacalhau-project/bacalhau/pkg/publicapi/apimodels"
	clientv2 "github.com/bacalhau-project/bacalhau/pkg/publicapi/client/v2"
)

type BacalhauClient struct {
	api clientv2.API
}

func newBacalhauClient(apiHost string) (*BacalhauClient, error) {
	client := clientv2.NewHTTPClient(apiHost)
	apiClient := clientv2.NewAPI(client)
	return &BacalhauClient{api: apiClient}, nil
}

func (c *BacalhauClient) getID() (string, error) {
	getNodeRequest := apimodels.GetAgentNodeRequest{}
	response, err := c.api.Agent().Node(context.Background(), &getNodeRequest)
	if err != nil {
		return "", err
	}
	return response.NodeInfo.ID(), nil
}

func (c *BacalhauClient) alive() (bool, error) {

	response, err := c.api.Agent().Alive(context.Background())
	if err != nil {
		return false, err
	}

	return response.IsReady(), nil
}

func (c *BacalhauClient) getVersion() (string, error) {
	response, err := c.api.Agent().Version(context.Background())
	if err != nil {
		return "", fmt.Errorf("error getting bacalhau version %s", err.Error())
	}

	return response.BuildVersionInfo.GitVersion, nil
}

func (c *BacalhauClient) postJob(job bacalhau.Job) (*apimodels.PutJobResponse, error) {

	translatedJob := translateJob(job)

	putJobRequest := apimodels.PutJobRequest{
		Job: translatedJob,
	}

	return c.api.Jobs().Put(context.Background(), &putJobRequest)
}

func (c *BacalhauClient) getJob(jobID string) (*apimodels.GetJobResponse, error) {
	getJobRequest := apimodels.GetJobRequest{
		JobID:   jobID,
		Include: "executions",
	}

	response, err := c.api.Jobs().Get(context.Background(), &getJobRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *BacalhauClient) getJobResult(jobId string) (string, error) {
	getJobResultsRequest := apimodels.ListJobResultsRequest{
		JobID: jobId,
	}

	response, err := c.api.Jobs().Results(context.Background(), &getJobResultsRequest)
	if err != nil {
		return "", err
	}
	if len(response.Items) == 0 {
		return "", fmt.Errorf("no results found for job %s", jobId)
	}

	return response.Items[0].Params["URL"].(string), nil
}

func (c *BacalhauClient) getNodes() ([]*models.NodeState, error) {

	getNodesRequest := apimodels.ListNodesRequest{}
	response, err := c.api.Nodes().List(context.Background(), &getNodesRequest)
	if err != nil {
		return nil, err
	}
	return response.Nodes, nil
}

func (c *BacalhauClient) getMachineSpecs() ([]data.MachineSpec, error) {
	nodes, err := c.getNodes()
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

func translateJob(job bacalhau.Job) *models.Job {
	return &models.Job{
		ID:        job.Metadata.ID,
		Name:      job.Metadata.ID,
		Namespace: "default",
		Type:      models.JobTypeBatch,
		Priority:  0,
		Count:     1,
		Meta:      make(map[string]string),
		Labels:    make(map[string]string),
		Tasks: []*models.Task{
			{
				Name: "main",
				Engine: &models.SpecConfig{
					Type: "docker",
					Params: map[string]interface{}{
						"Image":                job.Spec.Docker.Image,
						"Entrypoint":           job.Spec.Docker.Entrypoint,
						"EnvironmentVariables": job.Spec.Docker.EnvironmentVariables,
					},
				},
				Publisher: &models.SpecConfig{
					Type: "local",
				},

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
				ResultPaths: translateOutputSources(job.Spec.Outputs),
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
				Type: source.StorageSource.String(),
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
				"Bucket":         source.S3.Bucket,
				"Key":            source.S3.Key,
				"Endpoint":       source.S3.Endpoint,
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
