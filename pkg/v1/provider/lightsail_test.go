package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// type lightsailMock struct {
// 	mock.Mock
// }

// func (m *lightsailMock)AllocateStaticIp( input *lightsail.AllocateStaticIpInput) (*lightsail.AllocateStaticIpOutput, error){
// return nil, nil
// }
// func (m *lightsailMock)AllocateStaticIpWithContext(context aws.Context, input *lightsail.AllocateStaticIpInput, options ...request.Option) (output *lightsail.AllocateStaticIpOutput, err error){

// }
// func (m *lightsailMock)AllocateStaticIpRequest(input *lightsail.AllocateStaticIpInput) ( req *request.Request, output *lightsail.AllocateStaticIpOutput){

// }

// func (m *lightsailMock)AttachDisk(input *lightsail.AttachDiskInput) (output *lightsail.AttachDiskOutput, err error){

// }
// func (m *lightsailMock)AttachDiskWithContext(aws.Context, *lightsail.AttachDiskInput, ...request.Option) (*lightsail.AttachDiskOutput, error)
// func (m *lightsailMock)AttachDiskRequest(*lightsail.AttachDiskInput) (*request.Request, *lightsail.AttachDiskOutput)

// func (m *lightsailMock)AttachInstancesToLoadBalancer(*lightsail.AttachInstancesToLoadBalancerInput) (*lightsail.AttachInstancesToLoadBalancerOutput, error)
// func (m *lightsailMock)AttachInstancesToLoadBalancerWithContext(aws.Context, *lightsail.AttachInstancesToLoadBalancerInput, ...request.Option) (*lightsail.AttachInstancesToLoadBalancerOutput, error)
// func (m *lightsailMock)AttachInstancesToLoadBalancerRequest(*lightsail.AttachInstancesToLoadBalancerInput) (*request.Request, *lightsail.AttachInstancesToLoadBalancerOutput)

// func (m *lightsailMock)AttachLoadBalancerTlsCertificate(*lightsail.AttachLoadBalancerTlsCertificateInput) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error)
// func (m *lightsailMock)AttachLoadBalancerTlsCertificateWithContext(aws.Context, *lightsail.AttachLoadBalancerTlsCertificateInput, ...request.Option) (*lightsail.AttachLoadBalancerTlsCertificateOutput, error)
// func (m *lightsailMock)AttachLoadBalancerTlsCertificateRequest(*lightsail.AttachLoadBalancerTlsCertificateInput) (*request.Request, *lightsail.AttachLoadBalancerTlsCertificateOutput)

// func (m *lightsailMock)AttachStaticIp(*lightsail.AttachStaticIpInput) (*lightsail.AttachStaticIpOutput, error)
// func (m *lightsailMock)AttachStaticIpWithContext(aws.Context, *lightsail.AttachStaticIpInput, ...request.Option) (*lightsail.AttachStaticIpOutput, error)
// func (m *lightsailMock)AttachStaticIpRequest(*lightsail.AttachStaticIpInput) (*request.Request, *lightsail.AttachStaticIpOutput)

// func (m *lightsailMock)CloseInstancePublicPorts(*lightsail.CloseInstancePublicPortsInput) (*lightsail.CloseInstancePublicPortsOutput, error)
// func (m *lightsailMock)CloseInstancePublicPortsWithContext(aws.Context, *lightsail.CloseInstancePublicPortsInput, ...request.Option) (*lightsail.CloseInstancePublicPortsOutput, error)
// func (m *lightsailMock)CloseInstancePublicPortsRequest(*lightsail.CloseInstancePublicPortsInput) (*request.Request, *lightsail.CloseInstancePublicPortsOutput)

// func (m *lightsailMock)CreateDisk(*lightsail.CreateDiskInput) (*lightsail.CreateDiskOutput, error)
// func (m *lightsailMock)CreateDiskWithContext(aws.Context, *lightsail.CreateDiskInput, ...request.Option) (*lightsail.CreateDiskOutput, error)
// func (m *lightsailMock)CreateDiskRequest(*lightsail.CreateDiskInput) (*request.Request, *lightsail.CreateDiskOutput)

// func (m *lightsailMock)CreateDiskFromSnapshot(*lightsail.CreateDiskFromSnapshotInput) (*lightsail.CreateDiskFromSnapshotOutput, error)
// func (m *lightsailMock)CreateDiskFromSnapshotWithContext(aws.Context, *lightsail.CreateDiskFromSnapshotInput, ...request.Option) (*lightsail.CreateDiskFromSnapshotOutput, error)
// func (m *lightsailMock)CreateDiskFromSnapshotRequest(*lightsail.CreateDiskFromSnapshotInput) (*request.Request, *lightsail.CreateDiskFromSnapshotOutput)

// func (m *lightsailMock)CreateDiskSnapshot(*lightsail.CreateDiskSnapshotInput) (*lightsail.CreateDiskSnapshotOutput, error)
// func (m *lightsailMock)CreateDiskSnapshotWithContext(aws.Context, *lightsail.CreateDiskSnapshotInput, ...request.Option) (*lightsail.CreateDiskSnapshotOutput, error)
// func (m *lightsailMock)CreateDiskSnapshotRequest(*lightsail.CreateDiskSnapshotInput) (*request.Request, *lightsail.CreateDiskSnapshotOutput)

// func (m *lightsailMock)CreateDomain(*lightsail.CreateDomainInput) (*lightsail.CreateDomainOutput, error)
// func (m *lightsailMock)CreateDomainWithContext(aws.Context, *lightsail.CreateDomainInput, ...request.Option) (*lightsail.CreateDomainOutput, error)
// func (m *lightsailMock)CreateDomainRequest(*lightsail.CreateDomainInput) (*request.Request, *lightsail.CreateDomainOutput)

// func (m *lightsailMock)CreateDomainEntry(*lightsail.CreateDomainEntryInput) (*lightsail.CreateDomainEntryOutput, error)
// func (m *lightsailMock)CreateDomainEntryWithContext(aws.Context, *lightsail.CreateDomainEntryInput, ...request.Option) (*lightsail.CreateDomainEntryOutput, error)
// func (m *lightsailMock)CreateDomainEntryRequest(*lightsail.CreateDomainEntryInput) (*request.Request, *lightsail.CreateDomainEntryOutput)

// func (m *lightsailMock)CreateInstanceSnapshot(*lightsail.CreateInstanceSnapshotInput) (*lightsail.CreateInstanceSnapshotOutput, error)
// func (m *lightsailMock)CreateInstanceSnapshotWithContext(aws.Context, *lightsail.CreateInstanceSnapshotInput, ...request.Option) (*lightsail.CreateInstanceSnapshotOutput, error)
// func (m *lightsailMock)CreateInstanceSnapshotRequest(*lightsail.CreateInstanceSnapshotInput) (*request.Request, *lightsail.CreateInstanceSnapshotOutput)

// func (m *lightsailMock)CreateInstances(*lightsail.CreateInstancesInput) (*lightsail.CreateInstancesOutput, error)
// func (m *lightsailMock)CreateInstancesWithContext(aws.Context, *lightsail.CreateInstancesInput, ...request.Option) (*lightsail.CreateInstancesOutput, error)
// func (m *lightsailMock)CreateInstancesRequest(*lightsail.CreateInstancesInput) (*request.Request, *lightsail.CreateInstancesOutput)

// func (m *lightsailMock)CreateInstancesFromSnapshot(*lightsail.CreateInstancesFromSnapshotInput) (*lightsail.CreateInstancesFromSnapshotOutput, error)
// func (m *lightsailMock)CreateInstancesFromSnapshotWithContext(aws.Context, *lightsail.CreateInstancesFromSnapshotInput, ...request.Option) (*lightsail.CreateInstancesFromSnapshotOutput, error)
// func (m *lightsailMock)CreateInstancesFromSnapshotRequest(*lightsail.CreateInstancesFromSnapshotInput) (*request.Request, *lightsail.CreateInstancesFromSnapshotOutput)

// func (m *lightsailMock)CreateKeyPair(*lightsail.CreateKeyPairInput) (*lightsail.CreateKeyPairOutput, error)
// func (m *lightsailMock)CreateKeyPairWithContext(aws.Context, *lightsail.CreateKeyPairInput, ...request.Option) (*lightsail.CreateKeyPairOutput, error)
// func (m *lightsailMock)CreateKeyPairRequest(*lightsail.CreateKeyPairInput) (*request.Request, *lightsail.CreateKeyPairOutput)

// func (m *lightsailMock)CreateLoadBalancer(*lightsail.CreateLoadBalancerInput) (*lightsail.CreateLoadBalancerOutput, error)
// func (m *lightsailMock)CreateLoadBalancerWithContext(aws.Context, *lightsail.CreateLoadBalancerInput, ...request.Option) (*lightsail.CreateLoadBalancerOutput, error)
// func (m *lightsailMock)CreateLoadBalancerRequest(*lightsail.CreateLoadBalancerInput) (*request.Request, *lightsail.CreateLoadBalancerOutput)

// func (m *lightsailMock)CreateLoadBalancerTlsCertificate(*lightsail.CreateLoadBalancerTlsCertificateInput) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error)
// func (m *lightsailMock)CreateLoadBalancerTlsCertificateWithContext(aws.Context, *lightsail.CreateLoadBalancerTlsCertificateInput, ...request.Option) (*lightsail.CreateLoadBalancerTlsCertificateOutput, error)
// func (m *lightsailMock)CreateLoadBalancerTlsCertificateRequest(*lightsail.CreateLoadBalancerTlsCertificateInput) (*request.Request, *lightsail.CreateLoadBalancerTlsCertificateOutput)

// func (m *lightsailMock)CreateRelationalDatabase(*lightsail.CreateRelationalDatabaseInput) (*lightsail.CreateRelationalDatabaseOutput, error)
// func (m *lightsailMock)CreateRelationalDatabaseWithContext(aws.Context, *lightsail.CreateRelationalDatabaseInput, ...request.Option) (*lightsail.CreateRelationalDatabaseOutput, error)
// func (m *lightsailMock)CreateRelationalDatabaseRequest(*lightsail.CreateRelationalDatabaseInput) (*request.Request, *lightsail.CreateRelationalDatabaseOutput)

// func (m *lightsailMock)CreateRelationalDatabaseFromSnapshot(*lightsail.CreateRelationalDatabaseFromSnapshotInput) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error)
// func (m *lightsailMock)CreateRelationalDatabaseFromSnapshotWithContext(aws.Context, *lightsail.CreateRelationalDatabaseFromSnapshotInput, ...request.Option) (*lightsail.CreateRelationalDatabaseFromSnapshotOutput, error)
// func (m *lightsailMock)CreateRelationalDatabaseFromSnapshotRequest(*lightsail.CreateRelationalDatabaseFromSnapshotInput) (*request.Request, *lightsail.CreateRelationalDatabaseFromSnapshotOutput)

// func (m *lightsailMock)CreateRelationalDatabaseSnapshot(*lightsail.CreateRelationalDatabaseSnapshotInput) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error)
// func (m *lightsailMock)CreateRelationalDatabaseSnapshotWithContext(aws.Context, *lightsail.CreateRelationalDatabaseSnapshotInput, ...request.Option) (*lightsail.CreateRelationalDatabaseSnapshotOutput, error)
// func (m *lightsailMock)CreateRelationalDatabaseSnapshotRequest(*lightsail.CreateRelationalDatabaseSnapshotInput) (*request.Request, *lightsail.CreateRelationalDatabaseSnapshotOutput)

// func (m *lightsailMock)DeleteDisk(*lightsail.DeleteDiskInput) (*lightsail.DeleteDiskOutput, error)
// func (m *lightsailMock)DeleteDiskWithContext(aws.Context, *lightsail.DeleteDiskInput, ...request.Option) (*lightsail.DeleteDiskOutput, error)
// func (m *lightsailMock)DeleteDiskRequest(*lightsail.DeleteDiskInput) (*request.Request, *lightsail.DeleteDiskOutput)

// func (m *lightsailMock)DeleteDiskSnapshot(*lightsail.DeleteDiskSnapshotInput) (*lightsail.DeleteDiskSnapshotOutput, error)
// func (m *lightsailMock)DeleteDiskSnapshotWithContext(aws.Context, *lightsail.DeleteDiskSnapshotInput, ...request.Option) (*lightsail.DeleteDiskSnapshotOutput, error)
// func (m *lightsailMock)DeleteDiskSnapshotRequest(*lightsail.DeleteDiskSnapshotInput) (*request.Request, *lightsail.DeleteDiskSnapshotOutput)

// func (m *lightsailMock)DeleteDomain(*lightsail.DeleteDomainInput) (*lightsail.DeleteDomainOutput, error)
// func (m *lightsailMock)DeleteDomainWithContext(aws.Context, *lightsail.DeleteDomainInput, ...request.Option) (*lightsail.DeleteDomainOutput, error)
// func (m *lightsailMock)DeleteDomainRequest(*lightsail.DeleteDomainInput) (*request.Request, *lightsail.DeleteDomainOutput)

// func (m *lightsailMock)DeleteDomainEntry(*lightsail.DeleteDomainEntryInput) (*lightsail.DeleteDomainEntryOutput, error)
// func (m *lightsailMock)DeleteDomainEntryWithContext(aws.Context, *lightsail.DeleteDomainEntryInput, ...request.Option) (*lightsail.DeleteDomainEntryOutput, error)
// func (m *lightsailMock)DeleteDomainEntryRequest(*lightsail.DeleteDomainEntryInput) (*request.Request, *lightsail.DeleteDomainEntryOutput)

// func (m *lightsailMock)DeleteInstance(*lightsail.DeleteInstanceInput) (*lightsail.DeleteInstanceOutput, error)
// func (m *lightsailMock)DeleteInstanceWithContext(aws.Context, *lightsail.DeleteInstanceInput, ...request.Option) (*lightsail.DeleteInstanceOutput, error)
// func (m *lightsailMock)DeleteInstanceRequest(*lightsail.DeleteInstanceInput) (*request.Request, *lightsail.DeleteInstanceOutput)

// func (m *lightsailMock)DeleteInstanceSnapshot(*lightsail.DeleteInstanceSnapshotInput) (*lightsail.DeleteInstanceSnapshotOutput, error)
// func (m *lightsailMock)DeleteInstanceSnapshotWithContext(aws.Context, *lightsail.DeleteInstanceSnapshotInput, ...request.Option) (*lightsail.DeleteInstanceSnapshotOutput, error)
// func (m *lightsailMock)DeleteInstanceSnapshotRequest(*lightsail.DeleteInstanceSnapshotInput) (*request.Request, *lightsail.DeleteInstanceSnapshotOutput)

// func (m *lightsailMock)DeleteKeyPair(*lightsail.DeleteKeyPairInput) (*lightsail.DeleteKeyPairOutput, error)
// func (m *lightsailMock)DeleteKeyPairWithContext(aws.Context, *lightsail.DeleteKeyPairInput, ...request.Option) (*lightsail.DeleteKeyPairOutput, error)
// func (m *lightsailMock)DeleteKeyPairRequest(*lightsail.DeleteKeyPairInput) (*request.Request, *lightsail.DeleteKeyPairOutput)

// func (m *lightsailMock)DeleteLoadBalancer(*lightsail.DeleteLoadBalancerInput) (*lightsail.DeleteLoadBalancerOutput, error)
// func (m *lightsailMock)DeleteLoadBalancerWithContext(aws.Context, *lightsail.DeleteLoadBalancerInput, ...request.Option) (*lightsail.DeleteLoadBalancerOutput, error)
// func (m *lightsailMock)DeleteLoadBalancerRequest(*lightsail.DeleteLoadBalancerInput) (*request.Request, *lightsail.DeleteLoadBalancerOutput)

// func (m *lightsailMock)DeleteLoadBalancerTlsCertificate(*lightsail.DeleteLoadBalancerTlsCertificateInput) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error)
// func (m *lightsailMock)DeleteLoadBalancerTlsCertificateWithContext(aws.Context, *lightsail.DeleteLoadBalancerTlsCertificateInput, ...request.Option) (*lightsail.DeleteLoadBalancerTlsCertificateOutput, error)
// func (m *lightsailMock)DeleteLoadBalancerTlsCertificateRequest(*lightsail.DeleteLoadBalancerTlsCertificateInput) (*request.Request, *lightsail.DeleteLoadBalancerTlsCertificateOutput)

// func (m *lightsailMock)DeleteRelationalDatabase(*lightsail.DeleteRelationalDatabaseInput) (*lightsail.DeleteRelationalDatabaseOutput, error)
// func (m *lightsailMock)DeleteRelationalDatabaseWithContext(aws.Context, *lightsail.DeleteRelationalDatabaseInput, ...request.Option) (*lightsail.DeleteRelationalDatabaseOutput, error)
// func (m *lightsailMock)DeleteRelationalDatabaseRequest(*lightsail.DeleteRelationalDatabaseInput) (*request.Request, *lightsail.DeleteRelationalDatabaseOutput)

// func (m *lightsailMock)DeleteRelationalDatabaseSnapshot(*lightsail.DeleteRelationalDatabaseSnapshotInput) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error)
// func (m *lightsailMock)DeleteRelationalDatabaseSnapshotWithContext(aws.Context, *lightsail.DeleteRelationalDatabaseSnapshotInput, ...request.Option) (*lightsail.DeleteRelationalDatabaseSnapshotOutput, error)
// func (m *lightsailMock)DeleteRelationalDatabaseSnapshotRequest(*lightsail.DeleteRelationalDatabaseSnapshotInput) (*request.Request, *lightsail.DeleteRelationalDatabaseSnapshotOutput)

// func (m *lightsailMock)DetachDisk(*lightsail.DetachDiskInput) (*lightsail.DetachDiskOutput, error)
// func (m *lightsailMock)DetachDiskWithContext(aws.Context, *lightsail.DetachDiskInput, ...request.Option) (*lightsail.DetachDiskOutput, error)
// func (m *lightsailMock)DetachDiskRequest(*lightsail.DetachDiskInput) (*request.Request, *lightsail.DetachDiskOutput)

// func (m *lightsailMock)DetachInstancesFromLoadBalancer(*lightsail.DetachInstancesFromLoadBalancerInput) (*lightsail.DetachInstancesFromLoadBalancerOutput, error)
// func (m *lightsailMock)DetachInstancesFromLoadBalancerWithContext(aws.Context, *lightsail.DetachInstancesFromLoadBalancerInput, ...request.Option) (*lightsail.DetachInstancesFromLoadBalancerOutput, error)
// func (m *lightsailMock)DetachInstancesFromLoadBalancerRequest(*lightsail.DetachInstancesFromLoadBalancerInput) (*request.Request, *lightsail.DetachInstancesFromLoadBalancerOutput)

// func (m *lightsailMock)DetachStaticIp(*lightsail.DetachStaticIpInput) (*lightsail.DetachStaticIpOutput, error)
// func (m *lightsailMock)DetachStaticIpWithContext(aws.Context, *lightsail.DetachStaticIpInput, ...request.Option) (*lightsail.DetachStaticIpOutput, error)
// func (m *lightsailMock)DetachStaticIpRequest(*lightsail.DetachStaticIpInput) (*request.Request, *lightsail.DetachStaticIpOutput)

// func (m *lightsailMock)DownloadDefaultKeyPair(*lightsail.DownloadDefaultKeyPairInput) (*lightsail.DownloadDefaultKeyPairOutput, error)
// func (m *lightsailMock)DownloadDefaultKeyPairWithContext(aws.Context, *lightsail.DownloadDefaultKeyPairInput, ...request.Option) (*lightsail.DownloadDefaultKeyPairOutput, error)
// func (m *lightsailMock)DownloadDefaultKeyPairRequest(*lightsail.DownloadDefaultKeyPairInput) (*request.Request, *lightsail.DownloadDefaultKeyPairOutput)

// func (m *lightsailMock)GetActiveNames(*lightsail.GetActiveNamesInput) (*lightsail.GetActiveNamesOutput, error)
// func (m *lightsailMock)GetActiveNamesWithContext(aws.Context, *lightsail.GetActiveNamesInput, ...request.Option) (*lightsail.GetActiveNamesOutput, error)
// func (m *lightsailMock)GetActiveNamesRequest(*lightsail.GetActiveNamesInput) (*request.Request, *lightsail.GetActiveNamesOutput)

// func (m *lightsailMock)GetBlueprints(*lightsail.GetBlueprintsInput) (*lightsail.GetBlueprintsOutput, error)
// func (m *lightsailMock)GetBlueprintsWithContext(aws.Context, *lightsail.GetBlueprintsInput, ...request.Option) (*lightsail.GetBlueprintsOutput, error)
// func (m *lightsailMock)GetBlueprintsRequest(*lightsail.GetBlueprintsInput) (*request.Request, *lightsail.GetBlueprintsOutput)

// func (m *lightsailMock)GetBundles(*lightsail.GetBundlesInput) (*lightsail.GetBundlesOutput, error)
// func (m *lightsailMock)GetBundlesWithContext(aws.Context, *lightsail.GetBundlesInput, ...request.Option) (*lightsail.GetBundlesOutput, error)
// func (m *lightsailMock)GetBundlesRequest(*lightsail.GetBundlesInput) (*request.Request, *lightsail.GetBundlesOutput)

// func (m *lightsailMock)GetDisk(*lightsail.GetDiskInput) (*lightsail.GetDiskOutput, error)
// func (m *lightsailMock)GetDiskWithContext(aws.Context, *lightsail.GetDiskInput, ...request.Option) (*lightsail.GetDiskOutput, error)
// func (m *lightsailMock)GetDiskRequest(*lightsail.GetDiskInput) (*request.Request, *lightsail.GetDiskOutput)

// func (m *lightsailMock)GetDiskSnapshot(*lightsail.GetDiskSnapshotInput) (*lightsail.GetDiskSnapshotOutput, error)
// func (m *lightsailMock)GetDiskSnapshotWithContext(aws.Context, *lightsail.GetDiskSnapshotInput, ...request.Option) (*lightsail.GetDiskSnapshotOutput, error)
// func (m *lightsailMock)GetDiskSnapshotRequest(*lightsail.GetDiskSnapshotInput) (*request.Request, *lightsail.GetDiskSnapshotOutput)

// func (m *lightsailMock)GetDiskSnapshots(*lightsail.GetDiskSnapshotsInput) (*lightsail.GetDiskSnapshotsOutput, error)
// func (m *lightsailMock)GetDiskSnapshotsWithContext(aws.Context, *lightsail.GetDiskSnapshotsInput, ...request.Option) (*lightsail.GetDiskSnapshotsOutput, error)
// func (m *lightsailMock)GetDiskSnapshotsRequest(*lightsail.GetDiskSnapshotsInput) (*request.Request, *lightsail.GetDiskSnapshotsOutput)

// func (m *lightsailMock)GetDisks(*lightsail.GetDisksInput) (*lightsail.GetDisksOutput, error)
// func (m *lightsailMock)GetDisksWithContext(aws.Context, *lightsail.GetDisksInput, ...request.Option) (*lightsail.GetDisksOutput, error)
// func (m *lightsailMock)GetDisksRequest(*lightsail.GetDisksInput) (*request.Request, *lightsail.GetDisksOutput)

// func (m *lightsailMock)GetDomain(*lightsail.GetDomainInput) (*lightsail.GetDomainOutput, error)
// func (m *lightsailMock)GetDomainWithContext(aws.Context, *lightsail.GetDomainInput, ...request.Option) (*lightsail.GetDomainOutput, error)
// func (m *lightsailMock)GetDomainRequest(*lightsail.GetDomainInput) (*request.Request, *lightsail.GetDomainOutput)

// func (m *lightsailMock)GetDomains(*lightsail.GetDomainsInput) (*lightsail.GetDomainsOutput, error)
// func (m *lightsailMock)GetDomainsWithContext(aws.Context, *lightsail.GetDomainsInput, ...request.Option) (*lightsail.GetDomainsOutput, error)
// func (m *lightsailMock)GetDomainsRequest(*lightsail.GetDomainsInput) (*request.Request, *lightsail.GetDomainsOutput)

// func (m *lightsailMock)GetInstance(*lightsail.GetInstanceInput) (*lightsail.GetInstanceOutput, error)
// func (m *lightsailMock)GetInstanceWithContext(aws.Context, *lightsail.GetInstanceInput, ...request.Option) (*lightsail.GetInstanceOutput, error)
// func (m *lightsailMock)GetInstanceRequest(*lightsail.GetInstanceInput) (*request.Request, *lightsail.GetInstanceOutput)

// func (m *lightsailMock)GetInstanceAccessDetails(*lightsail.GetInstanceAccessDetailsInput) (*lightsail.GetInstanceAccessDetailsOutput, error)
// func (m *lightsailMock)GetInstanceAccessDetailsWithContext(aws.Context, *lightsail.GetInstanceAccessDetailsInput, ...request.Option) (*lightsail.GetInstanceAccessDetailsOutput, error)
// func (m *lightsailMock)GetInstanceAccessDetailsRequest(*lightsail.GetInstanceAccessDetailsInput) (*request.Request, *lightsail.GetInstanceAccessDetailsOutput)

// func (m *lightsailMock)GetInstanceMetricData(*lightsail.GetInstanceMetricDataInput) (*lightsail.GetInstanceMetricDataOutput, error)
// func (m *lightsailMock)GetInstanceMetricDataWithContext(aws.Context, *lightsail.GetInstanceMetricDataInput, ...request.Option) (*lightsail.GetInstanceMetricDataOutput, error)
// func (m *lightsailMock)GetInstanceMetricDataRequest(*lightsail.GetInstanceMetricDataInput) (*request.Request, *lightsail.GetInstanceMetricDataOutput)

// func (m *lightsailMock)GetInstancePortStates(*lightsail.GetInstancePortStatesInput) (*lightsail.GetInstancePortStatesOutput, error)
// func (m *lightsailMock)GetInstancePortStatesWithContext(aws.Context, *lightsail.GetInstancePortStatesInput, ...request.Option) (*lightsail.GetInstancePortStatesOutput, error)
// func (m *lightsailMock)GetInstancePortStatesRequest(*lightsail.GetInstancePortStatesInput) (*request.Request, *lightsail.GetInstancePortStatesOutput)

// func (m *lightsailMock)GetInstanceSnapshot(*lightsail.GetInstanceSnapshotInput) (*lightsail.GetInstanceSnapshotOutput, error)
// func (m *lightsailMock)GetInstanceSnapshotWithContext(aws.Context, *lightsail.GetInstanceSnapshotInput, ...request.Option) (*lightsail.GetInstanceSnapshotOutput, error)
// func (m *lightsailMock)GetInstanceSnapshotRequest(*lightsail.GetInstanceSnapshotInput) (*request.Request, *lightsail.GetInstanceSnapshotOutput)

// func (m *lightsailMock)GetInstanceSnapshots(*lightsail.GetInstanceSnapshotsInput) (*lightsail.GetInstanceSnapshotsOutput, error)
// func (m *lightsailMock)GetInstanceSnapshotsWithContext(aws.Context, *lightsail.GetInstanceSnapshotsInput, ...request.Option) (*lightsail.GetInstanceSnapshotsOutput, error)
// func (m *lightsailMock)GetInstanceSnapshotsRequest(*lightsail.GetInstanceSnapshotsInput) (*request.Request, *lightsail.GetInstanceSnapshotsOutput)

// func (m *lightsailMock)GetInstanceState(*lightsail.GetInstanceStateInput) (*lightsail.GetInstanceStateOutput, error)
// func (m *lightsailMock)GetInstanceStateWithContext(aws.Context, *lightsail.GetInstanceStateInput, ...request.Option) (*lightsail.GetInstanceStateOutput, error)
// func (m *lightsailMock)GetInstanceStateRequest(*lightsail.GetInstanceStateInput) (*request.Request, *lightsail.GetInstanceStateOutput)

// func (m *lightsailMock)GetInstances(*lightsail.GetInstancesInput) (*lightsail.GetInstancesOutput, error)
// func (m *lightsailMock)GetInstancesWithContext(aws.Context, *lightsail.GetInstancesInput, ...request.Option) (*lightsail.GetInstancesOutput, error)
// func (m *lightsailMock)GetInstancesRequest(*lightsail.GetInstancesInput) (*request.Request, *lightsail.GetInstancesOutput)

// func (m *lightsailMock)GetKeyPair(*lightsail.GetKeyPairInput) (*lightsail.GetKeyPairOutput, error)
// func (m *lightsailMock)GetKeyPairWithContext(aws.Context, *lightsail.GetKeyPairInput, ...request.Option) (*lightsail.GetKeyPairOutput, error)
// func (m *lightsailMock)GetKeyPairRequest(*lightsail.GetKeyPairInput) (*request.Request, *lightsail.GetKeyPairOutput)

// func (m *lightsailMock)GetKeyPairs(*lightsail.GetKeyPairsInput) (*lightsail.GetKeyPairsOutput, error)
// func (m *lightsailMock)GetKeyPairsWithContext(aws.Context, *lightsail.GetKeyPairsInput, ...request.Option) (*lightsail.GetKeyPairsOutput, error)
// func (m *lightsailMock)GetKeyPairsRequest(*lightsail.GetKeyPairsInput) (*request.Request, *lightsail.GetKeyPairsOutput)

// func (m *lightsailMock)GetLoadBalancer(*lightsail.GetLoadBalancerInput) (*lightsail.GetLoadBalancerOutput, error)
// func (m *lightsailMock)GetLoadBalancerWithContext(aws.Context, *lightsail.GetLoadBalancerInput, ...request.Option) (*lightsail.GetLoadBalancerOutput, error)
// func (m *lightsailMock)GetLoadBalancerRequest(*lightsail.GetLoadBalancerInput) (*request.Request, *lightsail.GetLoadBalancerOutput)

// func (m *lightsailMock)GetLoadBalancerMetricData(*lightsail.GetLoadBalancerMetricDataInput) (*lightsail.GetLoadBalancerMetricDataOutput, error)
// func (m *lightsailMock)GetLoadBalancerMetricDataWithContext(aws.Context, *lightsail.GetLoadBalancerMetricDataInput, ...request.Option) (*lightsail.GetLoadBalancerMetricDataOutput, error)
// func (m *lightsailMock)GetLoadBalancerMetricDataRequest(*lightsail.GetLoadBalancerMetricDataInput) (*request.Request, *lightsail.GetLoadBalancerMetricDataOutput)

// func (m *lightsailMock)GetLoadBalancerTlsCertificates(*lightsail.GetLoadBalancerTlsCertificatesInput) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error)
// func (m *lightsailMock)GetLoadBalancerTlsCertificatesWithContext(aws.Context, *lightsail.GetLoadBalancerTlsCertificatesInput, ...request.Option) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error)
// func (m *lightsailMock)GetLoadBalancerTlsCertificatesRequest(*lightsail.GetLoadBalancerTlsCertificatesInput) (*request.Request, *lightsail.GetLoadBalancerTlsCertificatesOutput)

// func (m *lightsailMock)GetLoadBalancers(*lightsail.GetLoadBalancersInput) (*lightsail.GetLoadBalancersOutput, error)
// func (m *lightsailMock)GetLoadBalancersWithContext(aws.Context, *lightsail.GetLoadBalancersInput, ...request.Option) (*lightsail.GetLoadBalancersOutput, error)
// func (m *lightsailMock)GetLoadBalancersRequest(*lightsail.GetLoadBalancersInput) (*request.Request, *lightsail.GetLoadBalancersOutput)

// func (m *lightsailMock)GetOperation(*lightsail.GetOperationInput) (*lightsail.GetOperationOutput, error)
// func (m *lightsailMock)GetOperationWithContext(aws.Context, *lightsail.GetOperationInput, ...request.Option) (*lightsail.GetOperationOutput, error)
// func (m *lightsailMock)GetOperationRequest(*lightsail.GetOperationInput) (*request.Request, *lightsail.GetOperationOutput)func (m *lightsailMock)
// func (m *lightsailMock)GetOperations(*lightsail.GetOperationsInput) (*lightsail.GetOperationsOutput, error)
// func (m *lightsailMock)GetOperationsWithContext(aws.Context, *lightsail.GetOperationsInput, ...request.Option) (*lightsail.GetOperationsOutput, error)
// func (m *lightsailMock)GetOperationsRequest(*lightsail.GetOperationsInput) (*request.Request, *lightsail.GetOperationsOutput)

// func (m *lightsailMock)GetOperationsForResource(*lightsail.GetOperationsForResourceInput) (*lightsail.GetOperationsForResourceOutput, error)
// func (m *lightsailMock)GetOperationsForResourceWithContext(aws.Context, *lightsail.GetOperationsForResourceInput, ...request.Option) (*lightsail.GetOperationsForResourceOutput, error)
// func (m *lightsailMock)GetOperationsForResourceRequest(*lightsail.GetOperationsForResourceInput) (*request.Request, *lightsail.GetOperationsForResourceOutput)

// 	GetRegions(*lightsail.GetRegionsInput) (*lightsail.GetRegionsOutput, error)
// 	GetRegionsWithContext(aws.Context, *lightsail.GetRegionsInput, ...request.Option) (*lightsail.GetRegionsOutput, error)
// 	GetRegionsRequest(*lightsail.GetRegionsInput) (*request.Request, *lightsail.GetRegionsOutput)

// 	GetRelationalDatabase(*lightsail.GetRelationalDatabaseInput) (*lightsail.GetRelationalDatabaseOutput, error)
// 	GetRelationalDatabaseWithContext(aws.Context, *lightsail.GetRelationalDatabaseInput, ...request.Option) (*lightsail.GetRelationalDatabaseOutput, error)
// 	GetRelationalDatabaseRequest(*lightsail.GetRelationalDatabaseInput) (*request.Request, *lightsail.GetRelationalDatabaseOutput)

// 	GetRelationalDatabaseBlueprints(*lightsail.GetRelationalDatabaseBlueprintsInput) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error)
// 	GetRelationalDatabaseBlueprintsWithContext(aws.Context, *lightsail.GetRelationalDatabaseBlueprintsInput, ...request.Option) (*lightsail.GetRelationalDatabaseBlueprintsOutput, error)
// 	GetRelationalDatabaseBlueprintsRequest(*lightsail.GetRelationalDatabaseBlueprintsInput) (*request.Request, *lightsail.GetRelationalDatabaseBlueprintsOutput)

// 	GetRelationalDatabaseBundles(*lightsail.GetRelationalDatabaseBundlesInput) (*lightsail.GetRelationalDatabaseBundlesOutput, error)
// 	GetRelationalDatabaseBundlesWithContext(aws.Context, *lightsail.GetRelationalDatabaseBundlesInput, ...request.Option) (*lightsail.GetRelationalDatabaseBundlesOutput, error)
// 	GetRelationalDatabaseBundlesRequest(*lightsail.GetRelationalDatabaseBundlesInput) (*request.Request, *lightsail.GetRelationalDatabaseBundlesOutput)

// 	GetRelationalDatabaseEvents(*lightsail.GetRelationalDatabaseEventsInput) (*lightsail.GetRelationalDatabaseEventsOutput, error)
// 	GetRelationalDatabaseEventsWithContext(aws.Context, *lightsail.GetRelationalDatabaseEventsInput, ...request.Option) (*lightsail.GetRelationalDatabaseEventsOutput, error)
// 	GetRelationalDatabaseEventsRequest(*lightsail.GetRelationalDatabaseEventsInput) (*request.Request, *lightsail.GetRelationalDatabaseEventsOutput)

// 	GetRelationalDatabaseLogEvents(*lightsail.GetRelationalDatabaseLogEventsInput) (*lightsail.GetRelationalDatabaseLogEventsOutput, error)
// 	GetRelationalDatabaseLogEventsWithContext(aws.Context, *lightsail.GetRelationalDatabaseLogEventsInput, ...request.Option) (*lightsail.GetRelationalDatabaseLogEventsOutput, error)
// 	GetRelationalDatabaseLogEventsRequest(*lightsail.GetRelationalDatabaseLogEventsInput) (*request.Request, *lightsail.GetRelationalDatabaseLogEventsOutput)

// 	GetRelationalDatabaseLogStreams(*lightsail.GetRelationalDatabaseLogStreamsInput) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error)
// 	GetRelationalDatabaseLogStreamsWithContext(aws.Context, *lightsail.GetRelationalDatabaseLogStreamsInput, ...request.Option) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error)
// 	GetRelationalDatabaseLogStreamsRequest(*lightsail.GetRelationalDatabaseLogStreamsInput) (*request.Request, *lightsail.GetRelationalDatabaseLogStreamsOutput)

// 	GetRelationalDatabaseMasterUserPassword(*lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error)
// 	GetRelationalDatabaseMasterUserPasswordWithContext(aws.Context, *lightsail.GetRelationalDatabaseMasterUserPasswordInput, ...request.Option) (*lightsail.GetRelationalDatabaseMasterUserPasswordOutput, error)
// 	GetRelationalDatabaseMasterUserPasswordRequest(*lightsail.GetRelationalDatabaseMasterUserPasswordInput) (*request.Request, *lightsail.GetRelationalDatabaseMasterUserPasswordOutput)

// 	GetRelationalDatabaseMetricData(*lightsail.GetRelationalDatabaseMetricDataInput) (*lightsail.GetRelationalDatabaseMetricDataOutput, error)
// 	GetRelationalDatabaseMetricDataWithContext(aws.Context, *lightsail.GetRelationalDatabaseMetricDataInput, ...request.Option) (*lightsail.GetRelationalDatabaseMetricDataOutput, error)
// 	GetRelationalDatabaseMetricDataRequest(*lightsail.GetRelationalDatabaseMetricDataInput) (*request.Request, *lightsail.GetRelationalDatabaseMetricDataOutput)

// 	GetRelationalDatabaseParameters(*lightsail.GetRelationalDatabaseParametersInput) (*lightsail.GetRelationalDatabaseParametersOutput, error)
// 	GetRelationalDatabaseParametersWithContext(aws.Context, *lightsail.GetRelationalDatabaseParametersInput, ...request.Option) (*lightsail.GetRelationalDatabaseParametersOutput, error)
// 	GetRelationalDatabaseParametersRequest(*lightsail.GetRelationalDatabaseParametersInput) (*request.Request, *lightsail.GetRelationalDatabaseParametersOutput)

// 	GetRelationalDatabaseSnapshot(*lightsail.GetRelationalDatabaseSnapshotInput) (*lightsail.GetRelationalDatabaseSnapshotOutput, error)
// 	GetRelationalDatabaseSnapshotWithContext(aws.Context, *lightsail.GetRelationalDatabaseSnapshotInput, ...request.Option) (*lightsail.GetRelationalDatabaseSnapshotOutput, error)
// 	GetRelationalDatabaseSnapshotRequest(*lightsail.GetRelationalDatabaseSnapshotInput) (*request.Request, *lightsail.GetRelationalDatabaseSnapshotOutput)

// 	GetRelationalDatabaseSnapshots(*lightsail.GetRelationalDatabaseSnapshotsInput) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error)
// 	GetRelationalDatabaseSnapshotsWithContext(aws.Context, *lightsail.GetRelationalDatabaseSnapshotsInput, ...request.Option) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error)
// 	GetRelationalDatabaseSnapshotsRequest(*lightsail.GetRelationalDatabaseSnapshotsInput) (*request.Request, *lightsail.GetRelationalDatabaseSnapshotsOutput)

// 	GetRelationalDatabases(*lightsail.GetRelationalDatabasesInput) (*lightsail.GetRelationalDatabasesOutput, error)
// 	GetRelationalDatabasesWithContext(aws.Context, *lightsail.GetRelationalDatabasesInput, ...request.Option) (*lightsail.GetRelationalDatabasesOutput, error)
// 	GetRelationalDatabasesRequest(*lightsail.GetRelationalDatabasesInput) (*request.Request, *lightsail.GetRelationalDatabasesOutput)

// 	GetStaticIp(*lightsail.GetStaticIpInput) (*lightsail.GetStaticIpOutput, error)
// 	GetStaticIpWithContext(aws.Context, *lightsail.GetStaticIpInput, ...request.Option) (*lightsail.GetStaticIpOutput, error)
// 	GetStaticIpRequest(*lightsail.GetStaticIpInput) (*request.Request, *lightsail.GetStaticIpOutput)

// 	GetStaticIps(*lightsail.GetStaticIpsInput) (*lightsail.GetStaticIpsOutput, error)
// 	GetStaticIpsWithContext(aws.Context, *lightsail.GetStaticIpsInput, ...request.Option) (*lightsail.GetStaticIpsOutput, error)
// 	GetStaticIpsRequest(*lightsail.GetStaticIpsInput) (*request.Request, *lightsail.GetStaticIpsOutput)

// 	ImportKeyPair(*lightsail.ImportKeyPairInput) (*lightsail.ImportKeyPairOutput, error)
// 	ImportKeyPairWithContext(aws.Context, *lightsail.ImportKeyPairInput, ...request.Option) (*lightsail.ImportKeyPairOutput, error)
// 	ImportKeyPairRequest(*lightsail.ImportKeyPairInput) (*request.Request, *lightsail.ImportKeyPairOutput)

// 	IsVpcPeered(*lightsail.IsVpcPeeredInput) (*lightsail.IsVpcPeeredOutput, error)
// 	IsVpcPeeredWithContext(aws.Context, *lightsail.IsVpcPeeredInput, ...request.Option) (*lightsail.IsVpcPeeredOutput, error)
// 	IsVpcPeeredRequest(*lightsail.IsVpcPeeredInput) (*request.Request, *lightsail.IsVpcPeeredOutput)

// 	OpenInstancePublicPorts(*lightsail.OpenInstancePublicPortsInput) (*lightsail.OpenInstancePublicPortsOutput, error)
// 	OpenInstancePublicPortsWithContext(aws.Context, *lightsail.OpenInstancePublicPortsInput, ...request.Option) (*lightsail.OpenInstancePublicPortsOutput, error)
// 	OpenInstancePublicPortsRequest(*lightsail.OpenInstancePublicPortsInput) (*request.Request, *lightsail.OpenInstancePublicPortsOutput)

// 	PeerVpc(*lightsail.PeerVpcInput) (*lightsail.PeerVpcOutput, error)
// 	PeerVpcWithContext(aws.Context, *lightsail.PeerVpcInput, ...request.Option) (*lightsail.PeerVpcOutput, error)
// 	PeerVpcRequest(*lightsail.PeerVpcInput) (*request.Request, *lightsail.PeerVpcOutput)

// 	PutInstancePublicPorts(*lightsail.PutInstancePublicPortsInput) (*lightsail.PutInstancePublicPortsOutput, error)
// 	PutInstancePublicPortsWithContext(aws.Context, *lightsail.PutInstancePublicPortsInput, ...request.Option) (*lightsail.PutInstancePublicPortsOutput, error)
// 	PutInstancePublicPortsRequest(*lightsail.PutInstancePublicPortsInput) (*request.Request, *lightsail.PutInstancePublicPortsOutput)

// 	RebootInstance(*lightsail.RebootInstanceInput) (*lightsail.RebootInstanceOutput, error)
// 	RebootInstanceWithContext(aws.Context, *lightsail.RebootInstanceInput, ...request.Option) (*lightsail.RebootInstanceOutput, error)
// 	RebootInstanceRequest(*lightsail.RebootInstanceInput) (*request.Request, *lightsail.RebootInstanceOutput)

// 	RebootRelationalDatabase(*lightsail.RebootRelationalDatabaseInput) (*lightsail.RebootRelationalDatabaseOutput, error)
// 	RebootRelationalDatabaseWithContext(aws.Context, *lightsail.RebootRelationalDatabaseInput, ...request.Option) (*lightsail.RebootRelationalDatabaseOutput, error)
// 	RebootRelationalDatabaseRequest(*lightsail.RebootRelationalDatabaseInput) (*request.Request, *lightsail.RebootRelationalDatabaseOutput)

// 	ReleaseStaticIp(*lightsail.ReleaseStaticIpInput) (*lightsail.ReleaseStaticIpOutput, error)
// 	ReleaseStaticIpWithContext(aws.Context, *lightsail.ReleaseStaticIpInput, ...request.Option) (*lightsail.ReleaseStaticIpOutput, error)
// 	ReleaseStaticIpRequest(*lightsail.ReleaseStaticIpInput) (*request.Request, *lightsail.ReleaseStaticIpOutput)

// 	StartInstance(*lightsail.StartInstanceInput) (*lightsail.StartInstanceOutput, error)
// 	StartInstanceWithContext(aws.Context, *lightsail.StartInstanceInput, ...request.Option) (*lightsail.StartInstanceOutput, error)
// 	StartInstanceRequest(*lightsail.StartInstanceInput) (*request.Request, *lightsail.StartInstanceOutput)

// 	StartRelationalDatabase(*lightsail.StartRelationalDatabaseInput) (*lightsail.StartRelationalDatabaseOutput, error)
// 	StartRelationalDatabaseWithContext(aws.Context, *lightsail.StartRelationalDatabaseInput, ...request.Option) (*lightsail.StartRelationalDatabaseOutput, error)
// 	StartRelationalDatabaseRequest(*lightsail.StartRelationalDatabaseInput) (*request.Request, *lightsail.StartRelationalDatabaseOutput)

// 	StopInstance(*lightsail.StopInstanceInput) (*lightsail.StopInstanceOutput, error)
// 	StopInstanceWithContext(aws.Context, *lightsail.StopInstanceInput, ...request.Option) (*lightsail.StopInstanceOutput, error)
// 	StopInstanceRequest(*lightsail.StopInstanceInput) (*request.Request, *lightsail.StopInstanceOutput)

// 	StopRelationalDatabase(*lightsail.StopRelationalDatabaseInput) (*lightsail.StopRelationalDatabaseOutput, error)
// 	StopRelationalDatabaseWithContext(aws.Context, *lightsail.StopRelationalDatabaseInput, ...request.Option) (*lightsail.StopRelationalDatabaseOutput, error)
// 	StopRelationalDatabaseRequest(*lightsail.StopRelationalDatabaseInput) (*request.Request, *lightsail.StopRelationalDatabaseOutput)

// 	UnpeerVpc(*lightsail.UnpeerVpcInput) (*lightsail.UnpeerVpcOutput, error)
// 	UnpeerVpcWithContext(aws.Context, *lightsail.UnpeerVpcInput, ...request.Option) (*lightsail.UnpeerVpcOutput, error)
// 	UnpeerVpcRequest(*lightsail.UnpeerVpcInput) (*request.Request, *lightsail.UnpeerVpcOutput)

// 	UpdateDomainEntry(*lightsail.UpdateDomainEntryInput) (*lightsail.UpdateDomainEntryOutput, error)
// 	UpdateDomainEntryWithContext(aws.Context, *lightsail.UpdateDomainEntryInput, ...request.Option) (*lightsail.UpdateDomainEntryOutput, error)
// 	UpdateDomainEntryRequest(*lightsail.UpdateDomainEntryInput) (*request.Request, *lightsail.UpdateDomainEntryOutput)

// 	UpdateLoadBalancerAttribute(*lightsail.UpdateLoadBalancerAttributeInput) (*lightsail.UpdateLoadBalancerAttributeOutput, error)
// 	UpdateLoadBalancerAttributeWithContext(aws.Context, *lightsail.UpdateLoadBalancerAttributeInput, ...request.Option) (*lightsail.UpdateLoadBalancerAttributeOutput, error)
// 	UpdateLoadBalancerAttributeRequest(*lightsail.UpdateLoadBalancerAttributeInput) (*request.Request, *lightsail.UpdateLoadBalancerAttributeOutput)

// 	UpdateRelationalDatabase(*lightsail.UpdateRelationalDatabaseInput) (*lightsail.UpdateRelationalDatabaseOutput, error)
// 	UpdateRelationalDatabaseWithContext(aws.Context, *lightsail.UpdateRelationalDatabaseInput, ...request.Option) (*lightsail.UpdateRelationalDatabaseOutput, error)
// 	UpdateRelationalDatabaseRequest(*lightsail.UpdateRelationalDatabaseInput) (*request.Request, *lightsail.UpdateRelationalDatabaseOutput)

// 	UpdateRelationalDatabaseParameters(*lightsail.UpdateRelationalDatabaseParametersInput) (*lightsail.UpdateRelationalDatabaseParametersOutput, error)
// 	UpdateRelationalDatabaseParametersWithContext(aws.Context, *lightsail.UpdateRelationalDatabaseParametersInput, ...request.Option) (*lightsail.UpdateRelationalDatabaseParametersOutput, error)
// 	UpdateRelationalDatabaseParametersRequest(*lightsail.UpdateRelationalDatabaseParametersInput) (*request.Request, *lightsail.UpdateRelationalDatabaseParametersOutput)

func TestInitLightsailProvider(t *testing.T) {
	p, err := InitLightsailProvider("us-west-2")
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestGetRecord(t *testing.T) {
	p, err := InitLightsailProvider("us-west-2")
	assert.Nil(t, err)
	assert.NotNil(t, p)
	p.svc = new(LightsailAPI)
	value, err := p.GetRecord("test", "example.com")
	assert.Nil(t, err)
	assert.Equal(t, "10.0.0.1", value)

}
